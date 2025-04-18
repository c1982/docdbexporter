package repo

type Repoer interface {
	Connect() error
	Disconnect() error
	CollectionStat(name string) (CollStatsOutput, error)
	CollectionNames() ([]string, error)
}
