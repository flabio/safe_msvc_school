package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_user/clients/statesstruct"
	"github.com/safe_msvc_user/core"
	"github.com/safe_msvc_user/insfractruture/entities"
	"github.com/safe_msvc_user/insfractruture/helpers"
	"github.com/safe_msvc_user/insfractruture/utils"

	"github.com/safe_msvc_user/insfractruture/ui/global"
	"github.com/safe_msvc_user/insfractruture/ui/uicore"
	"github.com/ulule/deepcopier"
)

type SchoolService struct {
	UiSchool uicore.UISchoolCore
}

func NewSchoolService() global.UISchoolGlobal {
	return &SchoolService{UiSchool: core.NewSchoolRepository()}
}

func (s *SchoolService) GetSchoolFindAll(c *fiber.Ctx) error {
	page, begin := Pagination(c, int(utils.LIMIT))
	results, countSchools, err := s.UiSchool.GetSchoolFindAll(begin)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		utils.STATUS:      fiber.StatusOK,
		utils.DATA:        results,
		utils.TOTAL_COUNT: countSchools,
		utils.PAGE_COUNT:  page,
		utils.BEGIN:       begin,
	})
}

func (s *SchoolService) GetSchoolFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.UiSchool.GetSchoolFindById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(result)
	}
	if result.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  fiber.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
func (s *SchoolService) CreateSchool(c *fiber.Ctx) error {
	var schoolCreate entities.School
	schoolDto, msgError := ValidateSchool(0, s, c)
	if msgError != utils.EMPTY {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
			utils.DATA:    msgError,
		})
	}
	// Manejar el archivo subido
	fileHeader, err := c.FormFile(utils.FILE)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
			utils.DATA:    "",
		})
	}
	// Guardar el archivo (opcional)
	filePath := fmt.Sprintf(utils.UPLOADS_FILE, fileHeader.Filename)

	err = c.SaveFile(fileHeader, filePath)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})

	}

	urlFileName, err := helpers.UploadFileToS3(utils.AWS_BUCKET_NAME, fileHeader.Filename)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: err.Error(),
		})
	}
	deepcopier.Copy(schoolDto).To(&schoolCreate)
	schoolCreate.Url = urlFileName
	data, err := s.UiSchool.CreateSchool(schoolCreate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ERROR_UPDATE,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.CREATED,
		utils.DATA:    data,
	})
}

func (s *SchoolService) UpdateSchool(c *fiber.Ctx) error {
	var updatedSchool entities.School
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.UiSchool.GetSchoolFindById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	if result.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  fiber.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	deepcopier.Copy(result).To(&updatedSchool)
	stateExit, _ := statesstruct.MsvcStateFindById(result.StateId, c)
	if stateExit.Id == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.STATE_NOT_FOUND,
		})
	}
	schoolDto, msgError := ValidateSchool(uint(id), s, c)
	if msgError != utils.EMPTY {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(schoolDto).To(&updatedSchool)
	user, err := s.UiSchool.UpdateSchool(uint(id), updatedSchool)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ERROR_UPDATE,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.UPDATED,
		utils.DATA:    user,
	})
}
func (s *SchoolService) DeleteSchool(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	schoolFindById, err := s.UiSchool.GetSchoolFindById(uint(id))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	if schoolFindById.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  fiber.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	helpers.RemoveFileToS3(utils.AWS_BUCKET_NAME, schoolFindById.Url)
	result, err := s.UiSchool.DeleteSchool(uint(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ERROR_DELETE,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.REMOVED,
		utils.DATA:    result,
	})
}

func Pagination(c *fiber.Ctx, limit int) (int, int) {
	pageParam := c.Query(utils.PAGE)

	if pageParam == "" {
		return 1, 0
	}
	page, _ := strconv.Atoi(c.Query(utils.PAGE))
	if page < 1 {
		return 1, 0
	}

	begin := (limit * page) - limit
	return page, begin
}
