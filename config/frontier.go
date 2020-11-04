package config

type FrontierConfig struct {
	FrontierId       string `json:"frontierId"`
	LogLevel         int    `json:"logLevel"`
	HeartbeatTimeout int64  `json:"heartbeatTimeout"`
	WriterBufferSize int    `json:"writerBufferSize"`
	ReaderBufferSize int    `json:"readerBufferSize"`
}
