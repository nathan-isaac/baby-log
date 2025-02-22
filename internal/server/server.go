package server

import (
	"baby-log/internal/domain"
	"baby-log/internal/gateway"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type AdminUser struct {
	Username string
	Password string
}

type Server struct {
	db            *sql.DB
	ctx           context.Context
	queries       *gateway.Queries
	domain        *domain.App
	admin         AdminUser
	sentryEnabled bool
}

type Options struct {
	Port          int
	Host          string
	DatabaseUrl   string
	AdminUsername string
	AdminPassword string
}

func DefaultOptions() *Options {
	return &Options{
		Port:          8080,
		Host:          "localhost",
		DatabaseUrl:   ":memory:",
		AdminUsername: "",
		AdminPassword: "",
	}
}

type OptionsFunc func(options *Options) error

func WithServerAddress(options *Options) error {
	port, err := strconv.Atoi(os.Getenv("PORT"))

	if err != nil {
		return err
	}

	options.Port = port
	options.Host = os.Getenv("HOST")

	return nil
}

func WithAdmin(options *Options) error {
	username := os.Getenv("ADMIN_USERNAME")
	password := os.Getenv("ADMIN_PASSWORD")

	if username == "" {
		return fmt.Errorf("ADMIN_USERNAME environment variable is required")
	}
	if password == "" {
		return fmt.Errorf("ADMIN_PASSWORD environment variable is required")
	}

	options.AdminUsername = username
	options.AdminPassword = password
	return nil
}

func WithOptions(options *Options, opts ...OptionsFunc) error {
	for _, opt := range opts {
		err := opt(options)

		if err != nil {
			return err
		}
	}

	return nil
}

func NewServer(db *sql.DB) (*http.Server, error) {
	options := DefaultOptions()
	err := WithOptions(options, WithServerAddress, WithAdmin)

	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	queries := gateway.New(db)

	NewServer := &Server{
		db:      db,
		ctx:     ctx,
		queries: queries,
		domain: &domain.App{
			Queries: queries,
			Ctx:     ctx,
		},
		admin: AdminUser{
			Username: options.AdminUsername,
			Password: options.AdminPassword,
		},
		sentryEnabled: os.Getenv("SENTRY_DSN") != "",
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", options.Host, options.Port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server, nil
}
