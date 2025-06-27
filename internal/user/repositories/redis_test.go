package repositories

import (
	"encoding/json"
	"testing"

	"github.com/pusrenk/auth-service/internal/user/entities"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// UserRedisRepositoryTestSuite defines the test suite
type UserRedisRepositoryTestSuite struct {
	suite.Suite
	repository UserRedisRepository
}

func (suite *UserRedisRepositoryTestSuite) SetupTest() {
	// For integration tests, you would use a real Redis client
	// For unit tests, we'll test the logic with a simpler approach
}

func (suite *UserRedisRepositoryTestSuite) TestNewUserRedisRepository() {
	// Test the constructor
	mockClient := redis.NewClient(&redis.Options{})
	repo := NewUserRedisRepository(mockClient)

	assert.NotNil(suite.T(), repo)
	assert.IsType(suite.T(), &userRedisRepository{}, repo)
}

// Run the test suite
func TestUserRedisRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRedisRepositoryTestSuite))
}

// Unit tests for JSON marshaling/unmarshaling logic
func TestUserRedisRepository_JSONOperations(t *testing.T) {
	user := &entities.User{
		ID:       "user-123",
		Username: "testuser",
		Email:    "test@example.com",
		Password: "hashed-password",
		Role:     "user",
	}

	// Test JSON marshaling
	jsonData, err := json.Marshal(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, jsonData)

	// Test JSON unmarshaling
	var unmarshaledUser entities.User
	err = json.Unmarshal(jsonData, &unmarshaledUser)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, unmarshaledUser.ID)
	assert.Equal(t, user.Username, unmarshaledUser.Username)
	assert.Equal(t, user.Email, unmarshaledUser.Email)
}

func TestUserRedisRepository_InvalidJSON(t *testing.T) {
	invalidJSON := "invalid-json"
	var user entities.User

	err := json.Unmarshal([]byte(invalidJSON), &user)
	assert.Error(t, err)
}

// Integration test functions (these would require a running Redis instance)
func TestUserRedisRepository_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}
	// TODO: to use real redis

	// This is an example of how you would do integration testing
	// You would need to have Redis running and configured properly

	/*
		client := redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			DB:   1, // Use a test database
		})

		repo := NewUserRedisRepository(client)
		ctx := context.Background()

		// Test data
		user := &entities.User{
			ID:       "test-user-123",
			Username: "testuser",
			Email:    "test@example.com",
			Password: "hashed-password",
			Role:     "user",
		}

		// Store user session
		err := repo.StoreUserSession(ctx, user)
		assert.NoError(t, err)

		// Retrieve user by session (this would need the actual session ID)
		// For integration tests, you'd need to modify the repository to return the session ID
	*/
}
