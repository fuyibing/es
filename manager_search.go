// author: wsfuyibing <websearch@163.com>
// date: 2021-08-19

package es

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type (
	SearchCond map[string]map[string]interface{}
	SearchSort map[string]string
)

type SearchManager struct {
	body map[string]interface{}
	req  esapi.SearchRequest
}

func NewSearchManager() *SearchManager {
	return &SearchManager{
		body: make(map[string]interface{}),
		req:  esapi.SearchRequest{},
	}
}

func (o *SearchManager) Do(ctx context.Context) (*SearchResponse, error) {

	if err := o.Validate(); err != nil {
		return nil, err
	}

	buf, err := json.Marshal(o.body)
	if err != nil {
		return nil, err
	}

	o.req.Body = bytes.NewBuffer(buf)

	body, err := Conn.Send(ctx, o.req)
	if err != nil {
		return nil, err
	}

	println("body: ", string(body))

	res := &SearchResponse{}
	if err = json.Unmarshal(body, res); err != nil {
		return nil, err
	}

	return res, nil
}

func (o *SearchManager) Index(documentIndex ...string) *SearchManager {
	o.req.Index = documentIndex
	return o
}

func (o *SearchManager) Query(query *SearchCondition) *SearchManager {
	if query != nil {
		o.body["query"] = query.Generated()
	}
	return o
}

func (o *SearchManager) Type(documentType ...string) *SearchManager {
	o.req.DocumentType = documentType
	return o
}

func (o *SearchManager) Validate() error {

	if o.req.Index == nil || len(o.req.Index) == 0 {
		return ErrorNoDocumentIndex
	}

	if o.req.DocumentType == nil || len(o.req.DocumentType) == 0 {
		return ErrorNoDocumentType
	}

	return nil
}

// 查询条件.
type SearchCondition struct {
	must    []SearchCond
	mustNot []SearchCond
	should  []SearchCond
	sorts   []SearchSort
}

func NewSearchQuery() *SearchCondition {
	return &SearchCondition{
		must:    make([]SearchCond, 0),
		mustNot: make([]SearchCond, 0),
		should:  make([]SearchCond, 0),
		sorts:   make([]SearchSort, 0),
	}
}

// 导出结果.
//
//   {
//	     "bool" : {
//		     "must" : {
//			     "term" : { "user" : "kimchy" }
//		     },
//		     "must_not" : {
//			     "range" : {
//				     "age" : { "from" : 10, "to" : 20 }
//			     }
//		     },
//           "filter" : [
//               ...
//           ]
//		    "should" : [
//			    {
//				    "term" : { "tag" : "wow" }
//			    },
//			    {
//				    "term" : { "tag" : "elasticsearch" }
//			    }
//		     ],
//		     "minimum_should_match" : 1,
//		     "boost" : 1.0
//	     }
//   }
func (o *SearchCondition) Generated() interface{} {
	// 1. 布尔查询.
	bs := make(map[string]interface{})
	// 1.1 Must条件.
	if len(o.must) > 0 {
		bs["must"] = o.must
	}
	// 1.2 Not条件.
	if len(o.mustNot) > 0 {
		bs["must_not"] = o.mustNot
	}
	// 1.3 Should条件.
	if len(o.should) > 0 {
		bs["should"] = o.should
	}

	// 0. 结果.
	res := map[string]interface{}{}
	// 0.1 Boolean查询.
	if len(bs) > 0 {
		res["bool"] = bs
	}
	// 0.2 结果排序.
	if len(o.sorts) > 0 {
		res["sort"] = o.sorts
	}

	// n. 输出结果.
	return res
}

func (o *SearchCondition) MustExists(ks ...string) *SearchCondition {
	for _, k := range ks {
		o.must = append(o.must, SearchCond{
			"exists": {
				"field": k,
			},
		})
	}
	return o
}
func (o *SearchCondition) MustFuzzy(k string, v, me interface{}) *SearchCondition {
	o.must = append(o.must, SearchCond{"fuzzy": {k: map[string]interface{}{"value": v, "max_expansions": me}}})
	return o
}
func (o *SearchCondition) MustMatch(k string, v interface{}) *SearchCondition {
	o.must = append(o.must, SearchCond{"match": {k: v}})
	return o
}
func (o *SearchCondition) MustPrefix(k, v string) *SearchCondition {
	o.must = append(o.must, SearchCond{"prefix": {k: v}})
	return o
}
func (o *SearchCondition) MustRange(k string, gt, gte, lt, lte interface{}) *SearchCondition {
	rs := make(map[string]interface{})
	if lt != nil {
		rs["lt"] = lt
	} else if lte != nil {
		rs["lte"] = lte
	}
	if gt != nil {
		rs["gt"] = gt
	} else if gte != nil {
		rs["gte"] = gte
	}
	o.must = append(o.must, SearchCond{"range": {k: rs}})
	return o
}
func (o *SearchCondition) MustQueryString(k string, v interface{}) *SearchCondition {
	o.must = append(o.must, SearchCond{"query_string": {k: map[string]interface{}{"default_field": k, "query": v}}})
	return o
}
func (o *SearchCondition) MustTerm(k string, v interface{}) *SearchCondition {
	o.must = append(o.must, SearchCond{"term": {k: v}})
	return o
}
func (o *SearchCondition) MustText(k string, v interface{}) *SearchCondition {
	o.must = append(o.must, SearchCond{"text": {k: v}})
	return o
}
func (o *SearchCondition) MustWildcard(k string, v interface{}) *SearchCondition {
	o.must = append(o.must, SearchCond{"wildcard": {k: v}})
	return o
}

