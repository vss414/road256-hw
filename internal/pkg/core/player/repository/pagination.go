package repository

import (
	"github.com/pkg/errors"
	pb "gitlab.ozon.dev/vss414/hw-1/pkg/api"
	"strings"
)

type Pagination struct {
	Limit, Offset    uint64
	Order, Direction string
}

func NewPagination(limit, page uint64, order, direction string) (Pagination, error) {
	if limit == 0 {
		limit = 50
	}

	if page == 0 {
		page = 1
	}

	offset := limit * (page - 1)

	direction = strings.ToUpper(direction)
	if direction == "" {
		direction = pb.Direction_DIRECTION_ASC.String()
	}
	if v, ok := pb.Direction_value[direction]; !ok || v == 0 {
		return Pagination{}, errors.Wrapf(ErrListParameter, "wrong direction: %s", direction)
	}

	order = strings.ToUpper(order)
	if order == "" {
		order = pb.Order_ORDER_ID.String()
	}
	if v, ok := pb.Order_value[order]; !ok || v == 0 {
		return Pagination{}, errors.Wrapf(ErrListParameter, "wrong order: %s", order)
	}

	return Pagination{
		Limit:     limit,
		Offset:    offset,
		Order:     convert(order),
		Direction: convert(direction),
	}, nil
}

func convert(s string) string {
	parts := strings.Split(s, "_")

	return parts[len(parts)-1]
}
