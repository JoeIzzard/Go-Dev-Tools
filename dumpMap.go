package godevtools

import (
	"encoding/json"
	"fmt"
)

func DumpMap(m interface{}) {
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
}
