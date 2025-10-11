Start-Process -NoNewWindow -FilePath "swag" -ArgumentList "init"
Start-Process -NoNewWindow -FilePath "go" -ArgumentList "build -o .\output\mercadinho.exe"

.\output\mercadinho.exe
