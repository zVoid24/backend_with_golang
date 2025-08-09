package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/zVoid24/backend_with_golang/internal/types"
	"github.com/zVoid24/backend_with_golang/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, err.Error())
			return
		}
		slog.Info("Creating a student")
		response.WriteJson(w, http.StatusCreated, map[string]string{
			"success": "ok",
		})
	}
}
