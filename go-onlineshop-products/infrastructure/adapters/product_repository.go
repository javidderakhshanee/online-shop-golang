package adapters

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"onlineshopproduct/config"
	domain "onlineshopproduct/domain"
)

type ProductRepository struct {
	db            *gorm.DB
	configuration config.Configuration
}

func NewProductRepository(ctx context.Context) *ProductRepository {
	configuration := config.NewConfiguration()

	db, err := gorm.Open(mysql.Open(configuration.MySql.ConnectionString), &gorm.Config{})

	if err != nil {
		return nil
	}

	repo := &ProductRepository{
		db:            db,
		configuration: configuration,
	}

	repo.Migrate(ctx)

	return repo

}

func (repo *ProductRepository) Migrate(ctx context.Context) error {
	return repo.db.WithContext(ctx).AutoMigrate(&domain.Product{})
}

func (repo *ProductRepository) GetProduct(ctx context.Context, id int) (domain.Product, error) {

	var Product domain.Product

	err := repo.db.WithContext(ctx).First(&Product, id).Error

	return Product, err
}

func (repo *ProductRepository) GetProducts(ctx context.Context, categoryId int) ([]domain.Product, error) {
	var products []domain.Product

	err := repo.db.WithContext(ctx).Find(&products, "category_Id = ?", categoryId).Error

	return products, err
}

func (repo *ProductRepository) AddProduct(ctx context.Context, newProduct domain.Product) error {

	err := repo.db.WithContext(ctx).Create(&newProduct).Error

	return err
}

func (repo *ProductRepository) UpdateProduct(ctx context.Context, newProduct domain.Product) error {

	oldProduct, err := repo.GetProduct(ctx, newProduct.Id)
	if err != nil {
		return err
	}

	oldProduct.Name = newProduct.Name
	errResult := repo.db.WithContext(ctx).Save(&oldProduct).Error

	return errResult
}

func (repo *ProductRepository) DeleteProduct(ctx context.Context, id int) error {

	err := repo.db.WithContext(ctx).Delete(&domain.Product{}, id).Error

	return err
}
