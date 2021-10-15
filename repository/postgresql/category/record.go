package category

import (
	"github.com/daniel5u/suisei/domain/category"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string
}

func repositoryToDomain(categoryRepository Category) category.Domain {
	return category.Domain{
		ID:        int(categoryRepository.ID),
		Name:      categoryRepository.Name,
		CreatedAt: categoryRepository.CreatedAt,
		UpdatedAt: categoryRepository.UpdatedAt,
	}
}

func domainToRepository(categoryDomain category.Domain) Category {
	return Category{
		Name: categoryDomain.Name,
	}
}
