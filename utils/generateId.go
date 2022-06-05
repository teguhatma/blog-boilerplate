package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateNewID() string {
	newID := uuid.NewString()
	fmtID := strings.Replace(newID, "-", "", -1)
	return fmtID
}
