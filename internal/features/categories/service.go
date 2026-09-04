package categories

import (
	"errors"
	"sprig/internal/model"
	"strings"
)

type CategoryService struct {
	repository CategoryRepository
}

func NewCategoryService(repository CategoryRepository) *CategoryService {
	return &CategoryService{
		repository: repository,
	}
}

func (s *CategoryService) CreateCategory(req CreateCategoryRequest) (*model.Category, error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("Nama harus diisi!")
	}

	category := &model.Category{
		UserID: req.UserID,
		Name:   req.Name,
	}

	err := s.repository.AddCategory(category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) FindCategory(name string) (*model.Category, error) {
	if strings.TrimSpace(name) == "" {
		return nil, errors.New("Nama tidak boleh kosong!")
	}

	category, err := s.repository.FindCategoryByName(name)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *CategoryService) FindAllCategory() ([]model.Category, error) {
	categories, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *CategoryService) UpdateCategory(id uint64, req CreateCategoryUpdateRequest) (*model.Category, error) {
	existingCategory, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil && strings.TrimSpace(*req.Name) != "" {
		existingCategory.Name = *req.Name
	}

	// existingCategory = &model.Category{
	// 	Name: *req.Name,
	// }

	if err = s.repository.Save(existingCategory); err != nil {
		return nil, err
	}

	return existingCategory, nil
}

func (s *CategoryService) DeleteCategory(name string) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("Nama tidak boleh kosong!")
	}

	err := s.repository.DeleteCategoryByName(name)
	if err != nil {
		return err
	}

	return nil
}
