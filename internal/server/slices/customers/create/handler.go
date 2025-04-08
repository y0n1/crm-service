package create

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/y0n1/crm-service/internal/models/aggregates"
	"github.com/y0n1/crm-service/internal/models/entities"
	store_pkg "github.com/y0n1/crm-service/internal/store"
)

const UrlPattern = "/customers"

func MakeHandler(ctx context.Context, store store_pkg.Storable[*aggregates.CustomerAggregate], logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var request CreateCustomerRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", err.Error()), http.StatusBadRequest)
			return
		}

		customer := entities.NewCustomer(
			request.FirstName,
			request.LastName,
			request.Role,
			request.Email,
			request.Phone,
		)

		aggregate, err := aggregates.NewCustomerAggregate(customer)
		if err != nil {
			http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", err.Error()), http.StatusBadRequest)
			return
		}

		if err := store.Create(aggregate); err != nil {
			if err == store_pkg.ErrCustomerAlreadyExists || err == store_pkg.ErrCustomerEmailAlreadyExists {
				http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", err.Error()), http.StatusConflict)
			} else {
				http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", err.Error()), http.StatusInternalServerError)
			}
			return
		}

		response := newCreateCustomerResponseFromAggregate(aggregate)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, fmt.Sprintf("{ \"error\": \"%s\" }", err.Error()), http.StatusInternalServerError)
			return
		}
	}
}
