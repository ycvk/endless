//go:build macos

package signals

import "syscall"

const (
	SIGUSR1 = syscall.SIGUSR1
	SIGUSR2 = syscall.SIGUSR2
	SIGTSTP = syscall.SIGTSTP
)
