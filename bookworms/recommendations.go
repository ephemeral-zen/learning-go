package main

type bookRecommendations map[Book]map[Book]uint

func listOtherBooksOnShelves(indx int, bookworms []Bookworm) []Book {
	var books []Book
	for i, bookworm := range bookworms {
		if i != indx {
			for _, book := range bookworm.Books {
				books = append(books, book)
			}
		}
	}
	return books
}

func registerBookRecommendations(similarbooks bookRecommendations, book Book, otherbooks []Book) {

}

func recommendOtherBooks(bookworms []Bookworm) []Bookworm {
	sb := make(bookRecommendations)

	for i, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			otherBooksOnShelves := listOtherBooksOnShelves(i, bookworms)
			registerBookRecommendations(sb, book, otherBooksOnShelves)
		}
	}
	recommendations := make([]Bookworm, len(bookworms))
	for i, bookworm := range bookworms {
		recommendations[i] = Bookworm{
			Name:  bookworm.Name,
			Books: recommendBooks(sb, bookworm.Books),
		}
	}
	return recommendations
}

type set map[Book]struct{}

func (s set) Contains(b Book) bool {
	_, ok := s[b]
	return ok
}
