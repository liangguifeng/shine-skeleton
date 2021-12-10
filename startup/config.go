package startup

import (
	"git.myscrm.cn/2c/antman-union/vars"
	"git.myscrm.cn/golang/stark/v4/config"
)

// LoadConfig 加载配置对象映射
func LoadConfig() error {
	config.MapConfig("App", vars.AppSetting, true)
	config.MapConfig("Messenger", vars.MessengerSetting, true)
	config.MapConfig("FileBox", vars.FileBoxSetting, true)
	config.MapConfig("UserCenterAppId", vars.UserCenterAppIdSetting, true)
	config.MapConfig("OpenApi", vars.OpenApiSetting, true)
	config.MapConfig("ExtApi", vars.ExtApiSetting, true)
	return nil
}
