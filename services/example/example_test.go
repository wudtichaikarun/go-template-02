package services_example

import (
	"testing"

	"github.com/stretchr/testify/assert"
	mocks "github.com/wudtichaikarun/poc-go-template-02/mocks/example"
	"github.com/wudtichaikarun/poc-go-template-02/repositories/example/models"
)

func TestListBanner(t *testing.T) {
	mockRepository := &mocks.ExampleRepository{}

	mockRepository.On("List").Return([]models.Example{}, nil)

	service := New(mockRepository)
	_, err := service.List()

	if err != nil {
		assert.Error(t, err)
	}

	mockRepository.AssertExpectations(t)
}
