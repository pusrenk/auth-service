# Testing Guide for Auth Service

This guide covers the comprehensive testing setup for the auth service microservice.

## Test Structure

```
auth-service/
├── test/
│   └── mocks/                    # Generated mocks using Mockery
│       ├── UserServiceMock.go
│       └── UserRedisRepositoryMock.go
├── internal/user/
│   ├── handlers/
│   │   ├── user.go
│   │   └── user_test.go          # Handler tests
│   ├── services/
│   │   ├── user.go
│   │   └── user_test.go          # Service tests
│   └── repositories/
│       ├── redis.go
│       └── redis_test.go         # Repository tests
└── .mockery.yaml                 # Mockery configuration
```

## Tools and Dependencies

### Testing Libraries
- **testify**: Comprehensive testing toolkit with assertions and test suites
- **mockery**: Mock generation for Go interfaces

### Dependencies Added
```bash
go get github.com/stretchr/testify
go install github.com/vektra/mockery/v2@latest
```

## Mock Generation

### Configuration (`.mockery.yaml`)
```yaml
with-expecter: true
mockname: "{{.InterfaceName}}Mock"
filename: "{{.MockName}}.go"
outpkg: mocks
dir: "test/mocks"
packages:
  github.com/pusrenk/auth-service/internal/user/services:
    interfaces:
      UserService:
  github.com/pusrenk/auth-service/internal/user/repositories:
    interfaces:
      UserRedisRepository:
```

### Generate Mocks
```bash
# Generate mocks from configuration
mockery

# Or generate all mocks
mockery --all
```

## Test Categories

### 1. Unit Tests
- **Repository Layer**: Tests Redis operations with mocked Redis client
- **Service Layer**: Tests business logic with mocked repositories
- **Handler Layer**: Tests gRPC handlers with mocked services

### 2. Integration Tests
- Tests with actual Redis connection (requires running Redis)
- Marked with `testing.Short()` check to skip in CI/CD

### 3. Test Suites
Using testify suites for organized testing with setup/teardown:

```go
type UserServiceTestSuite struct {
    suite.Suite
    mockRepo *mocks.UserRedisRepositoryMock
    service  UserService
}

func (suite *UserServiceTestSuite) SetupTest() {
    suite.mockRepo = mocks.NewUserRedisRepositoryMock(suite.T())
    suite.service = NewUserService(suite.mockRepo)
}
```

## Running Tests

### All Tests
```bash
go test ./... -v
```

### Specific Package
```bash
go test ./internal/user/services -v
go test ./internal/user/handlers -v
go test ./internal/user/repositories -v
```

### Skip Integration Tests
```bash
go test ./... -short -v
```

### With Coverage
```bash
go test ./... -cover
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Test Examples

### Service Test with Mock
```go
func (suite *UserServiceTestSuite) TestGetUserBySessionID_Success() {
    // Arrange
    ctx := context.Background()
    sessionID := "test-session-id"
    expectedUser := &entities.User{
        ID:       "user-123",
        Username: "testuser",
        Email:    "test@example.com",
        Role:     "user",
    }
    
    suite.mockRepo.EXPECT().GetUserBySessionID(ctx, sessionID).Return(expectedUser, nil)
    
    // Act
    result, err := suite.service.GetUserBySessionID(ctx, sessionID)
    
    // Assert
    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), expectedUser, result)
}
```

### Handler Test
```go
func (suite *UserHandlerTestSuite) TestStoreUserSession_Success() {
    ctx := context.Background()
    user := &entities.User{
        ID:       "user-123",
        Username: "testuser",
        Email:    "test@example.com",
        Role:     "user",
    }
    
    suite.mockService.EXPECT().StoreUserSession(ctx, user).Return(nil)
    
    err := suite.handler.userService.StoreUserSession(ctx, user)
    
    assert.NoError(suite.T(), err)
}
```

## Best Practices

### 1. Test Organization
- Use test suites for complex scenarios
- Group related tests together
- Use descriptive test names

### 2. Mocking
- Mock external dependencies (Redis, databases)
- Use generated mocks for consistency
- Verify mock expectations

### 3. Test Data
- Use realistic test data
- Create helper functions for common test data
- Avoid hard-coded values

### 4. Coverage
- Aim for high test coverage
- Test both success and error cases
- Test edge cases and boundary conditions

### 5. Integration Tests
- Use separate test database/Redis instances
- Clean up test data after each test
- Use docker-compose for local testing environment

## CI/CD Integration

### GitHub Actions Example
```yaml
name: Tests
on: [push, pull_request]
jobs:
  test:
    runs-on: ubuntu-latest
    services:
      redis:
        image: redis:alpine
        ports:
          - 6379:6379
    steps:
      - uses: actions/checkout@v2
      - uses: actions/setup-go@v2
        with:
          go-version: 1.21
      - name: Run tests
        run: go test ./... -v -cover
```

## Debugging Tests

### Verbose Output
```bash
go test ./... -v
```

### Run Specific Test
```bash
go test -run TestUserService ./internal/user/services
```

### Debug with Print Statements
```go
func TestSomething(t *testing.T) {
    t.Logf("Debug info: %+v", someData)
    // ... test code
}
```

This testing setup provides comprehensive coverage for your gRPC microservice with proper mocking and separation of concerns. 