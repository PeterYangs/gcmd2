package gcmd2

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
)

type gcmd2 struct {
	cxt     context.Context
	command string
	cmd     *exec.Cmd
}

func NewCommand(command string, cxt context.Context) *gcmd2 {

	sysType := runtime.GOOS

	var cmd *exec.Cmd

	// linux/mac
	if sysType == "linux" || sysType == "darwin" {

		cmd = exec.CommandContext(cxt, "bash", "-c", command)
	}

	// windows
	if sysType == "windows" {

		cmd = exec.CommandContext(cxt, "cmd", "/c", command)

	}

	return &gcmd2{
		cxt:     cxt,
		command: command,
		cmd:     cmd,
	}
}

// SetSystemEnv 设置系统环境变量
func (g *gcmd2) SetSystemEnv() *gcmd2 {

	g.cmd.Env = os.Environ()

	return g
}

// SetEnv 设置环境变量
func (g *gcmd2) SetEnv(env []string) *gcmd2 {

	g.cmd.Env = env

	return g
}

func (g *gcmd2) CombinedOutput() ([]byte, error) {

	if g.cmd == nil {

		return []byte{}, errors.New("cmd is nil")
	}

	return g.cmd.CombinedOutput()

}

// Start StartBlock 阻塞等待并输出
func (g *gcmd2) Start() error {

	outIo, err := g.cmd.StdoutPipe()

	if err != nil {

		return err
	}

	go getOut(outIo)

	errIo, err := g.cmd.StderrPipe()

	if err != nil {

		return err
	}

	go getOut(errIo)

	err = g.cmd.Start()

	if err != nil {

		return err
	}

	return g.cmd.Wait()
}

func (g *gcmd2) GetOutPipe() (io.ReadCloser, error) {

	return g.cmd.StdoutPipe()
}

func (g *gcmd2) GetErrPipe() (io.ReadCloser, error) {

	return g.cmd.StderrPipe()
}

func (g *gcmd2) StartNotOut() error {

	err := g.cmd.Start()

	if err != nil {

		return err
	}

	return g.cmd.Wait()

}

// StartNoWait 非阻塞，相当于后台运行
func (g *gcmd2) StartNoWait() error {

	err := g.cmd.Start()

	return err

}

func getOut(st io.ReadCloser) {

	defer st.Close()

	buf := make([]byte, 1024)

	for {

		n, readErr := st.Read(buf)

		if readErr != nil {

			if readErr == io.EOF {

				return
			}

			log.Println(readErr)

			return
		}

		fmt.Print(string(buf[:n]))

	}

}
