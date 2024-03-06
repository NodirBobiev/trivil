package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
)

const (
	trivilPath = "/home/cyrus/trivil/src/trivil"
	jasminPath = "/home/cyrus/Downloads/jasmin-2.4/jasmin.jar"
	javaPath   = "/usr/bin/java"
)

type Test interface {
	GetName() string
	GetID() string
	Run() error
}

type TestBase struct {
	Name string
	ID   string
}

func (t *TestBase) GetName() string {
	return t.Name
}
func (t *TestBase) GetID() string {
	return t.ID
}

func (t *TestBase) Run() error {
	return fmt.Errorf("unimplemented")
}

type SimpleTest struct {
	TestBase
	ProgramName string
	PackagePath string
	OutputPath  string
}

func (t *SimpleTest) Run() error {

	genOutput := fmt.Sprintf("/tmp/trivil-outputs/%s-%s", t.ProgramName, time.Now().Format("20060102150405"))

	fmt.Printf("genOutput: %s\n", genOutput)
	stderr, _, err := runCommand(exec.Command(trivilPath, "-out", genOutput, t.PackagePath))
	if err != nil {
		return fmt.Errorf("trivil error: %s", err)
	}
	if stderr != "" {
		return fmt.Errorf("trivil stderr: %s", stderr)
	}
	jasminFiles, err := getJasminFiles(genOutput)
	if err != nil {
		return fmt.Errorf("get jasmin files: %s", err)
	}
	args := append([]string{"-jar", jasminPath, "-d", genOutput}, jasminFiles...)
	cmd := exec.Command(javaPath, args...)

	stderr, _, err = runCommand(cmd)
	//fmt.Println("stdout:", stdout)
	//fmt.Println("stderr:", stderr)
	if err != nil {
		return fmt.Errorf("jasmin error: %s", err)
	}
	if stderr != "" {
		return fmt.Errorf("jasmin stderr: %s", stderr)
	}

	return nil
}

func runCommand(cmd *exec.Cmd) (stderr, stdout string, err error) {
	fmt.Printf("CMD: %q\n", cmd.String())
	var (
		stdoutPipe    io.ReadCloser
		stderrPipe    io.ReadCloser
		stdoutScanner *bufio.Scanner
		stderrScanner *bufio.Scanner
		wg            sync.WaitGroup
	)

	stdoutPipe, err = cmd.StdoutPipe()
	if err != nil {
		return
	}
	stderrPipe, err = cmd.StderrPipe()
	if err != nil {
		return
	}

	stdoutScanner = bufio.NewScanner(stdoutPipe)
	stderrScanner = bufio.NewScanner(stderrPipe)
	err = cmd.Start()
	wg = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for stderrScanner.Scan() {
			stderr += stderrScanner.Text() + "\n"
		}
		for stdoutScanner.Scan() {
			stdout += stdoutScanner.Text() + "\n"
		}
	}()
	wg.Wait()
	err = cmd.Wait()
	if err != nil {
		var exitErr *exec.ExitError
		ok := errors.As(err, &exitErr)
		if ok {
			err = fmt.Errorf("status code: %d: %s", exitErr.ExitCode(), err)
		}
		err = fmt.Errorf("cmd run: %s: %s", cmd.String(), err)
	}

	return
}

func getJasminFiles(dir string) ([]string, error) {
	dirEntries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("os read dir: %s", err)
	}
	files := make([]string, 0)
	for _, d := range dirEntries {
		files = append(files, filepath.Join(dir, d.Name()))
	}
	return files, nil
}

var tests = []Test{
	&SimpleTest{
		ProgramName: "counter",
		PackagePath: "/home/cyrus/trivil/examples/simples/counter",
	},
	&SimpleTest{
		ProgramName: "local_decl",
		PackagePath: "/home/cyrus/trivil/examples/simples/local_decl",
	},
}

func main() {
	for _, t := range tests {
		err := t.Run()
		if err != nil {
			panic(err)
		}
	}
}
