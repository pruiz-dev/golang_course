package storage_volume

import (
	"errors"
	model "practice_01/model"
)

var ErrIncorrectStorageVolumeOrType = errors.New("Incorrect storage volume parameters: size must be positive or incorrect type")

type StorageVolume struct {
	model.BaseResource
	Volume int
	Type   string
}

func NewStorageVolume(id, createdAt, volume int, name, region, sType string) (StorageVolume, error) {
	if volume <= 0 || !(sType == "SSD" || sType == "HDD") {
		return StorageVolume{}, ErrIncorrectStorageVolumeOrType
	}

	br, err := model.NewBaseResource(id, createdAt, name, region)
	if err != nil {
		return StorageVolume{}, err
	}

	return StorageVolume{
		BaseResource: br,
		Volume:       volume,
		Type:         sType,
	}, nil
}
