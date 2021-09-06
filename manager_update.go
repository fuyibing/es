// author: wsfuyibing <websearch@163.com>
// date: 2021-08-19

package es

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type UpdateManager struct {
	err error
	req esapi.UpdateRequest
}

func NewUpdateManager() *UpdateManager {
	return &UpdateManager{}
}

func (o *UpdateManager) Do(ctx context.Context) (*UpdateResponse, error) {
	if err := o.Validate(); err != nil {
		return nil, err
	}

	body, err := Conn.Send(ctx, o.req)
	if err != nil {
		return nil, err
	}

	res := &UpdateResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}

// 设置新数据.
//
//   .Set(map[string]interface{}{
//      "k": 1,
//      "key": "value",
//   })
func (o *UpdateManager) Set(data map[string]interface{}) *UpdateManager {
	buf, err := json.Marshal(map[string]interface{}{
		"doc": data,
	})

	// 2. 字段错误.
	if err != nil {
		o.err = err
		return o
	}

	// 3. 设置字段.
	o.req.Body = bytes.NewBuffer(buf)
	return o
}

func (o *UpdateManager) Index(documentIndex string) *UpdateManager {
	o.req.Index = documentIndex
	return o
}

func (o *UpdateManager) Id(documentId string) *UpdateManager {
	o.req.DocumentID = documentId
	return o
}

func (o *UpdateManager) Type(documentType string) *UpdateManager {
	o.req.DocumentType = documentType
	return o
}

func (o *UpdateManager) Validate() error {
	if o.err != nil {
		return o.err
	}

	if o.req.DocumentID == "" {
		return ErrorNoDocumentId
	}

	if o.req.Index == "" {
		return ErrorNoDocumentIndex
	}

	if o.req.DocumentType == "" {
		return ErrorNoDocumentType
	}

	if o.req.Body == nil {
		return ErrorNoDocumentBody
	}

	o.req.Refresh = "true"
	return nil
}

type UpdateResponse struct {
	DocumentId    string `json:"_id"`
	DocumentIndex string `json:"_index"`
	DocumentType  string `json:"_type"`
	Version       int    `json:"_version"`
	Result        string `json:"result"`
}

func (o *UpdateResponse) Succeed() bool {
	return o.Result == "updated"
}
