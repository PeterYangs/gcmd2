// +build linux

package gcmd2

import (
	"os/user"
	"runtime"
	"strconv"
	"syscall"
)

func (g *Gcmd2) withSysProcAttr(userName string) error {
	// 检测用户是否存在
	user, err := user.Lookup(userName)

	if err != nil {
		return err
	}

	// set process attr
	// 获取用户 id
	uid, err := strconv.ParseUint(user.Uid, 10, 32)
	if err != nil {
		return err
	}
	// 获取用户组 id
	gid, err := strconv.ParseUint(user.Gid, 10, 32)
	if err != nil {
		return err
	}

	attr := getSysProcAttr()
	//设置进程执行用户
	attr.Credential = &syscall.Credential{
		Uid:         uint32(uid),
		Gid:         uint32(gid),
		NoSetGroups: true,
	}

	g.cmd.SysProcAttr = attr
	return nil
}

// 设置进程组属性
func getSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		Setpgid: true,
	}
}

// SetUser 设置运行用户
func (g *Gcmd2) SetUser(user string) *Gcmd2 {

	//g.user = user
	sysType := runtime.GOOS

	if sysType == "linux" {

		g.withSysProcAttr(user)
	}

	return g
}
