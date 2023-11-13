package handlers

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player/model"
)

func (h *BotHandlers) create(params string) error {
	data := strings.Split(params, " ")

	if len(data) < 2 || len(data) > 5 {
		return errors.Wrapf(BadArguments, "%d items: <%v>", len(data), data)
	}

	games, err := parseInput(data, 2)
	if err != nil {
		return errors.Wrapf(BadArguments, "wrong games number <%v>: %s", data[2], err.Error())
	}
	goals, err := parseInput(data, 3)
	if err != nil {
		return errors.Wrapf(BadArguments, "wrong goals number <%v>: %s", data[2], err.Error())
	}
	assists, err := parseInput(data, 4)
	if err != nil {
		return errors.Wrapf(BadArguments, "wrong assists number <%v>: %s", data[2], err.Error())
	}

	p, err := model.New(data[0], data[1], games, goals, assists)
	if err != nil {
		return err
	}

	if _, err := h.p.Create(context.Background(), &p); err != nil {
		return err
	}

	fmt.Printf("player %s added\n", p.Name)

	return nil
}

func parseInput(data []string, i int) (uint, error) {
	if len(data) > i {
		t, err := strconv.ParseUint(data[i], 0, 64)
		if err != nil {
			return 0, err
		}
		return uint(t), nil
	}

	return 0, nil
}
