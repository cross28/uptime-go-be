package health

import (
	"encoding/json"
	"net/http"
)

func Healthcheck(w http.ResponseWriter, r *http.Request) {
	health := &Health{
		Status: "healthy",
	}
	json.NewEncoder(w).Encode(health)
	// w.Write([]byte("healthy"))
}
