package mapper

import (
	"backend/dto"
	"backend/models"
)

func CategoryDtoToModel(categoryDto dto.Category) *models.Category {
	return &models.Category{
		ID:          categoryDto.ID,
		Name:        categoryDto.Name,
		Description: categoryDto.Description,
	}
}

func CategoryModelToDto(category *models.Category) *dto.Category {
	return &dto.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}

func CategoryModelToDtoList(categoryList []models.Category) []dto.Category {
	var category []dto.Category
	for _, m := range categoryList {
		category = append(category, *CategoryModelToDto(&m))
	}
	return category
}
