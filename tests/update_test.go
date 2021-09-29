// author: wsfuyibing <websearch@163.com>
// date: 2021-08-23

package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"testing"

	"github.com/fuyibing/es"
	"github.com/fuyibing/log/v2"
)

// import (
// 	"testing"
//
// 	"github.com/fuyibing/log/v2"
// 	"github.com/google/uuid"
//
// 	"github.com/fuyibing/es"
// )
//
func TestUpdateById(t *testing.T) {
	ctx := log.NewContext()

	docIndex, docType, docId := "finance", "fin", "2101-DS20210902100023"

	wait := new(sync.WaitGroup)
	wait.Add(4)
	go func() {
		defer wait.Done()
		testUpdateByQuery(ctx, docIndex, docType, docId, map[string]interface{}{"left:1":21})
	}()

	go func() {
		defer wait.Done()
		testUpdateByQuery(ctx, docIndex, docType, docId, map[string]interface{}{"left:2":22})
	}()

	go func() {
		defer wait.Done()
		testUpdateByQuery(ctx, docIndex, docType, docId, map[string]interface{}{"right:1":23})
	}()

	go func() {
		defer wait.Done()
		testUpdateByQuery(ctx, docIndex, docType, docId, map[string]interface{}{"right:2":24})
	}()

	wait.Wait()
}

func testUpdateByQuery(ctx context.Context, docIndex, docType, docId string, data map[string]interface{}) {
	res, err := es.Conn.Update().
		Index(docIndex).
		Type(docType).
		Id(docId).
		Set(data).
		Do(ctx)
	if err != nil {
		fmt.Printf("update err: %v.\n", err)
		return
	}

	body, _ := json.Marshal(res)
	fmt.Printf("update res: %s.\n", body)
}
