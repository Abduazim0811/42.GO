package main

import (
	"Homework_42/models"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)


func callRPC(client *rpc.Client, method string, args *models.Args) float64{
	var reply float64
	err:=client.Call(method,args,&reply)
	if err!=nil{
		fmt.Println("Error calling:", err)
		return 0
	}
	return reply

}

func main(){
	conn,err:=net.Dial("tcp", "localhost:1234")
	if err!=nil{
		log.Fatal("Dialing error:",err)
	}
	client:=rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	args:=models.Args{A:10, B:2}
	fmt.Println("10 + 2 =", callRPC(client, "Arith.Add", &args))
    fmt.Println("10 - 2 =", callRPC(client, "Arith.Subtract", &args))
    fmt.Println("10 * 2 =", callRPC(client, "Arith.Multiplication", &args))
    fmt.Println("10 / 2 =", callRPC(client, "Arith.Divide", &args))
}