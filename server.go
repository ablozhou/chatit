//chat server
//Author: Andy Zhou
//Date: 2014.1.19
//Email:ablozhou@gmail.com

package main

import (
	"fmt"
	"net"
	"strings"
)


func init() {
	server.clients = make(map[string]Client, 10240)
    server.msgchan = make(chan Msg,1000)
}

func (server *Server) Recv(client *Client) {

	fmt.Println(client.user.name, " entered")

	defer client.conn.Close()
	buf := make([]byte, 1024)
	var msg string
	for {
		length, err := client.conn.Read(buf)
		if checkError(err, "Connection") == false {
			fmt.Println(client.user.name, client.user.addr, " leaved")
			break
		}
		if length > 0 {
			buf[length] = 0
		} else {
			fmt.Println(client.user.name, " recv nothing,continue")
			continue
		}

		if buf[0] == '/' {
			msg = string(buf[1 : length-1])
			server.CmdParse(msg, client)
		} else {
			msg = string(buf[0:length])
			var climsg = Msg{client.user, nil, msg}
			server.msgchan <- climsg
		}
	}

}

func (server *Server) CmdParse(cmdstr string, client *Client) (cmd []string, err error) {

	cmd = strings.Split(cmdstr, " ")
    var msg string

	switch cmd[0] {
	case "username":
		if cmd[1] == "ROOT" || cmd[1] == "ME" {

			msg = fmt.Sprintf("%s changed name to %s failed, continu.\n", client.user.name, cmd[1])
			//err = error.Error()
		}
		msg = fmt.Sprintf("%s changed name to %s\n", client.user.name, cmd[1])
		client.user.name = cmd[1]

		var climsg = Msg{client.user, nil, msg}
		server.msgchan <- climsg
    case "to":
        var user User
        user.addr=cmd[1]
        msg = strings.Join(cmd[2:]," ")+"\n"
		var climsg = Msg{client.user, &user, msg}
		server.msgchan <- climsg

	}

    return
}

func (server *Server)GetClient(user *User) (client Client,ok bool) {
    if user == nil {
        ok = false
        return
    }
    client,ok = server.clients[user.addr]
    return
}

func (server *Server)Send(climsg *Msg) (err error) {

    if client,ok := server.GetClient(climsg.to); ok {
        conn := client.conn
        _, err = conn.Write([]byte(climsg.msg))
        if err != nil {
            fmt.Println(err.Error())
            fmt.Print(climsg.to.name, " leaved.\n")
            conn.Close()
            delete(server.clients, climsg.to.addr)
        }
    }
    return
}

func (server *Server) Listen(addr string)(err error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", addr)
	checkError(err, "ResolveTCPAddr")
	l, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, "ListenTCP")
	fmt.Println("Listening ...")
	for {
		conn, err := l.Accept()
		checkError(err, "Accept")
		var client Client
        var user User
        client.user = &user
		client.user.addr = conn.RemoteAddr().String()
		client.conn = conn
		client.user.name = client.user.addr
		server.clients[client.user.addr] = client

		fmt.Println("Accept ", client.user.name)
		go server.Recv(&client)
	}
}

//服务器发送消息
func (server *Server) ProcessMsg() {

	for {
		climsg := <-server.msgchan

		str := fmt.Sprintf("[%s]:%s", climsg.from.name, climsg.msg)
		fmt.Print(str)
        climsg.msg = str
        if climsg.to != nil {
            server.Send(&climsg)
        } else {
            for _,client := range server.clients {
                climsg.to = client.user
                server.Send(&climsg)
            }
        }
	}
}

//启动服务器
func StartServer(port string)  {

	addr := ":" + port
	go server.ProcessMsg()
	go server.Listen(addr)

}
