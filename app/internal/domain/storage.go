package domain

// Storager интерфейс для хранения и получения номера
type Storager interface {
	Get() int
	Put(int)
}
