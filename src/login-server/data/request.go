package data

import (
	"encoding/binary"
	"io"
)

type Request struct {
	Cmd    byte   // 指令
	Length int16  // 数据长度
	Data   []byte // 数据
}

// 打包
func (p *Request) Pack(writer io.Writer) (err error) {
	if err = binary.Write(writer, binary.BigEndian, &p.Cmd); err != nil {
		return
	}
	if err = binary.Write(writer, binary.BigEndian, &p.Length); err != nil {
		return
	}
	if err = binary.Write(writer, binary.BigEndian, &p.Data); err != nil {
		return
	}
	return
}

// 解包
func (p *Request) UnPack(reader io.Reader) (err error) {
	if err = binary.Read(reader, binary.BigEndian, &p.Cmd); err != nil {
		return
	}
	if err = binary.Read(reader, binary.BigEndian, &p.Length); err != nil {
		return
	}
	p.Data = make([]byte, p.Length)
	if err = binary.Read(reader, binary.BigEndian, &p.Data); err != nil {
		return
	}
	return
}

// 注册请求
type RegisterRequest struct {
	Username string
	Password string
}

// 登录请求
type LoginRequest struct {
	Username string
	Password string
}
