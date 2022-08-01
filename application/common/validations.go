package common

import "github.com/asaskevich/govalidator"

func SetGoValidator() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func Validation(strct interface{}) (bool, error) {
	isValid, err := govalidator.ValidateStruct(strct)

	return isValid, err
}