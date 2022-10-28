package refreshStream

import (
	"context"
)

type RefreshStreamCommon interface {
	Get(ctx context.Context) ([]RefreshStreamWithNull, error)
	GetId(ctx context.Context, id interface{}) (RefreshStreamWithNull, error)
	Insert(ctx context.Context, rs *RefreshStreamWithNull) error
	// Insert(ctx context.Context, nameColumns string, values string) error
	Update(ctx context.Context, rs *RefreshStreamWithNull) error
	Delete(ctx context.Context, id interface{}) error
}

type RefreshStreamUseCase interface {
	RefreshStreamCommon
}

type RefreshStreamRepository interface {
	RefreshStreamCommon
}
