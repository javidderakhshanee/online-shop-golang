package adapters

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"onlineshopproduct/config"
	domain "onlineshopproduct/domain"
)

var configuration config.Configuration

type CategoryRepository struct {
	db            *gorm.DB
	configuration config.Configuration
}

func NewCategoryRepository() *CategoryRepository {
	configuration = config.NewConfiguration()

	db, err := gorm.Open(mysql.Open(configuration.MySql.ConnectionString), &gorm.Config{})

	if err != nil {
		return nil
	}

	return &CategoryRepository{
		db:            db,
		configuration: configuration,
	}
}

func (repo *CategoryRepository) Migrate(ctx context.Context) error {
	m := &domain.Category{}
	return repo.db.WithContext(ctx).AutoMigrate(&m)
}

func (repo *CategoryRepository) GetCategory(ctx context.Context, id int) (domain.Category, error) {
	return domain.Category{}, nil
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

func (repo *CategoryRepository) UpdateCategory(ctx context.Context, category domain.Category) error {
	return nil
}

func (repo *CategoryRepository) DeleteCategory(ctx context.Context, id int) error {
	return nil
}
