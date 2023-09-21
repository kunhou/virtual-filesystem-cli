package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github/kunhou/virtual-filesystem-cli/internal/mocks"
)

type userTestSuite struct {
	suite.Suite
	ctrl     *gomock.Controller
	mockRepo *mocks.MockIRepository

	usecase *Usecase
}

func (suite *userTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockRepo = mocks.NewMockIRepository(suite.ctrl)

	suite.usecase = NewUsecase(suite.mockRepo)
}

func (suite *userTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(userTestSuite))
}

func (suite *userTestSuite) TestCreateUser() {
	username := "test"

	suite.mockRepo.EXPECT().CreateUser(newUser("test")).Return(nil)

	err := suite.usecase.CreateUser(username)
	suite.NoError(err)
}
