package refreshStream

import (
	"context"
)

type RefreshStreamCommon interface {
	Get(ctx context.Context) ([]RefreshStreamWithNull, error)
	GetId(ctx context.Context, value interface{}) (RefreshStreamWithNull, error)
	Insert(ctx context.Context, nameColumns string, values string) error
	Update(ctx context.Context, setCol string, setValue, whereValue interface{}) error
	Delete(ctx context.Context, value int) error
}

type RefreshStreamUseCase interface {
	RefreshStreamCommon
}

type RefreshStreamRepository interface {
	RefreshStreamCommon
}
