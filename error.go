// author: wsfuyibing <websearch@163.com>
// date: 2021-08-20

package es

import (
	"fmt"
)

var (
	ErrorDocumentNotFound = fmt.Errorf("document not found")
	ErrorNoDocumentBody   = fmt.Errorf("document body not specified")
	ErrorNoDocumentId     = fmt.Errorf("document id not specified")
	ErrorNoDocumentIndex  = fmt.Errorf("document index name not specified")
	ErrorNoDocumentQuery  = fmt.Errorf("document query not specified")
	ErrorNoDocumentScript = fmt.Errorf("document script not specified")
	ErrorNoDocumentType   = fmt.Errorf("document type not specified")
)
