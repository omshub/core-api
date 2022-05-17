package api

type Config struct {
	Port string `yaml:"port" env-description:"Server port" env-default:"1927"`
}
