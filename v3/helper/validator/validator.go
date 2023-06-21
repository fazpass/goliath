package validator

import (
	"reflect"

	govalidator "github.com/go-playground/validator/v10"
)

func Validate(s interface{}) map[string]string {
	var (
		structValue         = reflect.ValueOf(s)
		structValueIndirect = reflect.Indirect(structValue)
		structType          = structValueIndirect.Type()
		validator           = govalidator.New()
		err                 = validator.Struct(s)
		errs                = map[string]string{}
	)

	if err == nil {
		return nil
	}

	validatorErrs := err.(govalidator.ValidationErrors)
	for _, e := range validatorErrs {
		var (
			structfField, _ = structType.FieldByName(e.Field())
			field           = structfField.Tag.Get("json")
		)
		if field == "" {
			field = e.Field()
		}
		errs[field] = GetMessage(field, e.ActualTag())
	}

	return errs
}
