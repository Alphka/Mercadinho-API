Start-Process -NoNewWindow -FilePath "swag" -ArgumentList "init" -Wait
Start-Process -NoNewWindow -FilePath "go" -ArgumentList "build -o .\output\mercadinho.exe" -Wait

.\output\mercadinho.exe
