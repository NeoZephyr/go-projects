package store

import (
	"bookstore/store"
	"bookstore/store/factory"
	"sync"
)

func init() {
	factory.Register("mem", &MemStore{
		books: make(map[string]*store.Book),
	})
}

type MemStore struct {
	sync.RWMutex
	books map[string]*store.Book
}

func (m *MemStore) Create(book *store.Book) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.books[book.Id]; ok {
		return store.ErrExist
	}

	newBook := *book
	m.books[book.Id] = &newBook

	return nil
}

func (m *MemStore) Update(book *store.Book) error {
	m.Lock()
	defer m.Unlock()

	oldBook, ok := m.books[book.Id]

	if !ok {
		return store.ErrNotFound
	}

	newBook := *oldBook

	if book.Name != "" {
		newBook.Name = book.Name
	}

	if book.Authors != nil {
		newBook.Authors = book.Authors
	}

	if book.Publisher != "" {
		newBook.Publisher = book.Publisher
	}

	m.books[book.Id] = &newBook
	return nil
}

func (m *MemStore) Get(id string) (store.Book, error) {
	m.RLock()
	defer m.RUnlock()

	book, ok := m.books[id]

	if ok {
		return *book, nil
	}

	return store.Book{}, store.ErrNotFound
}

func (m *MemStore) GetAll() ([]store.Book, error) {
	m.RLock()
	defer m.RUnlock()

	allBooks := make([]store.Book, 0, len(m.books))

	for _, book := range m.books {
		allBooks = append(allBooks, *book)
	}

	return allBooks, nil
}

func (m *MemStore) Delete(id string) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.books[id]; !ok {
		return store.ErrNotFound
	}

	delete(m.books, id)
	return nil
}

