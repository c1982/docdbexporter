package main

import (
	"docdbexporter/exporter"
	"docdbexporter/repo"
	"fmt"
	"net/url"
	"os"
	"time"
)

var (
	docdb           repo.Repoer
	metrics         *exporter.PrometheusExporter
	interval        time.Duration
	defaultInterval = "1m"
)

func main() {
	uri := os.Getenv("DOCDB_EXPORTER_MONGODB_URI")
	database := os.Getenv("DOCDB_EXPORTER_MONGODB_DBNAME")
	metricPrefix := os.Getenv("DOCDB_EXPORTER_METRIC_PREFIX")
	httpAddr := os.Getenv("DOCDB_EXPORTER_HTTP_ADDR")
	collectInterval := os.Getenv("DOCDB_EXPORTER_COLLECT_INTERVAL")

	if uri == "" {
		fmt.Println("DOCDB_EXPORTER_MONGODB_URI environment variable is not set")
		return
	}
	if database == "" {
		fmt.Println("DOCDB_EXPORTER_MONGODB_DBNAME environment variable is not set")
		return
	}
	if metricPrefix == "" {
		fmt.Println("DOCDB_EXPORTER_METRIC_PREFIX environment variable is not set")
		return
	}
	if httpAddr == "" {
		httpAddr = ":8080"
	}
	if collectInterval == "" {
		collectInterval = defaultInterval
	}

	interval, err := time.ParseDuration(collectInterval)
	if err != nil {
		fmt.Printf("Error parsing DOCDB_EXPORTER_COLLECT_INTERVAL: %v\n", err)
		return
	}

	fmt.Printf("Connecting to MongoDB at %s\n", getHost(uri))
	fmt.Printf("Using database: %s\n", database)
	fmt.Printf("Using metric prefix: %s\n", metricPrefix)
	fmt.Printf("Using HTTP address: %s\n", httpAddr)
	fmt.Printf("Using collection interval: %s\n", collectInterval)
	fmt.Println("Starting metrics exporter...")

	metrics = exporter.NewPrometheusExporter(metricPrefix)
	docdb = repo.NewDocdb(uri, database)
	err = docdb.Connect()
	if err != nil {
		panic(err)
	}

	go metrics.CollectMetricsPeriodically(interval, collect)

	fmt.Println("Connected to MongoDB")
	metrics.ListenAndServe(httpAddr)
}

func collect() {
	dbStats()
	collectionStats()
	collectTop()
	fmt.Printf("Metrics collected %s\n", time.Now().Format(time.RFC3339))
}

func dbStats() {
	stats, err := docdb.DatabaseStat()
	if err != nil {
		fmt.Printf("Error getting database stats: %v\n", err)
		return
	}

	exporter.DatabaseNameMetric.Set(1, stats.DB)
	exporter.DatabaseCollectionCountMetric.Set(float64(stats.Collections), stats.DB)
	exporter.DatabaseObjectsCountMetric.Set(stats.Objects, stats.DB)
	exporter.DatabaseStorageSizeMetric.Set(stats.StorageSize, stats.DB)
	exporter.DatabaseIndexesCountMetric.Set(float64(stats.Indexes), stats.DB)
	exporter.DatabaseIndexSizeMetric.Set(stats.IndexSize, stats.DB)
	exporter.DatabaseFileSizeMetric.Set(stats.FileSize, stats.DB)
}

