package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/moisesPompilio/api-rest-go/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := r.ParseForm()
	if err != nil {
		log.Println("Error ao fazer o decode do json: $v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("inserido com sucesso. id: %v", id),
		}
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
