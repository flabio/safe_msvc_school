package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_user/insfractruture/ui/global"
	"github.com/safe_msvc_user/usecase/service"
)

type schoolHandler struct {
	school global.UISchoolGlobal
}

func NewSchoolHandler() global.UISchoolGlobal {
	return &schoolHandler{school: service.NewSchoolService()}
}

func (h *schoolHandler) GetSchoolFindAll(c *fiber.Ctx) error {
	return h.school.GetSchoolFindAll(c)
}

func (h *schoolHandler) GetSchoolFindById(c *fiber.Ctx) error {
	return h.school.GetSchoolFindById(c)
}

func (h *schoolHandler) CreateSchool(c *fiber.Ctx) error {
	return h.school.CreateSchool(c)
}

func (h *schoolHandler) UpdateSchool(c *fiber.Ctx) error {
	return h.school.UpdateSchool(c)
}

func (h *schoolHandler) DeleteSchool(c *fiber.Ctx) error {
	return h.school.DeleteSchool(c)
}
