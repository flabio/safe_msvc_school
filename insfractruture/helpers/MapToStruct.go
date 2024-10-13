package helpers

import (
	"strconv"

	utils "github.com/flabio/safe_constants"
	"github.com/safe_msvc_user/usecase/dto"
)

func MapToStructSchool(schoolDto *dto.SchoolDTO, dataMap map[string]string) {
	stateID, _ := strconv.Atoi(dataMap[utils.STATE_ID])
	isActive, _ := strconv.ParseBool(dataMap[utils.ACTIVE])

	school := dto.SchoolDTO{
		Name:           dataMap[utils.NAME],
		Address:        dataMap[utils.ADDRESS],
		Phone:          dataMap[utils.PHONE],
		Email:          dataMap[utils.EMAIL],
		ZipCode:        dataMap[utils.ZIP_CODE],
		StateId:        uint(stateID),
		ProviderNumber: dataMap[utils.PROVIDER_NUMBER],
		Active:         isActive,
	}

	*schoolDto = school
}
