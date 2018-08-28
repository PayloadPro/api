package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/mongodb/mongo-go-driver/mongo"
	"golang.org/x/net/context"

	"github.com/andrew-waters/pro.payload.api/configs"
	"github.com/andrew-waters/pro.payload.api/deps"
	"github.com/andrew-waters/pro.payload.api/rpc"
	"github.com/andrew-waters/pro.payload.api/services"
)

func main() {

	var err error

	sa := getFlagConfig()

	// Services
	services := &deps.Services{
		Payload: &services.PayloadService{},
	}

	// Config
	config := &deps.Config{
		App: &configs.AppConfig{},
		DB:  &configs.DatabaseConfig{},
	}
	config.DB.Setup()

	// Create a DB Connection
	dbc, err := mongo.NewClient(config.DB.ConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	err = dbc.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	defer dbc.Disconnect(nil)

	// Add the DB to the Service
	database := config.DB.BinDatabase
	collection := config.DB.BinRequestCollection
	services.Payload.Collection = dbc.Database(database).Collection(collection)

	// Context
	rand.Seed(time.Now().UnixNano())
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	defer cancel()

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		JSONEndpointHandler(w, r, func() (interface{}, int, error) {
			return rpc.NewLandingPayload(services)(ctx, r)
		})
	})

	router.HandleFunc("/bins/{id}", func(w http.ResponseWriter, r *http.Request) {
		JSONEndpointHandler(w, r, func() (interface{}, int, error) {
			return rpc.NewCreatePayload(services)(ctx, r)
		})
	})

	log.Fatal(http.ListenAndServe(*sa, router))
}

// getFlagConfig sets the runtime variables
func getFlagConfig() *string {

	fs := flag.NewFlagSet("", flag.ExitOnError)
	server := fs.String("server", "0.0.0.0", "HTTP server")
	port := fs.String("port", "8081", "HTTP server port")
	flag.Usage = fs.Usage

	fs.Parse(os.Args[1:])

	si := make([]string, 0)
	si = append(si, *server, *port)

	sa := strings.Join(si, ":")

	return &sa
}
