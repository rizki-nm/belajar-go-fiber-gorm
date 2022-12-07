package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/rizki-nm/belajar-go-fiber-gorm/model/web"
)

func Validate(request any) []*web.ErrorValidateResponse {
	var validate = validator.New()

	var errors []*web.ErrorValidateResponse

	err := validate.Struct(request)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element web.ErrorValidateResponse
			element.FailedField = err.StructField()
			element.Tag = err.Tag()
			element.Value = err.Value()
			errors = append(errors, &element)
		}
	}

	return errors
}
