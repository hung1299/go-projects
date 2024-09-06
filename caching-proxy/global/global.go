package global

type GlobalConfig struct {
	Port   int32
	Origin string
}

var Config = GlobalConfig{}
