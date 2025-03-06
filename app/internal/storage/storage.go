package storage

import (
	"app/app/internal/domain"
	"math/rand"
)

// Storage хранилище номера
type Storage struct {
	Number int
}

func (s *Storage) Get() int {
	return s.Number
}

func (s *Storage) Put(number int) {
	s.Number = number
}

// GenerateStorages генерирует n хранилищ, где каждое хранилище получит случайный номер от 1 до n.
// Номера не повторяются. Ключи хранилищ идут от 1 до n включтельно
func GenerateStorages(n int) map[int]domain.Storager {
	uniqeNumbers := make(map[int]struct{})
	storages := make(map[int]domain.Storager)

	for i := 1; i <= n; i++ {
		for {
			number := rand.Intn(n) + 1
			if _, ok := uniqeNumbers[number]; !ok {
				uniqeNumbers[number] = struct{}{}
				newStorage := &Storage{Number: number}
				storages[i] = newStorage
				break
			}
		}
	}
	return storages
}
