package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-firebase/entity"
	"testing"
)

type mockRepository struct {
	mock.Mock
}

func (mock *mockRepository) Save(*entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *mockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(mockRepository)

	var identifier int64 = 1

	post := entity.Post{
		ID:    1,
		Title: "A",
		Text:  "B",
	}
	// Setup expectations
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	// Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "A", result[0].Title)
	assert.Equal(t, "B", result[0].Text)
}

func TestCreate(t *testing.T) {
	mockRepo := new(mockRepository)

	post := entity.Post{
		Title: "A", Text: "B",
	}

	// Set expectations
	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	mockRepo.AssertExpectations(t)

	assert.Nil(t, err)
	assert.NotNil(t, result.ID)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)

	assert.Equal(t, err.Error(), "The post is empty ")
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{
		ID:    1,
		Title: "",
		Text:  "",
	}

	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err)

	assert.Equal(t, err.Error(), "The post title is empty ")
}
