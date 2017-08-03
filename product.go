package productshelf

// Book holds metadata about a book.
type Product struct {
	ID            int64
	Title         string
	Author        string
	PublishedDate string
	Description   string
}

// BookDatabase provides thread-safe access to a database of books.
type ProductDatabase interface {

	// AddBook saves a given book, assigning it a new ID.
	AddProduct(b *Product) (id int64, err error)

	// Close closes the database, freeing up any available resources.
	// TODO(cbro): Close() should return an error.
	Close()
}
