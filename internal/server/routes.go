package server

import (
	"baby-log/cmd/web"
	"baby-log/internal/domain"
	"baby-log/internal/gateway"
	"baby-log/internal/views"
	"database/sql"
	"fmt"
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

	admin.GET("/person/new", s.PersonCreateHandler)
	admin.GET("/person", s.PersonListHandler)
	admin.POST("/person", s.PersonPostHandler)
	admin.GET("/person/:id/activity", s.PersonActivityHandler)
	admin.GET("/person/:id/sleep/new", s.SleepCreateHandler)
	admin.GET("/person/:id/nurse/new", s.NurseCreateHandler)
	admin.POST("/sleep", s.SleepPostHandler)
	admin.POST("/nurse", s.NursePostHandler)

	return e
}

type SleepPostForm struct {
	PersonID string `form:"person_id"`
	Start    string `form:"start"`
	End      string `form:"end"`
	Notes    string `form:"notes"`
}

type NursePostForm struct {
	PersonID string `form:"person_id"`
	Start    string `form:"start"`
	End      string `form:"end"`
	Notes    string `form:"notes"`
}

const (
	FormTimeLayout = "2006-01-02T15:04"
)

func (s *Server) NurseCreateHandler(c echo.Context) error {
	personId := c.Param("id")

	person, err := s.queries.FindPerson(s.ctx, personId)

	if err != nil {
		return err
	}

	return Render(c, views.NurseNew(person))
}

func parseNullableTime(value string) (sql.NullTime, error) {
	if value == "" {
		return sql.NullTime{
			Valid: false,
		}, nil
	}

	parsedValue, err := time.Parse(FormTimeLayout, value)
	if err != nil {
		return sql.NullTime{
			Valid: false,
		}, err
	}

	return sql.NullTime{
		Time:  parsedValue,
		Valid: true,
	}, err
}

func (s *Server) NursePostHandler(c echo.Context) error {
	var form SleepPostForm
	err := c.Bind(&form)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	person, err := s.queries.FindPerson(s.ctx, form.PersonID)

	if err != nil {
		return err
	}

	id, err := domain.GenerateId()

	if err != nil {
		return err
	}

	startAt, err := time.Parse(FormTimeLayout, form.Start)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed to parse start time")
	}

	endAt, err := parseNullableTime(form.End)
	if err != nil {
		return err
	}

	err = s.queries.CreateNurse(s.ctx, gateway.CreateNurseParams{
		NurseID:              id,
		PersonID:             person.PersonID,
		RightDurationSeconds: 0,
		LeftDurationSeconds:  0,
		RightVolumeMl:        0,
		LeftVolumeMl:         0,
		StartedAt:            startAt,
		EndedAt:              endAt,
		Notes:                form.Notes,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	})

	return HxRedirect(c, fmt.Sprintf("/admin/person/%s/activity", person.PersonID))
}

func (s *Server) SleepPostHandler(c echo.Context) error {
	var form SleepPostForm
	err := c.Bind(&form)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	person, err := s.queries.FindPerson(s.ctx, form.PersonID)

	if err != nil {
		return err
	}

	id, err := domain.GenerateId()

	if err != nil {
		return err
	}

	startAt, err := time.Parse(FormTimeLayout, form.Start)
	if err != nil {
		return c.String(http.StatusBadRequest, "failed to parse start time")
	}

	endAt, err := parseNullableTime(form.End)
	if err != nil {
		return err
	}

	err = s.queries.CreateSleep(s.ctx, gateway.CreateSleepParams{
		SleepID:   id,
		PersonID:  person.PersonID,
		StartedAt: startAt,
		EndedAt:   endAt,
		Notes:     form.Notes,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return HxRedirect(c, fmt.Sprintf("/admin/person/%s/activity", person.PersonID))
}

func (s *Server) SleepCreateHandler(c echo.Context) error {
	personId := c.Param("id")

	person, err := s.queries.FindPerson(s.ctx, personId)

	if err != nil {
		return err
	}

	return Render(c, views.SleepNew(person))
}

func (s *Server) PersonActivityHandler(c echo.Context) error {
	personId := c.Param("id")

	person, err := s.queries.FindPerson(s.ctx, personId)

	if err != nil {
		return err
	}

	options := []views.NewLogOption{
		{
			Initial:   "SL",
			Name:      "Sleep",
			Color:     "bg-slate-600",
			Href:      fmt.Sprintf("/admin/person/%s/sleep/new", person.PersonID),
			TimeSince: "2 hours ago",
		},
		{
			Initial:   "BO",
			Name:      "Bottle",
			Color:     "bg-blue-600",
			Href:      "",
			TimeSince: "2 hours ago",
		},
		{
			Initial:   "NU",
			Name:      "Nursing",
			Color:     "bg-pink-600",
			Href:      fmt.Sprintf("/admin/person/%s/nurse/new", person.PersonID),
			TimeSince: "2 hours ago",
		},
		{
			Initial:   "WE",
			Name:      "Weight",
			Color:     "bg-emerald-600",
			Href:      "",
			TimeSince: "2 hours ago",
		},
		{
			Initial:   "DI",
			Name:      "Diaper",
			Color:     "bg-amber-600",
			Href:      "",
			TimeSince: "2 hours ago",
		},
		{
			Initial:   "PU",
			Name:      "Pumping",
			Color:     "bg-violet-600",
			Href:      "",
			TimeSince: "2 hours ago",
		},
		{
			Initial:   "MI",
			Name:      "Miscellaneous",
			Color:     "bg-fuchsia-600",
			Href:      "",
			TimeSince: "2 hours ago",
		},
	}

	sleeps, err := s.queries.FilterSleep(s.ctx, person.PersonID)

	if err != nil {
		return err
	}

	entries := []views.LogEnty{}

	for _, sleep := range sleeps {
		entries = append(entries, views.LogEnty{
			Name:    "Sleep",
			Date:    fmt.Sprintf("%s %s", sleep.StartedAt.Format(time.DateOnly), formatTimeRange(sleep.StartedAt, sleep.EndedAt)),
			Value:   "",
			Initial: "SL",
			Color:   "bg-slate-600",
		})
	}

	nurses, err := s.queries.FilterNurse(s.ctx, person.PersonID)

	if err != nil {
		return err
	}

	for _, nurse := range nurses {
		entries = append(entries, views.LogEnty{
			Name:    "Nurse",
			Date:    fmt.Sprintf("%s %s", nurse.StartedAt.Format(time.DateOnly), formatTimeRange(nurse.StartedAt, nurse.EndedAt)),
			Value:   "",
			Initial: "NU",
			Color:   "bg-pink-600",
		})
	}

	return Render(c, views.ActivityView(options, entries))
}

func formatTimeRange(start time.Time, end sql.NullTime) string {
	if end.Valid {
		return fmt.Sprintf("%s - %s", start.Format(time.Kitchen), end.Time.Format(time.Kitchen))
	}

	return start.Format(time.Kitchen)
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

func (s *Server) PingHandler(c echo.Context) error {
	resp := map[string]string{
		"status": "OK",
	}

	return c.JSON(http.StatusOK, resp)
}
