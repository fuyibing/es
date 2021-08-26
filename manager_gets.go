// author: wsfuyibing <websearch@163.com>
// date: 2021-08-19

package es

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type GetsManager struct {
	err error
	ids []string
	req esapi.MgetRequest
}

// 读取管理.
func NewGetsManager() *GetsManager {
	return &GetsManager{
		req: esapi.MgetRequest{},
	}
}

func (o *GetsManager) Do(ctx context.Context) (*GetsResponse, error) {
	if err := o.Validate(); err != nil {
		return nil, err
	}

	body, err := Conn.Send(ctx, o.req)
	if err != nil {
		return nil, err
	}

	res := &GetsResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Specify index name.
func (o *GetsManager) Index(documentIndex string) *GetsManager {
	o.req.Index = documentIndex
	return o
}

// Specify document id.
func (o *GetsManager) Id(ids ...string) *GetsManager {
	o.ids = ids
	return o
}

// Specify document type.
func (o *GetsManager) Type(documentType string) *GetsManager {
	o.req.DocumentType = documentType
	return o
}

// Validate struct.
func (o *GetsManager) Validate() error {

	if o.req.Index == "" {
		return ErrorNoDocumentIndex
	}

	if o.req.DocumentType == "" {
		return ErrorNoDocumentType
	}

	docs := make([]map[string]interface{}, 0)

	for _, id := range o.ids {
		docs = append(docs, map[string]interface{}{
			"_index": o.req.Index,
			"_type":  o.req.DocumentType,
			"_id":    id,
		})
	}

	buf, _ := json.Marshal(map[string]interface{}{
		"docs": docs,
	})

	o.req.Body = bytes.NewBuffer(buf)
	return nil
}

type GetsResponse struct {
	Docs []*GetsResponseHit `json:"docs"`
}

func (o *GetsResponse) Count() int {
	return len(o.Docs)
}

func (o *GetsResponse) Each(fn func(index int, res *GetsResponseHit) error) {
	for i, Gets := range o.Docs {
		if fn(i, Gets) != nil {
			break
		}
	}
}

// 读取结果.
//
// ES成功读到数据时返回结果.
type GetsResponseHit struct {
	DocumentId    string          `json:"_id"`
	DocumentIndex string          `json:"_index"`
	DocumentType  string          `json:"_type"`
	Version       int             `json:"_version"`
	Found         bool            `json:"found"`
	Source        json.RawMessage `json:"_source"`
}

func (o *GetsResponseHit) Unmarshal(ptr interface{}) error {
	if o.Found {
		return json.Unmarshal(o.Source, ptr)
	}
	return ErrorDocumentNotFound
}
