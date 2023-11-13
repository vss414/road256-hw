package handlers

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	playerPkg "gitlab.ozon.dev/vss414/hw-1/internal/pkg/core/player"
)

const (
	helpCmd   = "help"
	listCmd   = "list"
	getCmd    = "get"
	createCmd = "create"
	updateCmd = "update"
	deleteCmd = "delete"
)

var (
	UnknownCommand = errors.New("unknown command")
	BadArguments   = errors.New("bad arguments")
)

type IHandler interface {
	Call(cmd string, params string) (string, error)
	help() string
	list(string) (string, error)
	get(string) (string, error)
	create(string) error
	update(string) error
	delete(string) error
}

type BotHandlers struct {
	p playerPkg.IPlayer
}

func New(pool *pgxpool.Pool) IHandler {
	return &BotHandlers{
		p: playerPkg.New(pool),
	}
}

func (h *BotHandlers) Call(cmd string, params string) (string, error) {
	switch cmd {
	case helpCmd:
		return h.help(), nil
	case listCmd:
		return h.list(params)
	case getCmd:
		return h.get(params)
	case createCmd:
		return "", h.create(params)
	case updateCmd:
		return "", h.update(params)
	case deleteCmd:
		return "", h.delete(params)
	default:
		return "", UnknownCommand
	}
}
