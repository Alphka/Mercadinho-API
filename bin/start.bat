@echo off

swag init
go build -o .\output\mercadinho.exe -ldflags "-s -w"

.\output\mercadinho.exe
