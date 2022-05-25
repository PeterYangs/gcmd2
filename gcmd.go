package gcmd2

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
)

type Gcmd2 struct {
	cxt     context.Context
	command string
	cmd     *exec.Cmd
	user    string
}

func NewCommand(command string, cxt context.Context) *Gcmd2 {

	sysType := runtime.GOOS

	var cmd *exec.Cmd

	// linux/mac
	if sysType == "linux" || sysType == "darwin" {

		cmd = exec.CommandContext(cxt, "/bin/bash", "-c", command)
	}

	// windows
	if sysType == "windows" {

		cmd = exec.CommandContext(cxt, "cmd", "/c", command)

	}

	return &Gcmd2{
		cxt:     cxt,
		command: command,
		cmd:     cmd,
	}
}

// SetSystemEnv 设置系统环境变量
func (g *Gcmd2) SetSystemEnv() *Gcmd2 {

	g.cmd.Env = os.Environ()

	return g
}

// SetEnv 设置环境变量
func (g *Gcmd2) SetEnv(env []string) *Gcmd2 {

	g.cmd.Env = env

	return g
}

func (g *Gcmd2) CombinedOutput() ([]byte, error) {

	if g.cmd == nil {

		return []byte{}, errors.New("cmd is nil")
	}

	return g.cmd.CombinedOutput()

}

// Start StartBlock 阻塞等待并输出
func (g *Gcmd2) Start() error {

	outIo, err := g.cmd.StdoutPipe()

	if err != nil {

		return err
	}

	//实时输出
	go getOut(outIo)

	errIo, err := g.cmd.StderrPipe()

	if err != nil {

		return err
	}

	//实时输出
	go getOut(errIo)

	err = g.cmd.Start()

	if err != nil {

		return err
	}

	return g.cmd.Wait()
}

func (g *Gcmd2) GetOutPipe() (io.ReadCloser, error) {

	return g.cmd.StdoutPipe()
}

func (g *Gcmd2) GetErrPipe() (io.ReadCloser, error) {

	return g.cmd.StderrPipe()
}

func (g *Gcmd2) StartNotOut() error {

	err := g.cmd.Start()

	if err != nil {

		return err
	}

	return g.cmd.Wait()

}

// StartNoWait 非阻塞，相当于后台运行
func (g *Gcmd2) StartNoWait() error {

	err := g.cmd.Start()

	return err

}

func (g Gcmd2) StartNoWaitOutErr() error {

	errIo, err := g.cmd.StderrPipe()

	if err != nil {

		return err
	}

	//实时输出
	go getOut(errIo)

	err = g.cmd.Start()

	if err != nil {

		return err
	}

	return nil

}

func (g *Gcmd2) GetCmd() *exec.Cmd {

	return g.cmd
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

			//log.Println(readErr)

			return
		}

		fmt.Print(string(buf[:n]))

	}

}
