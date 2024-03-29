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
	GetGroup() string
	GetID() string
	Run() error
}

type TestBase struct {
	Name  string
	Group string
	ID    string
}

func (t *TestBase) GetName() string {
	return t.Name
}
func (t *TestBase) GetID() string {
	return t.ID
}
func (t *TestBase) GetGroup() string {
	return t.Group
}

func (t *TestBase) Run() error {
	return errors.New("unimplemented")
}

type SimpleTest struct {
	TestBase
	ProgramName string
	PackagePath string
	OutputPath  string
}

func (t *SimpleTest) Run() error {

	genDir := fmt.Sprintf("/tmp/trivil-outputs/%s-%s", t.ProgramName, time.Now().Format("20060102150405"))

	fmt.Printf("generated directory: %s\n", genDir)
	// run trivil to generate jasmin files
	_, err := runAndCheck("trivil", exec.Command(trivilPath, "-out", genDir, t.PackagePath))
	if err != nil {
		return err
	}

	// get generated jasmin files
	jasminFiles, err := getJasminFiles(genDir)
	if err != nil {
		return fmt.Errorf("get jasmin files: %s", err)
	}

	// run jasmin to generate java class files
	args := append([]string{"-jar", jasminPath, "-d", genDir}, jasminFiles...)
	cmd := exec.Command(javaPath, args...)
	_, err = runAndCheck("jasmin", cmd)

	if err != nil {
		return err
	}

	cmd = exec.Command(javaPath, "Main", "*")
	cmd.Dir = genDir
	stdout, err := runAndCheck("java", cmd)
	if err != nil {
		return err
	}
	fmt.Printf("STDOUT:\n%s\n------\n", stdout)
	return nil
}

func runAndCheck(name string, cmd *exec.Cmd) (string, error) {
	stderr, stdout, err := runCommand(cmd)
	if err != nil {
		return stdout, fmt.Errorf("%s error: %s", name, err)
	}
	if stderr != "" {
		return stdout, fmt.Errorf("%s stderr: %s", name, stderr)
	}
	return stdout, nil
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
	stdoutScanner.Split(bufio.ScanBytes)
	stderrScanner = bufio.NewScanner(stderrPipe)
	stderrScanner.Split(bufio.ScanBytes)
	err = cmd.Start()
	wg = sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for stderrScanner.Scan() {
			stderr += stderrScanner.Text()
		}
		for stdoutScanner.Scan() {
			stdout += stdoutScanner.Text()
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
		ProgramName: "cчётчик",
		PackagePath: "/home/cyrus/trivil/examples/simples/cчётчик",
		OutputPath:  "/home/cyrus/trivil/examples/simples/cчётчик/output.txt",
	},
	&SimpleTest{
		ProgramName: "арифметика",
		PackagePath: "/home/cyrus/trivil/examples/simples/арифметика",
		OutputPath:  "/home/cyrus/trivil/examples/simples/арифметика/output.txt",
	},
	&SimpleTest{
		ProgramName: "вещ64-тест",
		PackagePath: "/home/cyrus/trivil/examples/simples/вещ64-тест",
		OutputPath:  "/home/cyrus/trivil/examples/simples/вещ64-тест/output.txt",
	},
	&SimpleTest{
		ProgramName: "вывод",
		PackagePath: "/home/cyrus/trivil/examples/simples/вывод",
		OutputPath:  "/home/cyrus/trivil/examples/simples/вывод/output.txt",
	},
	&SimpleTest{
		ProgramName: "импортер",
		PackagePath: "/home/cyrus/trivil/examples/simples/импортер",
		OutputPath:  "/home/cyrus/trivil/examples/simples/импортер/output.txt",
	},
	&SimpleTest{
		ProgramName: "перемена",
		PackagePath: "/home/cyrus/trivil/examples/simples/перемена",
		OutputPath:  "/home/cyrus/trivil/examples/simples/перемена/output.txt",
	},
	&SimpleTest{
		ProgramName: "простой-класс",
		PackagePath: "/home/cyrus/trivil/examples/simples/простой-класс",
		OutputPath:  "/home/cyrus/trivil/examples/simples/простой-класс/output.txt",
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
