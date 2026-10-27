package compute_node

import (
	"errors"
	model "practice_01/model"
)

var ErrIncorrectComputerCPUorRAM = errors.New("Incorrect computer node parameters: CPU and RAM fields must be positive")

type ComputeNode struct {
	model.BaseResource
	CPU    int
	RAM    int
	HasGPU bool
}

func NewComputeNode(id, createdAt, cpu, ram int, name, region string, hasGPU bool) (ComputeNode, error) {
	// Валидация входных данных...
	if cpu <= 0 || ram <= 0 {
		return ComputeNode{}, ErrIncorrectComputerCPUorRAM
	}

	br, err := model.NewBaseResource(id, createdAt, name, region)
	if err != nil {
		return ComputeNode{}, err
	}

	return ComputeNode{
		// Вот здесь: BaseResource — это ключ поля!
		/*BaseResource: model.BaseResource{
			ID:        1,
			Name:      "worker-node-1",
			Region:    "eu-west-1",
			CreatedAt: 1700000000,
		},*/
		BaseResource: br,
		CPU:          cpu,
		RAM:          ram,
		HasGPU:       hasGPU,
	}, nil
}
