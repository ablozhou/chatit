//chat server common struct and function
//Author: Andy Zhou
//Date: 2014.1.19
//Email:ablozhou@gmail.com

package main

import (
	"fmt"
	"net"
)

type User struct {
	addr string //ip:port
	name string
	passwd string
}
type Client struct {
    user *User
	conn net.Conn
}

type Msg struct {
	from *User
	to   *User //default nil, broadcast
	msg  string
}

type Server struct {
	clients map[string]Client
    msgchan chan Msg
}

var server Server

func checkError(err error, info string) (res bool) {

	if err != nil {
		fmt.Println(info + "  " + err.Error())
		return false
	}
	return true
}
