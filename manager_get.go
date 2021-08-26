// author: wsfuyibing <websearch@163.com>
// date: 2021-08-19

package es

import (
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type GetManager struct {
	err error
	req esapi.GetRequest
}

func NewGetManager() *GetManager {
	return &GetManager{
		req: esapi.GetRequest{},
	}
}

func (o *GetManager) Do(ctx context.Context) (*GetResponse, error) {
	if err := o.Validate(); err != nil {
		return nil, err
	}

	body, err := Conn.Send(ctx, o.req)
	if err != nil {
		return nil, err
	}

	res := &GetResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (o *GetManager) Id(documentId string) *GetManager {
	o.req.DocumentID = documentId
	return o
}

func (o *GetManager) Index(documentIndex string) *GetManager {
	o.req.Index = documentIndex
	return o
}

func (o *GetManager) Type(documentType string) *GetManager {
	o.req.DocumentType = documentType
	return o
}

func (o *GetManager) Validate() error {
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

	return nil
}

type GetResponse struct {
	DocumentId    string          `json:"_id"`
	DocumentIndex string          `json:"_index"`
	DocumentType  string          `json:"_type"`
	Version       int             `json:"_version"`
	Found         bool            `json:"found"`
	Source        json.RawMessage `json:"_source"`
}
