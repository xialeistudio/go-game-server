// 登录服务器
// 本服务器完成用户登录注册以及凭证的分发
package main

import (
	"bufio"
	"bytes"
	"common/cfg"
	"common/db"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"login-server/data"
	"login-server/service"
	"net"
	"os"
)

var (
	err      error
	config   *cfg.Configuration
	database *gorm.DB
	// 业务类
	userService *service.User
	listener    net.Listener
)

const (
	_           = iota
	CmdRegister
	CmdLogin
)
const version = "1.0.0"

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	log.Printf("登录服务器, 版本: %s", version)
	loadConfig()
	loadDatabase()
	loadService()
	launchServer()
}

// 启动服务器
func launchServer() {
	log.Printf("[连接模块] 初始化开始...")
	address := fmt.Sprintf("%s:%d", config.Server.LoginServer.Host, config.Server.LoginServer.Port)
	listener, err = net.Listen("tcp", address)
	if err != nil {
		log.Printf("[连接模块] 初始化失败: %s", err.Error())
	}
	log.Printf("[连接模块] 初始化成功. 监听地址 %s", address)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("[连接模块] 接收连接失败: %s", err.Error())
			continue
		}
		go handleClientConn(conn)
	}
}

// 接收客户端连接
func handleClientConn(conn net.Conn) {
	log.Printf("[连接模块] 接收连接: %s", conn.RemoteAddr().String())
	scanner := bufio.NewScanner(conn)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if atEOF || len(data) < 3 {
			return
		}
		length := int16(0)
		if err = binary.Read(bytes.NewReader(data[1:3]), binary.BigEndian, &length); err != nil {
			return
		}
		if int(length)+3 <= len(data) {
			return int(length) + 3, data[:int(length)+3], nil
		}
		return
	})
	for scanner.Scan() {
		request := new(data.Request)
		request.UnPack(bytes.NewBuffer(scanner.Bytes()))
		switch request.Cmd {
		case CmdRegister:
			decoder := gob.NewDecoder(bytes.NewReader(request.Data))
			register := new(data.RegisterRequest)
			if err := decoder.Decode(register); err != nil {
				log.Printf("[协议模块] 注册报文解析失败: %s", err.Error())
				continue
			}
			user, err := userService.Register(register.Username, register.Password)
			if err != nil {
				log.Printf("[用户模块] 注册失败: %s", err.Error())
				continue
			}
			log.Printf("[用户模块] 注册成功, ID: %d", user.Id)
		}
	}
}

// 加载配置
func loadConfig() {
	log.Printf("[配置模块] 加载配置开始...")
	config, err = cfg.New("cfg/server.yml")
	if err != nil {
		log.Fatalf("[配置模块] 加载配置失败: %s", err.Error())
	}
	log.Printf("[配置模块] 加载配置成功")
}

// 连接数据库
func loadDatabase() {
	log.Printf("[数据库模块] 连接数据库开始...")
	database, err = db.New(config)
	if err != nil {
		log.Fatalf("[数据库模块] 连接数据库失败: %s", err.Error())
	}
	log.Printf("[数据库模块] 连接数据库成功")
}

// 初始化业务类
func loadService() {
	userService = service.NewUser(database)
}
