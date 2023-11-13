package handlers

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

func (h *BotHandlers) delete(param string) error {
	if len(param) == 0 {
		return BadArguments
	}

	id, err := strconv.ParseUint(param, 0, 64)
	if err != nil {
		return errors.Wrapf(BadArguments, "id: %v", param)
	}

	if err := h.p.Delete(context.Background(), id); err != nil {
		return err
	}

	fmt.Printf("player %d removed\n", id)

	return nil
}
