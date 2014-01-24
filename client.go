//chat client
//Author: Andy Zhou
//Date: 2014.1.19
//Email:ablozhou@gmail.com

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func (client *Client) Input(inputchan chan string) {

	reader := bufio.NewReader(os.Stdin)
	for {

		input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("error readstring",err.Error())
            continue
        }
		if input == "/quit\n" {
			fmt.Println("ByeBye..")
			client.conn.Close()
			os.Exit(0)
		}

        inputchan <- input
	}
}

func (client *Client) Send(inputchan chan string) {
	for {
        input := <-inputchan
        _, err := client.conn.Write([]byte(input))
		if err != nil {
			fmt.Println(err.Error())
			client.conn.Close()
			os.Exit(2)
		}
	}
}

func (client *Client) Recv(outputchan chan string) {

	buf := make([]byte, 1024)
	for {
		length, err := client.conn.Read(buf)
		if checkError(err, "Connection") == false {
			fmt.Println("Server is dead ...ByeBye")
			os.Exit(0)
		}
        outputchan <-string(buf[0:length])
	}

}

func (client *Client) Output(outputchan chan string) {
    for {
        outstr := <-outputchan
        fmt.Print(outstr)
    }
}
func StartClient(tcpAddr *net.TCPAddr) (cli *Client) {

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err, "DialTCP")
	var client Client
    var user User
	user.addr = "[Me]"
	user.name = user.addr
    client.user = &user
	client.conn = conn
    inputchan := make(chan string)
    outputchan := make(chan string)
	fmt.Println("start client send go routine")
    go client.Input(inputchan)
	go client.Send(inputchan)
    go client.Recv(outputchan)
    go client.Output(outputchan)
    return  &client
}
