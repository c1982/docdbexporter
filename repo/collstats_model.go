package repo

import "go.mongodb.org/mongo-driver/bson/primitive"

type CollStatsOutput struct {
	Ns             string                `bson:"ns"`
	Count          int64                 `bson:"count"`
	Size           float64               `bson:"size"`
	AvgObjSize     float64               `bson:"avgObjSize"`
	StorageSize    float64               `bson:"storageSize"`
	Capped         bool                  `bson:"capped"`
	Nindexes       int                   `bson:"nindexes"`
	TotalIndexSize float64               `bson:"totalIndexSize"`
	IndexSizes     map[string]float64    `bson:"indexSizes"`
	CollScans      int64                 `bson:"collScans"`
	IdxScans       int64                 `bson:"idxScans"`
	OpCounter      OpCounter             `bson:"opCounter"`
	CacheStats     CacheStats            `bson:"cacheStats"`
	Sharded        bool                  `bson:"sharded"`
	Shards         map[string]ShardStats `bson:"shards"`
	Ok             int                   `bson:"ok"`
	OperationTime  primitive.Timestamp   `bson:"operationTime"`
}

type OpCounter struct {
	NumDocsIns int64 `bson:"numDocsIns"`
	NumDocsUpd int64 `bson:"numDocsUpd"`
	NumDocsDel int64 `bson:"numDocsDel"`
}

type CacheStats struct {
	CollBlksHit  int64   `bson:"collBlksHit"`
	CollBlksRead int64   `bson:"collBlksRead"`
	CollHitRatio float64 `bson:"collHitRatio"`
	IdxBlksHit   int64   `bson:"idxBlksHit"`
	IdxBlksRead  int64   `bson:"idxBlksRead"`
	IdxHitRatio  float64 `bson:"idxHitRatio"`
}

type ShardStats struct {
	Ns             string             `bson:"ns"`
	Count          int64              `bson:"count"`
	Size           float64            `bson:"size"`
	AvgObjSize     float64            `bson:"avgObjSize"`
	StorageSize    float64            `bson:"storageSize"`
	Capped         bool               `bson:"capped"`
	Nindexes       int                `bson:"nindexes"`
	TotalIndexSize float64            `bson:"totalIndexSize"`
	IndexSizes     map[string]float64 `bson:"indexSizes"`
	CollScans      int64              `bson:"collScans"`
	IdxScans       int64              `bson:"idxScans"`
	OpCounter      OpCounter          `bson:"opCounter"`
	CacheStats     CacheStats         `bson:"cacheStats"`
	LastReset      string             `bson:"lastReset"`
}
