package main

import "github.com/app"

type User struct {
	Username *string
	Password *string
}

func main(){
	db,err := app.Gorm()
	if err != nil {
		panic(err)
	}

	username := "abc"
	password := "pwd"

	res := db.Create(&User{
		Username: &username,
		Password: &password,
	})
	if res.Error != nil{
		panic(res.Error)
	}
}

