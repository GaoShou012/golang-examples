package main

import (
	"github.com/golang/glog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	count := 0
	go func() {
		for{
			time.Sleep(time.Second)
			count++
			if count > 5 {
				os.Exit(1)
			}
		}
	}()

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
		glog.Errorln("exit")
		break
	}
}
