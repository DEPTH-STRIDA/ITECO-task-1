package domain

// Experiment функция эксперимента. Принимает на вход количество попыток, студента и коробки.
// Возвращает true, если студент нашел свой номер в коробках
type Experiment func(repeat int, student Storager, boxes map[int]Storager) bool
