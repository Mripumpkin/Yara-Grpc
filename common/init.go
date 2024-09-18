package common

import (
	"yara/setting"
)

var (
	GDB         *DBconfig
	GCFG        *ConfigFile
	GConfigInfo *RuntimeConfig
)

/*
@初始化加载配置
*/
func loadConfigHandle() {
	db_cfg, _ := ViperConfigFileHandle(setting.Str_conf_setting_path, "yml")
	GCFG.handle_yaml = db_cfg
}

func init() {
	Logger = MainLog() //先调用
	GDB = newDBConfig()
	GCFG = newCFGConfig()
	//配置文件相关初始化
	loadConfigHandle()
	initConfigInfo()

	InitCache()
	Logger.Info("配置读取成功")
}
