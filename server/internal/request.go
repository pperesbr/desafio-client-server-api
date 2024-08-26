package internal

import "context"

type Request interface {
	Do(ctx context.Context, code string) (*Quote, error)
}
