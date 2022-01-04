package data

import (
	"backend/app/user/internal/conf"
	"context"
	"fmt"
	"testing"
)

func Test_openerRecordRepo_getListPaginate(t *testing.T) {
	var conff = &conf.Config{
		FaunaDBSecret: "fnAEbfitSAACVKRgPF0ZYX-Q3zZiIE3jQpr_9km0",
	}
	d, _ := NewData(conff, logger)
	orr := NewOpenerRecordRepo(d, logger)
	ctx := context.Background()

	mr, err := orr.GetListPaginateBefore(ctx, 2, 1)
	fmt.Println(mr, err)
}
