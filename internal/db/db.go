package db

import (
	"context"

	"github.com/cockroachdb/pebble"
)

type DeviceDB interface {
	GetDevices(ctx context.Context) (error, []Device)
	AddDevice(ctx context.Context, device Device) error
	DeleteDevice(ctx context.Context, ieeeAddress uint64) error
	UpdateDevice(ctx context.Context, device Device) error
}

func NewDeviceDB(filename string) (DeviceDB, error) {
	db, err := pebble.Open("demo", &pebble.Options{})
	if err != nil {
		return nil, err
	}

	return &deviceDB{}, nil
}

type deviceDB struct {
}

func (d *deviceDB) GetDevices(ctx context.Context) (error, []Device) {
	panic("implement me!")
}

func (d *deviceDB) AddDevice(ctx context.Context, device Device) error {
	panic("implement me!")
}

func (d *deviceDB) DeleteDevice(ctx context.Context, ieeeAddress uint64) error {
	panic("implement me!")
}

func (d *deviceDB) UpdateDevice(ctx context.Context, device Device) error {
	panic("implement me!")
}
