package util

import (
	"log"
	"reflect"
	"strings"

	gobperror "github.com/mrpandey/gobp/src/internal/core/domain/error"
	fdom "github.com/mrpandey/gobp/src/internal/core/domain/furniture"

	english "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

const (
	FURNITURE_TYPE_VALIDATOR_TAG string = "valid_furniture_type"
)

type Validator struct {
	Validate  *validator.Validate
	Translate ut.Translator
}

func NewRequestBodyValidator(logger *StandardLogger) *Validator {
	eng := english.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	validate.RegisterTagNameFunc(jsonTagNameFunc)

	err := registerCustomValidations(validate)
	if err != nil {
		logger.Panicf("Could not register custom validations. %v", err)
	}

	err = en_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		logger.Panicf("Could not register default transalations for validator. %v", err)
	}

	err = registerCustomTranslations(validate, trans)
	if err != nil {
		logger.Panicf("Could not register custom transalations for validator. %v", err)
	}

	return &Validator{
		Validate:  validate,
		Translate: trans,
	}
}

func (v *Validator) ValidateStruct(logger *StandardLogger, s interface{}) error {
	err := v.Validate.Struct(s)
	if err != nil {
		failedValidation, transErr := translateValidationErrs(err, v.Translate)
		if transErr != nil {
			logger.Errorf("Could not translate validation error(s) for struct %+v. err = %v", s, err)
			return transErr
		}

		validationErr := gobperror.NewValidationError(failedValidation)
		return validationErr
	}
	return nil
}

// Registers struct field tags for custom validators.
func registerCustomValidations(v *validator.Validate) error {
	return v.RegisterValidation(FURNITURE_TYPE_VALIDATOR_TAG, fdom.FurnitureTypeValidator)
}

// Translates validation errors to more readable form. Returns only the first one.
func translateValidationErrs(err error, trans ut.Translator) (string, error) {
	errs, ok := err.(validator.ValidationErrors) //nolint:errorlint
	if !ok || len(errs) == 0 {
		return err.Error(), nil
	}

	translatedErr := errs[0].Translate(trans)
	return translatedErr, nil
}

type customTranslation struct {
	tag             string
	translation     string
	override        bool
	customRegisFunc validator.RegisterTranslationsFunc
	customTransFunc validator.TranslationFunc
}

// Each element of the slice contains arguments needed by Translator methods.
var customTranslations = []customTranslation{
	{
		tag:         FURNITURE_TYPE_VALIDATOR_TAG,
		translation: "{0}: invalid furniture type",
		override:    false,
	},
}

// Iterates over customTranslations and registers them with the Universal Translator.
func registerCustomTranslations(v *validator.Validate, trans ut.Translator) error {
	var err error
	for _, t := range customTranslations {

		switch {
		case t.customTransFunc != nil && t.customRegisFunc != nil:
			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, t.customTransFunc)
		case t.customTransFunc != nil && t.customRegisFunc == nil:
			err = v.RegisterTranslation(t.tag, trans, registrationFunc(t), t.customTransFunc)
		case t.customTransFunc == nil && t.customRegisFunc != nil:
			err = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, translateFunc)
		default:
			err = v.RegisterTranslation(t.tag, trans, registrationFunc(t), translateFunc)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

// Returns registration function for given custom translation.
func registrationFunc(t customTranslation) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) (err error) {
		if err = ut.Add(t.tag, t.translation, t.override); err != nil {
			return
		}
		return
	}
}

// Translates FieldError to human-readable message.
func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		log.Printf("warning: error translating FieldError: %#v", fe)
		return fe.(error).Error()
	}

	return t
}

// Returns json tag name of a struct field.
func jsonTagNameFunc(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	// skip if tag key says it should be ignored
	if name == "-" {
		return ""
	}
	return name
}
