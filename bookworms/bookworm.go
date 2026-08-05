package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

func loadBookworms(filePath string) ([]Bookworm, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var bookworms []Bookworm

	err = json.NewDecoder(file).Decode(&bookworms)
	if err != nil {
		return nil, err
	}
	//fmt.Println(bookworms)
	return bookworms, nil
}

func booksCount(bookworms []Bookworm) map[Book]uint {
	//use make to initialize a map
	count := make(map[Book]uint)
	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		}
	}
	return count
}

//func sortBooks(books []Book) []Book {
//	sort.Slice(books, func(i, j int) bool {
//		if books[i].Author != books[j].Author {
//			return books[i].Author < books[j].Author
//		}
//		return books[i].Title < books[j].Title
//	})
//	return books
//}

type ByAuthor []Book

func (b ByAuthor) Len() int { return len(b) }

func (b ByAuthor) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByAuthor) Less(i, j int) bool {
	if b[i].Author != b[j].Author {
		return b[i].Author < b[j].Author
	}
	return b[i].Title < b[j].Title
}

func sortBooks(books []Book) []Book {
	sort.Sort(ByAuthor(books))
	return books
}

func findCommonBooks(bookworms []Bookworm) []Book {
	booksOnShelves := booksCount(bookworms)
	var commonBooks []Book

	for book, count := range booksOnShelves {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}
	fmt.Println(sortBooks(commonBooks))
	return sortBooks(commonBooks)
}
