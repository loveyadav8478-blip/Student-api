package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/loveyadav8478-blip/student-api/internal/types"
	"github.com/loveyadav8478-blip/student-api/internal/utils/response"
	"io"
	"log/slog"
	"net/http"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		slog.Info("creating a student")

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralResponse(fmt.Errorf("empty body!!")))
			return
		}
		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralResponse(err))
			return
		}

		// Request validator
		if err := validator.New().Struct(student); err != nil {
			validationErrors := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest,response.ValidationError(validationErrors))
			return
		}


		// w.Write([]byte("Welcome to student"))
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})

	}
}
