package internal

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// test
func RunServer() {
	//Initialize database
	go func() {
		initDB()
	}()
	//контрольный запуск
	//Run webserver
	//test
	log.Info().Msgf("Running PIG Dummy Service on port %s", viper.GetString("Port"))
	r := mux.NewRouter()
	r.HandleFunc("/internal/healthz", healthcheck)
	r.HandleFunc("/database", getDBData)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("resources/"))))
	r.Use(loggingMiddleware)
	http.Handle("/", r)
	port := ":" + viper.GetString("Port")

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal().Err(err).Msg("Startup failed")
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Debug().Msg(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func getDBData(w http.ResponseWriter, r *http.Request) {
	var response API
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Server", "PIG Dummy Service")

	var data postgresSamples

	if len(viper.GetString("Database")) > 1 {
		data = getPGData()
	} else {
		data = postgresSamples{
			ID:      0,
			Message: "NO DATA IN DB",
		}
	}

	if r.Method != http.MethodGet {
		response = API{
			Code:    405,
			Message: "Method not allowed",
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		response = API{
			Code:    200,
			Message: data.Message,
		}
		w.WriteHeader(http.StatusOK)
	}
	js, err := json.Marshal(response)
	if err != nil {
		log.Debug().Msg(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	var response API
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Server", "PIG Dummy Service")
	if r.Method != http.MethodGet {
		response = API{
			Code:    405,
			Message: "Method not allowed",
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		response = API{
			Code:    200,
			Message: "healthy",
		}
		w.WriteHeader(http.StatusOK)
	}
	js, err := json.Marshal(response)
	if err != nil {
		log.Debug().Msg(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}
