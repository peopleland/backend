package data

import (
	"backend/app/user/internal/conf"
	"context"
	"fmt"
	"testing"
)

func Test_mintRecordRepo_FindLastMintRecord(t *testing.T) {
	var conff = &conf.Config{
		FaunaDBSecret: "fnAEbfitSAACVKRgPF0ZYX-Q3zZiIE3jQpr_9km0",
	}
	d, _ := NewData(conff, logger)
	mrr := NewMintRecordRepo(d, logger)
	ctx := context.Background()

	mr, err := mrr.FindLastMintRecord(ctx, "m1", "1", "1", 1641104969)
	fmt.Println(mr, err)
}
