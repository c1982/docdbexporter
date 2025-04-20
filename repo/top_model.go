package repo

import "go.mongodb.org/mongo-driver/bson/primitive"

type TopOutput struct {
	Note          string              `bson:"note"`
	Top           []ShardTopStats     `bson:"top"`
	Ok            int                 `bson:"ok"`
	OperationTime primitive.Timestamp `bson:"operationTime"`
}

type ShardTopStats struct {
	ShardName       string                        `bson:"shardName"`
	UptimeInSeconds int64                         `bson:"uptime(in seconds)"`
	Totals          map[string]CollectionTopStats `bson:"totals"`
}

type CollectionTopStats struct {
	Total    OperationStats `bson:"total"`
	Insert   OperationStats `bson:"insert"`
	Queries  OperationStats `bson:"queries"`
	Update   OperationStats `bson:"update"`
	Remove   OperationStats `bson:"remove"`
	Getmore  OperationStats `bson:"getmore"`
	Commands OperationStats `bson:"commands"`
}

type OperationStats struct {
	Time  int64 `bson:"time"`
	Count int64 `bson:"count"`
}
