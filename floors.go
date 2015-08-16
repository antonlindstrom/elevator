package main

import (
	"sort"
	"sync"
)

type Floors struct {
	data []int
	sync.Mutex
}

func (f *Floors) Len() int {
	return len(f.data)
}

func (f *Floors) Less(i, j int) bool {
	return f.data[i] < f.data[j]
}

func (f *Floors) Swap(i, j int) {
	f.data[i], f.data[j] = f.data[j], f.data[i]
}

func (f *Floors) Add(x int) {
	f.Lock()
	defer f.Unlock()

	f.data = append(f.data, x)
	if len(f.data) > 0 {
		sort.Sort(f)
	}
}

func (f *Floors) PeekHighest() int {
	return (f.data)[len(f.data)-1]
}

func (f *Floors) Highest() int {
	f.Lock()
	defer f.Unlock()

	old := f.data
	n := len(old)

	floor := old[n-1]
	f.data = old[0 : n-1]

	return floor
}

func (f *Floors) PeekLowest() int {
	return (f.data)[0]
}

func (f *Floors) Lowest() int {
	f.Lock()
	defer f.Unlock()

	old := f.data
	n := len(old)

	floor := old[0]
	f.data = old[1:n]

	return floor
}

func (f *Floors) Clear() {
	f.Lock()
	defer f.Unlock()
	f.data = make([]int, 0)
}
