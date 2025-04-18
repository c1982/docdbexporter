package repo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Docdb struct {
	uri               string
	dbname            string
	client            *mongo.Client
	db                *mongo.Database
	connectionTimeOut time.Duration
}

func NewDocdb(uri string, dbname string) Repoer {
	return &Docdb{
		uri:               uri,
		dbname:            dbname,
		connectionTimeOut: 6 * time.Second,
	}
}

func (d *Docdb) Connect() error {
	ctx, _ := context.WithTimeout(context.Background(), d.connectionTimeOut)

	clientOptions := options.Client().ApplyURI(d.uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	d.client = client
	d.db = client.Database(d.dbname)

	return nil
}

func (d *Docdb) Disconnect() error {
	err := d.client.Disconnect(context.TODO())
	if err != nil {
		return err
	}
	return nil
}

func (d *Docdb) CollectionStat(name string) (output CollStatsOutput, err error) {
	command := bson.D{{"collStats", name}, {"scale", 1000000}}
	err = d.db.RunCommand(context.TODO(), command).Decode(&output)
	if err != nil {
		return output, err
	}
	return output, err
}

func (d *Docdb) CollectionNames() ([]string, error) {
	collections, err := d.db.ListCollectionNames(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	return collections, nil
}
