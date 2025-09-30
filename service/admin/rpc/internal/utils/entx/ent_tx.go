package entx

import (
	"context"
	"fmt"

	ent2 "github.com/suyuan32/simple-admin-core/service/admin/rpc/ent"
	"github.com/zeromicro/go-zero/core/logx"
)

// WithTx uses transaction in ent.
func WithTx(ctx context.Context, client *ent2.Client, fn func(tx *ent2.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		logx.Errorw("failed to start transaction", logx.Field("detail", err.Error()))
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rollBackErr := tx.Rollback(); rollBackErr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rollBackErr)
		}
		logx.Errorw("errors occur in transaction", logx.Field("detail", err.Error()))
		return err
	}
	if err := tx.Commit(); err != nil {
		logx.Errorw("failed to commit transaction", logx.Field("detail", err.Error()))
		return err
	}
	return nil
}
