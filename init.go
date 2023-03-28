package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/vertex-center/vertex-core-golang/console"
)

var logger = console.New("vertex-redis")

type Environment struct {
	Port string `env:"PORT" envDefault:"6379"`
}

var environment Environment

func main() {
	loadEnv()

	command := exec.Command("redis-server", "--port", environment.Port)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin

	err := command.Run()
	if err != nil {
		panic(err)
	}
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
