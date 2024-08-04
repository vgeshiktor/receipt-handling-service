package webserver

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
)

type Server interface {
	Start() error
}

type HttpServer struct {
	mux  *http.ServeMux
	addr string
	port int
}

// NewHTTPServer creates a new HTTP server
func NewHTTPServer(addr string, port int) (*HttpServer, error) {
	return &HttpServer{
		mux:  http.NewServeMux(),
		addr: addr,
		port: port,
	}, nil
}

func (s *HttpServer) Start() error {
	fmt.Printf("Starting HTTP server on %s:%d\n", s.addr, s.port)

	// init server multiplexer

	// add routes
	s.mux.HandleFunc("/healthz", healthz)
	s.mux.HandleFunc("/ready", ready)

	// add receipt routes
	s.mux.HandleFunc("POST /receipts", createReceipt)
	s.mux.HandleFunc("GET /receipts/{receiptId}", getReceipt)
	s.mux.HandleFunc("PUT /receipts/{receiptId}", updateReceipt)
	s.mux.HandleFunc("DELETE /receipts/{receiptId}", deleteReceipt)

	ctx, cancelCtx := context.WithCancel(context.Background())
	serverOne := &http.Server{
		Addr:    ":3333",
		Handler: s.mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	// serverTwo := &http.Server{
	// 	Addr:    ":4444",
	// 	Handler: s.mux,
	// 	BaseContext: func(l net.Listener) context.Context {
	// 		ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
	// 		return ctx
	// 	},
	// }

	// start listening
	go func() {
		err := serverOne.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server one closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server one: %s\n", err)
		}
		cancelCtx()
	}()

	// go func() {
	// 	err := serverTwo.ListenAndServe()
	// 	if errors.Is(err, http.ErrServerClosed) {
	// 		fmt.Printf("server two closed\n")
	// 	} else if err != nil {
	// 		fmt.Printf("error listening for server two: %s\n", err)
	// 	}
	// 	cancelCtx()
	// }()

	<-ctx.Done()

	// err := http.ListenAndServe(fmt.Sprintf("%s:%d", s.addr, s.port), s.mux)
	// if errors.Is(err, http.ErrServerClosed) {
	// 	fmt.Printf("server closed\n")
	// } else if err != nil {
	// 	fmt.Printf("error starting server: %s\n", err)
	// 	os.Exit(1)
	// }

	return nil
}
