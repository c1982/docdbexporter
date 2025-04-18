package exporter

import "github.com/prometheus/client_golang/prometheus"

var (
	DocumentCountMetric      *GaugeMetric
	DocumentSizeMetric       *GaugeMetric
	AvgObjectSizeMetric      *GaugeMetric
	StorageSizeMetric        *GaugeMetric
	IndexCountMetric         *GaugeMetric
	TotalIndexSizeMetric     *GaugeMetric
	IndexSizesMetric         *GaugeMetric
	CollScansMetric          *GaugeMetric
	IdxScansMetric           *GaugeMetric
	OpCountInsertMetric      *GaugeMetric
	OpCountUpdateMetric      *GaugeMetric
	OpCountDeleteMetric      *GaugeMetric
	CacheHitsMetric          *GaugeMetric
	CacheReadMetric          *GaugeMetric
	CacheRatioMetric         *GaugeMetric
	IndexCacheHistsMetric    *GaugeMetric
	IndexCacheReadMetric     *GaugeMetric
	IndexCacheRatioMetric    *GaugeMetric
	IsShardedMetric          *GaugeMetric
	ShardDocumentCountMetric *GaugeMetric
	ShardDocumentSizeMetric  *GaugeMetric
	ShardAvgObjectSizeMetric *GaugeMetric
	ShardStorageSizeMetric   *GaugeMetric
	ShardCollScansMetric     *GaugeMetric
	ShardIdxScansMetric      *GaugeMetric
)

type GaugeMetric struct {
	gauge prometheus.GaugeVec
}

func (g *GaugeMetric) Set(value float64, labels ...string) {
	g.gauge.WithLabelValues(labels...).Set(value)
}

func registerMetrics(metrics *PrometheusExporter) {
	DocumentCountMetric = metrics.registerGaugeWithLabels("collection_document_count", "Count of documents in the collection", []string{"name"})
	DocumentSizeMetric = metrics.registerGaugeWithLabels("collection_document_size", "Size of documents in the collection", []string{"name"})
	AvgObjectSizeMetric = metrics.registerGaugeWithLabels("collection_avg_object_size", "Average size of documents in the collection", []string{"name"})
	StorageSizeMetric = metrics.registerGaugeWithLabels("collection_storage_size", "Storage size of the collection", []string{"name"})
	IndexCountMetric = metrics.registerGaugeWithLabels("collection_index_count", "Number of indexes in the collection", []string{"name"})
	TotalIndexSizeMetric = metrics.registerGaugeWithLabels("collection_total_index_size", "Total size of indexes in the collection", []string{"name"})
	IndexSizesMetric = metrics.registerGaugeWithLabels("collection_index_sizes", "Size of each index in the collection", []string{"name", "index"})
	CollScansMetric = metrics.registerGaugeWithLabels("collection_collscans", "Number of collection scans", []string{"name"})
	IdxScansMetric = metrics.registerGaugeWithLabels("collection_idx_scans", "Number of index scans", []string{"name"})
	OpCountInsertMetric = metrics.registerGaugeWithLabels("collection_op_count_insert", "Number of insert operations", []string{"name"})
	OpCountUpdateMetric = metrics.registerGaugeWithLabels("collection_op_count_update", "Number of update operations", []string{"name"})
	OpCountDeleteMetric = metrics.registerGaugeWithLabels("collection_op_count_delete", "Number of delete operations", []string{"name"})
	CacheHitsMetric = metrics.registerGaugeWithLabels("collection_cache_hits", "Number of cache hits", []string{"name"})
	CacheReadMetric = metrics.registerGaugeWithLabels("collection_cache_read", "Number of cache reads", []string{"name"})
	CacheRatioMetric = metrics.registerGaugeWithLabels("collection_cache_ratio", "Cache hit ratio", []string{"name"})
	IndexCacheHistsMetric = metrics.registerGaugeWithLabels("collection_index_cache_hits", "Number of index cache hits", []string{"name"})
	IndexCacheReadMetric = metrics.registerGaugeWithLabels("collection_index_cache_read", "Number of index cache reads", []string{"name"})
	IndexCacheRatioMetric = metrics.registerGaugeWithLabels("collection_index_cache_ratio", "Index cache hit ratio", []string{"name"})
	IsShardedMetric = metrics.registerGaugeWithLabels("collection_sharded", "Is the collection sharded", []string{"name"})
	ShardDocumentCountMetric = metrics.registerGaugeWithLabels("collection_shard_document_count", "Shards of the collection", []string{"name", "shard"})
	ShardDocumentSizeMetric = metrics.registerGaugeWithLabels("collection_shard_document_size", "Size of the collection shards", []string{"name", "shard"})
	ShardAvgObjectSizeMetric = metrics.registerGaugeWithLabels("collection_shard_avg_object_size", "Average size of the collection shards", []string{"name", "shard"})
	ShardStorageSizeMetric = metrics.registerGaugeWithLabels("collection_shard_storage_size", "Storage size of the collection shards", []string{"name", "shard"})
	ShardCollScansMetric = metrics.registerGaugeWithLabels("collection_shard_collscans", "Number of collection scans in the shards", []string{"name", "shard"})
	ShardIdxScansMetric = metrics.registerGaugeWithLabels("collection_shard_idx_scans", "Number of index scans in the shards", []string{"name", "shard"})
}
