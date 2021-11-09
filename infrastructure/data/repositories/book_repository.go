package repositories

import "book-store/domain/entities"

var eramOsDeuses = entities.Book{Id: 1, Title: "Eram os deuses astronautas", Category: "Ficção", Author: "Erich von Däniken", Price: 19.20}
var desposuidos = entities.Book{Id: 2, Title: "Os Desposuídos", Category: "Ficção", Author: "Ursula K. Le Guin", Price: 19.20}
var homemDoCasteloAlto = entities.Book{Id: 3, Title: "O Homem do Castelo Alto", Category: "Romance", Author: "Philip K. Dick", Price: 19.20}
var fundacao = entities.Book{Id: 4, Title: "Fundação", Category: "Ficção", Author: "Isaac Asimov", Price: 19.20}
var ddd = entities.Book{Id: 5, Title: "Domain Driven Design", Category: "Tecnologia", Author: "Eric Evans", Price: 19.20}

var books = []entities.Book{eramOsDeuses, desposuidos, homemDoCasteloAlto, fundacao, ddd}

func List() []entities.Book {
	return books
}

func Get(id int) entities.Book {

	for _, v := range books {
		if v.Id == id {
			return v
		}
	}
	return entities.Book{}
}
