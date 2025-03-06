package experiment

import (
	"app/app/internal/domain"
	"math/rand"
)

// UseRandomTactics делает repeat попыток вытянуть свой номер из случайных коробок
func UseRandomTactics(repeat int, student domain.Storager, boxes map[int]domain.Storager) bool {
	openedBoxes := make(map[int]struct{})
	stdNumber := student.Get()

	// Делаем repeat попыток вытянуть свой номер из коробок
	for i := 0; i < repeat; i++ {
		for {
			// Выбираем случайную коробку
			boxNumber := rand.Intn(len(boxes)) + 1
			// Если коробка не была открыта, то открываем её
			if _, ok := openedBoxes[boxNumber]; !ok {
				// Помечаем коробку как открытую
				openedBoxes[boxNumber] = struct{}{}
				// Берем коробку
				box, ok := boxes[boxNumber]
				if !ok {
					continue
				}
				// Берем номер из коробки
				boxNumber := box.Get()
				// Если номер совпадает с номером студента, то возвращаем true
				if boxNumber == stdNumber {
					return true
				} else {
					break
				}
			}

		}

	}

	return false
}

// UseСonsistentTactics делает repeat попыток вытянуть свой номер из коробок, используя тактику студентов
func UseСonsistentTactics(repeat int, student domain.Storager, boxes map[int]domain.Storager) bool {
	stdNumber := student.Get()

	currentBoxNumber := stdNumber

	for i := 0; i < repeat; i++ {
		box, ok := boxes[currentBoxNumber]
		if !ok {
			return false
		}

		numberInBox := box.Get()

		if numberInBox == stdNumber {
			return true
		}

		currentBoxNumber = numberInBox
	}

	return false
}

// ConductExperiment проводит эксперимент для каждого студента и возвращает true, если все студенты нашли свои номера
func ConductExperiment(repeat int, students map[int]domain.Storager, boxes map[int]domain.Storager, experiment domain.Experiment) bool {
	// Проводим эксперимент для каждого студента
	for i := 1; i <= len(students); i++ {
		fond := experiment(repeat, students[i], boxes)
		if !fond {
			return false
		}

	}
	return true
}
