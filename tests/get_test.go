// author: wsfuyibing <websearch@163.com>
// date: 2021-08-23

package tests

// import (
// 	"encoding/json"
// 	"testing"
//
// 	"github.com/fuyibing/log/v2"
//
// 	"github.com/fuyibing/es"
// )
//
// func TestMget(t *testing.T) {
// 	ctx := log.NewContext()
//
// 	res, err := es.Conn.Gets().
// 		Index("finance").
// 		Type("fins").
// 		Id("bb6ea0c3f5f24ec5b62f28bdb66813ec", "d184090d449649049fbd256bd57089f8").
// 		Do(ctx)
//
// 	if err != nil {
// 		t.Errorf("get by id ... error, error=%v.", err)
// 		return
// 	}
//
// 	t.Logf("get by id count: %d.", res.Count())
//
//
//
// 	res.Each(func(index int, res *es.GetsResponseHit) error {
// 		t.Logf("      found: %v, id: %v.", res.Found, res.Source)
// 		return nil
// 	})
//
// }
//
// func TestGet(t *testing.T) {
// 	ctx := log.NewContext()
//
// 	res, err := es.Conn.Get().
// 		Index("finance").
// 		Type("fin").
// 		Id("4aa6196f05be431f948e73241b018962").
// 		Do(ctx)
//
// 	if err != nil {
// 		t.Errorf("get by id ... error, error=%v.", err)
// 		return
// 	}
//
// 	body, _ := json.Marshal(res)
//
// 	t.Logf("get by id: %s.", string(body))
// }
