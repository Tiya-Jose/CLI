package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var timeout = 5 * time.Minute

// Collection contains the mongo Collection to perform operations on a given collection
type Collection struct {
	*mongo.Collection      //Collection is a handle to a MongoDB collection.
	flag              bool // To log errors
}

// NewCollection returns a new mongo collection
func (c Client) NewCollection(db, collection string, flag bool) Collection {
	return Collection{
		c.Database(db).Collection(collection), flag,
	}
}

// Find finds documents based on a query condition.
// The data parameter must be a pointer to a slice, otherwise it will PANIC
func (c Collection) Find(query, projection, data interface{}) (err error) {

	opts := options.Find().SetProjection(projection)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cur, err := c.Collection.Find(ctx, query, opts)
	if err != nil {
		log.Println(err)
		return
	}
	if err = cur.All(ctx, data); err != nil {
		log.Println(err)
		logFindError(c, query, projection, data)
		return
	}
	return
}

// InsertOne creates a new document.
func (c Collection) InsertOne(data interface{}) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	res, err := c.Collection.InsertOne(ctx, data)
	if err != nil {
		log.Println(err)
		logResult(c, nil, data, res)
		return
	}
	return
}

// DeleteOne deletes a single document from the collection.
func (c Collection) DeleteOne(query interface{}) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	res, err := c.Collection.DeleteOne(ctx, query)
	if err != nil {
		log.Println(err)
		logResult(c, query, nil, res)
		return
	}
	return
}
