package wikipediacrawl

import (
	"dataMiningGolang/weblayer/wikipediaweblayer"
	"fmt"
)

func Handle() string {

	response, err := wikipediaweblayer.GetPage("https://wikipedia.org")

	if err != nil {
		return "Error"
	}

	return fmt.Sprintf("Wikipedia: %s", response)

}
