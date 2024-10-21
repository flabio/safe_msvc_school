package helpers

import (
	utils "github.com/flabio/safe_constants"
	"github.com/safe_msvc_user/usecase/dto"
)

func ValidateRequired(school dto.SchoolDTO) string {
	var msg string = utils.EMPTY
	if school.Name == utils.EMPTY {
		msg = utils.NAME_IS_REQUIRED
	}
	if school.Address == utils.EMPTY {
		msg = utils.ADDRESS_IS_REQUIRED
	}
	if school.Phone == utils.EMPTY {
		msg = utils.PHONE_IS_REQUIRED
	}
	if school.Email == utils.EMPTY {
		msg = utils.EMAIL_IS_REQUIRED
	}
	if school.ZipCode == utils.EMPTY {
		msg = utils.ZIP_CODE_IS_REQUIRED
	}
	if school.ProviderNumber == utils.EMPTY {
		msg = utils.PROVIDER_NUMBER_IS_REQUIRED
	}
	if school.StateId == 0 {
		msg = utils.STATE_ID_IS_REQUIRED
	}
	return msg
}

func ValidateField(value map[string]interface{}) string {
	var msg string = utils.EMPTY
	if value[utils.NAME] == nil {
		msg = utils.NAME_FIELD_IS_REQUIRED
	}
	if value[utils.ADDRESS] == nil {
		msg = utils.ADDRESS_IS_FIELD_REQUIRED
	}
	if value[utils.PHONE] == nil {
		msg = utils.PHONE_IS_FIELD_REQUIRED
	}
	if value[utils.EMAIL] == nil {
		msg = utils.EMAIL_IS_FIELD_REQUIRED
	}
	if value[utils.ZIP_CODE] == nil {
		msg = utils.ZIP_CODE_IS_FIELD_REQUIRED
	}
	if value[utils.PROVIDER_NUMBER] == nil {
		msg = utils.PROVIDER_NUMBER_IS_FIELD_REQUIRED
	}
	/*if value[utils.STATE_ID] == nil {
		msg = utils.STATE_ID_IS_FIELD_REQUIRED
	}*/
	return msg
}
