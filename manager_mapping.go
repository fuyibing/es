// author: wsfuyibing <websearch@163.com>
// date: 2021-08-19

package es

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type MappingManager struct {
	body map[string]interface{}
	req  esapi.IndicesPutMappingRequest
}

// 映射管理.
//
// todo: mapping manager.
func NewMappingManager() *MappingManager {
	t := true
	// f := true
	return &MappingManager{
		body: make(map[string]interface{}),
		req: esapi.IndicesPutMappingRequest{
			AllowNoIndices:    &t,
			IgnoreUnavailable: &t,
			IncludeTypeName:   &t,
		},
	}
}

// /////////////////////////////
// 浮点型                      //
// /////////////////////////////

func (o *MappingManager) AsDouble(ks ...string) *MappingManager { return o.UseSimple("double", ks...) }
func (o *MappingManager) AsFloat(ks ...string) *MappingManager  { return o.UseSimple("float", ks...) }

// /////////////////////////////
// 整型                        //
// /////////////////////////////

func (o *MappingManager) AsByte(ks ...string) *MappingManager    { return o.UseSimple("byte", ks...) }
func (o *MappingManager) AsInteger(ks ...string) *MappingManager { return o.UseSimple("integer", ks...) }
func (o *MappingManager) AsLong(ks ...string) *MappingManager    { return o.UseSimple("long", ks...) }
func (o *MappingManager) AsShort(ks ...string) *MappingManager   { return o.UseSimple("short", ks...) }

// /////////////////////////////
// 文本                        //
// /////////////////////////////

// 关键词.
//
// 不支持分词, 查询时必须全词匹配.
func (o *MappingManager) AsKeyword(ks ...string) *MappingManager {
	for _, k := range ks {
		o.body[k] = map[string]interface{}{
			"type": "text",
			"fields": map[string]map[string]interface{}{
				"keyword": {
					"type":         "keyword",
					"ignore_above": 256,
				},
			},
		}
	}
	return o
}

// 文本.
//
// 全文检索, 支持分词(如: IK).
func (o *MappingManager) AsText(ks ...string) *MappingManager {
	for _, k := range ks {
		o.body[k] = map[string]interface{}{
			"type": "text",
		}
	}
	return o
}

// 布尔.
func (o *MappingManager) AsBool(ks ...string) *MappingManager { return o.UseSimple("boolean", ks...) }

// 日期.
//
// 接受: 2006-01-02, 2006-01-02 15:04:05.
func (o *MappingManager) AsDate(ks ...string) *MappingManager {
	for _, k := range ks {
		o.body[k] = map[string]interface{}{
			"type":   "date",
			"format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis",
		}
	}
	return o
}

// /////////////////////////////
// 操作                        //
// /////////////////////////////

func (o *MappingManager) Do(ctx context.Context) (*MappingResponse, error) {
	var err error

	if err = o.Validate(); err != nil {
		return nil, err
	}

	var buf []byte
	if buf, err = json.Marshal(map[string]interface{}{
		"properties": o.body,
	}); err != nil {
		return nil, err
	}

	o.req.Body = bytes.NewBuffer(buf)

	var body []byte
	if body, err = Conn.Send(ctx, o.req); err != nil {
		return nil, err
	}

	res := &MappingResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (o *MappingManager) Index(documentIndex ...string) *MappingManager {
	o.req.Index = documentIndex
	return o
}

func (o *MappingManager) Type(documentType string) *MappingManager {
	o.req.DocumentType = documentType
	return o
}

// 简易模式.
//
// 整型: byte, short, integer, long.
// 浮点: double, float.
// 布尔: bool.
func (o *MappingManager) UseSimple(t string, ks ...string) *MappingManager {
	for _, k := range ks {
		o.body[k] = map[string]string{"type": t}
	}
	return o
}

func (o *MappingManager) Validate() error {
	if len(o.req.Index) == 0 {
		return ErrorNoDocumentIndex
	}

	if o.req.DocumentType == "" {
		return ErrorNoDocumentType
	}

	return nil
}

type MappingResponse struct {
	Ack bool `json:"acknowledged"`
}
