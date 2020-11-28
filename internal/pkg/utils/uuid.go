package utils

import (
	"fmt"

	"github.com/satori/go.uuid"
)

// NewUUID ..
// Creates a new UUID and returns string
func NewUUID() string {
	u := uuid.NewV4()

	// trim the uuid without hyphen
	return fmt.Sprintf("%x", u[0:])
}