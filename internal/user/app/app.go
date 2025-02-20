package app

import (
	"github.com/NewGlassbiller/gb-services-insurance/internal/user/domain"
	"gorm.io/gorm"
)

type Query struct {
	db   *gorm.DB
	repo domain.Repository
}

type Command struct {
	repo domain.Repository
}

func NewQuery(
	db *gorm.DB,
	repo domain.Repository,

) *Query {
	return &Query{
		db:   db,
		repo: repo,
	}
}

func NewCommand(
	repo domain.Repository,

) *Command {
	return &Command{
		repo: repo,
	}
}
