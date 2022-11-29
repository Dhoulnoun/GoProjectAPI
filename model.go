package GoProjectAPI

import (
	"context"
	"github.com/go-redis/redis/v9"
)

type sensor struct {
	ID           string `json:"id"`
	IATA         string `json:"iata"`
	MEASURETYPE  string `json:"measuretype"`
	MEASUREVALUE string `json:"measurevalue"`
	MEASUREDATE  string `json:"measuredate"`
}

type sensors []sensor

func (s *sensor) getSensor(client *redis.Client) error {
	return client.Get(context.Background(), "sensor:"+s.ID).Scan(s)
}

func (s *sensor) getSensors(client *redis.Client, start, count int) ([]string, uint64) {
	rows, err, _ := client.Scan(context.Background(), uint64(start), "sensor:*", int64(uint64(count))).Result()
	return rows, err
}
