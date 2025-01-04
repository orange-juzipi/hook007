package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func PrettyJSON(src []byte) string {
	prettyJSON := &bytes.Buffer{}
	err := json.Indent(prettyJSON, src, "", "\t")
	if err != nil {
		return err.Error()
	}
	return prettyJSON.String()
}

func JsonParser(param any) {
	body, _ := json.MarshalIndent(param, " ", "\t")
	fmt.Println(string(body))
}
