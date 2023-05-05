package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"log"
	"math/rand"
	"os"
	"time"
)

type Node struct {
	data *Commit
	prev *Node
	next *Node
}

type DoubleLinkedList struct {
	head *Node // Начальный элемент в списке
	tail *Node // Последний элемент в списке
	curr *Node // Текущий элемент меняется при использовании методов next, prev
	len  int   // Количество элементов в списке
}

type LinkedLister interface {
	Append(commit *Commit)
	Len() int
	Current() *Node
	Next() *Node
	Prev() *Node
	LoadData(path string) error
	Insert(n int, c Commit) error
	Delete(n int) error
	DeleteCurrent() error
	Index() (int, error)
	Pop() *Node
	Shift() *Node
	SearchUUID(uuID string) *Node
	Search(message string) *Node
	Reverse() *DoubleLinkedList
}

type Commit struct {
	Message string    `json:"message"`
	UUID    string    `json:"uuid"`
	Date    time.Time `json:"date"`
}

var list = &DoubleLinkedList{}

func QuickSort(data []Commit) *DoubleLinkedList {
	if len(data) == 1 {
		list.Append(&data[0])
	} else if len(data) == 0 {
		return list
	} else {
		pivot := data[rand.Intn(len(data))]
		lowerPart := make([]Commit, 0, len(data))
		higherPart := make([]Commit, 0, len(data))

		for _, value := range data {
			switch {
			case value.Date.Before(pivot.Date):
				lowerPart = append(lowerPart, value)
			default:
				higherPart = append(higherPart, value)
			}
		}

		QuickSort(lowerPart)
		QuickSort(higherPart)
	}
	return list
}

func (d *DoubleLinkedList) Append(commit *Commit) {
	if d.len == 0 {
		nodeToAppend := &Node{
			data: commit,
		}
		d.head = nodeToAppend
		d.curr = d.head
		d.tail = d.head
	} else {
		nodeToAppend := &Node{
			data: commit,
			prev: d.tail,
		}
		d.tail.next = nodeToAppend
		d.tail = nodeToAppend
	}
	d.len++
}

// LoadData загрузка данных из подготовленного json файла
func (d *DoubleLinkedList) LoadData(path string) error {
	jsonData, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error while reading a file, %v", err)
	}
	if json.Valid(jsonData) {
		var data []Commit
		if err = json.Unmarshal(jsonData, &data); err != nil {
			return errors.New("error while unmarshalling")
		}
		QuickSort(data)
	} else {
		return errors.New("invalid JSON")
	}
	return nil
}

// Len получение длины списка
func (d *DoubleLinkedList) Len() int {
	return d.len
}

// Current получение текущего элемента
func (d *DoubleLinkedList) Current() *Node {
	return d.curr
}

// Next получение следующего элемента
func (d *DoubleLinkedList) Next() *Node {
	if d.curr == nil {
		return nil
	}
	if d.curr.next != nil {
		d.curr = d.curr.next
		return d.curr
	}
	return nil
}

// Prev получение предыдущего элемента
func (d *DoubleLinkedList) Prev() *Node {
	if d.curr == nil {
		return nil
	}
	if d.curr.prev != nil {
		d.curr = d.curr.prev
		return d.curr
	}
	return nil
}

// Insert вставка элемента после n элемента
func (d *DoubleLinkedList) Insert(n int, c *Commit) error {
	if n >= d.len && d.len != 0 {
		return errors.New("too large n")
	} else if n < 0 {
		return errors.New("index must be higher than 0")
	} else if (n == d.len-1) || n == d.len {
		d.Append(c)
	} else {
		current := d.head
		for i := 0; i < n; i++ {
			current = current.next
		}
		nodeToAppend := &Node{
			data: c,
			prev: current,
			next: current.next,
		}
		current.next.prev = nodeToAppend
		current.next = nodeToAppend
		d.len++
	}
	return nil
}

