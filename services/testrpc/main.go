package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json"
)

const port = 8099

type Args struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type Reply struct {
	Result float64 `json:"result"`
}

type MathService struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

func (m *MathService) Add(r *http.Request, args *Args, reply *Reply) error {
	reply.Result = args.A + args.B
	return nil
}

func (m *MathService) Divide(r *http.Request, args *Args, reply *Reply) error {
	if args.B == 0 {
		return fmt.Errorf("division by zero")
	}
	reply.Result = args.A / args.B
	return nil
}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")

	mathService := new(MathService)
	s.RegisterService(mathService, "Math")

	http.Handle("/rpc", s)

	log.Printf("RPC started and is listening on :%v", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		panic(err)
	}
}
