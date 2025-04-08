package delete

import (
	"context"
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

		uuid, err := uuid.ParseFromUrlParam(r, "id", true)
		if err != nil {
			http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", err.Error()), http.StatusBadRequest)
			return
		}

		if err := store.Delete(uuid); err != nil {
			if err == store_pkg.ErrCustomerNotFound {
				http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", err.Error()), http.StatusNotFound)
			} else {
				http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", err.Error()), http.StatusBadRequest)
			}
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
