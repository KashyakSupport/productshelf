package productshelf

import (
	//"errors"
	"log"
	//"os"

	//"cloud.google.com/go/datastore"
	//"cloud.google.com/go/pubsub"
	//"cloud.google.com/go/storage"

	"gopkg.in/mgo.v2"

	//"github.com/gorilla/sessions"

	//"golang.org/x/net/context"
	//"golang.org/x/oauth2"
	//"golang.org/x/oauth2/google"
)

var (
	DB  ProductDatabase

)

//const PubsubTopicID = "fill-book-details"

func init() {
	var err error

	// To use the in-memory test database, uncomment the next line.
	DB = newMemoryDB()


	 var cred *mgo.Credential
	 DB, err = newMongoDB("mongodb://spiderman:spiderman@ds129153.mlab.com:29153/productshelf", cred)



	if err != nil {
		log.Fatal(err)
	}


}
