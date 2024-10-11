package common

import "syscall"

func Umask(mask int) int {
	return syscall.Umask(mask)
}
