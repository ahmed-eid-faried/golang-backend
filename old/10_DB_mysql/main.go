package main

import (
	"database/sql"
	"fmt"
	"log"
)

// تعريف نوع Employee لتمثيل بيانات الموظف
type Employee struct {
	ID   int
	Name string
	City string
}

var db *sql.DB
var err error

// dbConn تأسيس الاتصال بقاعدة البيانات
func dbConn() {
	// معلومات الاتصال بقاعدة البيانات
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "goblog"

	// فتح الاتصال بقاعدة البيانات
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Fatal(err)
	}
}

// getAllEmployees استرجاع كل الموظفين من قاعدة البيانات
func getAllEmployees(name string) []Employee {
	// استعلام SQL لاسترجاع كل الموظفين
	row, err := db.Query("SELECT * FROM name=?")
	// row, err := db.Query("SELECT * FROM employee")
	if err != nil {
		log.Fatal(err)
	}
	emp := Employee{}
	employees := []Employee{}
	// معالجة كل صف من النتائج
	for row.Next() {
		// قراءة بيانات الموظف من الصف
		err := row.Scan(&emp.ID, &emp.Name, &emp.City)
		if err != nil {
			log.Fatal(err)
		}
		// إضافة الموظف إلى قائمة الموظفين
		employees = append(employees, emp)
	}
	return employees
}

// insert إضافة موظف جديد إلى قاعدة البيانات
func insert(name string, city string) {
	// إعداد وتنفيذ الاستعلام لإدراج موظف جديد
	stmt, err := db.Prepare("INSERT INTO employee(Name,City) VALUES (?,?)")
	if err != nil {
		log.Fatal(err)
	}

	r, err := stmt.Exec(name, city)
	if err != nil {
		log.Fatal(err)
	}

	// الحصول على عدد الصفوف المتأثرة بالتنفيذ
	affectedRows, err := r.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("تأثير الاستعلام على %d صفوف\n", affectedRows)
	fmt.Printf("The statement affected %d rows\n", affectedRows)

}

// update تحديث بيانات موظف معين
func update(id int, name string, city string) {
	// إعداد وتنفيذ الاستعلام لتحديث بيانات الموظف
	stmt, err := db.Prepare("UPDATE employee SET Name=?, City=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}

	r, err := stmt.Exec(name, city, id)
	if err != nil {
		log.Fatal(err)
	}

	// الحصول على عدد الصفوف المتأثرة بالتنفيذ
	affectedRows, err := r.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("تأثير الاستعلام على %d صفوف\n", affectedRows)

	fmt.Printf("The statement affected %d rows\n", affectedRows)
}

// delete حذف موظف بناء على الرقم المعرف
func delete(id int, name string) {
	// إعداد وتنفيذ الاستعلام لحذف موظف بالرقم المعرف
	// stmt, err := db.Prepare("DELETE FROM employee WHERE id=?")
	stmt, err := db.Prepare("DELETE FROM name=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}

	r, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	// الحصول على عدد الصفوف المتأثرة بالتنفيذ
	affectedRows, err := r.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("تأثير الاستعلام على %d صفوف\n", affectedRows)
	fmt.Printf("The statement affected %d rows\n", affectedRows)
}

func main() {
	// تأسيس الاتصال بقاعدة البيانات
	dbConn()
	// استرجاع وطباعة كل الموظفين
	// fmt.Println(getAllEmployees())
	// يمكنك تفعيل الأسطر التالية لإضافة، تحديث، أو حذف موظف
	// insert("Zizou", "Algeria")
	// update(1, "test", "Egypt")
	// delete(1)
}
