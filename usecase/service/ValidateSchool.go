package service

import (
	"log"

	utils "github.com/flabio/safe_constants"
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_user/insfractruture/helpers"
	"github.com/safe_msvc_user/usecase/dto"
)

func ValidateSchool(id uint, s *SchoolService, c *fiber.Ctx) (dto.SchoolDTO, string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("controlando el panic", r)
		}
	}()
	var schoolDto dto.SchoolDTO
	var msg string = utils.EMPTY

	dataMap := make(map[string]string)
	fields := []string{
		"name", "email", "address", "phone", "zip_code", "provider_number", "state_id",
	}

	// Iterar sobre las claves y obtener el valor para cada una
	for _, field := range fields {
		value := c.FormValue(field)

		if value != "" {

			dataMap[field] = value
		} else {
			dataMap[field] = ""
		}
	}
	for field, value := range dataMap {
		if value == "" || len(value) == 0 {
			msg = field + " is required"
			return dto.SchoolDTO{}, msg
		}
	}
	helpers.MapToStructSchool(&schoolDto, dataMap)

	/*msgValid := helpers.ValidateField(dataMap)
	if msgValid != utils.EMPTY {
		return dto.SchoolDTO{}, msgValid
	}
	*/
	existEmail, _ := s.UiSchool.GetSchoolFindByEmail(id, schoolDto.Email)
	if existEmail.Email != utils.EMPTY {
		msg = utils.EMAIL_ALREADY_EXIST
	}
	existProviderNumber, _ := s.UiSchool.GetSchoolFindByProviderNumber(id, schoolDto.ProviderNumber)
	if existProviderNumber.ProviderNumber != utils.EMPTY {
		msg = "The Provider number ready exists "
	}
	return schoolDto, msg
}
