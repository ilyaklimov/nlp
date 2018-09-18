package tokenize

import (
	"io/ioutil"
	"encoding/json"
	"reflect"
	"testing"
)

type testDataTerm struct {
	Name 	string 		`json:"name"`
	Input 	string 		`json:"input"`
	Output 	[]string 	`json:"output"`
}

func TestToUniterms(t *testing.T) {
	var unitermTests []testDataTerm
	bs, _ := ioutil.ReadFile("./testdata/uniterms.json")
	_ = json.Unmarshal(bs, &unitermTests)
	for _, test := range unitermTests {
		actual, _ := ToUniterms(test.Input)
		if !reflect.DeepEqual(actual, test.Output) {
			t.Errorf("cannot get uniterms (\"%s\"):\ntest:\t%#v\ngot:\t%#v\nwant:\t%#v\n\n", test.Name, test.Input, actual, test.Output)		
		}
	}
}