// author: wsfuyibing <websearch@163.com>
// date: 2021-08-19

package es

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type CreateManager struct {
	err error
	req esapi.CreateRequest
}

func NewCreateManager() *CreateManager {
	return &CreateManager{
		req: esapi.CreateRequest{},
	}
}

func (o *CreateManager) Do(ctx context.Context) (*CreateResponse, error) {
	if err := o.Validate(); err != nil {
		return nil, err
	}

	body, err := Conn.Send(ctx, o.req)
	if err != nil {
		return nil, err
	}

	res := &CreateResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (o *CreateManager) Body(data map[string]interface{}) *CreateManager {
	buf, err := json.Marshal(data)
	if err != nil {
		o.err = err
	} else {
		o.req.Body = bytes.NewBuffer(buf)
	}
	return o
}

func (o *CreateManager) Id(documentId string) *CreateManager {
	o.req.DocumentID = documentId
	return o
}

func (o *CreateManager) Index(documentIndex string) *CreateManager {
	o.req.Index = documentIndex
	return o
}

func (o *CreateManager) Type(documentType string) *CreateManager {
	o.req.DocumentType = documentType
	return o
}

func (o *CreateManager) Validate() error {
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

type CreateResponse struct {
	DocumentId    string `json:"_id"`
	DocumentIndex string `json:"_index"`
	DocumentType  string `json:"_type"`
	Version       int    `json:"_version"`
	Result        string `json:"result"`
}

func (o *CreateResponse) Succeed() bool {
	return o.Result == "created"
}
