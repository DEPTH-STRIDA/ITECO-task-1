package main

import (
	"app/app/internal/experiment"
	"app/app/internal/storage"
	"fmt"
)

func main() {
	// Количество студентов и коробок
	n := 50
	// Количество коробок, которые может открыть студент
	repeats := 25
	// Количество экспериментов
	experiments := 10_000

	successCount := 0

	for i := 0; i < experiments; i++ {
		// Для каждого эксперимента генерируем новый набор
		boxes := storage.GenerateStorages(n)
		students := storage.GenerateStorages(n)

		if experiment.ConductExperiment(repeats, students, boxes, experiment.UseСonsistentTactics) {
			successCount++
		}
	}

	probability := float64(successCount) / float64(experiments)
	fmt.Printf("Вероятность успеха: %.6f%%\n", probability*100)

}
