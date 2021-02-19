package app

import (
	"../../internal/delivery/http/v1/handler"
	todo "../server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(configPath string) {
	handles := new(handler.Handler)
	srv := new(todo.Server)

	go func() {
		if err := srv.Run("8080", handles.InitRout()); err != nil {
			log.Fatalf("ERROR OCCURRED WHILE RUNNING: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {

	}

}
