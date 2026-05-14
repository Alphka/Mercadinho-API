#!/bin/bash

swag init
go build -o ./output/mercadinho -ldflags "-s -w"

./output/mercadinho
