package service

import (
	"kenykendf/digital-kiosk/internal/app/mocks"
	"kenykendf/digital-kiosk/internal/app/model"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetProductCategoriesLists(t *testing.T) {
	ctrl := gomock.NewController(t)

	repo := mocks.NewMockProductCategoryRepo(ctrl)
	repo.EXPECT().GetProductCategoriesLists().Return([]model.ProductCategories{
		{
			ID:          1,
			Name:        "Komputer",
			Description: "Komputer",
		},
		{
			ID:          2,
			Name:        "Handphone",
			Description: "Handphone",
		},
	}, nil)

	svc := NewProductCategoryService(repo)
	prodCategories, err := svc.GetProductCategoriesLists()
	total := len(prodCategories)

	t.Run("GetProductCategoriesLists_success", func(t *testing.T) {
		assert.Equal(t, total, 2)
		assert.NoError(t, err)
	})
}

func TestGetProductCategoryByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	repo := mocks.NewMockProductCategoryRepo(ctrl)
	repo.EXPECT().GetProductCategoryByID(gomock.Any()).Return(model.ProductCategories{
		ID:          1,
		Name:        "Komputer",
		Description: "Komputer",
	}, nil)

	svc := NewProductCategoryService(repo)
	prodCategories, err := svc.GetProductCategoryByID("1")

	t.Run("GetProductCategoryByID_success", func(t *testing.T) {
		assert.Equal(t, int(prodCategories.ID), 1)
		assert.NoError(t, err)
	})
}
