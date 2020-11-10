package main

import (
	"fmt"
	"sync"
	"time"
)

type sp interface {
	Put(key string, value interface{})
	Rd(key string, timeout time.Duration) interface{}
}

type Map struct {
	c   map[string]*entry
	rmx *sync.Mutex
}

type entry struct {
	ch      chan struct{}
	value   interface{}
	isExist bool
}

func (m *Map) Put(key string, value interface{}) {
	m.rmx.Lock()
	defer func() {
		m.rmx.Unlock()
	}()
	if e, ok := m.c[key]; ok {
		e.value = value
	} else {
		e := &entry{
			ch:      make(chan struct{}),
			isExist: true,
			value:   value,
		}
		m.c[key] = e
		close(e.ch)
	}
}

func (m *Map) Rd(key string, timeout time.Duration) interface{} {
	m.rmx.Lock()
	if e, ok := m.c[key]; ok && e.isExist {
		m.rmx.Unlock()
		return e.value
	} else if !ok {
		e := &entry{
			ch:      make(chan struct{}),
			isExist: false,
		}
		m.c[key] = e
		m.rmx.Unlock()
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			return nil
		}

	} else {
		m.rmx.Unlock()
		select {
		case <-e.ch:
			return e.value
		case <-time.After(timeout):
			return nil
		}
	}
}

type st struct {
	k string
}

func main() {
	sp := Map{
		c:   make(map[string]*entry),
		rmx: &sync.Mutex{},
	}
	sp.Put("aa", st{
		k: "aa",
	})

	v := sp.Rd("aa", time.Second*1)
	fmt.Printf("%v", v)
}
