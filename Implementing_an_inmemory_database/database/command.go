package database

import "strings"

// 系统全部指令
var cmdTable = make(map[string]*command)

type command struct {
	exector ExecFunc
	// 参数数量
	arity int
}

func RegisterCommand(name string, exector ExecFunc, arity int) {
	name = strings.ToLower(name)
	cmdTable[name] = &command{
		exector: exector,
		arity:   arity,
	}
}
