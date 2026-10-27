package model

import "errors"

var ErrIncorrectNextStatus = errors.New("Next status is incorrect")
var ErrIncorrectModelParameters = errors.New("Model parameters are not valid")

type BaseResource struct {
	ID        int
	Name      string
	Region    string
	CreatedAt int
	status    string
}

func NewBaseResource(id, createdAt int, name, region string) (BaseResource, error) {
	if id < 0 || createdAt < 0 {
		return BaseResource{}, ErrIncorrectModelParameters
	}
	if name == "" || region == "" {
		return BaseResource{}, ErrIncorrectModelParameters
	}

	return BaseResource{
		ID:        id,
		Name:      name,
		Region:    region,
		CreatedAt: createdAt,
		status:    "PROVISIONING",
	}, nil
}

func (br *BaseResource) GetStatus() string {
	return br.status
}

func (br *BaseResource) SetStatus(newStatus string) error {
	oldStatus := br.GetStatus()
	if oldStatus == "" {
		return ErrIncorrectNextStatus
	}

	switch oldStatus {
	case "PROVISIONING":
		if newStatus != "RUNNING" {
			return ErrIncorrectNextStatus
		}
	case "RUNNING":
		if newStatus != "STOPPED" && newStatus != "TERMINATED" {
			return ErrIncorrectNextStatus
		}
	case "STOPPED":
		if newStatus != "RUNNING" {
			return ErrIncorrectNextStatus
		}
	default:
		return ErrIncorrectNextStatus
	}

	br.status = newStatus

	return nil
}
