package main

import (
	"context"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/proto"
)

// rewrite this to take postgresData
func getAllUsersFromDB() []byte {
	defer Recovery()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := userDetailsCollection.Find(ctx, bson.M{})
	handleError(err)
	defer cur.Close(ctx)
	var usersProto []*UserProto
	for cur.Next(ctx) {
		var result UserDetails
		err := cur.Decode(&result)
		handleError(err)
		userProto := &UserProto{
			Name:  result.Name,
			Age:   int32(result.Age),
			Email: result.Email,
		}
		usersProto = append(usersProto, userProto)
	}
	userListProto, err := proto.Marshal(&UserListProto{Users: usersProto})
	handleError(err)
	return userListProto
}

func getUserByMail(email string) UserDetails {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	var user UserDetails
	err := userDetailsCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	handleError(err)
	defer Recovery()
	return user
}

// to be rewritten to support extra values
func createUserDetailsDB(name string, age string, email string) UserDetails {
	ageInt, err := strconv.ParseInt(age, 10, 16)
	handleError(err)
	user := UserDetails{
		Name:   name,
		Age:    int16(ageInt),
		Email:  email,
		Active: true,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = userDetailsCollection.InsertOne(ctx, user)
	handleError(err)
	return user
}

func deleteUserByEmailDB(email string) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := userDetailsCollection.UpdateOne(ctx, bson.M{"email": email}, bson.M{"$set": bson.M{"active": false}})
	handleError(err)
	return res.UpsertedCount
}
