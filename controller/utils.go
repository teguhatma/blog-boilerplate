package controller

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// func convertToInt32(l, o string) (int32, int32, error) {
// 	limit, err := strconv.Atoi(l)
// 	if err != nil {
// 		return -1, -1, err
// 	}
// 	offset, err := strconv.Atoi(o)
// 	if err != nil {
// 		return -1, -1, err
// 	}

// 	a := int32(limit)
// 	b := int32(offset)

// 	return a, b, nil
// }

func structValidator(request any) error {
	var errorMessages []string
	validate := validator.New()

	if err := validate.Struct(request); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("%s, condition %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
	}

	if errorMessages == nil {
		return nil
	}
	return fmt.Errorf(errorMessages[0])
}