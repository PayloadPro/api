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
	"golang.org/x/net/context"

	"github.com/andrew-waters/payload.pro/deps"
	"github.com/andrew-waters/payload.pro/rpc"
	"github.com/andrew-waters/payload.pro/services"
)

func main() {

	sa := getFlagConfig()

	// Services
	services := &deps.Services{
		Payload: &services.PayloadService{},
	}

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
