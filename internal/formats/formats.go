package formats

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(v interface{}) {
	data, _ := json.MarshalIndent(&v, " ", "  ")
	fmt.Println(string(data))
}
