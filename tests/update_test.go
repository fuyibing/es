// author: wsfuyibing <websearch@163.com>
// date: 2021-08-23

package tests

// import (
// 	"testing"
//
// 	"github.com/fuyibing/log/v2"
// 	"github.com/google/uuid"
//
// 	"github.com/fuyibing/es"
// )
//
// func TestUpdateById(t *testing.T) {
// 	ctx := log.NewContext()
//
// 	docIndex, docType, docId := "finance", "fin", "34184b9c487d47b28f5bb87a94e472a4"
//
// 	res, err := es.Conn.Update().
// 		Index(docIndex).
// 		Type(docType).
// 		Id(docId).
// 		Set(map[string]interface{}{
// 			"key": "value",
// 		}).
// 		Do(ctx)
// 	if err != nil {
// 		t.Errorf("update by id error, error=%v.", err)
// 		return
// 	}
//
// 	t.Logf("update by id succeed, id=%s.", res.DocumentId)
// }
//
// func TestUpdateByQuery(t *testing.T) {
// 	ctx := log.NewContext()
//
// 	docIndex, docType := "finance", "fin"
// 	res, err := es.Conn.UpdateByQuery().
// 		Index(docIndex).
// 		Type(docType).
// 		Set(map[string]interface{}{
// 			"key": uuid.New().String(),
// 		}).
// 		Where(es.NewSearchQuery().NotExists("key")).
// 		Do(ctx)
// 	if err != nil {
// 		t.Errorf("update by query error, error=%v.", err)
// 		return
// 	}
//
// 	t.Logf("update by query succeed, total=%d, updated=%d.", res.Total, res.Updated)
// }
