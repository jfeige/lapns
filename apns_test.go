package lapns

import (
	"testing"
	"time"
	"fmt"
)

func Test_Send(t *testing.T){

	msg := new(Msg)
	msg.token = "6c59637a8f0654bb5f5a5dcf6d3c821a6a60988641f12910285d70635be04f92"

	aps := make(map[string]interface{})
	aps["alert"] = "hello tips"   		//tips
	aps["badge"] = 12	  				//应用图标上的数字
	aps["sound"] = "default" 			//声音

	payload := make(map[string]interface{})

	payload["aps"] = aps
	payload["atime"] = time.Now().Unix()
	payload["asterisk"] = 1
	payload["uid"] = 10057

	msg.payload = payload

	content := make(map[string]interface{})
	content["atime"] = time.Now().Unix()
	content["sid"] = 10057
	content["status"] = 1

	msg.content = content

	serverurl := "gateway.push.apple.com:2195"
	certpem := "cert.pem"
	noencpem := "noenc.pem"
	client := NewClient(serverurl,certpem,noencpem)

	resp := client.send(msg)

	fmt.Printf("发送状态:%v",resp)
}
