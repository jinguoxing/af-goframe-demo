package main

type Conf struct {
	LogConf struct {
		OutputPath string `yaml:"outputPath" json:"outputPath"`
		Mode       string `yaml:"mode" json:"mode"`
	} `yaml:"log" json:"log"`
}

//type LogConf struct {
//	OutputPath string `yaml:"outputPath"`
//}

//func InitLogger(filename string) (c *Conf, err error) {
//
//	yamlFile, err := ioutil.ReadFile(filename)
//	if err != nil {
//		return nil, err
//	}
//
//	err = yaml.Unmarshal(yamlFile, &c)
//	if err != nil {
//		return nil, err
//	}
//
//	return c, nil
//}

//func CustomConfig(core zapcore.Core) zapcore.Core {
//	core, _ = CustomZapCore("./config/log.yaml")
//	return core
//}
//
//func CustomZapCore(lc Conf) (zapcore.Core, error) {
//	customCore := zapcore.NewCore(getEncoder(), getLogWriter(c.LogConf.OutputPath), zapcore.DebugLevel)
//	return customCore, nil
//}
//
//func getEncoder() zapcore.Encoder {
//	encoderConfig := zap.NewProductionEncoderConfig()
//	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
//	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
//	return zapcore.NewJSONEncoder(encoderConfig)
//}
//
//func getLogWriter(path string) zapcore.WriteSyncer {
//	file, _ := os.Create(path)
//	return zapcore.AddSync(file)
//}
