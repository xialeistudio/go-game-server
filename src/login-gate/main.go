// 登录网关，数据加解密操作
package main

import (
	"common/cfg"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

var (
	err             error
	config          *cfg.Configuration
	loginServerConn net.Conn
	privateKey      *rsa.PrivateKey
	publicKey       *rsa.PublicKey
)

const version = "1.0.0"

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	log.Printf("登录网关, 版本: %s", version)
	loadConfig()
	loadEncrypt()
	connectLoginServer()
	launchServer()
}

// 初始化加解密模块
func loadEncrypt() {
	log.Printf("[加解密模块] 加载私钥开始...")
	data, err := ioutil.ReadFile("cfg/private.key")
	if err != nil {
		log.Fatalf("[加解密模块] 加载私钥失败: %s", err.Error())
	}
	block, _ := pem.Decode(data)
	if block == nil {
		log.Printf("[加解密模块] 加载私钥失败")
	}
	privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatalf("[加解密模块] 加载私钥失败: %s", err.Error())
	}
	log.Printf("[加解密模块] 加载私钥成功")
	// 加载公钥
	log.Printf("[加解密模块] 加载公钥开始...")
	data, err = ioutil.ReadFile("cfg/public.key")
	if err != nil {
		log.Fatalf("[加解密模块] 加载公钥失败: %s", err.Error())
	}
	block, _ = pem.Decode(data)
	if block == nil {
		log.Printf("[加解密模块] 加载公钥失败")
	}
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatalf("[加解密模块] 加载公钥失败: %s", err.Error())
	}
	publicKey = pub.(*rsa.PublicKey)
	log.Printf("[加解密模块] 加载公钥成功")
}

// 连接登录服务器
func connectLoginServer() {
	log.Printf("[服务器模块] 连接登录服务器开始...")
	address := fmt.Sprintf("%s:%d", config.Server.LoginServer.Host, config.Server.LoginServer.Port)
	loginServerConn, err = net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("[服务器模块] 连接登录服务器失败: %s", err.Error())
	}
	log.Printf("[服务器模块] 连接登录服务器成功")
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

// 启动服务器
func launchServer() {
	log.Printf("[连接模块] 初始化开始...")
	address := fmt.Sprintf("%s:%d", config.Server.LoginGate.Host, config.Server.LoginGate.Port)
	listener, err := net.Listen("tcp", address)
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

}
