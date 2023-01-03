package shorteningo

type Configuration struct {
  Key string
}

func Add(key string) *Configuration { return &Configuration{Key: key} }
