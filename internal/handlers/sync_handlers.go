package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/bhavyavarshney-123/catalyst/internal/services/sync"
)

func Sync(s *sync.SyncService) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		param := r.URL.Query().Get("limit")
		limit, err := strconv.Atoi(param)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = s.Sync(limit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "Sync completed successfully")
	}

}
