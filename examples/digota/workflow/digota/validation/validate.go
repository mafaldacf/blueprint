package validation

/* import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/go-playground/validator.v9"
)

var v = validator.New()

// Validate is doing struct level validation using validator.v9
func Validate(req interface{}) error {
	// validate req message
	if err := v.Struct(req); err != nil {
		if v, ok := err.(validator.ValidationErrors); ok {
			err = v
		}
		return status.Errorf(codes.InvalidArgument, "Request validation failed: %s", err)
	}
	return nil
} */
