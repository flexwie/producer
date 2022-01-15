package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"felixwie.com/producer/api"
	"felixwie.com/producer/client"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/gin-contrib/cors"
	"github.com/gliderlabs/ssh"
)

const host = "localhost"
const port = 2222

func main() {
	router := api.GetRouter()
	router.Use(cors.Default())

	s, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf("%s:%d", host, port)),
		wish.WithMiddleware(
			authHandler(),
			bm.Middleware(tuiHandler),
		),
	)

	if err != nil {
		log.Fatalln(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Starting SSH server on %s:%d", host, port)
	go func() {
		if err = s.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	<-done
	log.Println("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}

	router.Run(":8080")
}

func authHandler() wish.Middleware {
	return func(h ssh.Handler) ssh.Handler {
		return ssh.DefaultHandler
	}
}

func tuiHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, active := s.Pty()
	if !active {
		fmt.Println("no active terminal, skipping")
		return nil, nil
	}
	m := client.GetClient(pty, s.User())
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}
