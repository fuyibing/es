// author: wsfuyibing <websearch@163.com>
// date: 2021-08-20

package tests

// import (
// 	"testing"
//
// 	"github.com/fuyibing/log/v2"
//
// 	"github.com/fuyibing/es"
// )
//
// func TestMapping(t *testing.T) {
//
// 	x := es.Conn.Mapping()
// 	x.Index("finance3")
// 	x.Type("fin")
//
// 	x.AsInteger("_catalog", "_deleted")
// 	x.AsBool("as_bool")
// 	x.AsByte("as_byte")
// 	x.AsDate("as_date")
// 	x.AsDouble("as_double")
// 	x.AsFloat("as_float")
// 	x.AsInteger("as_integer")
// 	x.AsKeyword("as_keyword")
// 	x.AsLong("as_long")
// 	x.AsShort("as_short")
// 	x.AsText("as_text")
//
// 	ctx := log.NewContext()
//
// 	r, err := x.Do(ctx)
//
// 	if err != nil {
// 		t.Errorf("mapping error: %v.", err)
// 		return
// 	}
//
// 	t.Logf("mapping field: %v.", r)
// }
//
// func TestMapping2(t *testing.T) {
//
// 	x := es.Conn.Mapping()
// 	x.Index("finance")
// 	x.Type("fin")
//
// 	x.AsInteger("_catalog", "_deleted")
// 	// x.AsBool("as_bool")
// 	// x.AsByte("as_byte")
// 	// x.AsDate("as_date")
// 	// x.AsDouble("as_double")
// 	// x.AsFloat("as_float")
// 	// x.AsInteger("as_integer")
// 	// x.AsKeyword("as_keyword")
// 	// x.AsLong("as_long")
// 	// x.AsShort("as_short")
// 	// x.AsText("as_text")
//
// 	ctx := log.NewContext()
//
// 	r, err := x.Do(ctx)
//
// 	if err != nil {
// 		t.Errorf("mapping error: %v.", err)
// 		return
// 	}
//
// 	t.Logf("mapping field: %v.", r)
// }
