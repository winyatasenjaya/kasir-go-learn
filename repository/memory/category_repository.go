package memory

import (
	"errors"
	"kasir-api/domain"
)

// CategoryRepository implements domain.CategoryRepository interface
type CategoryRepository struct {
	categories []domain.Category
}

// NewCategoryRepository creates a new instance of CategoryRepository with sample data
func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{
		categories: []domain.Category{
			{ID: 1, Name: "Electronics", Description: "Devices and gadgets"},
			{ID: 2, Name: "Home Appliances", Description: "Appliances for home use"},
			{ID: 3, Name: "Books", Description: "Various genres of books"},
		},
	}
}

// GetAll returns all categories
func (r *CategoryRepository) GetAll() ([]domain.Category, error) {
	return r.categories, nil
}

// GetByID returns a category by ID
func (r *CategoryRepository) GetByID(id int) (*domain.Category, error) {
	for _, c := range r.categories {
		if c.ID == id {
			return &c, nil
		}
	}
	return nil, errors.New("category not found")
}

// Create adds a new category
func (r *CategoryRepository) Create(category *domain.Category) error {
	category.ID = len(r.categories) + 1
	r.categories = append(r.categories, *category)
	return nil
}

// Update modifies an existing category
func (r *CategoryRepository) Update(id int, category *domain.Category) error {
	for i := range r.categories {
		if r.categories[i].ID == id {
			category.ID = id
			r.categories[i] = *category
			return nil
		}
	}
	return errors.New("category not found")
}

// Delete removes a category by ID
func (r *CategoryRepository) Delete(id int) error {
	for i, c := range r.categories {
		if c.ID == id {
			r.categories = append(r.categories[:i], r.categories[i+1:]...)
			return nil
		}
	}
	return errors.New("category not found")
}
