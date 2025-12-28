package smerch

type Smerch struct {
	Locker         *Locker
	Config         *Config
	RepositoryPool *RepositoryPool
}

func NewSmerch(locker *Locker, config *Config, pool *RepositoryPool) *Smerch {
	return &Smerch{
		Locker:         locker,
		Config:         config,
		RepositoryPool: pool,
	}
}
