package custom

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ParseError(err error) []string {
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		fmt.Println("err valid", validationErrs)
		errorMessages := make([]string, len(validationErrs))
		for i, e := range validationErrs {

			errorMessages[i] = fmt.Sprintf("The field %s is %s", e.Field(), e.Tag())
			// }
		}
		return errorMessages
	} else if marshallingErr, ok := err.(*json.UnmarshalTypeError); ok {
		return []string{fmt.Sprintf("The field %s must be a %s", marshallingErr.Field, marshallingErr.Type.String())}
	}
	return nil
}
