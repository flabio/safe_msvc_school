package uicore

import (
	"github.com/safe_msvc_user/insfractruture/entities"
	"github.com/safe_msvc_user/usecase/dto"
)

type UISchoolCore interface {
	GetSchoolFindAll(begin int) ([]dto.SchoolResponseDTO, int64, error)
	GetSchoolFindById(id uint) (entities.School, error)
	GetSchoolFindByEmail(id uint, email string) (entities.School, error)
	GetSchoolFindByProviderNumber(id uint, providerNumber string) (entities.School, error)
	CreateSchool(school entities.School) (entities.School, error)
	UpdateSchool(id uint, school entities.School) (entities.School, error)
	DeleteSchool(id uint) (bool, error)
}
