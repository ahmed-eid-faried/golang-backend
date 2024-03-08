package mysqldb

import (
	"database/sql"
	"fmt"
)

////////////////////////////////////examples////////////////////////////////////////
// ///////////////////////////// /////////// ///////////////////////////////////////
// ///////////////////////////// /////////// ///////////////////////////////////////

///User///
// CreateUserTable creates the users table in the database
func CreateUserTable() error {
	return CreateTable("users", "id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255)")
}

// InsertUser inserts a new user into the users table
func InsertUser(name string) error {
	return Insert("users", fmt.Sprintf("'%s'", name))
}

// GetAllUsers retrieves all users from the users table
func GetAllUsers() (*sql.Rows, error) {
	return ViewAll("users")
}

// UpdateUser updates a user record in the users table
func UpdateUser(userID int, name string) error {
	query := fmt.Sprintf("UPDATE users SET name='%s' WHERE id=%d", name, userID)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes a user record from the users table
func DeleteUser(userID int) error {
	return Delete("users", fmt.Sprintf("id=%d", userID))
}

// DropUsersTable drops the users table from the database if it exists
func DropUsersTable() error {
	return DropTableIfExists("users")
}

// ExecuteQuery executes an arbitrary SQL query
func ExecuteQueryUser(query string) (*sql.Rows, error) {
	return ExecuteQuery(query)

}

// SearchUsers executes a search query on the users table and returns matching records
func SearchUsers(condition string) (*sql.Rows, error) {
	return Search("users", condition)
}

// CreateTableIfNotExistsUsers creates the users table if it doesn't already exist
func CreateTableIfNotExistsUsers() error {
	return CreateTableIfNotExists("users", "id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255)")
}

// TruncateUsersTable truncates the users table, deleting all rows but keeping the table structure intact
func TruncateUsersTable() error {
	return TruncateTable("users")
}

// CountUsersRows counts the number of rows in the users table
func CountUsersRows() (int, error) {
	return CountRows("users")
}

// ExecuteNonQuery executes a non-query SQL statement (e.g., INSERT, UPDATE, DELETE) for the users table
func ExecuteNonQueryUser(query string) (sql.Result, error) {
	return ExecuteNonQuery(query)
}

// ///////////////////////////// /////////// ///////////////////////////////////////
// ///////////////////////////// /////////// ///////////////////////////////////////
// ///////////////////////////// /////////// ///////////////////////////////////////

// /Address///
// CreateAddressTable creates the addresses table in the database
func CreateAddressTable() error {
	return CreateTable("addresses", "id INT AUTO_INCREMENT PRIMARY KEY, user_id INT, street VARCHAR(255), city VARCHAR(255), state VARCHAR(255), zipcode VARCHAR(255)")
}

// InsertAddress inserts a new address into the addresses table
func InsertAddress(userID int, street, city, state, zipcode string) error {
	values := fmt.Sprintf("%d, '%s', '%s', '%s', '%s'", userID, street, city, state, zipcode)
	return Insert("addresses", values)
}

// GetAllAddresses retrieves all addresses from the addresses table
func GetAllAddresses() (*sql.Rows, error) {
	return ViewAll("addresses")
}
// ///////////////////////////// /////////// ///////////////////////////////////////
// ///////////////////////////// /////////// ///////////////////////////////////////
// ///////////////////////////// /////////// ///////////////////////////////////////
