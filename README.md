## MongoDB Exporter for Amazon DocumentDB

This is a Collection Data Exporter for Amazon DocumentDB Elastic Cluster verison. It is designed to scrape metrics from Amazon DocumentDB and expose them in a format that can be consumed by Prometheus.

## Metrics

The exporter collects the following metrics:

| Metric Name | Description |
|------------|-------------|
| docdb_collection_document_count | Count of documents in the collection |
| docdb_collection_document_size | Size of documents in the collection |
| docdb_collection_avg_object_size | Average size of documents in the collection |
| docdb_collection_storage_size | Size of the collection on disk |
| docdb_collection_index_count | Number of indexes in the collection |
| docdb_collection_total_index_size | Total size of indexes in the collection |
| docdb_collection_index_sizes | Size of each index in the collection |
| docdb_collection_collscans | Number of collection scans |
| docdb_collection_idx_scans | Number of index scans |
| docdb_collection_op_count_insert | Number of insert operations |
| docdb_collection_op_count_update | Number of update operations |
| docdb_collection_op_count_delete | Number of delete operations |
| docdb_collection_cache_hits | Number of cache hits |
| docdb_collection_cache_read | Number of cache reads |
| docdb_collection_cache_ratio | Cache hit ratio |
| docdb_collection_index_cache_hits | Number of index cache hits |
| docdb_collection_index_cache_read | Number of index cache reads |
| docdb_collection_index_cache_ratio | Index cache hit ratio |
| docdb_collection_sharded | Is the collection sharded |
| docdb_collection_shard_document_count | Shards of the collection |
| docdb_collection_shard_document_size | Size of documents in the shard |
| docdb_collection_shard_avg_object_size | Average size of documents in the shard |
| docdb_collection_shard_storage_size | Storage size of the collection shards |
| docdb_collection_shard_collscans | Number of collection scans in the shards |
| docdb_collection_shard_idx_scans | Number of index scans in the shards |