package rawresourcetypes

import (
	"encoding/json"
	"net/http"

	"github.com/guioliunb/Chain-Services/back-end/models"
	RawResourceTypesModel "github.com/guioliunb/Chain-Services/back-end/models/v1/rawresourcestypes"
)
func Store() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var rawresourcetype models.RawResourceType
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()
		err := decoder.Decode(&rawresourcetype)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		newRawResourceType , err := RawResourceTypesModel.Store(rawresourcetype.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packet , err := json.Marshal(newRawResourceType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(packet)
	}
}