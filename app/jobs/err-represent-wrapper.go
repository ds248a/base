package jobs

import (
	"fmt"
)

type errMarshal struct {
	error
}

// MarshalJSON
func (err errMarshal) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", err.String())), nil
}

// Error
func (err errMarshal) Error() string {
	if err.error == nil {
		return "null"
	}
	return err.error.Error()
}

// String
func (err errMarshal) String() string {
	if i, ok := err.error.(fmt.Stringer); ok && err.error != nil {
		return i.String()
	}
	return err.Error()
}
