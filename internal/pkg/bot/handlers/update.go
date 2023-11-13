package handlers

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func (h *BotHandlers) update(params string) error {
	ctx := context.Background()
	data := strings.Split(params, " ")

	if len(data) < 2 || len(data) > 6 {
		return errors.Wrapf(BadArguments, "%d items: <%v>", len(params), params)
	}

	id, err := strconv.ParseUint(data[0], 0, 64)
	if err != nil {
		return errors.Wrapf(BadArguments, "id: %v", data[0])
	}

	p, err := h.p.Get(ctx, id)
	if err != nil {
		return err
	}

	p.Name = data[1]
	ind := 2
	if len(data) > ind {
		p.Club = data[ind]
	}

	ind = 3
	if len(data) > ind {
		games, err := strconv.ParseUint(data[ind], 0, 64)
		if err != nil {
			return errors.Wrapf(BadArguments, "games: <%v>", data[ind])
		}
		p.Games = uint(games)
	}

	ind = 4
	if len(data) > ind {
		goals, err := strconv.ParseUint(data[ind], 0, 64)
		if err != nil {
			return errors.Wrapf(BadArguments, "goals: <%v>", data[ind])
		}
		p.Goals = uint(goals)
	}

	ind = 5
	if len(data) > ind {
		assists, err := strconv.ParseUint(data[ind], 0, 64)
		if err != nil {
			return errors.Wrapf(BadArguments, "assists: <%v>", data[ind])
		}
		p.Assists = uint(assists)
	}

	if err = p.Validate(); err != nil {
		return err
	}

	if err = h.p.Update(ctx, p); err != nil {
		return err
	}

	fmt.Printf("player %s updated\n", data[0])

	return nil
}
