package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	pb "github.com/rvecwxqz/url-shortener/internal/grpc/proto"
	server2 "github.com/rvecwxqz/url-shortener/internal/grpc/server"
	"github.com/rvecwxqz/url-shortener/internal/middleware"
	"github.com/rvecwxqz/url-shortener/internal/server"
	"github.com/rvecwxqz/url-shortener/internal/storage"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)
import _ "net/http/pprof"

func main() {
	ctx := context.Background()

	h := server.NewHandlers(ctx)
	r := chi.NewRouter()

	go listenGrpc(ctx, h.GetStorage())

	r.Use(middleware.GzipHandle)
	r.Use(middleware.CookieHandler)

	r.Route("/", func(r chi.Router) {
		r.Get("/{id}", h.GetHandler)
		r.Post("/", h.PostHandler)

		r.Post("/api/shorten", h.APIHandler)
		r.Get("/api/user/urls", h.URLsHandler)
		r.Post("/api/shorten/batch", h.BatchHandler)
		r.Delete("/api/user/urls", h.DeleteHandler)

		r.Get("/ping", h.PingHandler)
	})
	http.ListenAndServe(h.Config.ServerAddress, r)
}

func listenGrpc(ctx context.Context, db storage.Storage) {
	listen, err := net.Listen("tcp", ":3200")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(server2.UserIDInterceptor))
	pb.RegisterShortenerServer(s, server2.NewServer(ctx, db))

	fmt.Println("gRPC server started")
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}
