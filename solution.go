package main

import (
	"github.com/kyokomi/emoji/v2"
)

func main() {
	Get()
}
func Get() string {
	return emoji.Sprint("Hello :world_map:!")
}
