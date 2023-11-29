package tasks

// Задание:
// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action
// от родительской структуры Human (аналог наследования).

// Определеям тип Human
type Human struct {
	Tall   int64
	Weight float64
}

// Определяем метод для Human
func (h *Human) AddPoint(a int) int{
	a++
	return a
}

// Делаем структуру Action, которая встраивает методы Human
type Action struct {
	Human
	Name string
}

func Test1() bool {
	// Создадим объект класса Action, и проверим метод из Human
	a := Action{
		Human: Human{130, 13.3},
		Name: "TestTestov",
	}
	temp := 1
	return (temp + 1 == a.AddPoint(temp))
}