package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Check(value interface{}) bool {
	err := validate.Struct(value)
	if err != nil {
		logrus.Println("error validating struct : ", err)
	}
	return err != nil
}
