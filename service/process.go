package service

import (
	"os"
	"syscall"
)

const (
	SYSTEM_SIGNAL_CHANNEL_SIZE = 256
)

var (
	system_signal_channel = make(chan os.Signal, SYSTEM_SIGNAL_CHANNEL_SIZE)
)

func Process(f func()) {
	for system_signal := range system_signal_channel {
		if system_signal == syscall.SIGINT || system_signal == syscall.SIGTERM {
			f()
		}
	}
}