func (o *SearchCondition) NotExists(ks ...string) *SearchCondition {
	for _, k := range ks {
		o.mustNot = append(o.mustNot, SearchCond{
			"exists": {
				"field": k,
			},
		})
	}
	return o
}
func (o *SearchCondition) NotFuzzy(k string, v, me interface{}) *SearchCondition {
	o.mustNot = append(o.mustNot, SearchCond{"fuzzy": {k: map[string]interface{}{"value": v, "max_expansions": me}}})
	return o
}
func (o *SearchCondition) NotMatch(k string, v interface{}) *SearchCondition {
	o.mustNot = append(o.mustNot, SearchCond{"match": {k: v}})
	return o
}
func (o *SearchCondition) NotPrefix(k, v string) *SearchCondition {
	o.mustNot = append(o.mustNot, SearchCond{"prefix": {k: v}})
	return o
}
func (o *SearchCondition) NotRange(k string, gt, gte, lt, lte interface{}) *SearchCondition {
	rs := make(map[string]interface{})
	if lt != nil {
		rs["lt"] = lt
	} else if lte != nil {
		rs["lte"] = lte
	}
	if gt != nil {
		rs["gt"] = gt
	} else if gte != nil {
		rs["gte"] = gte
	}
	o.mustNot = append(o.mustNot, SearchCond{"range": {k: rs}})
	return o
}
func (o *SearchCondition) NotQueryString(k string, v interface{}) *SearchCondition {
	o.mustNot = append(o.mustNot, SearchCond{"query_string": {k: map[string]interface{}{"default_field": k, "query": v}}})
	return o
}
func (o *SearchCondition) NotTerm(k string, v interface{}) *SearchCondition {
	o.mustNot = append(o.mustNot, SearchCond{"term": {k: v}})
	return o
}
func (o *SearchCondition) NotText(k string, v interface{}) *SearchCondition {
	o.mustNot = append(o.mustNot, SearchCond{"text": {k: v}})
	return o
}
func (o *SearchCondition) NotWildcard(k string, v interface{}) *SearchCondition {
	o.mustNot = append(o.mustNot, SearchCond{"wildcard": {k: v}})
	return o
}

func (o *SearchCondition) ShouldExists(ks ...string) *SearchCondition {
	for _, k := range ks {
		o.should = append(o.should, SearchCond{
			"exists": {
				"field": k,
			},
		})
	}
	return o
}
func (o *SearchCondition) ShouldFuzzy(k string, v, me interface{}) *SearchCondition {
	o.should = append(o.should, SearchCond{"fuzzy": {k: map[string]interface{}{"value": v, "max_expansions": me}}})
	return o
}
func (o *SearchCondition) ShouldMatch(k string, v interface{}) *SearchCondition {
	o.should = append(o.should, SearchCond{"match": {k: v}})
	return o
}
func (o *SearchCondition) ShouldPrefix(k, v string) *SearchCondition {
	o.should = append(o.should, SearchCond{"prefix": {k: v}})
	return o
}
func (o *SearchCondition) ShouldRange(k string, gt, gte, lt, lte interface{}) *SearchCondition {
	rs := make(map[string]interface{})
	if lt != nil {
		rs["lt"] = lt
	} else if lte != nil {
		rs["lte"] = lte
	}
	if gt != nil {
		rs["gt"] = gt
	} else if gte != nil {
		rs["gte"] = gte
	}
	o.should = append(o.should, SearchCond{"range": {k: rs}})
	return o
}
func (o *SearchCondition) ShouldQueryString(k string, v interface{}) *SearchCondition {
	o.should = append(o.should, SearchCond{"query_string": {k: map[string]interface{}{"default_field": k, "query": v}}})
	return o
}
func (o *SearchCondition) ShouldTerm(k string, v interface{}) *SearchCondition {
	o.should = append(o.should, SearchCond{"term": {k: v}})
	return o
}
func (o *SearchCondition) ShouldText(k string, v interface{}) *SearchCondition {
	o.should = append(o.should, SearchCond{"text": {k: v}})
	return o
}
func (o *SearchCondition) ShouldWildcard(k string, v interface{}) *SearchCondition {
	o.should = append(o.should, SearchCond{"wildcard": {k: v}})
	return o
}

func (o *SearchCondition) SortAsc(ks ...string) *SearchCondition {
	for _, k := range ks {
		o.sorts = append(o.sorts, SearchSort{k: "asc"})
	}
	return o
}
func (o *SearchCondition) SortDesc(ks ...string) *SearchCondition {
	for _, k := range ks {
		o.sorts = append(o.sorts, SearchSort{k: "desc"})
	}
	return o
}

// 查询结果.
type SearchResponse struct {
	Took    int64               `json:"took,omitempty"` // search time in milliseconds
	Timeout bool                `json:"timed_out"`      // search timed out
	Hits    *SearchResponseHits `json:"hits,omitempty"` // the actual search hits
}

type SearchResponseHit struct {
	Index  string          `json:"_index,omitempty"`  // index name
	Type   string          `json:"_type,omitempty"`   // type meta field
	Id     string          `json:"_id,omitempty"`     // external or internal
	Source json.RawMessage `json:"_source,omitempty"` // stored document source
}

func (o *SearchResponseHit) Unmarshal(ptr interface{}) error {
	return json.Unmarshal(o.Source, ptr)
}

type SearchResponseHits struct {
	TotalHits int64                `json:"total"`               // total number of hits found
	MaxScore  *float64             `json:"max_score,omitempty"` // maximum score of all hits
	Hits      []*SearchResponseHit `json:"hits,omitempty"`      // the actual hits returned
}

func (o *SearchResponseHits) Each(fn func(index int, hit *SearchResponseHit) error) {
	for i, hit := range o.Hits {
		if fn(i, hit) != nil {
			break
		}
	}
}
