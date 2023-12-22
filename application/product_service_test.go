package application

import (
	"testing"
	"github.com/jeffersonmartins/go-hexagonal/mocks/application"
	gomock "github.com/golang/mock/gomock"
	testify "github.com/stretchr/testify/require"
)

func TestProductService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	product := NewMockProductInterface(ctrl)
	persistence := NewMockProductPersistenceInterface(ctrl)

	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service = &ProductService{persistence: persistence
	}

	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)
}