package list

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/y0n1/crm-service/internal/models/aggregates"
	"github.com/y0n1/crm-service/internal/models/dtos"
	store_pkg "github.com/y0n1/crm-service/internal/store"
)

const UrlPattern = "/customers"

func MakeHandler(ctx context.Context, store store_pkg.Storable[*aggregates.CustomerAggregate], logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming not supported by this connection", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Transfer-Encoding", "chunked")
		w.Write([]byte("["))
		flusher.Flush()
		
		firstItem := true
		for itemAggregate := range store.List() {
			if !firstItem {
				w.Write([]byte(","))
			} else {
				firstItem = false
			}

			item := dtos.NewCustomerDtoFromAggregate(itemAggregate)
			itemBytes, err := json.Marshal(item)
			if err != nil {
				logger.Log(ctx, slog.LevelWarn, fmt.Sprintf("failed to marshal item: %s", string(itemBytes)))
				continue
			}

			w.Write(itemBytes)
			flusher.Flush()
		}

		w.Write([]byte("]"))
		flusher.Flush()
	}
}
