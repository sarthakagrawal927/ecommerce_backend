package main

import "strconv"

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
