package repositories

import "book-store/domain/entities"

type BookRepository interface {
	Get(id int) *entities.Book
	List() []*entities.Book
}
