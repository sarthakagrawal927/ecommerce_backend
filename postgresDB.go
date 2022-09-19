package main

import (
	"strconv"

	"github.com/golang/protobuf/proto"
)

func createUserDB(name string, age string, email string) User {
	ageInt, err := strconv.ParseUint(age, 10, 8)
	handleError(err)
	user := &User{
		Name:  name,
		Age:   uint8(ageInt),
		Email: email,
	}
	postgresDB.Create(user)
	return *user
}

func getAllUsers() []byte {
	var users []User
	postgresDB.Find(&users)
	var usersProto []*UserProto
	for i, _ := range users {
		userProto := &UserProto{
			Name:   users[i].Name,
			Age:    int32(users[i].Age),
			Email:  users[i].Email,
			Active: users[i].Active,
		}
		usersProto = append(usersProto, userProto)
	}
	userListProto, err := proto.Marshal(&UserListProto{Users: usersProto})
	handleError(err)
	return userListProto
}
