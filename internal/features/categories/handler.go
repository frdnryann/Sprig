package categories

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	categoryService *CategoryService
}

func NewCategoryHandler(categoryService *CategoryService) *CategoryHandler {
	return &CategoryHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateCategoryRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "invalid JSON",
		})
		return
	}

	category, err := h.categoryService.CreateCategory(req)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{
			"message": err.Error(),
		})
		return
	}

	response := CategoryResponse{
		ID:     category.ID,
		UserID: category.UserID,
		Name:   category.Name,
	}

	writeJSON(w, http.StatusCreated, map[string]any{
		"message": "category created successfully",
		"data":    response,
	})
}

func (h *CategoryHandler) FindByName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	category, err := h.categoryService.FindCategory(name)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{
			"message": err.Error(),
		})
		return
	}

	response := CategoryResponse{
		ID:     category.ID,
		UserID: category.UserID,
		Name:   category.Name,
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"data": response,
	})
}

func (h *CategoryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	categories, err := h.categoryService.FindAllCategory()
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "invalid JSON",
		})
		return
	}

	responses := make([]CategoryResponse, 0, len(categories))

	for _, category := range categories {
		responses = append(responses, CategoryResponse{
			ID:     category.ID,
			UserID: category.UserID,
			Name:   category.Name,
		})
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"data": responses,
	})
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	var req CreateCategoryRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "invalid JSON",
		})
	}

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"message": "invalid JSON",
		})
	}

	category, err := h.categoryService.UpdateCategory(id, req)
	if err != nil {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{
			"message": err.Error(),
		})
	}

	response := CategoryResponse{
		ID:     category.ID,
		UserID: category.UserID,
		Name:   category.Name,
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"message": "category updated successfully",
		"data":    response,
	})
}

func (h *CategoryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	err := h.categoryService.DeleteCategory(name)
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

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}
