# **Renference**

- `This project is golang gin demo`

## running
    1. mkdir golang-gin.
    2. cd golang-gin.
    3. go mod init vigourzhao/gin-demo.
    4. create main.go and coding.
    5. CLI excute `go get github.com/gin-gonic/gin`
    6. go run .
    7. open browser `http://localhost:8080/`

## swagger config:
    1. go get -u github.com/swaggo/swag/cmd/swag
    tips: if you Command 'swag' not found, 
        check /usr/local/go/bin or $HOME/go/bin where those swag executable is present.
        Make sure you have that PATH in /etc/profile or $HOME/.profile
        e.g. export PATH=$PATH:$HOME/go/bin
    2. excute `swag init`
    3. excute `go get -u github.com/swaggo/gin-swagger`
    4. excute `go get -u github.com/swaggo/files`
    5. go run .
    6. open browser `http://localhost:8080/swagger/index.html#/`
