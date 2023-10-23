package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"strings"
	"time"
)

var tName = "shortener"

type DBStorage struct {
	pool *pgxpool.Pool
}

type rowDB struct {
	id       uint
	original string
	short    string
	uuid     string
}

type NotUniqueError struct {
	err error
}

func (e *NotUniqueError) Error() string {
	return fmt.Sprintf("%v", e.err)
}

func NewNotUniqueError() error {
	return &NotUniqueError{
		err: errors.New("found duplicate"),
	}
}

type DeletedError struct {
	err error
}

func (e *DeletedError) Error() string {
	return fmt.Sprintf("%v", e.err)
}
func NewDeletedError() error {
	return &DeletedError{
		err: errors.New("URL deleted"),
	}
}

func newDBStorage(DSN string) *DBStorage {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbpool, err := pgxpool.New(ctx, DSN)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	tExists, err := isTableExists(ctx, dbpool)
	if err != nil {
		fmt.Println(err)
	}
	if !tExists {
		err = createTable(ctx, dbpool)
		if err != nil {
			fmt.Println(err)
		}
	}
	return &DBStorage{pool: dbpool}
}

func createTable(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, "CREATE TABLE shortener (id serial primary key, original varchar, short varchar, uuid varchar, deleted boolean DEFAULT false)")
	if err != nil {
		return err
	}
	return nil
}

func isTableExists(ctx context.Context, pool *pgxpool.Pool) (bool, error) {
	_, err := pool.Query(ctx, "SELECT * FROM shortener LIMIT 1")
	if err != nil && strings.Contains(err.Error(), "не существует") {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func DBPing(DSN string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	pool, _ := pgxpool.New(ctx, DSN)
	if err := pool.Ping(context.Background()); err != nil {
		return err
	}
	return nil
}

func (s DBStorage) GetLongURL(ctx context.Context, key string) (string, error) {
	var longURL string
	var deleted bool
	err := s.pool.QueryRow(ctx, "SELECT original FROM shortener WHERE short = $1 LIMIT 1", key).Scan(&longURL)
	if err != nil {
		return "", errors.New("not found")
	}
	err = s.pool.QueryRow(ctx, "SELECT deleted FROM shortener WHERE short = $1 LIMIT 1", key).Scan(&deleted)
	if err != nil {
		fmt.Printf("Err in GetLongURL %v\n", err)
	}
	if deleted {
		return "", NewDeletedError()
	}
	return longURL, nil
}

func (s DBStorage) CheckExists(ctx context.Context, longURL string) (string, error) {
	var shortURL string
	err := s.pool.QueryRow(ctx, "SELECT short FROM shortener WHERE original = $1 LIMIT 1", longURL).Scan(&shortURL)
	if err != nil && err.Error() == "no rows in result set" {
		return "", nil
	} else if err != nil && err.Error() != "no rows in result set" {
		return "", err
	}
	return shortURL, nil
}

func (s DBStorage) AppendURL(ctx context.Context, shortURL string, longURL string, id string) error {
	_, err := s.pool.Exec(ctx, "INSERT INTO shortener (original, short, uuid) VALUES ($1, $2, $3)", longURL, shortURL, id)
	if err != nil {
		return err
	}
	return nil
}

func (s DBStorage) GetURLsByID(ctx context.Context, id string, baseURL string) ([]URLsData, error) {
	data := make([]URLsData, 0)
	rows, err := s.pool.Query(ctx, "SELECT original, short FROM shortener WHERE uuid = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var u URLsData
		err = rows.Scan(&u.OriginalURL, &u.ShortURL)
		u.ShortURL = baseURL + "/" + u.ShortURL
		if err != nil {
			return nil, err
		}
		data = append(data, u)
	}
	return data, nil
}

func (s DBStorage) PrintStorage() error {
	rows, err := s.pool.Query(context.Background(), "SELECT * FROM shortener")
	if err != nil {
		return err
	}

	fmt.Println("--------")
	for rows.Next() {
		var row rowDB
		_ = rows.Scan(&row.id, &row.original, &row.short, &row.uuid)
		fmt.Println(row)
	}
	fmt.Println("--------")
	return nil
}

func (s DBStorage) AppendBatch(ctx context.Context, data map[string]string, id string) error {
	b := &pgx.Batch{}
	for k, v := range data {
		b.Queue("INSERT INTO shortener (original, short, uuid) VALUES ($1, $2, $3)", v, k, id)
	}
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	res := conn.SendBatch(ctx, b)
	_, err = res.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (s DBStorage) DeleteBatch(ctx context.Context, URLs []string) error {
	b := &pgx.Batch{}
	for _, v := range URLs {
		b.Queue("UPDATE shortener SET deleted = true WHERE short = $1", v)
	}
	conn, err := s.pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	res := conn.SendBatch(ctx, b)
	_, err = res.Exec()
	if err != nil {
		return fmt.Errorf("error in DeleteBatch: %v", err.Error())
	}
	return nil
}

func (s DBStorage) checkShortURLExists(ctx context.Context, shortURL string) bool {
	err := s.pool.QueryRow(ctx, "SELECT short FROM shortener WHERE short = $1 LIMIT 1", shortURL).Scan(&shortURL)
	return err.Error() == "no rows in result set"
}
