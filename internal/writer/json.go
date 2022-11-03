package writer

import (
	"encoding/json"
	"fuzzygolang/internal/util"
)

type Result struct {
	N1          int    `json:"n1"`
	N2          int    `json:"n2"`
	Operation   string `json:"operation"`
	FinalResult int    `json:"final_result"`
}

func WriteJSON(n1, n2 int, operation string, result int) string {
	j, err := json.Marshal(Result{
		N1:          n1,
		N2:          n2,
		Operation:   operation,
		FinalResult: result,
	})

	util.HandleError(err)

	return string(j)
}
