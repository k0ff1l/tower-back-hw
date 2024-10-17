package main

import (
	"fmt"
	"os"
)

type Node struct {
	val  int
	next *Node
	prev *Node
}

type Deque struct {
	head *Node
	tail *Node
}

func NewDeque() *Deque {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &Deque{head: head, tail: tail}
}

func (d *Deque) AddFront(val int) {
	newNode := &Node{val: val}
	newNode.next = d.head.next
	newNode.prev = d.head
	d.head.next.prev = newNode
	d.head.next = newNode
}

func (d *Deque) AddBack(val int) {
	newNode := &Node{val: val}
	newNode.prev = d.tail.prev
	newNode.next = d.tail
	d.tail.prev.next = newNode
	d.tail.prev = newNode
}

func (d *Deque) PopFront() (int, error) {
	if d.head.next == d.tail {
		return 0, fmt.Errorf("EMPTY DEQUE")
	}
	val := d.head.next.val
	d.head.next = d.head.next.next
	d.head.next.prev = d.head
	return val, nil
}

func (d *Deque) PopBack() (int, error) {
	if d.head.next == d.tail {
		return 0, fmt.Errorf("EMPTY DEQUE")
	}
	val := d.tail.prev.val
	d.tail.prev = d.tail.prev.prev
	d.tail.prev.next = d.tail
	return val, nil
}

func (d *Deque) IsExist(val int) bool {
	current := d.head.next
	for current != d.tail {
		if current.val == val {
			return true
		}
		current = current.next
	}
	return false
}

// ДАЛЬШЕ ИДЕТ ЧАТ ГЭПЭТЭ ДЛЯ ТЕСТОВ

func assertEqual(actual, expected int, testName string) {
	if actual != expected {
		fmt.Printf("Тест '%s' провален: ожидается %d, получено %d\n", testName, expected, actual)
		os.Exit(1)
	}
}

func assertTrue(actual bool, testName string) {
	if !actual {
		fmt.Printf("Тест '%s' провален: ожидается true, получено false\n", testName)
		os.Exit(1)
	}
}

func assertFalse(actual bool, testName string) {
	if actual {
		fmt.Printf("Тест '%s' провален: ожидается false, получено true\n", testName)
		os.Exit(1)
	}
}

func assertNoError(err error, testName string) {
	if err != nil {
		fmt.Printf("Тест '%s' провален: ошибка: %v\n", testName, err)
		os.Exit(1)
	}
}

func assertError(err error, testName string) {
	if err == nil {
		fmt.Printf("Тест '%s' провален: ожидалась ошибка, но ее нет\n", testName)
		os.Exit(1)
	}
}

func main() {
	deque := NewDeque()

	// Тест 1: Добавление элементов в начало и конец
	deque.AddFront(10)
	deque.AddBack(20)
	deque.AddFront(5)
	deque.AddBack(30)

	// Тест 2: Проверка на существование элементов
	assertTrue(deque.IsExist(10), "Проверка существования 10")
	assertTrue(deque.IsExist(5), "Проверка существования 5")
	assertTrue(deque.IsExist(30), "Проверка существования 30")
	assertFalse(deque.IsExist(100), "Проверка отсутствия 100")

	// Тест 3: Удаление элементов с начала и конца
	val, err := deque.PopFront()
	assertNoError(err, "PopFront (не должно быть ошибки)")
	assertEqual(val, 5, "PopFront (значение)")

	val, err = deque.PopBack()
	assertNoError(err, "PopBack (не должно быть ошибки)")
	assertEqual(val, 30, "PopBack (значение)")

	// Тест 4: Попытка удаления из пустого deque
	deque = NewDeque()
	_, err = deque.PopFront()
	assertError(err, "PopFront на пустом deque")

	_, err = deque.PopBack()
	assertError(err, "PopBack на пустом deque")

	// Тест 5: Проверка работы после очистки
	deque.AddBack(15)
	val, err = deque.PopFront()
	assertNoError(err, "PopFront после очистки")
	assertEqual(val, 15, "PopFront после очистки (значение)")

	// Тест 6: Добавление и удаление одного элемента
	deque.AddFront(42)
	val, err = deque.PopBack()
	assertNoError(err, "PopBack с одним элементом")
	assertEqual(val, 42, "PopBack с одним элементом (значение)")

	// Если дошли до сюда, то все тесты прошли успешно
	fmt.Println("Все тесты успешно пройдены.")

	TreeTest()
}
