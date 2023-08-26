package persistence

import (
	"errors"
)

func (s *Service) WarmCache() {
	s.Logger().Traceln("Warming config value cache")

	res, err := s.db.GetDatabase().Query("SELECT key, value FROM datastore")
	if err != nil || res.Err() != nil {
		s.Logger().Fatalln("error warming config cache", err, res.Err())
	}
	defer res.Close()

	for res.Next() {
		var rowKey string
		var rowValue []byte
		if err = res.Scan(&rowKey, &rowValue); err != nil {
			s.Logger().Errorln("error pre-caching config row", err)
		}
		s.cache[rowKey] = rowValue
	}
}

// GetCachedValue will return a value for key from the cache.
func (s *Service) GetCachedValue(key string) ([]byte, error) {
	s.muxCache.Lock()
	defer s.muxCache.Unlock()

	// Check for a cached value
	if val, ok := s.cache[key]; ok {
		return val, nil
	}

	return nil, errors.New(key + " not found in cache")
}

// SetCachedValue will set a value for key in the cache.
func (s *Service) SetCachedValue(key string, b []byte) {
	s.muxCache.Lock()
	defer s.muxCache.Unlock()

	s.cache[key] = b
}
