package store

import (
	"context"
	"fmt"
	"github.com/jiang2084/bookstore/store"
	"github.com/jiang2084/bookstore/store/factory"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoStore struct {
	books *mongo.Collection
}

type Book struct {
	Id      string   `bson:"id"`      // 图书ISBN ID
	Name    string   `bson:"name"`    // 图书名称
	Authors []string `bson:"authors"` // 图书作者
	Press   string   `bson:"press"`   // 出版社
}

func (m *MongoStore) Create(book *store.Book) error {
	insertBook := Book{
		Id:      book.Id,
		Name:    book.Name,
		Authors: book.Authors,
		Press:   book.Press,
	}
	insertResult, err := m.books.InsertOne(context.TODO(), insertBook)
	if err != nil {
		return err
	}
	fmt.Println("insertResult:", insertResult)
	return nil
}

func (m *MongoStore) Update(book *store.Book) error {
	//TODO implement me
	panic("implement me")
}

func (m *MongoStore) Get(id string) (store.Book, error) {
	filter := bson.D{{"id", id}}
	var result Book
	err := m.books.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return store.Book{}, err
	}
	fmt.Println("id", result.Id)
	return store.Book{
		Id:      result.Id,
		Name:    result.Name,
		Authors: result.Authors,
		Press:   result.Press,
	}, nil
}

func (m *MongoStore) GetAll() ([]store.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MongoStore) Delete(s string) error {
	//TODO implement me
	panic("implement me")
}

func init() {
	var (
		client     *mongo.Client
		err        error
		db         *mongo.Database
		collection *mongo.Collection
	)
	//1.建立连接
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://admin:password@localhost:27017").SetConnectTimeout(5*time.Second)); err != nil {
		fmt.Print(err)
		return
	}
	//2.选择数据库 test
	db = client.Database("test")

	//3.选择表 my_collection
	collection = db.Collection("books")
	factory.Register("mongo", &MongoStore{
		books: collection,
	})
}
