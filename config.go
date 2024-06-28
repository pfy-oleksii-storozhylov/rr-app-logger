package app

// your configuration is here
type Config struct {
	// yaml, sample mapping to your `key` in the .rr.yaml
	parseJson bool `mapstructure:"parse_json"`
}

//// InitDefaults used to initialize default configuration values
//func (c *Config) InitDefaults() {
//	//if c.parseJson ==false {
//	//	c.parseJson = "some_value"
//	//}
//}
