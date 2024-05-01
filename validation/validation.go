package validation

import (
	"context"
	"reflect"
	"strings"

	"github.com/cilloparch/cillop/i18np"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Field     string      `json:"field"`
	Message   string      `json:"message"`
	Namespace string      `json:"namespace,omitempty"`
	Value     interface{} `json:"value"`
}

type Validator struct {
	validate *validator.Validate
	i18n     *i18np.I18n
}

func New(i *i18np.I18n) *Validator {
	return &Validator{
		validate: validator.New(),
		i18n:     i,
	}
}

func (v *Validator) ValidateStruct(ctx context.Context, s interface{}, languages ...string) []*ErrorResponse {
	var errors []*ErrorResponse
	err := v.validate.StructCtx(ctx, s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			ns := v.mapStructNamespace(err.Namespace())
			if ns != "" {
				element.Namespace = ns
			}
			element.Field = err.Field()
			element.Value = err.Value()
			element.Message = v.translateErrorMessage(err, languages...)
			errors = append(errors, &element)
		}
	}
	return errors
}

func (v *Validator) ValidateMap(ctx context.Context, m map[string]interface{}, rules map[string]interface{}, languages ...string) []*ErrorResponse {
	var errors []*ErrorResponse
	errMap := v.validate.ValidateMapCtx(ctx, m, rules)
	for key, err := range errMap {
		var element ErrorResponse
		if _err, ok := err.(validator.ValidationErrors); ok {
			for _, err := range _err {
				ns := v.mapStructNamespace(err.Namespace())
				if ns != "" {
					element.Namespace = ns
				}
				element.Field = err.Field()
				if element.Field == "" {
					element.Field = key
				}
				element.Value = err.Value()
				element.Message = v.translateErrorMessage(err, languages...)
				errors = append(errors, &element)
			}
			continue
		}
		element.Field = key
		element.Value = m[key]
		errors = append(errors, &element)
	}
	return errors
}

func (v *Validator) mapStructNamespace(ns string) string {
	str := strings.Split(ns, ".")
	return strings.Join(str[1:], ".")
}

func (v *Validator) ConnectCustom() {
	_ = v.validate.RegisterValidation("username", validateUserName)
	_ = v.validate.RegisterValidation("password", validatePassword)
	_ = v.validate.RegisterValidation("locale", validateLocale)
	_ = v.validate.RegisterValidation("object_id", validateObjectId)
	_ = v.validate.RegisterValidation("slug", validateSlug)
	_ = v.validate.RegisterValidation("gender", validateGender)
	_ = v.validate.RegisterValidation("phone", validatePhone)
	_ = v.validate.RegisterValidation("uuid", validateUUID)
}

func (v *Validator) Register(key string, fn validator.Func, callValidationEvenIfNull ...bool) {
	_ = v.validate.RegisterValidation(key, fn, callValidationEvenIfNull...)
}

func (v *Validator) RegisterTagName() {
	v.validate.RegisterTagNameFunc(func(f reflect.StructField) string {
		name := strings.SplitN(f.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
}
