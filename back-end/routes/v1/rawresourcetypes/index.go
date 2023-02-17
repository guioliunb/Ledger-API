package rawresourcetypes

import (
	"encoding/json"
	"fmt"
	"net/http"

	RawResourceTypesModel "github.com/guioliunb/Chain-Services/back-end/models/v1/rawresourcestypes"
)


func Index() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		rawresourcetypes, err := RawResourceTypesModel.Index()

		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		packet, err := json.Marshal(rawresourcetypes)

		if err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		
		fmt.Printf("Current Current RawResourceType: %s \n", packet );
		w.Write(packet)
	}
}