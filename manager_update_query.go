// author: wsfuyibing <websearch@163.com>
// date: 2021-08-19

package es

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type UpdateByQueryManager struct {
	body map[string]interface{}
	err  error
	req  esapi.UpdateByQueryRequest
}

func NewUpdateByQueryManager() *UpdateByQueryManager {
	bt := true
	return &UpdateByQueryManager{
		body: make(map[string]interface{}),
		req: esapi.UpdateByQueryRequest{
			AllowNoIndices:    &bt,
			IgnoreUnavailable: &bt,
			Refresh:           &bt,
		},
	}
}

func (o *UpdateByQueryManager) Do(ctx context.Context) (*UpdateByQueryResponse, error) {

	if err := o.Validate(); err != nil {
		return nil, err
	}

	buf, err := json.Marshal(o.body)
	if err != nil {
		return nil, err
	}

	o.req.Body = bytes.NewBuffer(buf)

	var body []byte
	if body, err = Conn.Send(ctx, o.req); err != nil {
		return nil, err
	}

	res := &UpdateByQueryResponse{}
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
func (o *UpdateByQueryManager) Set(data map[string]interface{}) *UpdateByQueryManager {
	if data == nil {
		o.err = ErrorNoDocumentScript
		return o
	}

	n := 0
	source, comma := "", ""
	params := map[string]interface{}{}
	for k, v := range data {
		source += comma + fmt.Sprintf("ctx._source.%s=params.%s", k, k)
		params[k] = v
		comma = ";"
		n++
	}
	if n == 0 {
		o.err = ErrorNoDocumentScript
		return o
	}

	o.body["script"] = map[string]interface{}{
		"source": source,
		"params": params,
	}

	return o
}

// 按条件设置.
//
//   .Where(condition)
func (o *UpdateByQueryManager) Where(query *SearchCondition) *UpdateByQueryManager {
	if query != nil {
		o.body["query"] = query.Generated()
	}
	return o
}

func (o *UpdateByQueryManager) Index(documentIndex ...string) *UpdateByQueryManager {
	o.req.Index = documentIndex
	return o
}

func (o *UpdateByQueryManager) Type(documentType ...string) *UpdateByQueryManager {
	o.req.DocumentType = documentType
	return o
}

func (o *UpdateByQueryManager) Validate() error {
	if o.err != nil {
		return o.err
	}

	if o.req.Index == nil || len(o.req.Index) < 1 {
		return ErrorNoDocumentIndex
	}

	if o.req.DocumentType == nil || len(o.req.DocumentType) < 1 {
		return ErrorNoDocumentType
	}

	if _, ok := o.body["query"]; !ok {
		return ErrorNoDocumentQuery
	}

	if _, ok := o.body["script"]; !ok {
		return ErrorNoDocumentScript
	}

	return nil
}

// 修改结果.
//
//   {
//       "took":4740,
//       "timed_out":false,
//       "total":5,
//       "updated":5,
//       "deleted":0,
//       "batches":1,
//       "version_conflicts":0,
//       "noops":0,
//       "retries":{"bulk":0,"search":0},
//       "throttled_millis":0,
//       "requests_per_second":-1.0,
//       "throttled_until_millis":0,
//       "failures":[]
//   }
type UpdateByQueryResponse struct {
	Total   int `json:"total"`
	Updated int `json:"updated"`
	Deleted int `json:"deleted"`
}
