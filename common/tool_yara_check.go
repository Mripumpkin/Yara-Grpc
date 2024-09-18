package common

import (
	"encoding/json"
	"os"
	"time"
	"yara/format"

	yara "github.com/hillu/go-yara/v4"
)

func YaraBuildData(strSamplePath string, structYaraResult *format.YaraResult) string {
	str_data, err := json.Marshal(structYaraResult)
	if err != nil {
		Logger.Error(err)
		str_data = []byte{}
	}
	GetGCacheHandle().SetWithExpire(strSamplePath, string(str_data), time.Second*60)
	return string(str_data)
}

func YaraCheckOneSample(samplepath string, yaraRules []string) string {

	//默认初始化为失败
	structAvlxResult := format.YaraResult{}
	structAvlxResult.SampleFilePath = samplepath
	structAvlxResult.Status = "Fail"
	_, err := os.Stat(samplepath)
	if err != nil {
		Logger.Errorf("文件不存在: %v", samplepath)
		return YaraBuildData(samplepath, &structAvlxResult)
	}

	if len(yaraRules) <= 0 {
		Logger.Errorf("无yara规则进行检测: %v", samplepath)
		return YaraBuildData(samplepath, &structAvlxResult)
	}
	
	result := "nil"
	if result == "nil" && err == nil {
		structAvlxResult.Status = "nil"
		Logger.Errorf("yara规则检测不匹配: %v", result)
		return YaraBuildData(samplepath, &structAvlxResult)
	}

	if err != nil {
		Logger.Errorf("yara规则检测失败: %v", err)
		return YaraBuildData(samplepath, &structAvlxResult)
	}
	structAvlxResult.Result = result
	structAvlxResult.Status = "True"
	return YaraBuildData(samplepath, &structAvlxResult)
}

func YaraToCheck(yaraRule []string, samplePath string) (yara.MatchRules, error) {
	c, err := yara.NewCompiler()
	if c != nil && err != nil {
		Logger.Errorf("yara规则获取失败: %v", samplePath)
		return nil, err
	}
	for _, yara := range yaraRule {
		if err := c.AddString(yara, ""); err != nil {
			Logger.Errorf("读取规则失败: %v", samplePath)
			return nil, err
		}
		c.Destroy()
	}
	rules, err := c.GetRules()
	if err != nil {
		Logger.Errorf("规则加载失败: %v", samplePath)
		return nil, err
	}
	s, err := yara.NewScanner(rules)
	if err != nil {
		return nil, err
	}
	var m yara.MatchRules
	if err := s.SetCallback(&m).ScanFile(samplePath); err != nil {
		Logger.Errorf("规则检测失败: %v", samplePath)
		return nil, err
	}
	return m, err
}

func YaraFileCheckOneSample(samplepath string, rulepath string) string {
	//缓存处理
	gcdata, gcerr := GetGCacheHandle().Get(samplepath)
	if gcerr == nil { //在缓存中找得到
		Logger.Infof("Get from cache %v", samplepath)
		return gcdata.(string)
	}
	//默认初始化为失败
	var structAvlxResult format.YaraResult
	structAvlxResult.SampleFilePath = samplepath
	structAvlxResult.Status = "Fail"
	_, err := os.Stat(samplepath)
	if err != nil {
		Logger.Errorf("样本文件不存在: %v", samplepath)
		return YaraBuildData(samplepath, &structAvlxResult)
	}

	_, err = os.Stat(rulepath)
	if err != nil {
		Logger.Errorf("yara文件不存在: %v", rulepath)
		return YaraBuildData(samplepath, &structAvlxResult)
	}

	// result, err := YaraFileToCheck(samplepath, rulepath)
	result := "nil"
	if result == "nil" && err == nil {
		structAvlxResult.Status = "nil"
		Logger.Errorf("yara文件检测不匹配: %v", result)
		return YaraBuildData(samplepath, &structAvlxResult)
	}

	if err != nil {
		Logger.Errorf("yara文件检测失败: %v", err)
		return YaraBuildData(samplepath, &structAvlxResult)
	}

	structAvlxResult.Result = result
	structAvlxResult.Status = "True"
	return YaraBuildData(samplepath, &structAvlxResult)
}
