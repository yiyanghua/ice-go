package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

type MyDataBucket struct {
	br     *bytes.Buffer
	rcond  *sync.Cond
}

func NewDataBucket() *MyDataBucket {
	buf := make([]byte, 0)
	db := &MyDataBucket{
		br:     bytes.NewBuffer(buf),
	}
	db.rcond = sync.NewCond(&sync.RWMutex{})
	return db
}

func (db *MyDataBucket) Read(i int) {
	db.rcond.L.Lock()
	defer db.rcond.L.Unlock()
	var data []byte
	var d byte
	var err error
	for {
		//读取一个字节
		if d, err = db.br.ReadByte(); err != nil {
			if err == io.EOF {
				if string(data) != "" {
					fmt.Printf("reader-%d: %s\n", i, data)
				}
				db.rcond.Wait()
				data = data[:0]
				continue
			}
		}
		data = append(data, d)
	}
}

func (db *MyDataBucket) Put(d []byte) (int, error) {
	db.rcond.L.Lock()
	defer db.rcond.L.Unlock()
	//写入一个数据块
	n, err := db.br.Write(d)
	db.rcond.Broadcast()
	return n, err
}

func main() {
	db := NewDataBucket()

	go db.Read(1)

	go db.Read(2)

	for i := 0; i < 10; i++ {
		go func(i int) {
			d := fmt.Sprintf("data-%d", i)
			db.Put([]byte(d))
		}(i)
		time.Sleep(100 * time.Millisecond)
	}
}
