package internal

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	log "github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// check
func initDB() {
	if len(viper.GetString("Database")) > 1 {
		dbpool, err := pgxpool.New(context.Background(), viper.GetString("Database"))
		if err != nil {
			log.Error().Err(err)
		}
		defer dbpool.Close()

		var resp string

		//Creating postgres samples table
		err = dbpool.QueryRow(context.Background(), "create table IF NOT EXISTS postgres_samples(id int PRIMARY KEY UNIQUE, message TEXT);").Scan(&resp)
		if err != nil {
			log.Error().Err(err)
		}

		log.Info().Msgf("Init table... DB response: %s", resp)

		//Inserting a row
		err = dbpool.QueryRow(context.Background(), "insert into postgres_samples values (1, 'SAMPLE DATA IS IN DB NOW');").Scan(&resp)
		if err != nil {
			log.Error().Err(err)
		}

		log.Info().Msgf("Init rows... DB response: %s", resp)
		log.Info().Msg("DB Initialized")
	}
}

func getPGData() postgresSamples {
	dbpool, err := pgxpool.New(context.Background(), viper.GetString("Database"))
	if err != nil {
		log.Error().Err(err)
	}
	defer dbpool.Close()

	var resp string

	err = dbpool.QueryRow(context.Background(), "select message from postgres_samples where id=1;").Scan(&resp)
	if err != nil {
		log.Error().Err(err)
	}

	data := postgresSamples{
		ID:      1,
		Message: resp,
	}

	return data
}
