package service

import (
	"fmt"
	"net/http"
	"strconv"

	constants "github.com/flabio/safe_constants"
	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_user/core"
	"github.com/safe_msvc_user/insfractruture/entities"
	"github.com/safe_msvc_user/insfractruture/helpers"
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
	page, begin := Pagination(c, int(constants.LIMIT))
	results, countSchools, err := s.UiSchool.GetSchoolFindAll(begin)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS: fiber.StatusBadRequest,
			constants.DATA:   constants.ERROR_QUERY,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		constants.STATUS:      fiber.StatusOK,
		constants.DATA:        results,
		constants.TOTAL_COUNT: countSchools,
		constants.PAGE_COUNT:  page,
		constants.BEGIN:       begin,
	})
}

func (s *SchoolService) GetSchoolFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(constants.ID))
	result, err := s.UiSchool.GetSchoolFindById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(result)
	}
	if result.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			constants.STATUS:  fiber.StatusNotFound,
			constants.MESSAGE: constants.ID_NO_EXIST,
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)
}
func (s *SchoolService) CreateSchool(c *fiber.Ctx) error {
	var schoolCreate entities.School
	schoolDto, msgError := ValidateSchool(0, s, c)
	if msgError != constants.EMPTY {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS:  http.StatusBadRequest,
			constants.MESSAGE: msgError,
			constants.DATA:    msgError,
		})
	}
	// Manejar el archivo subido
	fileHeader, err := c.FormFile(constants.FILE)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS:  http.StatusBadRequest,
			constants.MESSAGE: err.Error(),
			constants.DATA:    "",
		})
	}
	// Guardar el archivo (opcional)
	filePath := fmt.Sprintf(constants.UPLOADS_FILE, fileHeader.Filename)

	err = c.SaveFile(fileHeader, filePath)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS:  http.StatusBadRequest,
			constants.MESSAGE: err.Error(),
		})

	}

	urlFileName, err := helpers.UploadFileToS3(constants.AWS_BUCKET_NAME, fileHeader.Filename)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS:  http.StatusBadRequest,
			constants.MESSAGE: err.Error(),
		})
	}
	deepcopier.Copy(schoolDto).To(&schoolCreate)
	schoolCreate.Url = urlFileName
	data, err := s.UiSchool.CreateSchool(schoolCreate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS:  http.StatusBadRequest,
			constants.MESSAGE: constants.ERROR_UPDATE,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		constants.STATUS:  http.StatusOK,
		constants.MESSAGE: constants.CREATED,
		constants.DATA:    data,
	})
}

func (s *SchoolService) UpdateSchool(c *fiber.Ctx) error {
	var updatedSchool entities.School
	id, _ := strconv.Atoi(c.Params(constants.ID))
	result, err := s.UiSchool.GetSchoolFindById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS: fiber.StatusBadRequest,
			constants.DATA:   constants.ERROR_QUERY,
		})
	}
	if result.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			constants.STATUS:  fiber.StatusNotFound,
			constants.MESSAGE: constants.ID_NO_EXIST,
		})
	}
	// stateExit, _ := statesstruct.MsvcStateFindById(result.StateId, c)
	// if stateExit.Id == 0 {
	// 	return c.Status(http.StatusBadRequest).JSON(fiber.Map{
	// 		constants.STATUS:  http.StatusBadRequest,
	// 		constants.MESSAGE: constants.STATE_NOT_FOUND,
	// 	})
	// }
	schoolDto, msgError := ValidateSchool(uint(id), s, c)
	if msgError != constants.EMPTY {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS:  http.StatusBadRequest,
			constants.MESSAGE: msgError,
		})
	}

	// Manejar el archivo subido
	fileHeader, _ := c.FormFile(constants.FILE)

	if fileHeader != nil {
		// Guardar el archivo (opcional)
		filePath := fmt.Sprintf(constants.UPLOADS_FILE, fileHeader.Filename)

		err = c.SaveFile(fileHeader, filePath)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				constants.STATUS:  http.StatusBadRequest,
				constants.MESSAGE: err.Error(),
			})

		}
		urlFileName, err := helpers.UploadFileToS3(constants.AWS_BUCKET_NAME, fileHeader.Filename)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				constants.STATUS:  http.StatusBadRequest,
				constants.MESSAGE: err.Error(),
			})
		}
		if result.Url != "" {
			// Borrar el archivo anterior de S3
			helpers.RemoveFileToS3(constants.AWS_BUCKET_NAME, result.Url)

		}
		updatedSchool.Url = urlFileName
	}

	deepcopier.Copy(schoolDto).To(&updatedSchool)

	user, err := s.UiSchool.UpdateSchool(uint(id), updatedSchool)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS:  http.StatusBadRequest,
			constants.MESSAGE: constants.ERROR_UPDATE,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		constants.STATUS:  http.StatusOK,
		constants.MESSAGE: constants.UPDATED,
		constants.DATA:    user,
	})
}
func (s *SchoolService) DeleteSchool(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(constants.ID))
	schoolFindById, err := s.UiSchool.GetSchoolFindById(uint(id))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS: fiber.StatusBadRequest,
			constants.DATA:   constants.ERROR_QUERY,
		})
	}
	if schoolFindById.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			constants.STATUS:  fiber.StatusNotFound,
			constants.MESSAGE: constants.ID_NO_EXIST,
		})
	}
	helpers.RemoveFileToS3(constants.AWS_BUCKET_NAME, schoolFindById.Url)
	result, err := s.UiSchool.DeleteSchool(uint(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			constants.STATUS:  http.StatusBadRequest,
			constants.MESSAGE: constants.ERROR_DELETE,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		constants.STATUS:  http.StatusOK,
		constants.MESSAGE: constants.REMOVED,
		constants.DATA:    result,
	})
}

func Pagination(c *fiber.Ctx, limit int) (int, int) {
	pageParam := c.Query(constants.PAGE)

	if pageParam == "" {
		return 1, 0
	}
	page, _ := strconv.Atoi(c.Query(constants.PAGE))
	if page < 1 {
		return 1, 0
	}

	begin := (limit * page) - limit
	return page, begin
}
