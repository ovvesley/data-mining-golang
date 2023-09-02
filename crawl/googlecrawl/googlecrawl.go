package googlecrawl

import (
	"dataMiningGolang/weblayer/googleweblayer"
	"fmt"
)

func Handle() string {
	response, err := googleweblayer.GetPage("https://google.com")

	if err != nil {
		return "Error"
	}

	return fmt.Sprintf("Google: %s", response)
}
