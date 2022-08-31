package web

import (
	"net/http"
	"strconv"
	"time"
)

func ParseRequestIntParam(r *http.Request, paramname string, val *int) (bool, error) {
	has := r.Form.Has(paramname)
	requestVal := r.Form.Get(paramname)

	intVal, err := strconv.Atoi(requestVal)
	if err != nil {
		return has, err
	}

	*val = intVal
	return has, nil
}

func ParseRequestDateParam(r *http.Request, paramname string, val *time.Time) (bool, error) {
	has := r.Form.Has(paramname)
	requestVal := r.Form.Get(paramname)

	timeVal, err := time.Parse("2006-01-02", requestVal)
	if err != nil {
		return has, err
	}

	*val = timeVal

	return has, nil
}
