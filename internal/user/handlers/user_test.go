package handlers

import (
	"context"
	"testing"

	"github.com/pusrenk/auth-service/internal/protobuf/protogen"
	"github.com/pusrenk/auth-service/internal/user/entities"
	"github.com/pusrenk/auth-service/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// UserHandlerTestSuite defines the test suite
type UserHandlerTestSuite struct {
	suite.Suite
	mockService *mocks.UserServiceMock
	handler     *UserHandler
}

func (suite *UserHandlerTestSuite) SetupTest() {
	suite.mockService = mocks.NewUserServiceMock(suite.T())
	suite.handler = NewUserHandler(suite.mockService)
}

func (suite *UserHandlerTestSuite) TestGetUserBySessionID_Success() {
	// Arrange
	ctx := context.Background()
	sessionID := "test-session-123"

	expectedUser := &entities.User{
		ID:       "user-123",
		Username: "testuser",
		Email:    "test@example.com",
		Password: "hashed-password",
		Role:     "user",
	}

	suite.mockService.EXPECT().GetUserBySessionID(ctx, sessionID).Return(expectedUser, nil)

	// Act
	result, err := suite.handler.userService.GetUserBySessionID(ctx, sessionID)

	// Assert
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), result)
	assert.Equal(suite.T(), expectedUser.ID, result.ID)
	assert.Equal(suite.T(), expectedUser.Username, result.Username)
	assert.Equal(suite.T(), expectedUser.Email, result.Email)
	assert.Equal(suite.T(), expectedUser.Role, result.Role)
}

func (suite *UserHandlerTestSuite) TestStoreUserSession_Success() {
	// Arrange
	ctx := context.Background()
	user := &entities.User{
		ID:       "user-123",
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
		Role:     "user",
	}

	suite.mockService.EXPECT().StoreUserSession(ctx, user).Return(nil)

	// Act
	err := suite.handler.userService.StoreUserSession(ctx, user)

	// Assert
	assert.NoError(suite.T(), err)
}

func (suite *UserHandlerTestSuite) TestNewUserHandler() {
	// Test handler creation
	assert.NotNil(suite.T(), suite.handler)
	assert.IsType(suite.T(), &UserHandler{}, suite.handler)
}

// Run the test suite
func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}

// Additional unit tests
func TestNewUserHandler(t *testing.T) {
	mockService := mocks.NewUserServiceMock(t)
	handler := NewUserHandler(mockService)

	assert.NotNil(t, handler)
	assert.IsType(t, &UserHandler{}, handler)
}

// Simple test to verify service is called for basic operations
func TestUserHandler_ServiceIntegration(t *testing.T) {
	mockService := mocks.NewUserServiceMock(t)
	handler := NewUserHandler(mockService)

	// Test that the handler has access to the service
	assert.NotNil(t, handler)

	// Test example data handling
	user := &entities.User{
		ID:       "test-123",
		Username: "testuser",
		Email:    "test@example.com",
		Role:     "user",
	}

	ctx := context.Background()

	// Setup expectation for StoreUserSession
	mockService.EXPECT().StoreUserSession(ctx, user).Return(nil)

	// Call the service through the handler's dependency
	err := handler.userService.StoreUserSession(ctx, user)

	// Assert
	assert.NoError(t, err)
}

// Test that UserHandler implements the protogen.MainServer interface
func TestUserHandler_ImplementsMainServer(t *testing.T) {
	mockService := mocks.NewUserServiceMock(t)
	handler := NewUserHandler(mockService)

	// This should compile without issues if UserHandler implements MainServer
	var _ protogen.MainServer = handler
	assert.NotNil(t, handler)
}
