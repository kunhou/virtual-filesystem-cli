package usecase

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github/kunhou/virtual-filesystem-cli/internal/entity"
	"github/kunhou/virtual-filesystem-cli/internal/mocks"
)

type folderTestSuite struct {
	suite.Suite
	ctrl     *gomock.Controller
	mockRepo *mocks.MockIRepository

	usecase *Usecase
}

func (suite *folderTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockRepo = mocks.NewMockIRepository(suite.ctrl)

	suite.usecase = NewUsecase(suite.mockRepo)
}

func (suite *folderTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func TestFolderTestSuite(t *testing.T) {
	suite.Run(t, new(folderTestSuite))
}

func (suite *folderTestSuite) TestCreateFolder() {
	username := "test"
	folderName := "testFolder"
	description := "testDescription"

	suite.mockRepo.EXPECT().CreateFolder(username, &folderMatcher{newFolder("testFolder", "testDescription")}).Return(nil)

	err := suite.usecase.CreateFolder(entity.CreateFolderParam{
		Username:    username,
		Name:        folderName,
		Description: description,
	})
	suite.NoError(err)
}

func (suite *folderTestSuite) TestDeleteFolder() {
	username := "test"
	folderName := "testFolder"

	suite.mockRepo.EXPECT().DeleteFolder("test", "testFolder").Return(nil)

	err := suite.usecase.DeleteFolder(username, folderName)
	suite.NoError(err)
}

// folderMatcher is a custom matcher for comparing Folder entities
type folderMatcher struct {
	expected *entity.Folder
}

func (m *folderMatcher) Matches(x interface{}) bool {
	actual, ok := x.(*entity.Folder)
	if !ok {
		return false
	}

	// Here, we only compare the Name and Description fields without CreatedAt
	return actual.Name == m.expected.Name && actual.Description == m.expected.Description
}

func (m *folderMatcher) String() string {
	return fmt.Sprintf("is equal to folder having name: %s and description: %s", m.expected.Name, m.expected.Description)
}
