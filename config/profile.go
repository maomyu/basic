package config

type Profiles interface {
	GetInclude() string
}

type defaultProfiles struct {
	Include string `json:"include",yml:"include"`
}

// 获得待读取的配置文件名称
func (p defaultProfiles) GetInclude() string {
	return p.Include
}
