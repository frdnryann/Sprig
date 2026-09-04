package bootstrap

import (
	"database/sql"
	"sprig/internal/features/categories"
	"sprig/internal/features/users"
)

// List layer yang mau diinject
type Container struct {
	UserHandler     *users.UserHandler
	CategoryHandler *categories.CategoryHandler
}

// Setiap ada DI, tambahnya disini aja
func NewContainer(db *sql.DB) *Container {
	usersRepository := users.NewUserRepository(db)
	usersService := users.NewUserService(usersRepository)
	usersHandler := users.NewUserHandler(usersService)

	categoriesRepository := categories.NewCategoryRepository(db)
	categoriesService := categories.NewCategoryService(categoriesRepository)
	categoriesHandler := categories.NewCategoryHandler(categoriesService)

	return &Container{
		UserHandler:     usersHandler,
		CategoryHandler: categoriesHandler,
	}
}
