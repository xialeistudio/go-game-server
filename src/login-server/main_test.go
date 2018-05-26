package main

import (
	"bytes"
	"encoding/gob"
	"login-server/data"
	"net"
	"testing"
)

func TestSendData(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:11000")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.Close()
	// 编码注册数据包
	buff := new(bytes.Buffer)
	encoder := gob.NewEncoder(buff)
	registerRequest := &data.RegisterRequest{
		Username: "xialei",
		Password: "111111",
	}
	if err := encoder.Encode(registerRequest); err != nil {
		t.Error(err)
		return
	}
	// 编码协议报文
	request := &data.Request{
		Cmd:    CmdRegister,
		Length: int16(buff.Len()),
		Data:   buff.Bytes(),
	}
	t.Log(request)
	if err := request.Pack(conn); err != nil {
		t.Error(err)
		return
	}
	t.Log("编码完成")
}
