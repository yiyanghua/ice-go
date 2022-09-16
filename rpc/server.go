package rpc

import (
	"encoding/json"
	"github.com/yiyanghua/ice-go/rpc/codec"
	"io"
	"log"
	"reflect"
	"sync"
	"time"
)

const MagicNumber = 0x3bef5c

type Option struct {
	MagicNumber    int
	CodecType      codec.Type
	ConnectTimeout time.Duration
	HandleTimeout  time.Duration
}

var DefaultOption = &Option{
	MagicNumber:    MagicNumber,
	CodecType:      codec.GobType,
	ConnectTimeout: time.Second * 10,
}

type Server struct {
	serverMap sync.Map
}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

func (*Server) ServerConn(conn io.ReadWriteCloser) {
	var opt Option
	if err := json.NewDecoder(conn).Decode(opt); err != nil {
		log.Println("rpc server: options error", err)
	}
	if opt.MagicNumber != MagicNumber {
		log.Printf("rpc server: invalid magic number %x", opt.MagicNumber)
		return
	}
	f := codec.NewCodecFuncMap[opt.CodecType]
	if f == nil {
		log.Printf("rpc server: invalid codec type %x", opt.CodecType)
		return
	}

}

var invalidRequest = struct{}{}

func (server *Server) serverCodec(cc codec.Codec, opt *Option) {
	/*sending := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	for {
		req,err :=
	}*/
}

type request struct {
	h           *codec.Header
	argv, reply reflect.Value
}
