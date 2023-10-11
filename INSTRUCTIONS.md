Angular and Golang: A Rapid Guide - Advanced

Run on WSL Ubuntu - ang-go

```erlang
# VSC - Go
# Extention - Go

Fix 
go get -u golang.org/x/tools/gopls
```

1. Setup
    
    ```erlang
    # Windows Install 
    https://go.dev/dl/ - Microsoft Windows
    
    # WSL - Ubuntu Install
    https://medium.com/@benzbraunstein/how-to-install-and-setup-golang-development-under-wsl-2-4b8ca7720374
    
    wget https://dl.google.com/go/go1.16.1.linux-amd64.tar.gz
    sudo tar -xvf go1.16.1.linux-amd64.tar.gz
    sudo mv go /usr/local
    
    cd ~
    explorer.exe .
    
    # Edit .bashrc
    """
    ...
    export GOROOT=/usr/local/go
    export GOPATH=$HOME/go
    export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
    """
    
    go version
    
    # Go
    mkdir ang-go && cd ang-go
    mkdir go-ambassador && cd go-ambassador
    go mod init ambassador # create go.mod
    go get github.com/gofiber/fiber/v2
    go run main.go
    
    # http://localhost:3000/
    ```
        
    
3. Docker
    
    ```erlang
    sudo docker-compose up --build
    # http://localhost:8000
    ```
        
4. Database
    
    ```erlang
    sudo docker-compose up --build # twice
    # http://localhost:8000
    ```
        
5. Migrations
    
    ```erlang
    # HeidiSQL
    https://www.heidisql.com/download.php
    
    # Open to MYSQL
    """
    HOST: localhost
    USER/PW: root
    PORT: 33066
    DATABASE: ambassador
    """
    
    # Add table
    
    > Files
    
    sudo docker-compose up --build
    ```
        
6. Live Reloading
    
    ```erlang
    # Live Reload - https://thedevelopercafe.com/articles/live-reload-in-go-with-air-4eff64b7a642
    # Github - https://github.com/cosmtrek/air
    
    sudo docker-compose up --build
    ```
        
8. Routes
    
    ```erlang
    # Postman
    
    POST - http://localhost:8000/api/admin/register
    """
    {
        "message": "hello"
    }
    """
    ```
        
9. Register
    
    ```erlang
    # GO Package - Bcrypt
    https://pkg.go.dev/
    go get golang.org/x/crypto/bcrypt@v0.0.0-20210513164829-c07d793c2f9a
    
    # Validate - Postman
    POST - http://localhost:8000/api/admin/register
    {
        "first_name": "a",
        "last_name": "a",
        "email": "a@a.com",
        "password": "a",
        "password_confirm": "a2"
    }
    ...
    {
        "message": "passwords do not match"
    }
    
    {
        "first_name": "a",
        "last_name": "a",
        "email": "a@a.com",
        "password": "a",
        "password_confirm": "a"
    }
    ...
    {
        "Id": 1,
        "FirstName": "a",
        "LastName": "a",
        "Email": "a@a.com",
        "Password": "JDJhJDEyJHNpLm5ncTFEMkRlcnEuWlRBN1dTWk9BeXhYeHpodzFqc3dVWTVTVmYyM3FoS0NQcWlySkF1",
        "IsAmbassador": false
    }
    ```
        
10. Login
    ```erlang
    # Validate - Postman
    POST - http://localhost:8000/api/admin/login
    
    # Wrong password
    {
        "email": "a@a.com",
        "password": "b",
    }
    ...
    {
        "message": "Invalid credentials"
    }
    
    # Wrong email
    {
        "email": "a@a1.com",
        "password": "a",
    }
    ...
    {
        "message": "Invalid credentials"
    }
    
    # Right
    {
        "email": "a@a.com",
        "password": "a",
    }
    ...
    {
        "Id": 1,
        "FirstName": "a",
        "LastName": "a",
        "Email": "a@a.com",
        "Password": "JDJhJDEyJHNpLm5ncTFEMkRlcnEuWlRBN1dTWk9BeXhYeHpodzFqc3dVWTVTVmYyM3FoS0NQcWlySkF1",
        "IsAmbassador": false
    }
    ```
    
11. Jwt
    
    ```erlang
    POST - http://localhost:8000/api/admin/login
    
    # Token returned
    {
        "email": "a@a.com",
        "password": "a"
    }
    ...
    {  "message": "success" } - COOKIE
    ```
    
12. Methods (SetPassword, ComparePassword)
    
    ```erlang
    # Validate - Postman
    POST - http://localhost:8000/api/admin/register
    
    # Right
    {
        "first_name": "a2",
        "last_name": "a2",
        "email": "a@a2.com",
        "password": "a",
        "password_confirm": "a"
    }
    ...
    { "user": "userInfo..." }
    ```
        
13. Authenticated User
    
    ```erlang
    GET - http://localhost:8000/api/admin/user
    
    # Token returned
    {
        "email": "a@a.com",
        "password": "a"
    }
    ...
    {  "message": "success" } - COOKIE
    ```
        
14. Formatting Output - (user.go)
    
    ```erlang
    # GET - http://localhost:8000/api/admin/user
    
    {
        "id": 4,
        "first_name": "sa",
        "last_name": "a",
        "email": "a@a.com"
    }
    ```