// Delete удаление n элемента
func (d *DoubleLinkedList) Delete(n int) error {
	if n >= d.len && d.len != 0 {
		return errors.New("too large n")
	} else if n < 0 {
		return errors.New("index must be higher than 0")
	} else if d.len == 0 {
		return errors.New("cannot delete an element from empty list")
	} else if n == d.len-1 {
		if d.len == 1 {
			*d = DoubleLinkedList{}
			d.len = 1
		}
		if d.tail == d.curr {
			d.curr = d.tail.prev
		}
		d.tail.prev.next = nil
		d.tail = d.tail.prev
	} else if n == 0 {
		if d.head == d.curr {
			d.curr = d.head.next
		}
		d.head.next.prev = nil
		d.head = d.head.next
	} else {
		current := d.head
		for i := 0; i < n; i++ {
			current = current.next
		}
		if current == d.curr {
			d.curr = current.prev
		}
		current.prev.next = current.next
		current.next.prev = current.prev
	}
	d.len--
	return nil
}

// DeleteCurrent удаление текущего элемента
func (d *DoubleLinkedList) DeleteCurrent() error {
	if d.len == 0 {
		return errors.New("cannot delete an element from empty list")
	}
	if d.len == 1 && d.curr == d.head {
		*d = DoubleLinkedList{}
		return nil
	} else if d.curr == d.tail {
		return d.Delete(d.len - 1)
	} else if d.curr == d.head {
		return d.Delete(0)
	} else {
		d.curr.next.prev = d.curr.prev
		d.curr.prev.next = d.curr.next
		d.curr = d.curr.prev
		d.len--
		return nil
	}
}

// Index получение индекса текущего элемента
func (d *DoubleLinkedList) Index() (int, error) {
	if d.len == 0 {
		return 0, errors.New("list is empty")
	}
	current := d.head
	index := 0
	for current != nil {
		if current == d.curr {
			break
		}
		current = current.next
		index++
	}
	return index, nil
}

// Pop Операция Pop
func (d *DoubleLinkedList) Pop() *Node {
	if d.len == 0 {
		return &Node{}
	} else if d.len == 1 {
		result := d.head
		*d = DoubleLinkedList{}
		return result
	}
	if d.tail == d.curr {
		d.curr = d.tail.prev
	}
	result := d.tail
	d.tail.prev.next = nil
	d.tail = d.tail.prev
	d.len--
	return result
}

// Shift операция shift
func (d *DoubleLinkedList) Shift() *Node {
	if d.len == 0 {
		return &Node{}
	} else if d.len == 1 {
		return d.head
	}
	tail := d.tail

	d.head.prev = d.tail
	d.tail.next = d.head
	d.tail.prev = nil
	d.head.next = nil

	d.tail = d.head
	d.head = tail
	return tail
}

// SearchUUID поиск коммита по uuid
func (d *DoubleLinkedList) SearchUUID(uuID string) *Node {
	if d.len == 0 {
		return &Node{}
	}
	current := d.head
	for current != nil {
		if current.data.UUID == uuID {
			return current
		}
		current = current.next
	}
	return &Node{}
}

// Search поиск коммита по message
func (d *DoubleLinkedList) Search(message string) *Node {
	if d.len == 0 {
		return &Node{}
	}
	current := d.head
	for current != nil {
		if current.data.Message == message {
			return current
		}
		current = current.next
	}
	return &Node{}
}

// Reverse возвращает перевернутый список
func (d *DoubleLinkedList) Reverse() *DoubleLinkedList {
	result := &DoubleLinkedList{}
	current := d.tail
	for current != nil {
		result.Append(current.data)
		current = current.prev
	}
	return result
}

func GenerateCommit(n int) []Commit {
	result := make([]Commit, 0, n)
	for i := 0; i < n; i++ {
		var commit Commit
		err := gofakeit.Struct(&commit)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, commit)
	}
	return result
}

func GenerateData(n int) *DoubleLinkedList {
	// Дополнительное задание написать генератор данных
	// используя библиотеку gofakeit
	result := &DoubleLinkedList{}
	for i := 0; i < n; i++ {
		var commit Commit
		err := gofakeit.Struct(&commit)
		if err != nil {
			log.Fatal(err)
		}
		result.Append(&commit)
	}
	return result
}
