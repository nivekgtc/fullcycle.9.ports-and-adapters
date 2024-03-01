package application_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/nivekgtc/hexagonal/application"
	mock_application "github.com/nivekgtc/hexagonal/application/mocks"
	"github.com/stretchr/testify/require"
	// "github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProduct := mock_application.NewMockProductInterface(ctrl)
	mockPersistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	mockPersistence.EXPECT().Get(gomock.Any()).Return(mockProduct, nil)

	service := application.ProductService{
		Persistence: mockPersistence,
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, mockProduct, result)
}

func TestProductService_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProduct := mock_application.NewMockProductInterface(ctrl)
	mockPersistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	mockPersistence.EXPECT().Save(gomock.Any()).Return(mockProduct, nil)

	service := application.ProductService{
		Persistence: mockPersistence,
	}

	result, err := service.Create("Product 1", 10)
	require.Nil(t, err)
	require.Equal(t, mockProduct, result)
}

func TestProductService_EnableDisable(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockProduct := mock_application.NewMockProductInterface(ctrl)
	mockProduct.EXPECT().Enable().Return(nil)
	mockProduct.EXPECT().Disable().Return(nil)

	mockPersistence := mock_application.NewMockProductPersistenceInterface(ctrl)
	mockPersistence.EXPECT().Save(gomock.Any()).Return(mockProduct, nil).AnyTimes()

	service := application.ProductService{
		Persistence: mockPersistence,
	}

	result, err := service.Enable(mockProduct)
	require.Nil(t, err)
	require.Equal(t, mockProduct, result)

	result, err = service.Disable(mockProduct)
	require.Nil(t, err)
	require.Equal(t, mockProduct, result)
}
