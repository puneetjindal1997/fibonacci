package redis

import "time"

// local function for redis to insert the data using
// ttl tiunme duration of 15 minutes
func InsertDataTORedis(db Database, key string, value interface{}, ttl time.Duration) error {
	_, err := db.Set(key, value, ttl)
	if err != nil {
		return err
	}
	return nil
}

// get data from redis
// using key
func GetDataTORedis(db Database, key string) (interface{}, error) {
	value, err := db.Get(key)
	if err != nil {
		return 0, err
	}
	return value, nil
}
