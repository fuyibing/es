// author: wsfuyibing <websearch@163.com>
// date: 2021-08-19

package es

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

var Conn = &Connection{}

// ES连接.
type Connection struct {
	cli *elasticsearch.Client
	err error
}

// 创建文档.
func (o *Connection) Create() *CreateManager {
	return NewCreateManager()
}

// 读取文档.
func (o *Connection) Get() *GetManager {
	return NewGetManager()
}

// 读取文档.
func (o *Connection) Gets() *GetsManager {
	return NewGetsManager()
}

// 删除文档.
func (o *Connection) Delete() *DeleteManager {
	return NewDeleteManager()
}

// 文档查询.
func (o *Connection) Search() *SearchManager {
	return NewSearchManager()
}

// 按ID更新.
//
// 每次更新1条记录.
func (o *Connection) Update() *UpdateManager {
	return NewUpdateManager()
}

// 按条件更新.
//
// 批量更新, 每次请求可更新N条记录.
func (o *Connection) UpdateByQuery() *UpdateByQueryManager {
	return NewUpdateByQueryManager()
}

// 映射维护.
func (o *Connection) Mapping() *MappingManager {
	return NewMappingManager()
}

// 索引设置.
func (o *Connection) Setting() *SettingManager {
	return NewSettingManager()
}

// 发送请求.
func (o *Connection) Send(ctx context.Context, req esapi.Request) ([]byte, error) {
	// 1. 全局错误.
	//    创建连接时出错.
	if o.err != nil {
		return nil, o.err
	}

	// 2. 请求过程错误.
	res, err := req.Do(ctx, o.cli)
	if err != nil {
		return nil, err
	}

	// 3. 读取请求内容.
	//    结束时关闭缓冲区.
	var body []byte
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}
	defer func() {
		if res.Body != nil {
			_ = res.Body.Close()
		}
	}()

	// 4. 简易错误.
	//    {
	//        "error": "error reason"
	//    }
	er1 := &struct {
		Error string `json:"error"`
	}{}
	if json.Unmarshal(body, er1) == nil && er1.Error != "" {
		return nil, fmt.Errorf(er1.Error)
	}

	// 5. 标准错误.
	//    {
	//        "error": {
	//            "type": "exception",
	//            "reason": "error reason"
	//        }
	//    }
	er2 := &struct {
		Error *struct {
			Type   string `json:"type"`
			Reason string `json:"reason"`
		} `json:"error"`
	}{}
	if json.Unmarshal(body, er2) == nil && er2.Error != nil && er2.Error.Reason != "" {
		return nil, fmt.Errorf("[%s] %s", er2.Error.Type, er2.Error.Reason)
	}

	return body, nil
}

// 初始化连接.
func (o *Connection) init() {
	o.cli, o.err = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: Config.Address,
	})
}
