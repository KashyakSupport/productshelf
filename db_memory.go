package productshelf
import (
	//"errors"
	//"fmt"
	//"sort"
	"sync"
)

// Ensure memoryDB conforms to the BookDatabase interface.
var _ ProductDatabase = &memoryDB{}

// memoryDB is a simple in-memory persistence layer for books.
type memoryDB struct {
	mu     sync.Mutex
	nextID int64           // next ID to assign to a book.
	products  map[int64]*Product // maps from Book ID to Book.
}

func newMemoryDB() *memoryDB {
	return &memoryDB{
		products:  make(map[int64]*Product),
		nextID: 1,
	}
}

// Close closes the database.
func (db *memoryDB) Close() {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.products = nil
}


// AddBook saves a given book, assigning it a new ID.
func (db *memoryDB) AddProduct(b *Product) (id int64, err error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	b.ID = db.nextID
	db.products[b.ID] = b

	db.nextID++

	return b.ID, nil
}
