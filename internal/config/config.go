package config

import (
	"rest-api-server/pkg/logging"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

// создаём структуру которая соответствует конфигу в yml формате
type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8000"`
	} `yaml:"listen"`
}

// объявляем две переменные инстанс используем для записи туда конфига, а ванс
// для того что бы убедиться что наш код сработает только один раз
var instance *Config
var once sync.Once

func GetConfig() *Config {
	//ванс говорит о том что код выполнится только один раз, а остальные запуски
	//проигнорируются
	once.Do(func() {
		//передаём логер
		logger := logging.GetLogger()
		logger.Info("чтение конфига")
		instance = &Config{}
		//записываем в инстанс конфиг из файла, проверяем на ошибки, возвращаем конфиг выше
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info("ошибка чтения конфига: ", help)
			logger.Fatal(err)
		}
	})
	return instance
}
