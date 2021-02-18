package app

import (
	"context"
	todo "github.com/alicancho/my-go-invest-server/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(configPath string) {
	srv := new(todo.Server)

	go func() {
		if err := srv.Run("8080"); err != nil {
			log.Fatalf("ERROR OCCURRED WHILE RUNNING: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {

	}

}
