package store

import (
	mystore "github.com/jiang2084/bookstore/store"
	factory "github.com/jiang2084/bookstore/store/factory"
	"sync"
)

type MemStore struct {
	sync.RWMutex
	books map[string]*mystore.Book
}

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*mystore.Book),
	})
}

func (m *MemStore) Create(book *mystore.Book) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.books[book.Id]; ok {
		return mystore.ErrExist
	}

	nBook := *book
	m.books[book.Id] = &nBook
	return nil
}

func (m *MemStore) Update(book *mystore.Book) error {
	m.Lock()
	defer m.Unlock()

	oldBook, ok := m.books[book.Id]
	if !ok {
		return mystore.ErrNotFound
	}

	nBook := *oldBook
	if book.Name != "" {
		nBook.Name = book.Name
	}

	if book.Authors != nil {
		nBook.Authors = book.Authors
	}

	if book.Press != "" {
		nBook.Press = book.Press
	}

	m.books[book.Id] = &nBook
	return nil
}

func (m *MemStore) Get(id string) (mystore.Book, error) {
	m.Lock()
	defer m.Unlock()

	t, ok := m.books[id]
	if ok {
		return *t, nil
	}
	return mystore.Book{}, mystore.ErrNotFound
}

func (m *MemStore) GetAll() ([]mystore.Book, error) {
	m.Lock()
	defer m.Unlock()

	allBooks := make([]mystore.Book, 0, len(m.books))
	for _, book := range m.books {
		allBooks = append(allBooks, *book)
	}
	return allBooks, nil
}

func (m *MemStore) Delete(id string) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.books[id]; !ok {
		return mystore.ErrNotFound
	}
	delete(m.books, id)
	return nil
}
