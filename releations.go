package main

import (
	// sqldb "main/core/db/sql"
	user "main/features/user"
)

///releations///

// CreateTables creates both users and addresses tables and establishes a relationship between them
func InitDataBase() error {
	user.InitDataUsers()

	// sqldb.AddRelation("users", "id", "addresses", "user_id")

	return nil
}
