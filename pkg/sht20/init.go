package sht20

import (
	"atp-sht20/pkg/utils/domain"
	"context"
	"time"
)

type shtStruct struct {
	setting Setting
}

type Setting struct {
	Port     string
	Baudrate uint
	Timeout  time.Duration
}

func NewRepository(setting Setting) RepositoryI {
	return &shtStruct{
		setting: setting,
	}
}

type RepositoryI interface {
	Read(ctx context.Context, id uint8) (result domain.SHT20, err error)
	Write(ctx context.Context, prev, next uint8) (err error)
}
