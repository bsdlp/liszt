package apiutils

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestErrorJSON(t *testing.T) {
	i := ErrNotFound
	bs, err := json.Marshal(i)
	if err != nil {
		t.Error(err)
	}

	var errObj ErrorObject
	err = json.Unmarshal(bs, &errObj)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(i, errObj) {
		t.Fail()
	}
}
