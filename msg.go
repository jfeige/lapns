package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"math/rand"
	"time"
	"encoding/hex"
)

const (
	command = 2
	deviceTokenItemid = 1
	payloadItemid = 2
	identifierItemid = 3
	expirationDateItemid = 4
	priorityItemid = 5
	deviceTokenItemLength = 32
	identifierItemLength = 4
	expirationDateItemLength = 4
	priorityItemLength = 1
)

type Msg struct {
	payload map[string]interface{}			//消息内容
	content map[string]interface{}			//扩展参数
	token string							//device token
}



//command = 2 增强性通知格式
//command,identifier,expiry,token length,device length,payload length,payload

func (this *Msg)createPayload()[]byte{
	bpayload,_ := json.Marshal(this.payload) //消息体

	blength := len(bpayload)	//消息体长度

	btoken,_ := hex.DecodeString(this.token)  //token

	identifier := getIdentifier()   //消息唯一标示

	frameBuffer := bytes.NewBuffer([]byte{})
	binary.Write(frameBuffer,binary.BigEndian,uint8(deviceTokenItemid))
	binary.Write(frameBuffer,binary.BigEndian,uint16(deviceTokenItemLength))
	binary.Write(frameBuffer,binary.BigEndian,btoken)
	binary.Write(frameBuffer,binary.BigEndian,uint8(payloadItemid))
	binary.Write(frameBuffer,binary.BigEndian,uint16(blength))
	binary.Write(frameBuffer,binary.BigEndian,bpayload)
	binary.Write(frameBuffer,binary.BigEndian,uint8(identifierItemid))
	binary.Write(frameBuffer,binary.BigEndian,uint16(identifierItemLength))
	binary.Write(frameBuffer,binary.BigEndian,identifier)
	binary.Write(frameBuffer,binary.BigEndian,uint8(expirationDateItemid))
	binary.Write(frameBuffer,binary.BigEndian,uint16(expirationDateItemLength))
	binary.Write(frameBuffer,binary.BigEndian,uint32(0))
	binary.Write(frameBuffer,binary.BigEndian,uint8(priorityItemid))
	binary.Write(frameBuffer,binary.BigEndian,uint16(priorityItemLength))
	binary.Write(frameBuffer,binary.BigEndian,uint8(10))


	buffer := bytes.NewBuffer([]byte{})

	binary.Write(buffer, binary.BigEndian, uint8(command))
	binary.Write(buffer, binary.BigEndian, uint32(frameBuffer.Len()))
	binary.Write(buffer, binary.BigEndian, frameBuffer.Bytes())


	return buffer.Bytes()
}


//消息唯一标示
func getIdentifier()int{
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r.Intn(999999)
}