package main

import (
	// sqldb "main/core/db/sql"
	address "main/features/address"
	addressmongodb "main/features/addressmongodb"
	user "main/features/user"
)

///releations///

// CreateTables creates both users and addresses tables and establishes a relationship between them
func InitDataBase() error {
	user.InitData()
	address.InitData()
	addressmongodb.InitData()

	// sqldb.AddRelation("users", "id", "addresses", "user_id")

	return nil
}
