package domain

import (
	"baby-log/internal/gateway"
	"context"
)

type App struct {
	Queries *gateway.Queries
	Ctx     context.Context
}
