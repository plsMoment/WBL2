package pattern

// StrategySort provides an interface for sort algorithms.
type StrategySort interface {
	Sort([]int)
}

// BubbleSort implements bubble sort algorithm.
type BubbleSort struct {
}

// Sort sorts data.
func (s *BubbleSort) Sort(a []int) {
	_ = a
}

// InsertionSort implements insertion sort algorithm.
type InsertionSort struct {
}

// Sort sorts data.
func (s *InsertionSort) Sort(a []int) {
	_ = a
}

// Context provides a context for execution of a strategy.
type Context struct {
	strategy StrategySort
}

// Algorithm replaces strategies.
func (c *Context) Algorithm(a StrategySort) {
	c.strategy = a
}

// Sort sorts data according to the chosen strategy.
func (c *Context) Sort(s []int) {
	c.strategy.Sort(s)
}
