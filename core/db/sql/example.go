package sqldb

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
	_, err := DB.Exec(query)
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
// func example() {
// 	groupByRows, err := GroupBy("orders", "customer_id")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Process grouped rows...

// 	havingRows, err := Having("sales", "SUM(amount) > 1000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Process rows with having clause...

// 	exists, err := Exists("SELECT id FROM users WHERE age > 18")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if exists {
// 		fmt.Println("At least one user is above 18 years old")
// 	} else {
// 		fmt.Println("No user is above 18 years old")
// 	}

// 	insert, err := InsertIntoSelect("new_table", "SELECT * FROM old_table WHERE condition")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	caseRows, err := CaseStatement("employees", "salary", "WHEN salary > 50000")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Example usage with ANY syntax in SELECT statement
// 	anyRows, err := AnySyntaxWithSelect("products", "price", "<", "SELECT avg_price FROM average_prices")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Process rows...

// 	// Example usage with ALL syntax in SELECT statement
// 	allRows, err := AllSyntaxWithSelect("employees", "salary", ">", "SELECT min_salary FROM salary_ranges")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Process rows...

// 	// Example usage with ANY syntax in WHERE clause
// 	anyWhereRows, err := AnySyntaxWithWhere("orders", "total_amount", "<", "SELECT max_amount FROM thresholds")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Process rows...

// 	// Example usage with ALL syntax in WHERE clause
// 	allWhereRows, err := AllSyntaxWithWhere("transactions", "amount", ">", "SELECT min_amount FROM thresholds")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Example usage with LIKE operator
// 	likeRows, err := LikeOperator("products", "name", "App%")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Process rows...

// 	// Example usage with IN operator
// 	inRows, err := InOperator("employees", "department_id", []interface{}{1, 2, 3})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Process rows...

// 	// Example usage with BETWEEN operator
// 	betweenRows, err := BetweenOperator("orders", "total_amount", 100, 500)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Process rows...

// 	// Example usage with Aliases
// 	aliasRows, err := Aliases("employees", "EmpDetails", []string{"name", "salary"})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Process rows...

// 	// Example usage with Join
// 	joinRows, err := Join("orders", "customers", "orders.customer_id = customers.id")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Example usage of creating a database
// 	err3 := CreateDatabase("new_database")
// 	if err3 != nil {
// 		log.Fatal(err)
// 	}

// 	// Example usage of dropping a database
// 	err4 := DropDatabase("old_database")
// 	if err4 != nil {
// 		log.Fatal(err)
// 	}

// 	// Example usage of creating a table
// 	err5 := CreateTable("users", "id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), age INT")
// 	if err5 != nil {
// 		log.Fatal(err)
// 	}

// 	// Example usage of dropping a table
// 	err6 := DropTable("old_table")
// 	if err6 != nil {
// 		log.Fatal(err)
// 	}

// 	// Example usage of altering a table
// 	err7 := AlterTable("users", "ADD COLUMN email VARCHAR(255)")
// 	if err7 != nil {
// 		log.Fatal(err)
// 	}

// 	// Example usage of creating a view
// 	err8 := CreateView("view_name", "SELECT * FROM users WHERE age > 18")
// 	if err8 != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(groupByRows, havingRows, insert, caseRows, joinRows, aliasRows, betweenRows, inRows, likeRows, allWhereRows, anyWhereRows, anyRows, allRows, likeRows)

// }