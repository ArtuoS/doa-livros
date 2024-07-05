package controller

import (
	"strconv"

	"github.com/ArtuoS/doa-livros/internal/entity"
	"github.com/ArtuoS/doa-livros/internal/repository"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserRepo *repository.UserRepository
}

func NewUserController(userRepo *repository.UserRepository) *UserController {
	return &UserController{UserRepo: userRepo}
}

func (u *UserController) GetAuthentication(c *fiber.Ctx) error {
	return c.Render("auth", nil)
}

func (u *UserController) Authenticate(c *fiber.Ctx) error {
	var auth entity.Auth
	if err := c.BodyParser(&auth); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	user, err := u.UserRepo.GetUserByAuth(auth)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid email or password",
		})
	}

	token, err := entity.GenerateJWT(user.Id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate token",
		})
	}

	return c.JSON(fiber.Map{
		"token":        token,
		"loggedUserId": user.Id,
	})
}

func (u *UserController) GetUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	user, err := u.UserRepo.GetUser(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	data := fiber.Map{
		"user": user,
	}

	return c.Render("profile", data)
}
