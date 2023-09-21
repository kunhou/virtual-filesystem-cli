package usecase

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	"github/kunhou/virtual-filesystem-cli/internal/entity"
	"github/kunhou/virtual-filesystem-cli/internal/mocks"
)

type fileTestSuite struct {
	suite.Suite
	ctrl     *gomock.Controller
	mockRepo *mocks.MockIRepository

	usecase *Usecase
}

func (suite *fileTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockRepo = mocks.NewMockIRepository(suite.ctrl)

	suite.usecase = NewUsecase(suite.mockRepo)
}

func (suite *fileTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func TestFileTestSuite(t *testing.T) {
	suite.Run(t, new(fileTestSuite))
}

func (suite *fileTestSuite) TestCreateFile() {
	username := "test"
	folderName := "testfolder"
	fileName := "testfile"
	description := "testDescription"

	suite.mockRepo.EXPECT().CreateFile(username, folderName, &fileMatcher{newFile("testfile", "testDescription")}).Return(nil)

	err := suite.usecase.CreateFile(entity.CreateFileParam{
		Username:   username,
		FolderName: folderName,

		Name:        fileName,
		Description: description,
	})
	suite.NoError(err)
}

// fileMatcher is a custom matcher for comparing file entities
type fileMatcher struct {
	expected *entity.File
}

func (m *fileMatcher) Matches(x interface{}) bool {
	actual, ok := x.(*entity.File)
	if !ok {
		return false
	}

	// Here, we only compare the Name and Description fields without CreatedAt
	return actual.Name == m.expected.Name && actual.Description == m.expected.Description
}

func (m *fileMatcher) String() string {
	return fmt.Sprintf("is equal to file having name: %s and description: %s", m.expected.Name, m.expected.Description)
}
