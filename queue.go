// author: wsfuyibing <websearch@163.com>
// date: 2021-08-18

package es

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"github.com/fuyibing/log/v2"
)

var Queue = new(QueueManager)

// 操作类别.
const (
	QueueOperationUnknown QueueOperation = iota
	QueueOperationCreate
	QueueOperationDelete
	QueueOperationUpdate
)

// 队列类型.
type (
	// 队列数据.
	QueueData struct {
		Operation  QueueOperation
		DocumentId string
	}

	// 队列管理.
	QueueManager struct {
		mu                        *sync.RWMutex
		limit, offset, processing int64
		list                      map[int64]QueueData
		running                   bool
	}

	// 操作类别.
	QueueOperation int
)

// 入列.
func (o *QueueManager) Add(ctx context.Context, data QueueData) {
	// 1. 锁操作.
	//    开始时加锁, 完成后解锁并触发队列.
	o.mu.Lock()
	defer func() {
		o.mu.Unlock()
		go o.Run(ctx)
	}()

	// 2. 数据操作.
	index := atomic.LoadInt64(&o.limit)
	atomic.AddInt64(&o.limit, 1)
	o.list[index] = data
	log.Infofc(ctx, "[es][queue] 数据入列, index=%d.", index)
}

// 执行.
func (o *QueueManager) Run(ctx context.Context) {
	// 1. 执行唯一性.
	if o.running {
		return
	}

	// 2. 执行状态监控.
	o.running = true
	if log.Config.DebugOn() {
		log.Debug("[es][queue] 开始出列.")
	}
	defer func() {
		if r := recover(); r != nil {
			log.Panicf("[es][queue] 出列异常, fatal=%v.", r)
		} else if log.Config.DebugOn() {
			log.Debug("[es][queue] 出列完成.")
		}
		o.running = false
	}()

	// 3. 执行过程.
	for {
		limit := atomic.LoadInt64(&o.limit)
		offset := atomic.LoadInt64(&o.offset)

		// 3.1 空队列.
		if offset >= limit {
			break
		}

		// 3.2 并发限流.
		processing := atomic.LoadInt64(&o.processing)
		if processing >= Config.MaxConcurrency {
			if log.Config.DebugOn() {
				log.Debugf("[es][queue] 并发限流, processing=%d, maximum=%d, waiting=%d.", processing, Config.MaxConcurrency, limit-offset)
			}
			time.Sleep(time.Millisecond * 200)
			continue
		}

		// 3.3 取出数据.
		atomic.AddInt64(&o.offset, 1)
		atomic.AddInt64(&o.processing, 1)
		if wd, ok := o.pop(offset); ok {
			log.Infof("[es][queue] 取出数据, index=%d, processing=%d.", offset, processing)
			// 3.3.1 分发数据.
			go func(index int64, data QueueData) {
				if log.Config.DebugOn() {
					log.Debugf("[es][queue] 数据分发, index=%d.", index)
				}
				defer func() {
					if r := recover(); r != nil {
						log.Debugf("[es][queue] 分发异常, index=%d, fatal=%v.", index, r)
					} else if log.Config.DebugOn() {
						log.Debugf("[es][queue] 分发完成, index=%d.", index)
					}
					atomic.AddInt64(&o.processing, -1)
				}()
				o.dispatch(index, data)
			}(offset, wd)
		} else {
			// 3.3.2 已被取出.
			atomic.AddInt64(&o.processing, -1)
		}
	}
}

// 数据分发.
func (o *QueueManager) dispatch(index int64, ata QueueData) {
	time.Sleep(time.Second)
}

// 初始队列.
func (o *QueueManager) init() {
	o.list = make(map[int64]QueueData)
	o.mu = &sync.RWMutex{}
}

// 取出数据.
func (o *QueueManager) pop(index int64) (QueueData, bool) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if data, ok := o.list[index]; ok {
		delete(o.list, index)
		return data, ok
	}
	return QueueData{}, false
}
