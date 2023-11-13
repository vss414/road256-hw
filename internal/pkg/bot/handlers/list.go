package handlers

import (
	"context"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

func (h *BotHandlers) list(params string) (string, error) {
	data := strings.Split(params, " ")
	if len(data) > 5 {
		return "", errors.Wrapf(BadArguments, "%d items: <%v>", len(data), data)
	}

	var err error
	var page uint64
	ind := 0
	if len(data) > ind {
		page, err = strconv.ParseUint(data[ind], 0, 64)
		if err != nil {
			return "", errors.Wrapf(BadArguments, "wrong page: %v", data[ind])
		}
	}

	var limit uint64
	ind = 1
	if len(data) > ind {
		limit, err = strconv.ParseUint(data[ind], 0, 64)
		if err != nil {
			return "", errors.Wrapf(BadArguments, "wrong limit: %v", data[ind])
		}
	}

	players, err := h.p.List(context.Background(), limit, page, getParam(data, 2), getParam(data, 3))
	if err != nil {
		return "", err
	}

	if len(players) == 0 {
		return "No players", nil
	}

	res := make([]string, 0, len(players)+1)

	res = append(res, "Player | Club | Games played | Goals | Assists")

	for _, v := range players {
		res = append(res, v.String())
	}

	return strings.Join(res, "\n"), nil
}

func getParam(data []string, i int) string {
	if len(data) > i {
		return data[i]
	}

	return ""
}
