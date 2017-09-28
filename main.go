package main

import(
	"fmt"
)

func main() {
	fmt.Print(GetFeed().FeedItems[0].Description())
}
