package wikipediaweblayer

import (
	"fmt"
	"net/http"
)

func GetPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("Error getting page: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Error getting page: %v", resp.Status)
	}
	return resp.Status, nil
}
