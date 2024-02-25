package main

import "fmt"

func main() {
	// websites := []string{"https://google.com", "https://aws.com"}
	websites := map[string]string{
		"Google": "https://google.com",
		"Amazon": "https://aws.com",
	}
	
	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)
	delete(websites, "Linkedin")
}