package adapters

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"onlineshopproduct/config"
	domain "onlineshopproduct/domain"
)

type CategoryRepository struct {
	db            *gorm.DB
	configuration config.Configuration
}

func NewCategoryRepository(ctx context.Context) *CategoryRepository {
	configuration := config.NewConfiguration()

	db, err := gorm.Open(mysql.Open(configuration.MySql.ConnectionString), &gorm.Config{})

	if err != nil {
		return nil
	}

	repo := &CategoryRepository{
		db:            db,
		configuration: configuration,
	}

	repo.Migrate(ctx)

	return repo

}

func (repo *CategoryRepository) Migrate(ctx context.Context) error {
	return repo.db.WithContext(ctx).AutoMigrate(&domain.Category{})
}

func (repo *CategoryRepository) GetCategory(ctx context.Context, id int) (domain.Category, error) {

	var category domain.Category

	err := repo.db.WithContext(ctx).First(&category, id).Error

	return category, err
}

func (repo *CategoryRepository) GetCategories(ctx context.Context) ([]domain.Category, error) {
	var categories []domain.Category

	err := repo.db.WithContext(ctx).Find(&categories).Error

	return categories, err
}

func (repo *CategoryRepository) AddCategory(ctx context.Context, newCategory domain.Category) error {

	err := repo.db.WithContext(ctx).Create(&newCategory).Error

	return err
}

func (repo *CategoryRepository) UpdateCategory(ctx context.Context, newCategory domain.Category) error {

	oldCategory, err := repo.GetCategory(ctx, newCategory.Id)
	if err != nil {
		return err
	}

	oldCategory.Name = newCategory.Name
	errResult := repo.db.WithContext(ctx).Save(&oldCategory).Error

	return errResult
}

func (repo *CategoryRepository) DeleteCategory(ctx context.Context, id int) error {

	err := repo.db.WithContext(ctx).Delete(&domain.Category{}, id).Error

	return err
}
