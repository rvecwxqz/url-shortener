package server

import (
	"context"
	"errors"
	"github.com/rvecwxqz/url-shortener/internal/config"
	"github.com/rvecwxqz/url-shortener/internal/grpc/proto"
	pb "github.com/rvecwxqz/url-shortener/internal/grpc/proto"
	"github.com/rvecwxqz/url-shortener/internal/service"
	"github.com/rvecwxqz/url-shortener/internal/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/url"
)

type Server struct {
	proto.UnimplementedShortenerServer
	Storage      storage.Storage
	deleteWorker *service.DeleteWorker
	delCtx       context.Context
}

func NewServer(ctx context.Context, db storage.Storage) *Server {
	return &Server{
		Storage:      db,
		deleteWorker: service.NewDeleteWorker(ctx, db, config.NewConfig()),
		delCtx:       ctx,
	}
}

func (s Server) getIdFromCtx(ctx context.Context) string {
	var id string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("user_id")
		if len(values) > 0 {
			id = values[0]
		}
	}
	return id
}

func (s Server) GetLongUrl(ctx context.Context, in *pb.GetLongUrlRequest) (*emptypb.Empty, error) {
	var empty emptypb.Empty

	longURL, err := s.Storage.GetLongURL(ctx, in.Id)
	var e *storage.DeletedError
	if errors.As(err, &e) {
		return nil, status.Error(codes.PermissionDenied, "Value deleted")
	} else if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	header := metadata.New(map[string]string{"Location": longURL})
	if err := grpc.SetHeader(ctx, header); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &empty, nil
}

func (s Server) GetShortUrl(ctx context.Context, in *pb.GetShortUrlRequest) (*pb.GetShortUrlResponse, error) {
	var resp proto.GetShortUrlResponse
	id := s.getIdFromCtx(ctx)
	longURL := in.LongUrl
	longURL, _ = url.QueryUnescape(longURL)
	if !service.IsURLValid(longURL) {
		return nil, status.Error(codes.InvalidArgument, "URL not valid")
	}
	shortURL, err := s.Storage.CheckExists(ctx, longURL)
	if err != nil {
		return nil, status.Error(codes.Internal, "Internal error")
	}
	if shortURL == "" {
		shortURL = service.GetShortURL()
		err = s.Storage.AppendURL(ctx, shortURL, longURL, id)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	resp.ShortUrl = shortURL
	return &resp, nil
}

func (s Server) GetUrls(ctx context.Context, in *emptypb.Empty) (*pb.GetUrlsResponse, error) {
	var resp proto.GetUrlsResponse
	id := s.getIdFromCtx(ctx)
	data, err := s.Storage.GetURLsByID(ctx, id, config.NewConfig().BaseURL)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	urlDataSlc := make([]*pb.UrlData, 0, len(data))
	for _, v := range data {
		urlDataSlc = append(urlDataSlc, &pb.UrlData{ShortUrl: v.ShortURL, OriginalUrl: v.OriginalURL})
	}
	resp.Urls = urlDataSlc
	return &resp, err
}

func (s Server) DeleteUrls(ctx context.Context, in *pb.DeleteUrlsRequest) (*emptypb.Empty, error) {
	id := s.getIdFromCtx(ctx)
	go s.deleteWorker.AppendURLs(s.delCtx, in.Urls, config.NewConfig().BaseURL, id)
	return &emptypb.Empty{}, nil
}

func (s Server) PingDb(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	err := storage.DBPing(config.NewConfig().DatabaseDSN)
	if err != nil {
		return nil, status.Error(codes.Internal, "Database not working")
	}
	return &emptypb.Empty{}, nil
}
