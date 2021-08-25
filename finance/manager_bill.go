// author: wsfuyibing <websearch@163.com>
// date: 2021-08-23

package finance

import (
	"context"
	"fmt"

	"github.com/fuyibing/es"
)

// 开票单管理.
type BillManager struct {
}

func NewBillManager() *BillManager {
	return &BillManager{}
}

func (o *BillManager) Create(ctx context.Context, billNo string, body map[string]interface{}) (*es.CreateResponse, error) {
	return es.Conn.Create().
		Index(DocumentIndex).
		Type(DocumentType).
		Id(o.GenerateId(billNo)).
		Body(body).
		Do(ctx)
}

func (o *BillManager) Delete(ctx context.Context, billNo string) (*es.DeleteResponse, error) {
	return es.Conn.Delete().Index(DocumentIndex).
		Type(DocumentType).
		Id(o.GenerateId(billNo)).
		Do(ctx)
}

func (o *BillManager) Update(ctx context.Context, billNo string, body map[string]interface{}) (*es.UpdateResponse, error) {
	return es.Conn.Update().
		Index(DocumentIndex).
		Type(DocumentType).
		Id(o.GenerateId(billNo)).
		Set(body).
		Do(ctx)
}

// 生成开票单ID.
//
// billNo为开票单单号, 接受整型
// 或字符串格式入参.
//
// 格式: 类别:开票单单号
// 例如: 30:20210123198
func (o *BillManager) GenerateId(billNo string) string {
	return fmt.Sprintf("%d:%s", CatalogBill, billNo)
}
