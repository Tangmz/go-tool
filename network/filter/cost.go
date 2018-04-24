package filter

import "sync"

type Cost struct {
	Used map[string]CostDetail
	lock *sync.Mutex
}

type CostDetail struct {
	Max         int
	Min         int
	Avg         int
	Total       int
	Count       int
	Concurrence int
	PendingNum  int
	Pending     map[string]int64
}

func (c *Cost) Start() {

}

func (c *Cost) Done() {
}
