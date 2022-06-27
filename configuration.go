package shorteningo

type ConfigurationCuttly struct {
  Key string
}

func Add(key string) *ConfigurationCuttly { return &ConfigurationCuttly{Key: key} }