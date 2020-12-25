package config

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

var reloadsignal chan os.Signal

func InitReloadConfiguration() {
	reloadsignal = make(chan os.Signal, 1)
	signal.Notify(reloadsignal, syscall.SIGUSR1)
	go reloadConfiguration()
}

func reloadConfiguration() {
	for {
		select {
		case <-reloadsignal:
			err := Load()
			if err != nil {
				panic(err)
			}
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
