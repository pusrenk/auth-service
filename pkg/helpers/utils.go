package helpers

import (
	"time"

	"github.com/google/uuid"
)

const (
	Development = "development"
	Production  = "production"
	SessionExpiry = 24 * time.Hour // 24 hours
)

func GenerateSessionID() string {
	uuid := uuid.New()
	return uuid.String()
}