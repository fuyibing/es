// author: wsfuyibing <websearch@163.com>
// date: 2021-08-20

package tests

import (
	"testing"

	"github.com/fuyibing/log/v2"
	"github.com/olivere/elastic"
)

func TestUtil(t *testing.T) {

	ctx := log.NewContext()

	log.Info("sfing")
	log.Infof("sfing %d", 1)
	log.Infofc(ctx, "sfing %d", 1)

	(&elastic.Client{}).Mget()
	// Search().
	// 	Set("").
	// 	Where(nil).
	// 	Set(nil)
}
