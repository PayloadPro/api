package main

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"

	"github.com/andrew-waters/payload.pro/deps"
	"github.com/andrew-waters/payload.pro/rpc"
	"github.com/andrew-waters/payload.pro/services"
)

func main() {

	// Flags
	fs := flag.NewFlagSet("", flag.ExitOnError)
	httpAddr := fs.String("http.addr", ":8081", "Address for HTTP (JSON) server")
	flag.Usage = fs.Usage // only show our flags
	fs.Parse(os.Args[1:])

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

	router.HandleFunc("/bins/{id}", func(w http.ResponseWriter, r *http.Request) {
		JSONEndpointHandler(w, r, func() (interface{}, int, error) {
			return rpc.NewCreatePayload(services)(ctx, r)
		})
	})

	log.Fatal(http.ListenAndServe(*httpAddr, router))

}
