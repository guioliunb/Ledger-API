package users

import (
	"encoding/json"
	"net/http"

	"github.com/guioliunb/Chain-Services/back-end/models"
	UsersModel "github.com/guioliunb/Chain-Services/back-end/models/v1/users"
)
func Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user models.User
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		err := decoder.Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		newUser , err := UsersModel.Store(user.FirstName, user.LastName, user.Email, user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packet , err := json.Marshal(newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(packet)
	}
}