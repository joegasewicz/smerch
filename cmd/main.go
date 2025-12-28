package main

import (
	"fmt"

	"github.com/joegasewicz/smerch"
)

func main() {
	fmt.Println("here...")
	locker := smerch.NewLocker("examples/smerch.lock", nil)
	config := smerch.NewConfig()
	pool := smerch.NewRepositoryPool()

	app := smerch.NewSmerch(locker, config, pool)
	fmt.Println("Smerch initialised: ", app)

}
