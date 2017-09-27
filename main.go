package main

import(
	"fmt"
)

func main() {
	a := FeedItem{DescriptionHTML: "<font face='Helvetica Neue'><p align=\"center\" style=\"text-align: center;\"><b>Science Building closed<\\/b><br>A flood in the Science Building on the St. John's campus has closed the building until 12 today. An update will be provided by 11 a.m.<br>Sent: 13 hours ago<br><\\/font>"}
	fmt.Print(a.Description())
}
