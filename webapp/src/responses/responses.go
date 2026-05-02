package responses

import (
	"encoding/json"
	"log"
	"net/http"
	"webapp/src/cookies"
)

// ErroAPI struct is used to represent the api error response
type ErroAPI struct {
	Erro string `json:"erro"`
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

// HandleStatusCodeError is used to handle requests with status code error 400 or higher
func HandleStatusCodeError(w http.ResponseWriter, r *http.Response) {
	if r.StatusCode == http.StatusUnauthorized {
		cookies.Delete(w)
		http.Redirect(w, r.Request, "/login", http.StatusFound)
		return
	}

	var erro ErroAPI
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
