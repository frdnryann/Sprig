package users

import (
	"errors"
	"sprig/internal/model"
	"strings"
)

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) CreateUser(req CreateUserRequest) (*model.User, error) {
	if strings.TrimSpace(*req.Name) == "" {
		return nil, errors.New("Nama dibutuhkan, tolong diisi")
	}

	if strings.TrimSpace(*req.Email) == "" {
		return nil, errors.New("Email harus diisi!")
	}

	if strings.TrimSpace(*req.Password) == "" {
		return nil, errors.New("Password dibutuhkan, tolong diisi!")
	}

	user := &model.User{
		Name:     *req.Name,
		Email:    *req.Email,
		Password: *req.Password,
	}

	err := s.repository.AddUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) FindUserByID(id uint64) (*model.User, error) {
	if id <= 0 {
		return nil, errors.New("ID tidak boleh kosong")
	}

	user, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) FindAllUser() ([]model.User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) UpdateUser(id uint64, req CreateUserRequest) (*model.User, error) {
	existingUser, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		existingUser.Name = *req.Name
	}

	if req.Email != nil {
		existingUser.Email = *req.Email
	}

	if req.Password != nil {
		existingUser.Password = *req.Password
	}

	if err = s.repository.Save(existingUser); err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (s *UserService) DeleteUser(id uint64) error {
	if id <= 0 {
		return errors.New("ID tidak boleh kosong")
	}

	err := s.repository.DeleteByID(id)
	if err != nil {
		return err
	}

	return nil
}
