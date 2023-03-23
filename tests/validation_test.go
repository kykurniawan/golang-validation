package tests

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	validate := validator.New()

	if validate == nil {
		t.Error("validate is nil")
	}
}

func TestValidationVariable(t *testing.T) {
	validate := validator.New()

	user := ""

	err := validate.Var(user, "required")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidateTwoVariables(t *testing.T) {
	validate := validator.New()

	password := "rahasia"

	confirm := "salah"

	err := validate.VarWithValue(password, confirm, "eqfield")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleTag(t *testing.T) {
	validate := validator.New()

	user := "12345"

	err := validate.Var(user, "required,numeric")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestTagParameter(t *testing.T) {
	validate := validator.New()

	user := "222222222222222"

	err := validate.Var(user, "required,numeric,min=5,max=10")

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestStruct(t *testing.T) {
	type LoginRequest struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()

	loginRequest := LoginRequest{
		Email:    "rizky@example.com",
		Password: "password",
	}

	err := validate.Struct(loginRequest)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationErrors(t *testing.T) {
	type LoginRequest struct {
		Email    string `validate:"required,email"`
		Password string `validate:"required,min=5"`
	}

	validate := validator.New()

	loginRequest := LoginRequest{
		Email:    "rizk",
		Password: "pass",
	}

	err := validate.Struct(loginRequest)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			fmt.Println("error", fieldError.Field(), "on tag", fieldError.Tag(), "with error", fieldError.Error())
		}
	}
}

func TestStructCrossField(t *testing.T) {
	type RegisterRequest struct {
		Email           string `validate:"required,email"`
		Password        string `validate:"required,min=5"`
		ConfirmPassword string `validate:"required,eqfield=Password"`
	}

	validate := validator.New()

	registerRequest := RegisterRequest{
		Email:           "rizky@example.com",
		Password:        "password",
		ConfirmPassword: "password",
	}

	err := validate.Struct(registerRequest)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestNestedStruct(t *testing.T) {
	validate := validator.New()

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id      string  `validate:"required"`
		Name    string  `validate:"required"`
		Address Address `validate:"required"`
	}

	user := User{
		Id:   "",
		Name: "",
		Address: Address{
			City:    "",
			Country: "",
		},
	}

	err := validate.Struct(user)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestCollection(t *testing.T) {
	validate := validator.New()

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
	}

	user := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
	}

	err := validate.Struct(user)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestBasicCollection(t *testing.T) {
	validate := validator.New()

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type User struct {
		Id        string    `validate:"required"`
		Name      string    `validate:"required"`
		Addresses []Address `validate:"required,dive"`
		Hobbies   []string  `validate:"required,dive,required,min=3"`
	}

	user := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"",
			"X",
		},
	}

	err := validate.Struct(user)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMap(t *testing.T) {
	validate := validator.New()

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"required,dive,required,min=3"`
		Schools   map[string]School `validate:"required,dive,keys,required,min=2,endkeys,dive"`
	}

	user := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"",
			"X",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SDN 2 Bumi Indah",
			},
			"SMP": {
				Name: "SMPN 3 Katingan Kuala",
			},
			"SMA": {
				Name: "SMAN 1 Katingan Kuala",
			},
		},
	}

	err := validate.Struct(user)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestBasicMap(t *testing.T) {
	validate := validator.New()

	type Address struct {
		City    string `validate:"required"`
		Country string `validate:"required"`
	}

	type School struct {
		Name string `validate:"required"`
	}

	type User struct {
		Id        string            `validate:"required"`
		Name      string            `validate:"required"`
		Addresses []Address         `validate:"required,dive"`
		Hobbies   []string          `validate:"required,dive,required,min=3"`
		Schools   map[string]School `validate:"required,dive,keys,required,min=2,endkeys,dive"`
		Wallet    map[string]int    `validate:"dive,keys,required,endkeys,required,gt=0"`
	}

	user := User{
		Id:   "",
		Name: "",
		Addresses: []Address{
			{
				City:    "",
				Country: "",
			},
			{
				City:    "",
				Country: "",
			},
		},
		Hobbies: []string{
			"Gaming",
			"Coding",
			"",
			"X",
		},
		Schools: map[string]School{
			"SD": {
				Name: "SDN 2 Bumi Indah",
			},
			"SMP": {
				Name: "SMPN 3 Katingan Kuala",
			},
			"SMA": {
				Name: "SMAN 1 Katingan Kuala",
			},
		},
		Wallet: map[string]int{
			"BCA":     1000000,
			"MANDIRI": 0,
		},
	}

	err := validate.Struct(user)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestAlias(t *testing.T) {
	validate := validator.New()

	validate.RegisterAlias("varchar", "required,max=255")

	type Seller struct {
		Id     string `validate:"varchar,min=5"`
		Name   string `validate:"varchar"`
		Slogan string `validate:"varchar"`
	}

	seller := Seller{
		Id:     "mawar-melati",
		Name:   "Mawar Melati",
		Slogan: "Mawar melati semuanya indah",
	}

	err := validate.Struct(seller)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)

	if ok {
		if value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}

	return true
}

func TestCustomValidation(t *testing.T) {
	validate := validator.New()

	validate.RegisterValidation("username", MustValidUsername)

	type RegisterRequest struct {
		Username string `validate:"required,username"`
		Password string `validate:"required,min=5"`
	}

	request := RegisterRequest{
		Username: "RIZKY",
		Password: "passwordbanget",
	}

	err := validate.Struct(request)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())

	if err != nil {
		panic(err)
	}

	var regexNumber = regexp.MustCompile("^[0-9]+$")

	value := field.Field().String()

	if !regexNumber.MatchString(value) {
		return false
	}

	return len(value) == length
}

func TestCustomValidationParameter(t *testing.T) {
	validate := validator.New()

	validate.RegisterValidation("pin", MustValidPin)

	type Login struct {
		Phone string `validate:"required,number"`
		Pin   string `validate:"required,pin=6"`
	}

	request := Login{
		Phone: "082222222222",
		Pin:   "123123",
	}

	err := validate.Struct(request)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestOrRule(t *testing.T) {
	type Login struct {
		Username string `validate:"required,email|numeric"`
		Password string `validate:"required"`
	}

	validate := validator.New()

	request := Login{
		Username: "123456543",
		Password: "rizky",
	}

	err := validate.Struct(request)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func MustEqualsIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()

	if !ok {
		panic("field not ok")
	}

	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestCrossFieldValidation(t *testing.T) {
	validate := validator.New()

	validate.RegisterValidation("field_equals_ignore_case", MustEqualsIgnoreCase)

	type User struct {
		Username string `validate:"required,field_equals_ignore_case=Email|field_equals_ignore_case=Phone"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,numeric"`
		Name     string `validate:"required"`
	}

	user := User{
		Username: "rizky@example.com",
		Email:    "rizky@example.com",
		Phone:    "08888888",
		Name:     "Rizky",
	}

	err := validate.Struct(user)

	if err != nil {
		fmt.Println(err.Error())
	}
}

type RegisterRequest struct {
	Username string `validate:"required"`
	Phone    string `validate:"required,numeric"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func MustValidRegisterSuccess(field validator.StructLevel) {
	registerRequest := field.Current().Interface().(RegisterRequest)

	if registerRequest.Username == registerRequest.Email || registerRequest.Username == registerRequest.Phone {
		// success
	} else {
		field.ReportError(registerRequest.Username, "Username", "Username", "username", "")
	}
}

func TestStructLevelValidation(t *testing.T) {
	validate := validator.New()

	validate.RegisterStructValidation(MustValidRegisterSuccess, RegisterRequest{})

	request := RegisterRequest{
		Username: "rizky",
		Email:    "rizky@example.com",
		Phone:    "082222222222",
		Password: "password",
	}

	err := validate.Struct(request)

	if err != nil {
		fmt.Println(err.Error())
	}
}
