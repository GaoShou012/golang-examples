package main

import (
	"fmt"
	"github.com/golang/glog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func task() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now())
		}
	}
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		switch s := <-c; s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			glog.Infof("got signal %s; stop server", s)
		case syscall.SIGHUP:
			glog.Infof("got signal %s; go to deamon", s)
			continue
		}
		break
	}
}
