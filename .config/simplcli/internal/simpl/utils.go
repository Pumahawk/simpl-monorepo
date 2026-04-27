package simpl

import (
	"encoding/json"
	"net/http"
)

// Generic function for check healt
func microCheck(url string) (bool, error) {
	r, err := http.Get(url + "/actuator/health")
	if err != nil {
		return false, err
	}
	defer r.Body.Close()

	rj := &healtResponseDto{}
	if err := json.NewDecoder(r.Body).Decode(rj); err != nil {
		return false, err
	}

	return rj.Status == "UP", nil
}
