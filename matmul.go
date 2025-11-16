package main

import (
	"runtime"
	"sync"
)

// Optimized for Intel i5 7th gen (Kaby Lake)
// L1 cache: 32KB per core
// L2 cache: 256KB per core
// Block size 32: 32×32×4 bytes = 4KB (fits well in L1)

func MatmulOptimized(a, b [][]int) [][]int {
	n := len(a)
	c := make([][]int, n)
	for i := range c {
		c[i] = make([]int, n)
	}

	bt := make([][]int, n)
	for i := 0; i < n; i++ {
		bt[i] = make([]int, n)
		for j := 0; j < n; j++ {
			bt[i][j] = b[j][i]
		}
	}

	numWorkers := runtime.NumCPU()
	runtime.GOMAXPROCS(numWorkers)

	const block = 32

	var wg sync.WaitGroup
	rowsPerWorker := n / numWorkers
	if rowsPerWorker == 0 {
		rowsPerWorker = n
	}

	for w := 0; w < numWorkers; w++ {
		rstart := w * rowsPerWorker
		rend := rstart + rowsPerWorker
		if w == numWorkers-1 {
			rend = n
		}

		wg.Add(1)
		go func(rstart, rend int) {
			defer wg.Done()

			for ii := rstart; ii < rend; ii += block {
				iimax := ii + block
				if iimax > rend {
					iimax = rend
				}

				for kk := 0; kk < n; kk += block {
					kkmax := kk + block
					if kkmax > n {
						kkmax = n
					}

					for jj := 0; jj < n; jj += block {
						jjmax := jj + block
						if jjmax > n {
							jjmax = n
						}

						for i := ii; i < iimax; i++ {
							ci := c[i]
							ai := a[i]

							for k := kk; k < kkmax; k++ {
								aik := ai[k]
								btk := bt[k]

								j := jj
								for ; j < jjmax-3; j += 4 {
									ci[j] += aik * btk[j]
									ci[j+1] += aik * btk[j+1]
									ci[j+2] += aik * btk[j+2]
									ci[j+3] += aik * btk[j+3]
								}

								for ; j < jjmax; j++ {
									ci[j] += aik * btk[j]
								}
							}
						}
					}
				}
			}
		}(rstart, rend)
	}

	wg.Wait()
	return c
}
