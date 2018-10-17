package event

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

const (
	EXIT = iota
	WAIT
	MAX
)

var (
	Events = make([][]func(interface{}), MAX+1)
	Ename = map[int]string{
		EXIT: "exit",
		WAIT: "wait",
		MAX:  "max",
	}
)

func On(name int, fs ...func(interface{})) error {
	if name > MAX || name < 0 {
		return errors.New("bad name")
	}

	for _, f := range fs {
		fp := reflect.ValueOf(f).Pointer()
		for i := 0; i < len(Events[name]); i++ {
			if reflect.ValueOf(Events[name][i]).Pointer() == fp {
				return fmt.Errorf("%v func already exists in %s event ", fp, Ename[name])
			}
		}
		Events[name] = append(Events[name], f)
	}
	return nil
}

func Emit(name int, arg interface{}) error {
	if name > MAX || name < 0 {
		return errors.New("bad name")
	}
	for _, f := range Events[name] {
		f(arg)
	}
	return nil
}

func EmitAll(arg interface{}) {
	for _, fs := range Events {
		for _, f := range fs {
			f(arg)
		}
	}
	return
}

func Off(name int, f func(interface{})) error {
	if name > MAX || name < 0 {
		return errors.New("bad name")
	}
	fp := reflect.ValueOf(f).Pointer()
	for i := 0; i < len(Events[name]); i++ {
		if reflect.ValueOf(Events[name][i]).Pointer() == fp {
			Events[name] = append(Events[name][:i], Events[name][i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("%v func dones't exist in %s event ", fp, Ename[name])
}

func OffAll(name int) error {
	if name > MAX || name < 0 {
		return errors.New("bad name")
	}
	Events[name] = nil
	return nil
}

// 等待信号
// 如果信号参数为空，则会等待常见的终止信号
// SIGHUP 1 A 终端挂起或者控制进程终止
// SIGINT 2 A 键盘中断（如break键被按下）
// SIGQUIT 3 C 键盘的退出键被按下
// SIGKILL 9 AEF Kill信号
// SIGTERM 15 A 终止信号
// SIGSTOP 17,19,23 DEF 终止进程
func Wait(sig ...os.Signal) os.Signal {
	c := make(chan os.Signal, 1)
	if len(sig) == 0 {
		signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGSTOP)
	} else {
		signal.Notify(c, sig...)
	}
	return <-c
}

func WaitExit() {
	On(MAX, func(interface{}) {
		log.Print("all events has triggered")
		os.Exit(255)
	})
	Wait()
	EmitAll(nil)
}
