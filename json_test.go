package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestJSON(t *testing.T) {
	buf, err := ioutil.ReadFile("plugins.json")
	if err != nil {
		fmt.Println(err)
		t.Fail()

	}
	err = isJSON(buf)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}

func isJSON(buf []byte) error {
	var temp interface{}
	return json.Unmarshal(buf, &temp)
}
