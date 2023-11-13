package handlers

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
)

func (h *BotHandlers) get(param string) (string, error) {
	if len(param) == 0 {
		return "", BadArguments
	}

	id, err := strconv.ParseUint(param, 0, 64)
	if err != nil {
		return "", errors.Wrapf(BadArguments, "id: %v", param)
	}

	p, err := h.p.Get(context.Background(), id)
	if err != nil {
		return "", err
	}

	return p.String(), nil
}
