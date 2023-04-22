#!/bin/bash

go build -o vertex-redis

exec go run .
