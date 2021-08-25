// author: wsfuyibing <websearch@163.com>
// date: 2021-08-20

package tests

import (
	"encoding/json"
	"testing"

	"github.com/fuyibing/es"
)

func TestQuery(t *testing.T) {

	sq := es.NewSearchQuery().
		// MustExists("f1-1", "f1-2", "f1-3").
		// MustFuzzy("f2", "v1", 10).
		// MustRange("m", nil, 1, nil, 10).
		SortAsc("m", "c")

	ss := sq.Generated()

	body, err := json.Marshal(ss)
	if err != nil {
		t.Errorf("marshal error: %v.", err)
		return
	}

	t.Logf("marshal field: %s.", string(body))
}
