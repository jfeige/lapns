package lapns

import (
	"crypto/tls"
	"net"
	"time"
	"errors"
	"strings"
)

const(
	TimeoutSecond = 5
)

type Client struct {
	serverurl string
	certpem string
	noencpem string
}

/**
	serverurl -- gateway.push.apple.com:2195
 */
func NewClient(serverurl,cert,noenc string)*Client{
	return &Client{
		serverurl:serverurl,
		certpem:cert,
		noencpem:noenc,
	}
}

func (this Client)send(msg *Msg)(resp *Response){

	resp = &Response{}
	cert,err := tls.LoadX509KeyPair(this.certpem,this.noencpem)
	if err != nil{
		resp.Err = err
		return
	}
	servers := strings.Split(this.serverurl, ":")

	conf := &tls.Config{
		ServerName : servers[0],
		Certificates : []tls.Certificate{cert},
	}
	conn,err := net.Dial("tcp",this.serverurl)
	if err != nil{
		resp.Err = err
		return
	}
	defer conn.Close()
	tlsconn := tls.Client(conn,conf)

	err = tlsconn.Handshake()
	if err != nil{
		resp.Err = err
		return
	}

	//state := tlsconn.ConnectionState()

	defer tlsconn.Close()
	_,err = tlsconn.Write(msg.createPayload())
	if err != nil{
		resp.Err = err
		return
	}

	timeoutChan := make(chan bool,1)
	go func(){
		time.Sleep(TimeoutSecond * time.Second)
		timeoutChan <- true
	}()
	responseChan := make(chan []byte, 1)
	//conn.SetReadDeadline(time.Now().Add(10*time.Second))
	go func(){
		response := make([]byte, 6, 6)  // 错误信息长度为 6 字节
		tlsconn.Read(response)
		responseChan <- response
	}()

	select{
	case r :=<- responseChan:
		//收到错误，记录错误
		resp.Err = errors.New(PushResponseErrCode[r[1]])
	case <- timeoutChan:
		//没有错误
		resp.Sucess = true
		resp.Err = nil
	}
	return

}
