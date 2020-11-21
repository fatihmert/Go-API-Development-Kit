package controllers

import (
	"apidev.fatihmert.dev/models"
	"apidev.fatihmert.dev/states"
	"github.com/gofiber/fiber/v2"
)

type LoginUser struct {
	Email    string `json:"mail" xml:"mail" form:"mail"`
	Password string `json:"password" xml:"password" form:"password"`
}

type RegisterUser struct {
	Email    string `json:"mail" xml:"mail" form:"mail"`
	Password string `json:"password" xml:"password" form:"password"`
}

func Login(c *fiber.Ctx) error {
	vars := new(LoginUser)

	if err := c.BodyParser(vars); err != nil {
		return err
	}

	var user models.User
	userIsExist, _ := states.DB.Query("SELECT * FROM users WHERE email=? AND password=?", vars.Email, vars.Password)

	var usersCount int = 0
	for userIsExist.Next() {
		var id uint
		var mail, password, token string

		scanErr := userIsExist.Scan(&id, &mail, &password, &token)
		if scanErr != nil {
			panic(scanErr.Error())
		}

		usersCount++

		user.ID = id
		user.Mail = mail
		user.Password = password
		user.Token = token
	}

	if usersCount > 0 {
		return c.JSON(fiber.Map{
			"message": "Success",
		})
	}

	return c.JSON(fiber.Map{
		"redirect": 0,
		"message":  "Error",
	})
}

func Register(c *fiber.Ctx) error {
	vars := new(RegisterUser)

	if err := c.BodyParser(vars); err != nil {
		return err
	}

	var user models.User
	userIsExist, _ := states.DB.Query("SELECT * FROM users WHERE email=?", vars.Email)

	var usersCount int = 0

	for userIsExist.Next() {
		var id uint
		var mail, password, token string

		scanErr := userIsExist.Scan(&id, &mail, &password, &token)
		if scanErr != nil {
			panic(scanErr.Error())
		}

		usersCount++

		user.ID = id
		user.Mail = mail
		user.Password = password
		user.Token = token
	}

	if usersCount == 0 {
		insForm, err := states.DB.Prepare("INSERT INTO users(email, password, token) VALUES(?,?,?)")

		if err != nil {
			panic(err.Error())
		}

		insForm.Exec(vars.Email, vars.Password, "TokenBasic")

		return c.JSON(fiber.Map{
			"message": "Success",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Error",
	})
}
