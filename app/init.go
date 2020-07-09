package app


func init() {
	// 加载配置文件
	err := LoadConfig()
	if err != nil {
		panic(err)
	}

}
