package helpers

import (
	"strconv"

	constants "github.com/flabio/safe_constants"
	"github.com/safe_msvc_user/usecase/dto"
)

func MapToStructSchool(schoolDto *dto.SchoolDTO, dataMap map[string]string) {
	stateID, _ := strconv.Atoi(dataMap[constants.STATE_ID])
	isActive, _ := strconv.ParseBool(dataMap[constants.ACTIVE])

	school := dto.SchoolDTO{
		Name:           dataMap[constants.NAME],
		Address:        dataMap[constants.ADDRESS],
		Phone:          dataMap[constants.PHONE],
		Email:          dataMap[constants.EMAIL],
		ZipCode:        dataMap[constants.ZIP_CODE],
		StateId:        uint(stateID),
		ProviderNumber: dataMap[constants.PROVIDER_NUMBER],
		Active:         isActive,
	}

	*schoolDto = school
}
