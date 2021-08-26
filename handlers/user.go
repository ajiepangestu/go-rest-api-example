package handlers

import (
	"github.com/ajiepangestu/go-rest-api-example/database"
	"github.com/ajiepangestu/go-rest-api-example/models"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	return nil
}

func GetUserList(c *fiber.Ctx) error {
	rows, err := database.DB.Query("SELECT first_name, last_name, email FROM users")
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	defer rows.Close()
	result := models.Users{}
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.FirstName, &user.LastName, &user.Email); err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}
		result.Users = append(result.Users, user)
	}
	if err := c.JSON(&fiber.Map{
		"success": true,
		"data":    result,
		"message": "All user data returned successfully",
	}); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"message": err,
		})
	}
	return err
}

func GetUser(c *fiber.Ctx) error {
	row, err := database.DB.Query("SELECT first_name, last_name, email FROM users WHERE id = ?", c.Params("id"))
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	defer row.Close()
	user := models.User{}
	for row.Next() {
		if err := row.Scan(&user.FirstName, &user.LastName, &user.Email); err != nil {
			return c.Status(500).JSON(&fiber.Map{
				"success": false,
				"error":   err,
			})
		}
	}
	if c.JSON(&fiber.Map{
		"success": true,
		"data":    user,
		"message": "User data returned successfully",
	}); err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	return err
}

func UpdateUser(c *fiber.Ctx) error {
	return nil
}

func DeleteUser(c *fiber.Ctx) error {
	return nil
}
