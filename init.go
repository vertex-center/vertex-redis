package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/vertex-center/vertex-core-golang/console"
)

var logger = console.New("vertex-redis")

type Environment struct {
	Port string `env:"PORT" envDefault:"6379"`
}

var environment Environment

var cmd *exec.Cmd

func main() {
	handleSignals()

	loadEnv()

	cmd = exec.Command("redis-server", "--port", environment.Port)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func handleSignals() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-c
		_ = cmd.Process.Kill()
		os.Exit(0)
	}()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		logger.Error(fmt.Errorf("error loading .env file: %v", err))
	}

	err = env.Parse(&environment)
	if err != nil {
		logger.Error(fmt.Errorf("failed to parse .env to Config: %v", err))
	}
}
