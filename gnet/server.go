package main

import (
	"fmt"
	"github.com/panjf2000/gnet/v2"
	"github.com/panjf2000/gnet/v2/pkg/logging"
	bbPool "github.com/panjf2000/gnet/v2/pkg/pool/bytebuffer"
	goPool "github.com/panjf2000/gnet/v2/pkg/pool/goroutine"
	"time"
)

type telnetServer struct {
	*gnet.BuiltinEventEngine
	workerPool *goPool.Pool
}

func (t *telnetServer) OnBoot(eng gnet.Engine) (action gnet.Action) {
	return
}

func (t *telnetServer) OnTraffic(c gnet.Conn) (action gnet.Action) {
	/*buf, _ := c.Next(-1)
	ends := string(buf)
	if strings.Compare(ends, "quit\r\n") == 0 {
		c.Write([]byte("bye bye~"))
		c.Close()
	} else {
		c.Write(buf)
	}
	return*/
	buf := bbPool.Get()
	_, _ = c.WriteTo(buf)

	// just for test
	_ = c.InboundBuffered()
	_ = c.OutboundBuffered()
	_, _ = c.Discard(1)

	_ = t.workerPool.Submit(
		func() {
			mid := buf.Len() / 2
			bs := make([][]byte, 2)
			bs[0] = buf.B[:mid]
			bs[1] = buf.B[mid:]
			_ = c.AsyncWrite([]byte(fmt.Sprint("ack from ", buf.String())), func(c gnet.Conn, err error) error {
				logging.Debugf("conn=%s done writev: %v", c.RemoteAddr().String(), err)
				return nil
			})

			/*if t.writev {
				mid := buf.Len() / 2
				bs := make([][]byte, 2)
				bs[0] = buf.B[:mid]
				bs[1] = buf.B[mid:]
				_ = c.AsyncWritev(bs, func(c Conn, err error) error {
					logging.Debugf("conn=%s done writev: %v", c.RemoteAddr().String(), err)
					return nil
				})
			} else {
				_ = c.AsyncWrite(buf.Bytes(), func(c Conn, err error) error {
					logging.Debugf("conn=%s done write: %v", c.RemoteAddr().String(), err)
					return nil
				})
			}*/
		})

	return
}

func (t *telnetServer) OnTick() (delay time.Duration, action gnet.Action) {
	return
}

func main() {
	server := &telnetServer{
		workerPool: goPool.Default(),
	}
	var opts []gnet.Option
	opts = append(opts, func(opts *gnet.Options) {
		opts.ReuseAddr = true
		opts.TCPNoDelay = gnet.TCPNoDelay
	})
	gnet.Run(server, fmt.Sprintf("tcp://%v", "127.0.0.1:8080"), opts...)
}
