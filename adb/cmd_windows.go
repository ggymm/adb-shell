package adb

import (
	"bufio"
	"io"
	"os/exec"
	"syscall"
)

func Exec(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func ExecAsync(name string, args ...string) (chan string, chan error) {
	cmd := exec.Command(name, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{}

	errChan := make(chan error)
	outChan := make(chan string, 1024)
	go func() {
		var (
			err    error
			stderr io.ReadCloser
			stdout io.ReadCloser
		)

		// stderr
		stderr, err = cmd.StderrPipe()
		if err != nil {
			errChan <- err
			return
		}

		// stdout
		stdout, err = cmd.StdoutPipe()
		if err != nil {
			errChan <- err
			return
		}

		// run command
		err = cmd.Start()
		if err != nil {
			errChan <- err
			return
		}

		// read command output
		if stderr != nil {
			go func() {
				s := bufio.NewScanner(stderr)
				for s.Scan() {
					outChan <- s.Text()
				}
				err = s.Err()
				if err != nil {
					errChan <- err
				}
			}()
		}
		if stdout != nil {
			go func() {
				s := bufio.NewScanner(stdout)
				for s.Scan() {
					outChan <- s.Text()
				}
				err = s.Err()
				if err != nil {
					errChan <- err
				}
			}()
		}

		// wait for command to finish
		err = cmd.Wait()
		if err != nil {
			errChan <- err
		}
	}()
	return outChan, errChan
}
