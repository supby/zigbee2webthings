package db

import "context"

type DB interface {
	GetDevices(ctx context.Context) (error, []Device)
	AddDevice(ctx context.Context, device Device) error
	DeleteDevice(ctx context.Context, ieeeAddress uint64) error
	UpdateDevice(ctx context.Context, device Device) error
}
