package scenarious

import (
	"bufio"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
	"time"
	"trivil/genjava/tests/common"
)

const (
	trivilPath = "/home/cyrus/trivil/src/trivil"
	jasminPath = "/home/cyrus/Downloads/jasmin-2.4/jasmin.jar"
	javaPath   = "/usr/bin/java"
)

type SimpleTest struct {
	common.TestBase
	ProgramName string
	PackagePath string
	OutputPath  string
}

func (t *SimpleTest) Run() error {
	zap.S().Infow("test is starting")
	genDir := fmt.Sprintf("/tmp/trivil-outputs/%s-%s", t.ProgramName, time.Now().Format("20060102150405"))

	// run trivil to generate jasmin files
	_, err := runAndCheck("trivil", exec.Command(trivilPath, "-out", genDir, t.PackagePath))
	if err != nil {
		return err
	}
	zap.S().Infow("generated", "directory", genDir)

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
	return readAndCheckResults(t.OutputPath, stdout)
}
func readAndCheckResults(filePath string, actualResult string) error {
	expectedResult, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	if string(expectedResult) != actualResult {
		return fmt.Errorf("outputs don't match: expected %q but got %q", expectedResult, actualResult)
	}
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
	zap.S().Infow("running cli command", "cmd", cmd.String())
	var (
		stdoutPipe    io.ReadCloser
		stderrPipe    io.ReadCloser
		stdoutScanner *bufio.Scanner
		stderrScanner *bufio.Scanner
		wg            sync.WaitGroup
	)

	stdoutPipe, _ = cmd.StdoutPipe()
	stderrPipe, _ = cmd.StderrPipe()

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
