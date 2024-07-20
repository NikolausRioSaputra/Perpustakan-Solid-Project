package main

import (
	"Project_Niko/internal/domain"
	"Project_Niko/internal/handler"
	"Project_Niko/internal/repository"
	"Project_Niko/internal/usecase"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bookRepo := repository.NewBookRepository()
	bookUsecase := usecase.NewBookUsecase(bookRepo)
	bookHandler := handler.NewBookHandler(bookUsecase)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n========== Library Management System ==========")
		fmt.Println("1. Manage Book")
		fmt.Println("2. Exit")
		fmt.Println("===============================================")
		fmt.Println("Chose the menu: ")
		optionStr, _ := reader.ReadString('\n')
		option, _ := strconv.Atoi(strings.TrimSpace(optionStr))

		switch option {
		case 1:
			manageBooks(reader, bookHandler)
		case 2:
			fmt.Println("Exit, Thankyou")
			return
		default:
			fmt.Println("Invalid option, please try again brow")
		}
	}
}

func manageBooks(reader *bufio.Reader, bookHandler handler.BookHendlerInterface) {
	for {
		fmt.Println("\n========== Manage Books ==========")
		fmt.Println("1. Add Book")
		fmt.Println("2. Update Book")
		fmt.Println("3. Delete Book")
		fmt.Println("4. List All Books")
		fmt.Println("5. Back to Main Menu")
		fmt.Println("===================================")
		fmt.Print("Choose an option: ")
		optionStr, _ := reader.ReadString('\n')
		option, _ := strconv.Atoi(strings.TrimSpace(optionStr))

		switch option {
		case 1:
			addBook(reader, bookHandler)
		case 2:
			updateBook(reader, bookHandler)
		case 3:
			deleteBook(reader, bookHandler)
		case 4:
			listBooks(bookHandler)
		case 5:
			return
		default:
			fmt.Println("Invalid option. Please try again brow.")
		}
	}
}

func addBook(reader *bufio.Reader, bookHandler handler.BookHendlerInterface) {
	fmt.Print("Enter book ID: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("Invalid book ID: must be a positive integer")
		return
	}

	fmt.Print("Enter book title: ")
	title, _ := reader.ReadString('\n')
	

	fmt.Print("Enter book author: ")
	author, _ := reader.ReadString('\n')

	book := domain.Book{
		ID:     id,
		Title:  strings.TrimSpace(title),
		Author: strings.TrimSpace(author),
	}

	err = bookHandler.StoreNewBook(book)
	if err != nil {
		fmt.Println("Error storing book:", err)
	} else {
		fmt.Println("Book added successfully!")
	}
}

func updateBook(reader *bufio.Reader, bookHandler handler.BookHendlerInterface) {
	fmt.Print("Enter book ID to update: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("Invalid book ID: must be a positive integer")
	}

	fmt.Print("Enter new book title: ")
	title, _ := reader.ReadString('\n')

	fmt.Print("Enter new book author: ")
	author, _ := reader.ReadString('\n')

	book := domain.Book{
		ID:     id,
		Title:  strings.TrimSpace(title),
		Author: strings.TrimSpace(author),
	}

	err = bookHandler.UpdateBook(book)
	if err != nil {
		fmt.Println("Error updating book:", err)
	} else {
		fmt.Println("Book updated successfully!")
	}
}

func deleteBook(reader *bufio.Reader, bookHandler handler.BookHendlerInterface) {
	fmt.Print("Enter book ID to delete: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("Invalid book ID: must be a positive integer")
	}

	book := domain.Book{
		ID: id,
	}

	err = bookHandler.DeleteBook(book)
	if err != nil {
		fmt.Println("Error deleting book:", err)
	} else {
		fmt.Println("Book deleted successfully!")
	}
}

func listBooks(bookHandler handler.BookHendlerInterface) {
	books := bookHandler.ListBooks()
	fmt.Println("\nBooks in Library")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
	}
}
