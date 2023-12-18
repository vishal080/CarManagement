package datafetch

import (
	"encoding/json"
	"fmt"
	"net/http"
)
func FetchCarData() ([]byte, error) {
	resp, err := http.Get("http://localhost:8080/getCars")  
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var carData []byte
	if err := json.NewDecoder(resp.Body).Decode(&carData); err != nil {
		return nil, err
	}

	return carData, nil
}
