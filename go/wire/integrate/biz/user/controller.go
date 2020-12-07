package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service: service}
}

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(r.FormValue("user_id"), 10, 32)

	u, err := c.service.Get(r.Context(), userId)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		json.NewEncoder(w).Encode(u)
	} else {
		fmt.Printf("%v\n", err)
		http.NotFound(w, r)
	}
}
