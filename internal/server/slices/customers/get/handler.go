package get

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/y0n1/crm-service/internal/models/aggregates"
	"github.com/y0n1/crm-service/internal/models/dtos"
	store_pkg "github.com/y0n1/crm-service/internal/store"
	"github.com/y0n1/crm-service/internal/utils/uuid"
)

const UrlPattern = "/customers/{id}"

func MakeHandler(ctx context.Context, store store_pkg.Storable[*aggregates.CustomerAggregate], logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		uuid, err := uuid.ParseFromUrlParam(r, "id", false)
		if err != nil {
			http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", err.Error()), http.StatusBadRequest)
			return
		}

		aggregate, err := store.Get(uuid)
		if err != nil {
			if err == store_pkg.ErrCustomerNotFound {
				http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", err.Error()), http.StatusNotFound)
			} else {
				http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", err.Error()), http.StatusBadRequest)
			}
			return
		}

		response := dtos.NewCustomerDtoFromAggregate(aggregate)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", err.Error()), http.StatusInternalServerError)
			return
		}
	}
}
