package main

import (
	"fmt"
)

func main() {
	var clusterLogs = [][]int{
		{1, 1, 0, 1},
		{1, 0, 1, 1},
		{1, 1, -1, 1},
		{1, 1, 1, 1},
	}

	successCount := 0

ClusterCheck:
	for i := range clusterLogs {
		for j := range clusterLogs[i] {
			switch clusterLogs[i][j] {
			case 1:
				successCount++
			case 0:
				continue
			case -1:
				fmt.Printf("Фатальная ошибка на сервере"+
					"[%d]! Экстренная остановка.\n", i)
				break ClusterCheck
			}
		}
	}
	fmt.Printf("Success count is: %d\n", successCount)
}
