package gohelpers

import "sync"

// ParallelIterator Iterates a slice in parallel
func ParallelIterator(length int, exec func(i int) error) error {
	var wg sync.WaitGroup
	var mx sync.Mutex
	var out error
	wg.Add(length)
	for i := 0; i < length; i++ {
		go func(i int) {
			if err := exec(i); err != nil {
				mx.Lock()
				out = err
				mx.Unlock()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return out
}
