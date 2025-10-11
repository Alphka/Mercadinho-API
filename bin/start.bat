@echo off

swag init
go build -o .\output\mercadinho.exe

.\output\mercadinho.exe
