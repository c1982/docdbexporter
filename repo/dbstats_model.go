package repo

import "go.mongodb.org/mongo-driver/bson/primitive"

type DBStatsOutput struct {
	DB            string              `json:"db" bson:"db"`
	Collections   int                 `json:"collections" bson:"collections"`
	Objects       float64             `json:"objects" bson:"objects"`
	StorageSize   float64             `json:"storageSize" bson:"storageSize"`
	Indexes       int                 `json:"indexes" bson:"indexes"`
	IndexSize     float64             `json:"indexSize" bson:"indexSize"`
	FileSize      float64             `json:"fileSize" bson:"fileSize"`
	Ok            int                 `json:"ok" bson:"ok"`
	OperationTime primitive.Timestamp `json:"operationTime" bson:"operationTime"`
}
