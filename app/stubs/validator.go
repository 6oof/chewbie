package stubs

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

type StructValidator struct {
	Validate *validator.Validate
	Decoder  *schema.Decoder
}

type FormErrs struct {
	Field string
	Error string
}

func NewValidator() *StructValidator {
	return &StructValidator{
		Validate: validator.New(),
		Decoder:  schema.NewDecoder(),
	}
}

func (v *StructValidator) ValidateAndMapForm(w http.ResponseWriter, r *http.Request, data interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	if err := v.Decoder.Decode(data, r.Form); err != nil {
		return err
	}

	if err := v.Validate.Struct(data); err != nil {
		return err
	}

	return nil
}

func (v *StructValidator) ExtractValidationErrors(err error) []FormErrs {
	var fe []FormErrs
	for _, err := range err.(validator.ValidationErrors) {
		cer := FormErrs{
			Field: err.Namespace(),
			Error: err.Error(),
		}
		fe = append(fe, cer)

	}
	return fe
}