func collectionStats() {
	names, err := docdb.CollectionNames()
	if err != nil {
		fmt.Printf("Error getting collection names: %v\n", err)
		return
	}

	if len(names) == 0 {
		fmt.Println("No collections found")
		return
	}

	for _, name := range names {
		stats, err := docdb.CollectionStat(name)
		if err != nil {
			fmt.Printf("error getting collection stats: %v\n", err)
			continue
		}
		var isSharded float64

		exporter.DocumentCountMetric.Set(float64(stats.Count), name)
		exporter.DocumentSizeMetric.Set(stats.Size, name)
		exporter.AvgObjectSizeMetric.Set(stats.AvgObjSize, name)
		exporter.StorageSizeMetric.Set(stats.StorageSize, name)
		exporter.IndexCountMetric.Set(float64(stats.Nindexes), name)
		exporter.TotalIndexSizeMetric.Set(stats.TotalIndexSize, name)
		exporter.CollScansMetric.Set(float64(stats.CollScans), name)
		exporter.IdxScansMetric.Set(float64(stats.IdxScans), name)

		exporter.OpCountInsertMetric.Set(float64(stats.OpCounter.NumDocsIns), name)
		exporter.OpCountUpdateMetric.Set(float64(stats.OpCounter.NumDocsUpd), name)
		exporter.OpCountDeleteMetric.Set(float64(stats.OpCounter.NumDocsDel), name)

		exporter.CacheHitsMetric.Set(float64(stats.CacheStats.CollBlksHit), name)
		exporter.CacheReadMetric.Set(float64(stats.CacheStats.CollBlksRead), name)
		exporter.CacheRatioMetric.Set(stats.CacheStats.CollHitRatio, name)
		exporter.IndexCacheHistsMetric.Set(float64(stats.CacheStats.IdxBlksHit), name)
		exporter.IndexCacheReadMetric.Set(float64(stats.CacheStats.IdxBlksRead), name)
		exporter.IndexCacheRatioMetric.Set(stats.CacheStats.IdxHitRatio, name)

		if stats.Sharded {
			isSharded = 1
		} else {
			isSharded = 0
		}
		exporter.IsShardedMetric.Set(isSharded, name)

		for shard, shardStats := range stats.Shards {
			exporter.ShardDocumentCountMetric.Set(float64(shardStats.Count), name, shard)
			exporter.ShardDocumentSizeMetric.Set(shardStats.Size, name, shard)
			exporter.ShardAvgObjectSizeMetric.Set(shardStats.AvgObjSize, name, shard)
			exporter.ShardStorageSizeMetric.Set(shardStats.StorageSize, name, shard)
			exporter.ShardCollScansMetric.Set(float64(shardStats.CollScans), name, shard)
			exporter.ShardIdxScansMetric.Set(float64(shardStats.IdxScans), name, shard)
		}

		for index, size := range stats.IndexSizes {
			exporter.IndexSizesMetric.Set(size, name, index)
		}
	}
}

func collectTop() {
	top, err := docdb.Top()
	if err != nil {
		fmt.Printf("Error getting top: %v\n", err)
		return
	}
	for _, shard := range top.Top {
		exporter.ShardUptimeMetric.Set(float64(shard.UptimeInSeconds), shard.ShardName)

		for k, collection := range shard.Totals {
			exporter.ShardCollectionTotalOpTimeMetric.Set(float64(collection.Total.Time), shard.ShardName, k)
			exporter.ShardCollectionTotalOpCountMetric.Set(float64(collection.Total.Count), shard.ShardName, k)
			exporter.ShardCollectionInsertTimeMetric.Set(float64(collection.Insert.Time), shard.ShardName, k)
			exporter.ShardCollectionInsertOpCountMetric.Set(float64(collection.Insert.Count), shard.ShardName, k)
			exporter.ShardCollectionUpdateTimeMetric.Set(float64(collection.Update.Time), shard.ShardName, k)
			exporter.ShardCollectionUpdateOpCountMetric.Set(float64(collection.Update.Count), shard.ShardName, k)
			exporter.ShardCollectionRemoveTimeMetric.Set(float64(collection.Remove.Time), shard.ShardName, k)
			exporter.ShardCollectionRemoveOpCountMetric.Set(float64(collection.Remove.Count), shard.ShardName, k)
			exporter.ShardCollectionQueryTimeMetric.Set(float64(collection.Queries.Time), shard.ShardName, k)
			exporter.ShardCollectionQueryOpCountMetric.Set(float64(collection.Queries.Count), shard.ShardName, k)
			exporter.ShardCollectionGetMoreTimeMetric.Set(float64(collection.Getmore.Time), shard.ShardName, k)
			exporter.ShardCollectionGetMoreOpCountMetric.Set(float64(collection.Getmore.Count), shard.ShardName, k)
			exporter.ShardCollectionCommandTimeMetric.Set(float64(collection.Commands.Time), shard.ShardName, k)
			exporter.ShardCollectionCommandOpCountMetric.Set(float64(collection.Commands.Count), shard.ShardName, k)
		}
	}
}
func getHost(uri string) string {
	parsedURI, err := url.Parse(uri)
	if err != nil {
		return ""
	}

	host := parsedURI.Hostname()
	if host == "" {
		return ""
	}

	return host
}
