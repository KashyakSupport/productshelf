package productshelf

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

type mongoDB struct {
	conn *mgo.Session
	c    *mgo.Collection
}

// Ensure mongoDB conforms to the BookDatabase interface.
var _ ProductDatabase = &mongoDB{}

// newMongoDB creates a new BookDatabase backed by a given Mongo server,
// authenticated with given credentials.
func newMongoDB(addr string, cred *mgo.Credential) (ProductDatabase, error) {
	conn, err := mgo.Dial(addr)
	if err != nil {
		return nil, fmt.Errorf("mongo: could not dial: %v", err)
	}

	if cred != nil {
		if err := conn.Login(cred); err != nil {
			return nil, err
		}
	}

	return &mongoDB{
		conn: conn,
		c:    conn.DB("productshelf").C("products"),
	}, nil
}

// Close closes the database.
func (db *mongoDB) Close() {
	db.conn.Close()
}


var maxRand = big.NewInt(1<<63 - 1)

// randomID returns a positive number that fits within an int64.
func randomID() (int64, error) {
	// Get a random number within the range [0, 1<<63-1)
	n, err := rand.Int(rand.Reader, maxRand)
	if err != nil {
		return 0, err
	}
	// Don't assign 0.
	return n.Int64() + 1, nil
}

// AddBook saves a given book, assigning it a new ID.
func (db *mongoDB) AddProduct(b *Product) (id int64, err error) {
	id, err = randomID()
	if err != nil {
		return 0, fmt.Errorf("mongodb: could not assign an new ID: %v", err)
	}

	b.ID = id
	if err := db.c.Insert(b); err != nil {
		return 0, fmt.Errorf("mongodb: could not add product: %v", err)
	}
	return id, nil
}
