package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// 假设从用户输入中获取到的密码
	password := "user_password"

	// 生成盐值，可通过随机数生成器生成
	salt := generateSalt()

	// 使用bcrypt进行加密
	hashedPassword, err := GeneratePassword(password, salt)

	// 将哈希密码转为字符串，存储到数据库中
	hashedPasswordString := string(hashedPassword)
	fmt.Println("Hashed Password:", hashedPasswordString)

	// 假设从用户输入中获取到的密码进行验证
	passwordToCheck := "user_password"

	// 验证密码
	err = VertifyPassword(passwordToCheck, hashedPasswordString, salt)
	if err != nil {
		fmt.Println("Password is incorrect.")
	} else {
		fmt.Println("Password is correct.")
	}
}

// 生成密码
func GeneratePassword(rawPassword string, salt string) (string, error) {
	// 将密码和盐值组合
	passwordWithSalt := []byte(rawPassword + salt)
	// 使用bcrypt进行加密
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordWithSalt, bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// 校验密码
func VertifyPassword(rawPassword string, hashPassword string, salt string) error {
	passwordToCheckWithSalt := []byte(rawPassword + salt)
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), passwordToCheckWithSalt)
}

func generateSalt() string {
	// 在实际应用中，应使用适当的随机数生成器生成盐值
	// 这里只是简单返回一个示例值
	return "random_salt"
}
