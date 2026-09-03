package users

import "time"

// dto?. buat request response json

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUserUpdateRequest berisi field yang dapat diperbarui melalui PATCH /users/{id}.
//
// Pointer digunakan untuk membedakan antara:
//   - field tidak dikirim     → nil
//   - field dikirim           → pointer ke nilai
//
// Dengan demikian, service dapat melakukan partial update tanpa
// menimpa field yang tidak ingin diubah oleh client.
type CreateUserUpdateRequest struct {
	Name     *string `json:"name"`
	Email    *string `json:"email"`
	Password *string `json:"password"`
}

type UserResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
