package cfg

import (
	"testing"
)

func TestNew(t *testing.T) {
	config, err := New("../../../cfg/server.yml")
	if err != nil {
		t.Error(err)
		return
	}
	if config.Server.LoginGate.Host != "0.0.0.0" || config.Server.LoginGate.Port != 10000 {
		t.Error("加载配置失败")
	}
}
