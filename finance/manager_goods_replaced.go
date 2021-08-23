// author: wsfuyibing <websearch@163.com>
// date: 2021-08-23

package finance

import (
	"context"
	"fmt"

	"es"
)

type GoodsReplacedManager struct {
	m int
}

func NewGoodsReplacedManager(m int) *GoodsReplacedManager {
	return &GoodsReplacedManager{m: m}
}

func (o *GoodsReplacedManager) Create(ctx context.Context, m int, orderNo string, body map[string]interface{}) (*es.CreateResponse, error) {
	return es.Conn.Create().
		Index(DocumentIndex).
		Type(DocumentType).
		Id(o.GenerateId(m, orderNo)).
		Body(body).
		Do(ctx)
}

func (o *GoodsReplacedManager) Delete(ctx context.Context, m int, orderNo string) (*es.DeleteResponse, error) {
	return es.Conn.Delete().Index(DocumentIndex).
		Type(DocumentType).
		Id(o.GenerateId(m, orderNo)).
		Do(ctx)
}

func (o *GoodsReplacedManager) Update(ctx context.Context, m int, orderNo string, body map[string]interface{}) (*es.UpdateResponse, error) {
	return es.Conn.Update().
		Index(DocumentIndex).
		Type(DocumentType).
		Id(o.GenerateId(m, orderNo)).
		Set(body).
		Do(ctx)
}

// 生成唯一ID.
//
// 参数m为结算模型编号, orderNo为订单号.
// 例如: 12:101:38978501845.
func (o *GoodsReplacedManager) GenerateId(m int, orderNo string) string {
	return fmt.Sprintf("%d:%d:%s", CatalogGoodsReplaced, m, orderNo)
}
