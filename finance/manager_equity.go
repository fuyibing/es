// author: wsfuyibing <websearch@163.com>
// date: 2021-08-23

package finance

import (
	"context"
	"fmt"

	"github.com/fuyibing/es"
)

type EquityManager struct {
	m int
}

func NewEquityManager(m int) *EquityManager {
	return &EquityManager{m: m}
}

func (o *EquityManager) Create(ctx context.Context, m int, equityNo string, body map[string]interface{}) (*es.CreateResponse, error) {
	return es.Conn.Create().
		Index(DocumentIndex).
		Type(DocumentType).
		Id(o.GenerateId(m, equityNo)).
		Body(body).
		Do(ctx)
}

func (o *EquityManager) Delete(ctx context.Context, m int, equityNo string) (*es.DeleteResponse, error) {
	return es.Conn.Delete().Index(DocumentIndex).
		Type(DocumentType).
		Id(o.GenerateId(m, equityNo)).
		Do(ctx)
}

func (o *EquityManager) Update(ctx context.Context, m int, equityNo string, body map[string]interface{}) (*es.UpdateResponse, error) {
	return es.Conn.Update().
		Index(DocumentIndex).
		Type(DocumentType).
		Id(o.GenerateId(m, equityNo)).
		Set(body).
		Do(ctx)
}

// 生成唯一ID.
//
// 参数m为结算模型编号, equityNo为权益号号码.
// 例如: 20:101:38978501845.
func (o *EquityManager) GenerateId(m int, equityNo string) string {
	return fmt.Sprintf("%d:%d:%s", CatalogEquity, m, equityNo)
}
