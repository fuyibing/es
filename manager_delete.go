// author: wsfuyibing <websearch@163.com>
// date: 2021-08-19

package es

import (
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type DeleteManager struct {
	req esapi.DeleteRequest
}

func NewDeleteManager() *DeleteManager {
	return &DeleteManager{
		req: esapi.DeleteRequest{},
	}
}

func (o *DeleteManager) Do(ctx context.Context) (*DeleteResponse, error) {
	if err := o.Validate(); err != nil {
		return nil, err
	}

	body, err := Conn.Send(ctx, o.req)
	if err != nil {
		return nil, err
	}

	res := &DeleteResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (o *DeleteManager) Index(documentIndex string) *DeleteManager {
	o.req.Index = documentIndex
	return o
}

func (o *DeleteManager) Id(documentId string) *DeleteManager {
	o.req.DocumentID = documentId
	return o
}

func (o *DeleteManager) Type(documentType string) *DeleteManager {
	o.req.DocumentType = documentType
	return o
}

func (o *DeleteManager) Validate() error {
	if o.req.DocumentID == "" {
		return ErrorNoDocumentId
	}

	if o.req.Index == "" {
		return ErrorNoDocumentIndex
	}

	if o.req.DocumentType == "" {
		return ErrorNoDocumentType
	}

	o.req.Refresh = "true"
	return nil
}

type DeleteResponse struct {
	DocumentId    string `json:"_id"`
	DocumentIndex string `json:"_index"`
	DocumentType  string `json:"_type"`
	Version       int    `json:"_version"`
	Result        string `json:"result"`
}

func (o *DeleteResponse) Succeed() bool {
	return o.Result == "deleted"
}
