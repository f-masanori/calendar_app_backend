package httputil

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func ParseID(r *http.Request) (int64, error) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		return 0, errors.New("required parameter is missing")
	}
	return strconv.ParseInt(id, 10, 64)
}
