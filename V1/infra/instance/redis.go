package instance

import (
	"github.com/redis/rueidis"
	"github.com/rs/zerolog/log"
)

func NewRedis(urls ...string) (rueidis.Client, CloseFunc) {
	client, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: urls})
	if err != nil {
		log.Fatal().Msgf("can't create redis client: %v", err)
	}

	return client, func() error {
		client.Close()
		return nil
	}
}
