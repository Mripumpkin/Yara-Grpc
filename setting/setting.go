package setting

import (
	"os"
	"path"
)

var (
	Str_runtime_path, _   = os.Getwd()
	Str_conf_setting_path = path.Join(Str_runtime_path, "conf", "setting.yml")
)

const ()
