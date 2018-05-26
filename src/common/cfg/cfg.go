// 配置加载器
package cfg

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Configuration struct {
	Database struct {
		Dsn     string
		MaxConn int `yaml:"maxConn"`
	}
	Server struct {
		// 登录网关配置
		LoginGate struct {
			Host string
			Port int
		}
		// 登录服务器配置
		LoginServer struct {
			Host string
			Port int
		}
		// 游戏网关配置
		GameGate struct {
			Host string
			Port int
		}
		// 游戏服务器配置
		GameServer struct {
			Host string
			Port int
		}
	}
}

// 加载配置
func New(filename string) (config *Configuration, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	config = new(Configuration)
	err = yaml.Unmarshal(data, config)
	return
}
