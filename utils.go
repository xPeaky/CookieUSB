package main

import (
	"fmt"

	"github.com/fatih/color"
)

const (
	Success = iota
	Error
	Warning
	Task
	Normal
)

func Debug(message string, level int) {
	switch level {
	case Success:
		fmt.Println(color.HiGreenString("[+]"), message)
	case Error:
		fmt.Println(color.HiRedString("[-]"), message)
	case Warning:
		fmt.Println(color.HiYellowString("[!]"), message)
	case Task:
		fmt.Println(color.HiCyanString("[~]"), message)
	case Normal:
		fmt.Println(color.HiMagentaString("[*]"), message)
	}
}