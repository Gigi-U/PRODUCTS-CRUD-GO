package products

import (
	"context"

	"github.com/Gigi-U/PRODUCTS-CRUD-GO.git/internal/models"
)

// interfaz , que es un contrato
type Repository interface {
	Create(ctx context.Context, producto models.Producto) (models.Producto, error)
	GetAll(ctx context.Context) ([]models.Producto, error)
	GetByID(ctx context.Context, id int) (models.Producto, error)
	Update(ctx context.Context, producto models.Producto, id int) (models.Producto, error)
	Delete(ctx context.Context, id int) error
}