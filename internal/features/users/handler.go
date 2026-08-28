package users

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService *UserService
}

func NewUserHandler(userService *UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest // tampung dulu di DTO

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "invalid JSON",
		})
		return
	}

	user, err := h.userService.CreateUser(req)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{
			"message": err.Error(),
		})
		return
	}

	response := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	writeJSON(w, http.StatusCreated, map[string]any{
		"message": "user created successfully",
		"data":    response,
	})
}

func (h *UserHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "invalid user id",
		})
		return
	}

	user, err := h.userService.FindUserByID(id)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{
			"message": err.Error(),
		})
		return
	}

	response := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"data": response,
	})
}

func (h *UserHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.FindAllUser()
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "failed to get users",
		})
		return
	}

	responses := make([]UserResponse, 0, len(users))

	for _, user := range users {
		responses = append(responses, UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		})
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"data": responses,
	})
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	var req CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "invalid JSON",
		})
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "invalid JSON",
		})
	}

	user, err := h.userService.UpdateUser(id, req)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{
			"message": err.Error(),
		})
	}

	response := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"message": "user updated successfully",
		"data":    response,
	})
}

func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "invalid user id",
		})
		return
	}

	err = h.userService.DeleteUser(id)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{
			"message": err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{
		"message": "deleted successfully",
	})
}

// helper (nanti kalao udah banyak yang make pindahin aja)
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}
