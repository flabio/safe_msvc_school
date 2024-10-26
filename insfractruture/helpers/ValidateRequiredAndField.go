package helpers

import (
	constants "github.com/flabio/safe_constants"
	"github.com/safe_msvc_user/usecase/dto"
)

func ValidateRequired(school dto.SchoolDTO) string {
	var msg string = constants.EMPTY
	if school.Name == constants.EMPTY {
		msg = constants.NAME_IS_REQUIRED
	}
	if school.Address == constants.EMPTY {
		msg = constants.ADDRESS_IS_REQUIRED
	}
	if school.Phone == constants.EMPTY {
		msg = constants.PHONE_IS_REQUIRED
	}
	if school.Email == constants.EMPTY {
		msg = constants.EMAIL_IS_REQUIRED
	}
	if school.ZipCode == constants.EMPTY {
		msg = constants.ZIP_CODE_IS_REQUIRED
	}
	if school.ProviderNumber == constants.EMPTY {
		msg = constants.PROVIDER_NUMBER_IS_REQUIRED
	}
	if school.StateId == 0 {
		msg = constants.STATE_ID_IS_REQUIRED
	}
	return msg
}

func ValidateField(value map[string]interface{}) string {
	var msg string = constants.EMPTY
	if value[constants.NAME] == nil {
		msg = constants.NAME_FIELD_IS_REQUIRED
	}
	if value[constants.ADDRESS] == nil {
		msg = constants.ADDRESS_IS_FIELD_REQUIRED
	}
	if value[constants.PHONE] == nil {
		msg = constants.PHONE_IS_FIELD_REQUIRED
	}
	if value[constants.EMAIL] == nil {
		msg = constants.EMAIL_IS_FIELD_REQUIRED
	}
	if value[constants.ZIP_CODE] == nil {
		msg = constants.ZIP_CODE_IS_FIELD_REQUIRED
	}
	if value[constants.PROVIDER_NUMBER] == nil {
		msg = constants.PROVIDER_NUMBER_IS_FIELD_REQUIRED
	}
	/*if value[constants.STATE_ID] == nil {
		msg = constants.STATE_ID_IS_FIELD_REQUIRED
	}*/
	return msg
}
