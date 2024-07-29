package server

import (
	"baby-log/cmd/web"
	"baby-log/internal/domain"
	"baby-log/internal/gateway"
	"baby-log/internal/views"
	"database/sql"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	if s.sentryEnabled {
		e.Use(sentryecho.New(sentryecho.Options{
			Repanic: true,
		}))
	}

	fileServer := http.FileServer(http.FS(web.Files))
	e.GET("/assets/*", echo.WrapHandler(fileServer))

	e.GET("/_ping", s.PingHandler)

	admin := e.Group("/admin")

	admin.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Validator: func(username string, password string, context echo.Context) (bool, error) {
			return username == s.admin.Username && password == s.admin.Password, nil
		},
	}))

	admin.GET("/history", s.HistoryHandler)
	admin.GET("/person/new", s.PersonCreateHandler)
	admin.GET("/person", s.PersonListHandler)
	admin.POST("/person", s.PersonPostHandler)
	admin.GET("/person/:id/activity", s.PersonActivityHandler)

	return e
}

func (s *Server) PersonActivityHandler(c echo.Context) error {
	return Render(c, views.ActivityView())
}

func (s *Server) PersonPostHandler(c echo.Context) error {
	id, err := domain.GenerateId()

	if err != nil {
		return err
	}

	err = s.queries.CreatePerson(s.ctx, gateway.CreatePersonParams{
		PersonID: id,
		Name:     c.FormValue("name"),
		BirthDate: sql.NullTime{
			Valid: false,
		},
		BirthTime: sql.NullTime{
			Valid: false,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return err
	}

	return HxRedirect(c, "/admin/person")
}

func (s *Server) PersonListHandler(c echo.Context) error {
	people, err := s.queries.FilterPeople(s.ctx)

	if err != nil {
		return err
	}

	return Render(c, views.PersonListView(people))
}

func (s *Server) PersonCreateHandler(c echo.Context) error {
	return Render(c, views.PersonCreateView())
}
func (s *Server) HistoryHandler(c echo.Context) error {
	return Render(c, views.HistoryView())
}

func (s *Server) PingHandler(c echo.Context) error {
	resp := map[string]string{
		"status": "OK",
	}

	return c.JSON(http.StatusOK, resp)
}
