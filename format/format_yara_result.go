package format

type YaraResult struct {
	SampleFilePath string      `json:"sample_file"`
	Result         interface{} `json:"result"`
	FileHash       string      `json:"file_hash"`
	Status         string      `json:"status"`
}
