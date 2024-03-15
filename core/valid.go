package core

import (
	// "database/sql"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Vilad(mystruct interface{}) {
	validate := validator.New()

	err := validate.Struct(mystruct)
	validationErrors := err.(validator.ValidationErrors)
	if validationErrors != nil {
		fmt.Println(validationErrors, "::::::::::::::::::::", err)
		return
	}
}

//////////////////////////////////////////custom validation/////////////////////////////////////////////////////
//////////////////////////////////////////custom validation/////////////////////////////////////////////////////
//////////////////////////////////////////custom validation/////////////////////////////////////////////////////
// package main

// import (
// 	"fmt"

// 	"github.com/go-playground/validator/v10"
// )

// // MyStruct ..
// type MyStruct struct {
// 	String string `validate:"is-awesome"`
// }

// // use a single instance of Validate, it caches struct info
// var validate *validator.Validate

// func main() {

// 	validate = validator.New()
// 	validate.RegisterValidation("is-awesome", ValidateMyVal)

// 	s := MyStruct{String: "awesome"}

// 	err := validate.Struct(s)
// 	if err != nil {
// 		fmt.Printf("Err(s):\n%+v\n", err)
// 	}

// 	s.String = "not awesome"
// 	err = validate.Struct(s)
// 	if err != nil {
// 		fmt.Printf("Err(s):\n%+v\n", err)
// 	}
// }

// // ValidateMyVal implements validator.Func
// func ValidateMyVal(fl validator.FieldLevel) bool {
// 	return fl.Field().String() == "awesome"
// }
//////////////////////////////////////////custom validation/////////////////////////////////////////////////////
//////////////////////////////////////////custom validation/////////////////////////////////////////////////////
//////////////////////////////////////////custom validation/////////////////////////////////////////////////////
// package main

// import (
// 	"database/sql"
// 	"database/sql/driver"
// 	"fmt"
// 	"reflect"

// 	"github.com/go-playground/validator/v10"
// )

// // DbBackedUser User struct
// type DbBackedUser struct {
// 	Name sql.NullString `validate:"required"`
// 	Age  sql.NullInt64  `validate:"required"`
// }

// // use a single instance of Validate, it caches struct info
// var validate *validator.Validate

// func main() {

// 	validate = validator.New()

// 	// register all sql.Null* types to use the ValidateValuer CustomTypeFunc
// 	validate.RegisterCustomTypeFunc(ValidateValuer, sql.NullString{}, sql.NullInt64{}, sql.NullBool{}, sql.NullFloat64{})

// 	// build object for validation
// 	x := DbBackedUser{Name: sql.NullString{String: "", Valid: true}, Age: sql.NullInt64{Int64: 0, Valid: false}}

// 	err := validate.Struct(x)

// 	if err != nil {
// 		fmt.Printf("Err(s):\n%+v\n", err)
// 	}
// }

// // ValidateValuer implements validator.CustomTypeFunc
// func ValidateValuer(field reflect.Value) interface{} {

// 	if valuer, ok := field.Interface().(driver.Valuer); ok {

// 		val, err := valuer.Value()
// 		if err == nil {
// 			return val
// 		}
// 		// handle the error how you want
// 	}

// 	return nil
// }
//////////////////////////////////////////custom validation/////////////////////////////////////////////////////
//////////////////////////////////////////custom validation/////////////////////////////////////////////////////
//////////////////////////////////////////custom validation/////////////////////////////////////////////////////
