package persistence

// GetStringSlice will return the string slice value for a key.
func (s *Service) GetStringSlice(key string) ([]string, error) {
	configEntry, err := s.Get(key)
	if err != nil {
		return []string{}, err
	}
	return configEntry.GetStringSlice()
}

// SetStringSlice will set the string slice value for a key.
func (s *Service) SetStringSlice(key string, value []string) error {
	configEntry := KeyVal{key, value}
	return s.Save(configEntry)
}

// GetString will return the string value for a key.
func (s *Service) GetString(key string) (string, error) {
	configEntry, err := s.Get(key)
	if err != nil {
		return "", err
	}
	return configEntry.GetString()
}

// SetString will set the string value for a key.
func (s *Service) SetString(key string, value string) error {
	configEntry := KeyVal{key, value}
	return s.Save(configEntry)
}

// GetNumber will return the numeric value for a key.
func (s *Service) GetNumber(key string) (float64, error) {
	configEntry, err := s.Get(key)
	if err != nil {
		return 0, err
	}
	return configEntry.GetNumber()
}

// SetNumber will set the numeric value for a key.
func (s *Service) SetNumber(key string, value float64) error {
	configEntry := KeyVal{key, value}
	return s.Save(configEntry)
}

// GetBool will return the boolean value for a key.
func (s *Service) GetBool(key string) (bool, error) {
	configEntry, err := s.Get(key)
	if err != nil {
		return false, err
	}
	return configEntry.GetBool()
}

// SetBool will set the boolean value for a key.
func (s *Service) SetBool(key string, value bool) error {
	configEntry := KeyVal{key, value}
	return s.Save(configEntry)
}

// GetStringMap will return the string map value for a key.
func (s *Service) GetStringMap(key string) (map[string]string, error) {
	configEntry, err := s.Get(key)
	if err != nil {
		return map[string]string{}, err
	}
	return configEntry.GetStringMap()
}

// SetStringMap will set the string map value for a key.
func (s *Service) SetStringMap(key string, value map[string]string) error {
	configEntry := KeyVal{key, value}
	return s.Save(configEntry)
}
