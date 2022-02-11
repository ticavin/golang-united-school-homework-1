package main

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func main() {
	ss := emoji.Sprint("Hello :world_map:!")

	fmt.Println(ss)
}
