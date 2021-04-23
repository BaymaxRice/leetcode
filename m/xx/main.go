package main

import (
	"sync"
	"time"
)

// GO里面MAP如何实现key不存在 get操作等待 直到key存在或者超时
// 保证并发安全，且需要实现以下接口：
type sp interface {
	// 存入key /val，如果该key读取的goroutine是挂起状态，则唤醒。
	// 此方法不会阻塞，时刻都可以立即执行并返回
	Out(key string, val interface{})

	Rd(key string, timeout time.Duration) interface{}
	// 读取一个key，如果key不存在阻塞，等待key存在或者超时
}

type WaitMap struct {
	data map[string]interface{}
	lock sync.RWMutex
}

func (m *WaitMap) Out(key string, val interface{}) {
	m.lock.Lock()
	if m.data == nil {
		m.data = make(map[string]interface{})
	}
	m.data[key] = val
	m.lock.Unlock()
}

func (m *WaitMap) Rd(key string, timeout time.Duration) interface{} {
	m.lock.Lock()

	if m.data == nil {
		m.lock.Unlock()
		return nil
	}

	ret, ok := m.data[key]
	if ok {
		m.lock.Unlock()
		return ret
	}

	m.lock.Unlock()

	timeCh := time.After(timeout)
	for {
		m.lock.Lock()
		ret, ok = m.data[key]
		select {
		case <-timeCh:
			m.lock.Unlock()
			return nil
		default:
			if ok {
				m.lock.Unlock()
				return ret
			}
			m.lock.Unlock()
		}
	}
}
