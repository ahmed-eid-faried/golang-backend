package main

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func filterRequest(c *gin.Context, requestname string) string {
	return c.PostForm(requestname)
}

func sha1Request(c *gin.Context, requestname string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(c.PostForm(requestname))))
}

func GetAllData(c *gin.Context, table string, where string, values ...interface{}) int {
	data := make(map[string]interface{})
	var stmt *sql.Rows
	var err error
	if where == "" {
		stmt, err = db.Query("SELECT * FROM " + table)
	} else {
		stmt, err = db.Query("SELECT * FROM "+table+" WHERE "+where, values...)
	}
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer stmt.Close()
	count := 0
	for stmt.Next() {
		count++
		err = stmt.Scan(&data)
		if err != nil {
			fmt.Println(err)
			return 0
		}
	}
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": data})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "failure"})
	}
	return count
}

func getData(c *gin.Context, table string, where string, values ...interface{}) int {
	data := make(map[string]interface{})
	// var stmt *sql.Row
	err := db.QueryRow("SELECT * FROM "+table+" WHERE "+where, values...).Scan(&data)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": data})
	return 1
}

func InsertData(c *gin.Context, table string, data map[string]interface{}) int {
	fields := make([]string, len(data))
	placeholders := make([]string, len(data))
	values := make([]interface{}, len(data))
	i := 0
	for field, value := range data {
		fields[i] = field
		placeholders[i] = "?"
		values[i] = value
		i++
	}
	stmt, err := db.Prepare("INSERT INTO " + table + " (" + strings.Join(fields, ", ") + ") VALUES (" + strings.Join(placeholders, ", ") + ")")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	_, err = stmt.Exec(values...)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
	return 1
}

func updateData(c *gin.Context, table string, data map[string]interface{}, where string, values ...interface{}) int {
	set := make([]string, len(data))
	i := 0
	for field, value := range data {
		set[i] = field + " = ?"
		values = append(values, value)
		i++
	}
	stmt, err := db.Prepare("UPDATE " + table + " SET " + strings.Join(set, ", ") + " WHERE " + where)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	_, err = stmt.Exec(values...)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
	return 1
}

func DeleteData(c *gin.Context, table string, where string, values ...interface{}) int {
	stmt, err := db.Prepare("DELETE FROM " + table + " WHERE " + where)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	_, err = stmt.Exec(values...)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
	return 1
}

// func imageUpload(dir string, imageRequest string) string {
//     file, err := c.FormFile(imageRequest)
//     if err != nil {
//         fmt.Println(err)
//         return "empty"
//     }
//     if file.Size > 2*MB {
//         fmt.Println("File size too large")
//         return "fail"
//     }
//     ext := filepath.Ext(file.Filename)
//     if ext != ".jpg" && ext != ".png" && ext != ".gif" && ext != ".mp3" && ext != ".pdf" && ext != ".svg" {
//         fmt.Println("Invalid file extension")
//         return "fail"
//     }
//     imagename := strconv.Itoa(rand.Intn(99999-10000)+10000) + file.Filename
//     if err := c.SaveUploadedFile(file, filepath.Join(dir, imagename)); err != nil {
//         fmt.Println(err)
//         return "fail"
//     }
//     return imagename
// }

var MB int64 =  1048576

func imageUpload(dir string, imageRequest *http.Request) string {
	file, fileHeader, err := imageRequest.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return "empty"
	}
	defer file.Close()

	if fileHeader.Size > (2 * MB) {
		fmt.Println("File size too large")
		return "fail"
	}

	ext := filepath.Ext(fileHeader.Filename)
	allowedExts := map[string]bool{
		".jpg": true,
		".png": true,
		".gif": true,
		".mp3": true,
		".pdf": true,
		".svg": true,
	}

	if !allowedExts[ext] {
		fmt.Println("Invalid file extension")
		return "fail"
	}

	imageName := strconv.Itoa(rand.Intn(99999-10000)+10000) + fileHeader.Filename
	err = saveFile(dir, imageName, file)
	if err != nil {
		fmt.Println(err)
		return "fail"
	}

	return imageName
}

func saveFile(dir, fileName string, file multipart.File) error {
	dst, err := os.Create(filepath.Join(dir, fileName))
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}

	return nil
}
func deleteFile(dir string, imagename string) {
	err := os.Remove(filepath.Join(dir, imagename))
	if err != nil {
		fmt.Println(err)
	}
}

func checkAuthenticate(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()
	if !ok || username != "wael" || password != "wael12345" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func printFailure(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{"status": "failure", "message": message})
}

func printSuccess(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
}

func result(count int, c *gin.Context) {
	if count > 0 {
		printSuccess(c, "")
	} else {
		printFailure(c, "")
	}
}

func sendEmail(to string, title string, body string) {
	from := "ahmedmady@amadytech.com"
	// header := "From: " + from
	err := smtp.SendMail("smtp.example.com:587", smtp.PlainAuth(title, from, "password", "smtp.example.com"), from, []string{to}, []byte(body))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func sendGCM(title string, message string, topic string, pageid string, pagename string) {
	// Implement your GCM sending logic here
}

func insertNotify(userid string, title string, body string, topic string, pageid string, pagename string) int {
	stmt, err := db.Prepare("INSERT INTO notifications (notifications_userid, notifications_title, notifications_body, notifications_topic) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return 0
	}
	_, err = stmt.Exec(userid, title, body, topic)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	go sendGCM(title, body, topic, pageid, pagename)
	return 1
}

