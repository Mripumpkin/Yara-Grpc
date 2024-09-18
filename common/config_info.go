package common

import (
	"github.com/spf13/viper"
)

type ConfigFile struct {
	handle_yaml *viper.Viper
}

func newCFGConfig() *ConfigFile {
	gcfg := &ConfigFile{}
	return gcfg
}

func GetConfigFileHanlde() *viper.Viper {
	if GCFG.handle_yaml == nil {
		panic("GCFG.handle_yaml nil")
	}
	return GCFG.handle_yaml
}

// 需要再添加
type RuntimeConfig struct {
}

/**
 * @func: 读取初始化时加载的配置文件 //暂无用
 * @return {*}
 * @mark:
 */
func initConfigInfo() {
	Logger.Info("配置信息加载...")
	GConfigInfo = new(RuntimeConfig)

	Logger.Infof("配置信息加载完毕 %v", GConfigInfo)
}

/*
@重载配置控制
*/
func ReloadConfig() {
	loadConfigHandle()
	// do something
	GDB.gc.Purge()
}
