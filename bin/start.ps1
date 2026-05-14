Start-Process -NoNewWindow -FilePath "swag" -ArgumentList "init" -Wait
Start-Process -NoNewWindow -FilePath "go" -ArgumentList "build -o .\output\mercadinho.exe -ldflags ""-s -w""" -Wait

.\output\mercadinho.exe
