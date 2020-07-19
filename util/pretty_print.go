package util

import (
	"fmt"
	"github.com/tidwall/pretty"
)

func PrettyPrint(data []byte) {
	result := pretty.Pretty(data)
	fmt.Print(string(result))
}
