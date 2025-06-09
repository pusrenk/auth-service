package helpers

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	Development = "development"
	Production  = "production"
	SessionExpiry = 24 * time.Hour // 24 hours
	CustomerServiceHost = "localhost"
	CustomerServicePort = "50051"
	UserRole = "user"
)

var (
	CustomerServiceURL = fmt.Sprintf("%s:%s", CustomerServiceHost, CustomerServicePort)
)

func GenerateSessionID() string {
	uuid := uuid.New()
	return uuid.String()
}