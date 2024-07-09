package configs

var setting Setting

type Setting struct {
	AppName string `yaml:"app_name" json:"app_name"`
}

func (m Setting) DefaultConfig() any {
	return Setting{
		AppName: "billion",
	}
}

func Settings() Setting {
	return setting
}

func SetSettings(s Setting) {
	setting = s
}
