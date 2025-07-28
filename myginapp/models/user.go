package models

import (
    "errors"
    "github.com/google/uuid"
)

// 用户模型
type User struct {
    ID    string `json:"id"`
    Name  string `json:"name" binding:"required,max=100"`
    Email string `json:"email" binding:"required,email"`
}

// 模拟数据库
var users = make([]User, 0)

// 用户不存在错误
var ErrNotFound = errors.New("user not found")

// 新建用户，生成 UUID 并返回
func CreateUser(u User) User {
    u.ID = uuid.NewString()
    users = append(users, u)
    return u
}

// 返回所有用户
func GetAllUsers() []User {
    return users
}

// 根据 ID 查找用户
func GetUserByID(id string) (User, error) {
    for _, u := range users {
        if u.ID == id {
            return u, nil
        }
    }
    return User{}, ErrNotFound
}

// 更新指定 ID 的用户信息
func UpdateUser(id string, input User) (User, error) {
    for i, u := range users {
        if u.ID == id {
            users[i].Name = input.Name
            users[i].Email = input.Email
            return users[i], nil
        }
    }
    return User{}, ErrNotFound
}

// 从切片中删除指定 ID 的用户
func DeleteUser(id string) error {
    for i, u := range users {
        if u.ID == id {
            users = append(users[:i], users[i+1:]...)
            return nil
        }
    }
    return ErrNotFound
}
