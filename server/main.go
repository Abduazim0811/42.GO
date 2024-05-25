package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"Homework_42/models"
)



type Arith int


func (t *Arith) Add(args *models.Args, reply *float64) error {
	*reply = args.A + args.B
	return nil
}

func (t *Arith) Subtract(args *models.Args, reply *float64) error {
	*reply = args.A - args.B
	return nil
}

func (t *Arith) Multiplication(args *models.Args, reply *float64) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *models.Args, reply *float64) error {
	if args.B == 0 {
		return fmt.Errorf("division by zero!!!(nolga bo'lib bo'lamaydi!!!)")
	}
	*reply = args.A / args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Serving RPC server on port 1234")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error", err)
			continue
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
