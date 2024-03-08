package mysqldb

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB
var err error

// Init initializes the database connection
func Init()   {
	// Database connection parameters
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "goblog"

	// Open a connection to the database
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
 }

// Close closes the database connection
func Close() {
	db.Close()
}

// CreateTable creates a table in the database
func CreateTable(tableName string, columns string) error {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", tableName, columns)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a new record into the specified table
func Insert(tableName string, values string) error {
	query := fmt.Sprintf("INSERT INTO %s VALUES (%s)", tableName, values)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// Update updates a record in the specified table
func Update(tableName string, setValues string, condition string) error {
	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s", tableName, setValues, condition)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the specified table
func Delete(tableName string, condition string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", tableName, condition)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// ExecuteQuery executes an arbitrary SQL query
func ExecuteQuery(query string) (*sql.Rows, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Search executes a search query and returns matching records
func Search(tableName string, condition string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s", tableName, condition)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// ViewAll retrieves all records from the specified table
func ViewAll(tableName string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// AddRelation creates a relationship between two tables based on foreign key constraints
func AddRelation(parentTable, parentColumn, childTable, childColumn string) error {
	// Construct the SQL query to add a foreign key constraint
	query := fmt.Sprintf("ALTER TABLE %s ADD CONSTRAINT fk_%s_%s FOREIGN KEY (%s) REFERENCES %s(%s) ON DELETE CASCADE",
		childTable, parentTable, childTable, childColumn, parentTable, parentColumn)
	// Execute the query
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

// CreateTableIfNotExists creates a table in the database if it doesn't already exist
func CreateTableIfNotExists(tableName string, columns string) error {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", tableName, columns)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// DropTableIfExists drops a table from the database if it exists
func DropTableIfExists(tableName string) error {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// TruncateTable truncates a table in the database (deletes all rows but keeps the table structure)
func TruncateTable(tableName string) error {
	query := fmt.Sprintf("TRUNCATE TABLE %s", tableName)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

// CountRows counts the number of rows in a table
func CountRows(tableName string) (int, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)
	var count int
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// ExecuteNonQuery executes a non-query SQL statement (e.g., INSERT, UPDATE, DELETE)
func ExecuteNonQuery(query string) (sql.Result, error) {
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// MaxValue returns the maximum value of a column in the specified table
func MaxValue(tableName, columnName string) (interface{}, error) {
	var maxVal interface{}
	query := fmt.Sprintf("SELECT MAX(%s) FROM %s", columnName, tableName)
	err := db.QueryRow(query).Scan(&maxVal)
	if err != nil {
		return nil, err
	}
	fmt.Println("Max ", columnName, ":", maxVal)

	return maxVal, nil
}

// MinValue returns the minimum value of a column in the specified table
func MinValue(tableName, columnName string) (interface{}, error) {
	var minVal interface{}
	query := fmt.Sprintf("SELECT MIN(%s) FROM %s", columnName, tableName)
	err := db.QueryRow(query).Scan(&minVal)
	if err != nil {
		return nil, err
	}
	fmt.Println("Min ", columnName, ":", minVal)

	return minVal, nil
}

// SumValue returns the sum of a column in the specified table
func SumValue(tableName, columnName string) (float64, error) {
	var sumVal float64
	query := fmt.Sprintf("SELECT SUM(%s) FROM %s", columnName, tableName)
	err := db.QueryRow(query).Scan(&sumVal)
	if err != nil {
		return 0, err
	}
	fmt.Println("Sum ", columnName, ":", sumVal)

	return sumVal, nil
}

// AverageValue returns the average value of a column in the specified table
func AverageValue(tableName, columnName string) (float64, error) {
	var avgVal float64
	query := fmt.Sprintf("SELECT AVG(%s) FROM %s", columnName, tableName)
	err := db.QueryRow(query).Scan(&avgVal)
	if err != nil {
		return 0, err
	}
	fmt.Println("AvgVal ", columnName, ":", avgVal)

	return avgVal, nil
}

// OrderBy performs a query and orders the result by the specified column
func OrderBy(tableName, orderByColumn string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s ORDER BY %s", tableName, orderByColumn)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Limit performs a query and limits the number of rows returned
func LimitSELECT(tableName string, limit int) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s LIMIT %d", tableName, limit)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// InnerJoin performs an inner join between two tables based on the specified condition
func InnerJoin(table1, table2, condition string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s INNER JOIN %s ON %s", table1, table2, condition)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// LeftJoin performs a left join between two tables based on the specified condition
func LeftJoin(table1, table2, condition string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s LEFT JOIN %s ON %s", table1, table2, condition)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// RightJoin performs a right join between two tables based on the specified condition
func RightJoin(table1, table2, condition string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s RIGHT JOIN %s ON %s", table1, table2, condition)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// CrossJoin performs a cross join between two tables
func CrossJoin(table1, table2 string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s CROSS JOIN %s", table1, table2)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// SelfJoin performs a self join on a table based on the specified condition
func SelfJoin(tableName, condition string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s AS t1 JOIN %s AS t2 ON %s", tableName, tableName, condition)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Union performs a union operation between two queries
func Union(query1, query2 string) (*sql.Rows, error) {
	unionQuery := fmt.Sprintf("%s UNION %s", query1, query2)
	rows, err := db.Query(unionQuery)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// GroupBy performs a query and groups the result by the specified column
func GroupBy(tableName, groupByColumn string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s GROUP BY %s", tableName, groupByColumn)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Having performs a query with the HAVING clause based on the specified condition
func Having(tableName, condition string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s HAVING %s", tableName, condition)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Exists checks if a subquery returns any rows
func Exists(subQuery string) (bool, error) {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (%s)", subQuery)
	err := db.QueryRow(query).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// InsertIntoSelect inserts data from one table into another based on a select query
func InsertIntoSelect(targetTable, selectQuery string) (sql.Result, error) {
	query := fmt.Sprintf("INSERT INTO %s %s", targetTable, selectQuery)
	result, err := db.Exec(query)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CaseStatement performs a query with a CASE statement
func CaseStatement(tableName, columnName, condition string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT *, CASE %s WHEN %s THEN 'Condition True' ELSE 'Condition False' END AS result FROM %s", columnName, condition, tableName)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// LikeOperator performs a query with the LIKE operator
func LikeOperator(tableName, columnName, pattern string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s LIKE '%s'", tableName, columnName, pattern)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// InOperator performs a query with the IN operator
func InOperator(tableName, columnName string, values []interface{}) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s IN (", tableName, columnName)
	for i, val := range values {
		if i != 0 {
			query += ","
		}
		query += fmt.Sprintf("'%v'", val)
	}
	query += ")"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// BetweenOperator performs a query with the BETWEEN operator
func BetweenOperator(tableName, columnName string, lower, upper interface{}) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s BETWEEN '%v' AND '%v'", tableName, columnName, lower, upper)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Aliases performs a query with column aliases
func Aliases(tableName, aliasName string, columnNames []string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT %s FROM %s", aliasName, tableName)
	for _, col := range columnNames {
		query += fmt.Sprintf(", %s AS %s", col, col)
	}
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Join performs a simple join between two tables
func Join(table1, table2, joinCondition string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s JOIN %s ON %s", table1, table2, joinCondition)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// AnySyntaxWithSelect performs a query with ANY syntax in SELECT statement
func AnySyntaxWithSelect(tableName, columnName, operator, subQuery string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s %s ANY (%s)", tableName, columnName, operator, subQuery)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// AllSyntaxWithSelect performs a query with ALL syntax in SELECT statement
func AllSyntaxWithSelect(tableName, columnName, operator, subQuery string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s %s ALL (%s)", tableName, columnName, operator, subQuery)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// AnySyntaxWithWhere performs a query with ANY syntax in WHERE clause
func AnySyntaxWithWhere(tableName, columnName, operator, values string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s %s ANY (%s)", tableName, columnName, operator, values)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// AllSyntaxWithWhere performs a query with ALL syntax in WHERE clause
func AllSyntaxWithWhere(tableName, columnName, operator, values string) (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s %s ALL (%s)", tableName, columnName, operator, values)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// CreateDatabase creates a new database with the given name
func CreateDatabase(dbName string) error {
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
	_, err := db.Exec(query)
	return err
}

// DropDatabase deletes the database with the given name
func DropDatabase(dbName string) error {
	query := fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbName)
	_, err := db.Exec(query)
	return err
}


// DropTable deletes the table with the given name
func DropTable(tableName string) error {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)
	_, err := db.Exec(query)
	return err
}

// AlterTable modifies the structure of an existing table
func AlterTable(tableName, alterStatement string) error {
	query := fmt.Sprintf("ALTER TABLE %s %s", tableName, alterStatement)
	_, err := db.Exec(query)
	return err
}

// CreateView creates a new view in the database
func CreateView(viewName, query string) error {
	query = fmt.Sprintf("CREATE VIEW %s AS %s", viewName, query)
	_, err := db.Exec(query)
	return err
}

func example() {
	groupByRows, err := GroupBy("orders", "customer_id")
	if err != nil {
		log.Fatal(err)
	}
	// Process grouped rows...

	havingRows, err := Having("sales", "SUM(amount) > 1000")
	if err != nil {
		log.Fatal(err)
	}
	// Process rows with having clause...

	exists, err := Exists("SELECT id FROM users WHERE age > 18")
	if err != nil {
		log.Fatal(err)
	}
	if exists {
		fmt.Println("At least one user is above 18 years old")
	} else {
		fmt.Println("No user is above 18 years old")
	}

	insert, err := InsertIntoSelect("new_table", "SELECT * FROM old_table WHERE condition")
	if err != nil {
		log.Fatal(err)
	}

	caseRows, err := CaseStatement("employees", "salary", "WHEN salary > 50000")
	if err != nil {
		log.Fatal(err)
	}
	// Example usage with ANY syntax in SELECT statement
	anyRows, err := AnySyntaxWithSelect("products", "price", "<", "SELECT avg_price FROM average_prices")
	if err != nil {
		log.Fatal(err)
	}
	// Process rows...

	// Example usage with ALL syntax in SELECT statement
	allRows, err := AllSyntaxWithSelect("employees", "salary", ">", "SELECT min_salary FROM salary_ranges")
	if err != nil {
		log.Fatal(err)
	}
	// Process rows...

	// Example usage with ANY syntax in WHERE clause
	anyWhereRows, err := AnySyntaxWithWhere("orders", "total_amount", "<", "SELECT max_amount FROM thresholds")
	if err != nil {
		log.Fatal(err)
	}
	// Process rows...

	// Example usage with ALL syntax in WHERE clause
	allWhereRows, err := AllSyntaxWithWhere("transactions", "amount", ">", "SELECT min_amount FROM thresholds")
	if err != nil {
		log.Fatal(err)
	}
	// Example usage with LIKE operator
	likeRows, err := LikeOperator("products", "name", "App%")
	if err != nil {
		log.Fatal(err)
	}
	// Process rows...

	// Example usage with IN operator
	inRows, err := InOperator("employees", "department_id", []interface{}{1, 2, 3})
	if err != nil {
		log.Fatal(err)
	}
	// Process rows...

	// Example usage with BETWEEN operator
	betweenRows, err := BetweenOperator("orders", "total_amount", 100, 500)
	if err != nil {
		log.Fatal(err)
	}
	// Process rows...

	// Example usage with Aliases
	aliasRows, err := Aliases("employees", "EmpDetails", []string{"name", "salary"})
	if err != nil {
		log.Fatal(err)
	}
	// Process rows...

	// Example usage with Join
	joinRows, err := Join("orders", "customers", "orders.customer_id = customers.id")
	if err != nil {
		log.Fatal(err)
	}
	// Example usage of creating a database
	err3 := CreateDatabase("new_database")
	if err3 != nil {
		log.Fatal(err)
	}

	// Example usage of dropping a database
	err4 := DropDatabase("old_database")
	if err4 != nil {
		log.Fatal(err)
	}

	// Example usage of creating a table
	err5 := CreateTable("users", "id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), age INT")
	if err5 != nil {
		log.Fatal(err)
	}

	// Example usage of dropping a table
	err6 := DropTable("old_table")
	if err6 != nil {
		log.Fatal(err)
	}

	// Example usage of altering a table
	err7 := AlterTable("users", "ADD COLUMN email VARCHAR(255)")
	if err7 != nil {
		log.Fatal(err)
	}

	// Example usage of creating a view
	err8 := CreateView("view_name", "SELECT * FROM users WHERE age > 18")
	if err8 != nil {
		log.Fatal(err)
	}

	fmt.Println(groupByRows, havingRows, insert, caseRows, joinRows, aliasRows, betweenRows, inRows, likeRows, allWhereRows, anyWhereRows, anyRows, allRows, likeRows)

}