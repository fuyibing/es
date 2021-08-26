// author: wsfuyibing <websearch@163.com>
// date: 2021-08-19

package tests

import (
	"strings"
	"testing"
	"time"

	"github.com/fuyibing/log/v2"
	"github.com/google/uuid"

	"github.com/fuyibing/es"
)

func TestConnCreate(t *testing.T) {

	ctx := log.NewContext()
	res, err := es.Conn.Create().
		Id(strings.ReplaceAll(uuid.New().String(), "-", "")).
		Index("finance").
		Type("fin").
		Body(map[string]interface{}{
			"c": 4, "m": 5, "dateline": time.Now().Format("2006-01-02"),
		}).
		Do(ctx)
	if err != nil {
		t.Errorf("create error: %v.", err)
		return
	}

	if res.Succeed() {
		t.Logf("create succeed, index=%s, id=%v.", res.DocumentIndex, res.DocumentId)
		return
	}

	t.Logf("create result: %s.", res.Result)
}

func TestConnDelete(t *testing.T) {

	ctx := log.NewContext()
	res, err := es.Conn.Delete().
		Index("finance").
		Type("fin").
		Id("example").
		Do(ctx)
	if err != nil {
		t.Errorf("del error: %v.", err)
		return
	}

	if res.Succeed() {
		t.Logf("del succeed, index=%s, id=%s.", res.DocumentIndex, res.DocumentId)
		return
	}

	t.Logf("del not found, index=%s.", res.DocumentIndex)

}

func TestConnSearch(t *testing.T) {

	ctx := log.NewContext()
	res, err := es.Conn.Search().
		// Index("logstash-uniondrug-log-2021.08.23").
		Index("finance").
		// Type("log").
		Type("fin").
		Query(
			es.NewSearchQuery().MustTerm("version.keyword", "0.95.5"),
		).
		Do(ctx)
	if err != nil {
		t.Errorf("search error: %v.", err)
		return
	}

	// if res.Found {
	// 	t.Logf("get succeed, index=%s, source=%v.", res.DocumentIndex, res.Source)
	// 	return
	// }

	t.Logf("search found: %d.", res.Hits.TotalHits)
	res.Hits.Each(func(index int, hit *es.SearchResponseHit) error {

		ptr := &PHPLog{}
		if he := hit.Unmarshal(ptr); he != nil {
			return he
		}

		t.Logf("hit: %d -> level=%s and time=%v.", index, ptr.Level, ptr.Time)
		return nil
	})

	// body, _ := json.Marshal(res)
	// t.Logf("search field, index=%s.", string(body))

}

func TestConnUpdate(t *testing.T) {

	ctx := log.NewContext()
	res, err := es.Conn.Update().
		Id("bb6ea0c3f5f24ec5b62f28bdb66813ec").
		Index("finance").
		Type("fin").
		Set(map[string]interface{}{
			"c": 11, "m": 12,
			"dateline": time.Now().Format("2006-01-02"),
			"time":     time.Now().Format("2006-01-02 15:04:05"),
		}).
		Do(ctx)
	if err != nil {
		t.Errorf("update error: %v.", err)
		return
	}

	if res.Succeed() {
		t.Logf("update succeed, index=%s, id=%v.", res.DocumentIndex, res.DocumentId)
		return
	}

	t.Logf("update result: %s.", res.Result)
}

func TestConnUpdateByQuery(t *testing.T) {

	ctx := log.NewContext()
	res, err := es.Conn.UpdateByQuery().
		Index("finance").
		Type("fin").
		Set(map[string]interface{}{
			"dateline": time.Now().Format("2006-01-02"),
			"time":     time.Now().Format("2006-01-02 15:04:05"),
		}).
		Where(
			es.NewSearchQuery().MustTerm("order_no.keyword", "20f7443b-890e-4573-af7e-e7d14198b8d7"),
		).
		Do(ctx)
	if err != nil {
		t.Errorf("update by error: %v.", err)
		return
	}

	// if res.Succeed() {
	// 	t.Logf("update succeed, index=%s, id=%v.", res.DocumentIndex, res.DocumentId)
	// 	return
	// }

	t.Logf("result: updated=%d, deleted=%d, total=%d.", res.Updated, res.Deleted, res.Total)
}

type PHPLog struct {
	TraceId string `json:"traceId"`
	Level   string `json:"level"`
	Time    es.Timeline
	// `json:"time"`
}
