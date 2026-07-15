package main

import (
	"log"

	"portal-news-blog/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
