package maps

import (
	"fmt"
)

func main() {
	link := map[string]string{
		"google": "https://www.google.com",
		"amazon": "https://www.amazon.com",
	}
	fmt.Println(link)
	fmt.Println(link["google"])

	//adding and deleting link
	link["linkedin"] = "https://www.linkedin.com"
	delete(link, "google")
	fmt.Println(link)
}
