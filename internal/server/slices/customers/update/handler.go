package update

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/y0n1/crm-service/internal/models/aggregates"
	store_pkg "github.com/y0n1/crm-service/internal/store"
	"github.com/y0n1/crm-service/internal/utils/uuid"
)

const UrlPattern = "/customers/{id}"

func MakeHandler(ctx context.Context, store store_pkg.Storable[*aggregates.CustomerAggregate], logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var requestBody UpdateCustomerRequest
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", err.Error()), http.StatusBadRequest)
			return
		}

		uuid, err := uuid.ParseFromUrlParam(r, "id", true)
		if err != nil {
			http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", err.Error()), http.StatusBadRequest)
			return
		}

		aggregate, err := store.Get(uuid)
		if err != nil {
			http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", err.Error()), http.StatusNotFound)
			return
		}

		if err := aggregate.Update(
			requestBody.FirstName,
			requestBody.LastName,
			requestBody.Role,
			requestBody.Email,
			requestBody.Phone,
			requestBody.Contacted,
		); err != nil {
			http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", err.Error()), http.StatusBadRequest)
			return
		}

		if err := store.Update(aggregate); err != nil {
			http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", err.Error()), http.StatusBadRequest)
			return
		}

		response := newUpdateCustomerResponse(*aggregate)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", err.Error()), http.StatusInternalServerError)
			return
		}
	}
}
