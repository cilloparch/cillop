package validation

import "github.com/go-playground/validator/v10"

func (v *Validator) translateErrorMessage(err validator.FieldError, languages ...string) string {
	return v.i18n.TranslateWithParams("validation_"+err.Tag(), map[string]interface{}{
		"Value": err.Value(),
		"Field": err.Field(),
		"Param": err.Param(),
	}, languages...)
}
