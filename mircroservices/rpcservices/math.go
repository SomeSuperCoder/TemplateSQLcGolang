package rpcservices

import (
	"fmt"
	"net/http"
)

type Args struct {
	A float64 `json:"a"`
	B float64 `json:"b"`
}

type Reply struct {
	Result float64 `json:"result"`
}

type MathService struct{}

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
