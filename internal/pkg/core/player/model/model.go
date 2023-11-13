package model

import (
	"fmt"
	"github.com/pkg/errors"
)

var ErrValidation = errors.New("invalid data")

type Player struct {
	Id      uint64
	Name    string
	Club    string
	Games   uint
	Goals   uint
	Assists uint
}

func New(name, club string, games, goals, assists uint) (Player, error) {
	p := Player{
		Name:    name,
		Club:    club,
		Games:   games,
		Goals:   goals,
		Assists: assists,
	}

	if err := p.Validate(); err != nil {
		return Player{}, err
	}

	return p, nil
}

func (p *Player) String() string {
	return fmt.Sprintf("%d: %s | %s | %d | %d | %d", p.Id, p.Name, p.Club, p.Games, p.Goals, p.Assists)
}

func (p *Player) Validate() error {
	if len(p.Name) == 0 || len(p.Name) > 30 {
		return errors.Wrap(ErrValidation, "field: [name] length should be between 0 and 30 symbols")
	}

	if len(p.Club) == 0 || len(p.Club) > 30 {
		return errors.Wrap(ErrValidation, "field: [club] length should be between 0 and 30 symbols")
	}

	return nil
}
