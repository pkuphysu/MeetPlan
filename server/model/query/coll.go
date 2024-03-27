package query

import (
	"context"
	"errors"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func init() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://meetplan:wyHT75Pm285hdpLi@192.168.2.5:27277/meetplan"))
	if err != nil {
		panic(err)
	}
	db = client.Database("meetplan")
}

type Coll[T any] struct {
	CollectionName string
}

func New[T any](name string) *Coll[T] {
	return &Coll[T]{CollectionName: name}
}

func (c *Coll[T]) Raw() *mongo.Collection {
	return db.Collection(c.CollectionName)
}

func (c *Coll[T]) FindAll(ctx context.Context, filter bson.M) ([]*T, error) {
	return c.FindOffset(ctx, filter, -1, false)
}

func (c *Coll[T]) FindOffset(ctx context.Context, filter bson.M, limit int, asc bool) ([]*T, error) {
	if limit == 0 {
		return nil, nil
	}

	opt := options.Find()
	if limit > 0 {
		opt.SetLimit(int64(limit))
	}
	if asc {
		opt.SetSort(bson.D{{"_id", 1}})
	} else {
		opt.SetSort(bson.D{{"_id", -1}})
	}

	res, err := c.Raw().Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}
	if res.Err() != nil {
		return nil, res.Err()
	}
	var result []*T
	err = res.All(ctx, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Coll[T]) FindOne(ctx context.Context, filter bson.M) (*T, error) {
	res := c.Raw().FindOne(ctx, filter)
	if res.Err() != nil {
		return nil, res.Err()
	}
	var result T
	err := res.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (c *Coll[T]) FindByID(ctx context.Context, id primitive.ObjectID) (*T, error) {
	return c.FindOne(ctx, bson.M{"_id": id})
}

func (c *Coll[T]) Count(ctx context.Context, filter bson.M) (int, error) {
	if len(filter) == 0 {
		count, err := c.Raw().EstimatedDocumentCount(ctx)
		return int(count), err
	}
	count, err := c.Raw().CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return int(count), nil
}

func (c *Coll[T]) Exists(ctx context.Context, filter bson.M) (bool, error) {
	err := c.Raw().FindOne(ctx, filter, options.FindOne().SetProjection(bson.M{"_id": 1})).Err()
	if errors.Is(err, mongo.ErrNoDocuments) {
		return false, nil
	}
	return err == nil, err
}

func (c *Coll[T]) Upsert(ctx context.Context, doc *T) error {
	id, ok := reflect.ValueOf(doc).Elem().FieldByName("ID").Interface().(primitive.ObjectID)
	if !ok {
		return errors.New("ID field not found")
	}
	if id.IsZero() {
		return errors.New("_id must not be zero")
	}
	_, err := c.Raw().ReplaceOne(ctx, bson.M{"_id": id}, doc, options.Replace().SetUpsert(true))
	return err
}

func (c *Coll[T]) UpsertMany(ctx context.Context, docs []*T) error {
	if len(docs) == 0 {
		return nil
	}
	sess, err := db.Client().StartSession()
	if err != nil {
		return err
	}
	defer sess.EndSession(ctx)
	sessCtx := mongo.NewSessionContext(ctx, sess)
	err = sess.StartTransaction()
	if err != nil {
		return err
	}
	defer func() {
		_ = sess.AbortTransaction(sessCtx)
	}()
	res, err := c.Raw().UpdateMany(sessCtx, docs, options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	if res.MatchedCount+res.UpsertedCount != int64(len(docs)) {
		return errors.New("not all documents are upserted")
	}
	err = sess.CommitTransaction(sessCtx)
	if err != nil {
		return err
	}
	return nil
}
