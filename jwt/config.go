package jwt

type Config struct {
	KeyFileDir string
}

// Create a new configuration object
// fileDir is the directory of key file(s)
func NewConfig(fileDir string) *Config {
	return &Config{
		KeyFileDir: fileDir,
	}
}

// Configuration to take effect
func (s *Config) Effect() error {
	if fs, cnt := findJwtKeyFiles(s.KeyFileDir); cnt <= 0 {
		return nil
	} else if hs, err := filesToHandlers(fs); err != nil {
		return err
	} else {
		return setJwtHandlers(hs)
	}
}
