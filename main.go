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

	"github.com/PayloadPro/pro.payload.api/configs"
	"github.com/PayloadPro/pro.payload.api/deps"
	"github.com/PayloadPro/pro.payload.api/rpc"
	"github.com/PayloadPro/pro.payload.api/services"
)

func main() {

	var err error

	sa := getFlagConfig()

	// Services
	services := &deps.Services{
		Request: &services.RequestService{},
		Bin:     &services.BinService{},
	}

	// Config
	config := &deps.Config{
		App: &configs.AppConfig{},
		DB:  &configs.DatabaseConfig{},
	}
	config.Setup()

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
	db := config.DB.BinDatabase
	rc := config.DB.BinRequestCollection
	bc := config.DB.BinCollection
	services.Request.Collection = dbc.Database(db).Collection(rc)
	services.Bin.Collection = dbc.Database(db).Collection(bc)

	router := createRouter(services, config)

	log.Fatal(http.ListenAndServe(*sa, router))
}

func createRouter(services *deps.Services, config *deps.Config) *mux.Router {

	// Context
	rand.Seed(time.Now().UnixNano())
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	defer cancel()

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		JSONEndpointHandler(w, r, func() (interface{}, int, error) {
			return rpc.NewGetRoot(services, config)(ctx, r)
		})
	}).Methods("GET")

	router.HandleFunc("/bins", func(w http.ResponseWriter, r *http.Request) {
		JSONEndpointHandler(w, r, func() (interface{}, int, error) {
			return rpc.NewCreateBin(services)(ctx, r)
		})
	}).Methods("POST")

	router.HandleFunc("/bins", func(w http.ResponseWriter, r *http.Request) {
		JSONEndpointHandler(w, r, func() (interface{}, int, error) {
			return rpc.NewGetBins(services)(ctx, r)
		})
	}).Methods("GET")

	router.HandleFunc("/bins/{id}", func(w http.ResponseWriter, r *http.Request) {
		JSONEndpointHandler(w, r, func() (interface{}, int, error) {
			return rpc.NewCreateRequest(services)(ctx, r)
		})
	})

	router.HandleFunc("/bins/{id}/requests", func(w http.ResponseWriter, r *http.Request) {
		JSONEndpointHandler(w, r, func() (interface{}, int, error) {
			return rpc.NewGetRequestsForBin(services)(ctx, r)
		})
	}).Methods("GET")

	return router
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
