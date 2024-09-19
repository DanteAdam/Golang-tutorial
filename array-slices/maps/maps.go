package maps

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites["Google"])
	websites["Youtube"] = "https//youtube.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)

}
