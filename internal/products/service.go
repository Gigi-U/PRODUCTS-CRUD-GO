package products

import (
	"context"
	"log"

	"github.com/Gigi-U/PRODUCTS-CRUD-GO.git/internal/models"
)

type Service interface {
	Create(ctx context.Context, producto models.Producto) (models.Producto, error)
	GetAll(ctx context.Context) ([]models.Producto, error)
	GetByID(ctx context.Context, id string) (models.Producto, error)
	Update(ctx context.Context, producto models.Producto, id string) (models.Producto, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	repository Repository
}

func NewServiceProduct(repository Repository) Service {
	return &service{repository: repository}
}

// Create ....
func (s *service) Create(ctx context.Context, producto models.Producto) (models.Producto, error) {
	producto, err := s.repository.Create(ctx, producto)
	if err != nil {
		log.Println("[ProductsService][Create] error creating product", err)
		return models.Producto{}, err
	}

	return producto, nil
}

// GetAll ...
func (s *service) GetAll(ctx context.Context) ([]models.Producto, error) {
	listProducts, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("[ProductsService][GetAll] error getting all products", err)
		return []models.Producto{}, err
	}

	return listProducts, nil
}

// GetByID ....
func (s *service) GetByID(ctx context.Context, id string) (models.Producto, error) {
	producto, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[ProductsService][GetByID] error getting product by ID", err)
		return models.Producto{}, err
	}

	return producto, nil
}

// Update ...
func (s *service) Update(ctx context.Context, producto models.Producto, id string) (models.Producto, error) {
	producto, err := s.repository.Update(ctx, producto, id)
	if err != nil {
		log.Println("[ProductsService][Update] error updating product by ID", err)
		return models.Producto{}, err
	}

	return producto, nil
}

// Delete ...
func (s *service) Delete(ctx context.Context, id string) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		log.Println("[ProductsService][Delete] error deleting product by ID", err)
		return err
	}

	return nil
}