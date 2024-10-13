package global

import "github.com/gofiber/fiber/v2"

type UISchoolGlobal interface {
	GetSchoolFindAll(c *fiber.Ctx) error
	GetSchoolFindById(c *fiber.Ctx) error
	CreateSchool(c *fiber.Ctx) error
	UpdateSchool(c *fiber.Ctx) error
	DeleteSchool(c *fiber.Ctx) error
}
