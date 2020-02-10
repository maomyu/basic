package config

type RabbitMQConfig interface {
	GetURL() string
	GetUser() string
	GetPassword() string
}
type defaultRabbitMQConfig struct {
	URL      string `json:"url",yaml:"url"`
	User     string `json:"user",yaml:"user"`
	Password string `json:"password",yaml:"password"`
}

func (d defaultRabbitMQConfig) GetURL() string {
	return d.URL
}
func (d defaultRabbitMQConfig) GetUser() string {
	return d.User
}
func (d defaultRabbitMQConfig) GetPassword() string {
	return d.Password
}
