package services

import (
	"context"
	"testing"

	"github.com/pusrenk/auth-service/internal/user/entities"
	"github.com/pusrenk/auth-service/internal/user/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// UserServiceTestSuite defines the test suite
type UserServiceTestSuite struct {
	suite.Suite
	mockRepo *mocks.UserRedisRepositoryMock
	service  UserService
}

func (suite *UserServiceTestSuite) SetupTest() {
	suite.mockRepo = mocks.NewUserRedisRepositoryMock(suite.T())
	suite.service = NewUserService(suite.mockRepo)
}

func (suite *UserServiceTestSuite) TestGetUserBySessionID_Success() {
	// Arrange
	ctx := context.Background()
	sessionID := "test-session-id"
	expectedUser := &entities.User{
		ID:       "user-123",
		Username: "testuser",
		Email:    "test@example.com",
		Password: "hashed-password",
		Role:     "user",
	}

	suite.mockRepo.EXPECT().GetUserBySessionID(ctx, sessionID).Return(expectedUser, nil)

	// Act
	result, err := suite.service.GetUserBySessionID(ctx, sessionID)

	// Assert
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedUser, result)
}

func (suite *UserServiceTestSuite) TestGetUserBySessionID_Error() {
	// Arrange
	ctx := context.Background()
	sessionID := "test-session-id"
	expectedError := assert.AnError

	suite.mockRepo.EXPECT().GetUserBySessionID(ctx, sessionID).Return(nil, expectedError)

	// Act
	result, err := suite.service.GetUserBySessionID(ctx, sessionID)

	// Assert
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Equal(suite.T(), expectedError, err)
}

func (suite *UserServiceTestSuite) TestStoreUserSession_Success() {
	// Arrange
	ctx := context.Background()
	user := &entities.User{
		ID:       "user-123",
		Username: "testuser",
		Email:    "test@example.com",
		Password: "hashed-password",
		Role:     "user",
	}

	suite.mockRepo.EXPECT().StoreUserSession(ctx, user).Return(nil)

	// Act
	err := suite.service.StoreUserSession(ctx, user)

	// Assert
	assert.NoError(suite.T(), err)
}

func (suite *UserServiceTestSuite) TestStoreUserSession_Error() {
	// Arrange
	ctx := context.Background()
	user := &entities.User{
		ID:       "user-123",
		Username: "testuser",
		Email:    "test@example.com",
		Password: "hashed-password",
		Role:     "user",
	}
	expectedError := assert.AnError

	suite.mockRepo.EXPECT().StoreUserSession(ctx, user).Return(expectedError)

	// Act
	err := suite.service.StoreUserSession(ctx, user)

	// Assert
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), expectedError, err)
}

// Additional test cases can be added here for other service methods

// Run the test suite
func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}

// Additional unit tests
func TestNewUserService(t *testing.T) {
	mockRepo := mocks.NewUserRedisRepositoryMock(t)
	service := NewUserService(mockRepo)

	assert.NotNil(t, service)
	assert.IsType(t, &userService{}, service)
}
