package main

import "fmt"

type summa struct { // структура с переменными
	num1 int
	num2 int
}

type st interface { //интерфейс с функциями
	sum() int
	del() int
	umn() int
}

func (n summa) sum() int { // функция суммы
	return n.num1 + n.num2
}

func (n summa) del() int { // деления
	return n.num1 / n.num2
}

func (n summa) umn() int { // умножения
	return n.num1 * n.num2
}

func main() {
	var i st                                 // переменная
	numb := summa{9, 8}                      // выдача значений
	i = numb                                 // переменная добавляет значения
	fmt.Printf("Cумма чисел: %d\n", i.sum()) // напечатал результат
	fmt.Printf("Cумма чисел: %d\n", i.umn())
	fmt.Printf("Cумма чисел: %f\n", i.del())
}
