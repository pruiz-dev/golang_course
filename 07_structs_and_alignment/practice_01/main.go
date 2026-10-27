package main

import (
	"fmt"
	"practice_01/compute_node"
	"practice_01/storage_volume"
)

func main() {
	// 1. Создаем несколько виртуалок
	node1, err := compute_node.NewComputeNode(1, 1, 8, 8, "name1", "region1", false)
	if err != nil {
		fmt.Println("Ошибка создания node1:", err)
		return
	}

	node2, err := compute_node.NewComputeNode(2, 2, 12, 12, "name2", "region2", true)
	if err != nil {
		fmt.Println("Ошибка создания node2:", err)
		return
	}

	// 2. Создаем сетевые диски
	vol1, err := storage_volume.NewStorageVolume(3, 3, 20, "name3", "region3", "SSD")
	if err != nil {
		fmt.Println("Ошибка создания vol1:", err)
		return
	}

	// 3. Для демонстрации валидации: пробуем создать некорректный ресурс
	_, errInvalid := compute_node.NewComputeNode(-1, 1, 8, 8, "name1", "region1", false)
	if errInvalid != nil {
		fmt.Println("Успешно перехвачена ошибка валидации:", errInvalid)
	}

	// Прямой доступ к полям базовой структуры через node1 (promoted fields):
	fmt.Printf("Узел: %s, Регион: %s, Статус: %s, CPU: %d\n",
		node1.Name, node1.Region, node1.GetStatus(), node1.CPU)

	var totalCPU, totalRAM int
	nodes := []compute_node.ComputeNode{node1, node2}
	volumes := []storage_volume.StorageVolume{vol1}

	for i := range nodes {
		totalCPU += nodes[i].CPU
		totalRAM += nodes[i].RAM
	}

	var totalStorage int
	for i := range volumes {
		totalStorage += volumes[i].Volume
	}

	// Вот она — анонимная структура прямо в месте вывода:
	report := struct {
		TotalNodes   int
		TotalVolumes int
		TotalCPU     int
		TotalRAM     int
		TotalStorage int
	}{
		TotalNodes:   len(nodes),
		TotalVolumes: len(volumes),
		TotalCPU:     totalCPU,
		TotalRAM:     totalRAM,
		TotalStorage: totalStorage,
	}

	fmt.Println("\n--- Сводка инфраструктуры ---")
	fmt.Printf("Всего узлов: %d (CPU: %d, RAM: %d GB)\n", report.TotalNodes, report.TotalCPU, report.TotalRAM)
	fmt.Printf("Всего дисков: %d (Объем: %d GB)\n", report.TotalVolumes, report.TotalStorage)
}
