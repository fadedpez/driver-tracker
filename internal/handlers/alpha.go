package handlers

import (
	"context"
	"errors"
	"github.com/fadedpez/driver-tracker/protos"
)

const (
	requiredConfig = "config is required"
)

type Alpha struct {
}

type AlphaConfig struct {
}

func NewAlpha(cfg *AlphaConfig) (*Alpha, error) {
	if cfg == nil {
		return nil, errors.New(requiredConfig)
	}
	return &Alpha{}, nil
}

func (a *Alpha) StoreDriver(ctx context.Context, req *protos.StoreDriverRequest) (*protos.StoreDriverResponse, error) {
	return nil, errors.New("not implemented")
}
