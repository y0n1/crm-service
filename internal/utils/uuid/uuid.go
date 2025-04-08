package uuid

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	uuid_pkg "github.com/google/uuid"
	"github.com/y0n1/crm-service/pkg/collections"
)

const (
	ByteLength = 16
	Pattern = "[0-9a-f]{8}(?:\\-[0-9a-f]{4}){3}-[0-9a-f]{12}"
)

func ParseFromBody(r *http.Request) (uuid_pkg.UUID, error) {
	bodyBytes := collections.NewList[byte](ByteLength)
	if _, err := r.Body.Read(bodyBytes); err != nil {
		return uuid_pkg.Nil, errors.New("failed to read body")
	}
	if id, err := uuid_pkg.ParseBytes(bodyBytes); err != nil {
		return uuid_pkg.Nil, err
	} else {
		return id, nil
	}
}

func ParseFromUrlParam(r *http.Request, param string, required bool) (uuid_pkg.UUID, error) {
	p := chi.URLParam(r, param)
	if p == "" {
		if required {
			return uuid_pkg.Nil, fmt.Errorf("missing URLParam: %s", param)
		} else {
			return uuid_pkg.Nil, nil
		}
	}

	uuid, err := uuid_pkg.Parse(p)
	if err != nil {
		return uuid_pkg.Nil, err
	}

	return uuid, nil
}
