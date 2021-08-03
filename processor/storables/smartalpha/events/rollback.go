package events

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func (s *Storable) Rollback(ctx context.Context, tx pgx.Tx) error {
	b := &pgx.Batch{}
	tables := []string{
		"user_join_entry_queue_events",
		"user_join_exit_queue_events",
		"user_redeem_tokens_events",
		"user_redeem_underlying_events",
		"transaction_history",
		"epoch_end_events",
	}
	for _, t := range tables {
		query := fmt.Sprintf(`delete from smart_alpha.%s where included_in_block = $1`, t)
		b.Queue(query, s.block.Number)
	}

	br := tx.SendBatch(ctx, b)
	_, err := br.Exec()
	if err != nil {
		return err
	}

	return br.Close()
}
