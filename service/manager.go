package service

import (
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"
)

/*User 用户*/
type User struct {
    Name                string
    Email               string
    Phone               string
}

/*READUSERS  读取用户*/
func READUSERS() (user []User) {
    _, err := os.Getwd()
    checkerr(err)
    b, err := ioutil.ReadFile("Database/Users.txt")
    checkerr(err)
    //json转变为对象
    var users []User
    json.Unmarshal(b, &users)
    return users
}

/*WRITEUSER 写回用户*/
func WRITEUSER(users []User) {
    _, err := os.Getwd()
    checkerr(err)
    data, err := json.Marshal(users)
    checkerr(err)
    b := []byte(data)
    err = ioutil.WriteFile("Database/Users.txt", b, 0777)
    checkerr(err)
}

/*check 检查并输出错误*/
func checkerr(err error) {
    if err != nil {
        fmt.Print(err)
    }
}

/*Manager 处理指令
    regist 注册，需要信息：名字 Email 电话
    chcek  查询，需要信息：查询名字 
    update 更新数据，需要信息：名字 Email 电话
    接收参数：指令 名字 email 电话
    返回参数：
    regist 注册成功/失败
    check  查询结果
    update 更新成功/失败
*/
func Manager(command string, name string, email string, phone string) string {
    users := READUSERS()
    switch command {
        case "regist" : {
            //遍历查询，没有找到才可以注册
            for _, user := range users {
                if user.Name == name {
                    return "This name has been regist!"
                }
            }
            newUser := User{Name: name, Email: email, Phone: phone}
            users = append(users, newUser)
            WRITEUSER(users)
            return "Success! Have regist!"
        }
        case "check" : {
            //遍历查询，找到名字就返回结果
            for _, user := range users {
                if user.Name == name {
                    return "Name: " + user.Name + " Email: " + user.Email + " Phone: " + user.Phone
                }
            }
            return "Not regist!"
        }
        case "update" : {
            //遍历查询，找到名字就更新信息
            for i, user := range users {
                if user.Name == name {
                    users[i].Email = email
                    users[i].Phone = phone
                    WRITEUSER(users)
                    return "Update Success!"
                }
                return "Wrong Name!"
            }
            return "Not regist!"
        }
        default : return " "
    }
}