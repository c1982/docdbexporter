package repo

type Repoer interface {
	Connect() error
	Disconnect() error
	CollectionStat(name string) (CollStatsOutput, error)
	DatabaseStat() (DBStatsOutput, error)
	CollectionNames() ([]string, error)
	Top() (TopOutput, error)
}
