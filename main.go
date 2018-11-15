package main

import (
    "os"
    "encoding/json"
    "fmt"
    "io/ioutil"

    "github.com/github-user/cloudgo/service"
    flag "github.com/spf13/pflag"
)

const (
    PORT string = "8080"
)

/*User 用户*/
type User struct {
    Username           string
    Password           string
    Email              string
    Phone              string
}

/*READUSERS  读取用户*/
func READUSERS() (user []User) {
    b, err := ioutil.ReadFile("/Database/Users.txt")
    check(err)
    //json转变为对象
    var users []User
    json.Unmarshal(b, &users)
    return users
}

/*WRITEUSER 写回数据*/
func WRITEUSER(users []User) {
    data := json.Marshal(users)
    b := []byte(data)
    err := ioutil.WriteFile("/Database/Users.txt", b, 0777)
    check(err)
}

func main() {
    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = PORT
    }

    pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
    flag.Parse()
    if len(*pPort) != 0 {
        port = *pPort
    }

    server := service.NewServer()
    server.Run(":" + port)
}