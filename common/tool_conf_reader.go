package common

import (
	"github.com/spf13/viper"
)

/**
 * @func:读取配置文件viper
 * @param 文件名称
 * @param 文件类型
 * @param 配置文件的文件夹
 * @return *viper,err
 * @mark:
 */
func ViperConfigHandle(fileName, fileType, filePath string) (conf *viper.Viper, err error) {
	config := viper.New()
	config.SetConfigName(fileName)
	config.SetConfigType(fileType)
	config.AddConfigPath(filePath)

	err = config.ReadInConfig()
	if err != nil {
		Logger.Fatalf("读取配置 %v 失败。 %v", fileName, err)
		panic(err)
	}
	return config, nil

}

/**
 * @func:读取配置文件viper
 * @param 配置文件完整路径
 * @param 文件类型
 * @return *viper,err
 * @mark:
 */
func ViperConfigFileHandle(str_file_path string, str_file_type string) (conf *viper.Viper, err error) {
	config := viper.New()
	config.SetConfigFile(str_file_path)
	config.SetConfigType(str_file_type)

	err = config.ReadInConfig()
	if err != nil {
		Logger.Fatalf("读取配置 %v 失败。 %v", str_file_path, err)
	}
	return config, nil

}
