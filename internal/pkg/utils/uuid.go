package utils

import (
	"fmt"

	"github.com/satori/go.uuid"
)

// NewUUID ..
// Creates a new UUID and returns string
func NewUUID() string {
	u := uuid.NewV4()
	// if err != nil {
	// 	return "", fmt.Errorf("Can't generate new UUID", err)
	// }

	// trim the uuid without hyphen
	return fmt.Sprintf("%x", u[0:])
}