package validation

import (
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func validateUserName(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(userNameRegexp, fl.Field().String())
	return matched
}

func validatePassword(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(passwordRegexp, fl.Field().String())
	return matched
}

func validateSlug(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(slugRegexp, fl.Field().String())
	return matched
}

func validateLocale(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(localeRegexp, fl.Field().String())
	return matched
}

func validateObjectId(fl validator.FieldLevel) bool {
	return primitive.IsValidObjectID(fl.Field().String())
}

func validateGender(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(genderRegexp, fl.Field().String())
	return matched
}

func validatePhone(f1 validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(PhoneWithCountryCodeRegexp, f1.Field().String())
	return matched
}

func validateUUID(f1 validator.FieldLevel) bool {
	switch f1.Field().Kind() {
	case reflect.String:
		_, err := uuid.Parse(f1.Field().String())
		return err == nil
	case reflect.Ptr:
		if f1.Field().Type().String() == "*uuid.UUID" {
			return f1.Field().Interface() != nil && f1.Field().Elem().Interface() != uuid.Nil
		}
		return false
	case reflect.Struct:
		if f1.Field().Type().String() == "uuid.UUID" {
			return f1.Field().Interface() != uuid.Nil
		}
		return false
	default:
		return false
	}
}
