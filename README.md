# Optimized Matrix Multiplication in Go

High-performance matrix multiplication implementation optimized for Intel i5 7th generation processors.

## Performance

- **Matrix Size:** 1024×1024
- **Original Code:** 391ms
- **Optimized Code:** 240ms
- **Speedup:** 1.6× (38% faster)

## Key Optimizations

1. **Matrix Transposition** - Eliminates cache misses by converting column access to row access
2. **Cache Blocking** - Processes 32×32 blocks that fit in L1 cache (4KB per block)
3. **Loop Unrolling** - Processes 8 elements per iteration to reduce overhead
4. **Parallel Processing** - Utilizes all CPU cores with goroutines
5. **Optimal Loop Ordering** - i-k-j ordering keeps scalar values in registers

## Usage

```bash
# Run the program
go run *.go

# Or compile first
go build -o matmul *.go
./matmul
```

## Output

```
Time: 240ms
```

Generates three files:
- `matrix_a.txt` - Random input matrix A
- `matrix_b.txt` - Random input matrix B  
- `matrix_result.txt` - Result matrix C = A × B

## Configuration

Adjust matrix size in `main.go`:
```go
const N = 1024  // Change to 512, 2048, etc.
```

## Architecture

```
main.go      - Entry point, timing, file I/O orchestration
matmul.go    - Core multiplication algorithm with optimizations
io.go        - File reading/writing utilities
```

## Requirements

- Go 1.16 or higher
- Multi-core CPU (optimized for Intel i5 7th gen)

## How It Works

The implementation uses **cache-aware blocked matrix multiplication**:

1. Transposes matrix B once (eliminates ~40% of cache misses)
2. Divides work into 32×32 blocks (fits in L1 cache)
3. Distributes blocks across CPU cores
4. Unrolls inner loops 8× (reduces loop overhead)
5. Uses i-k-j loop ordering (better register usage)

## Tuning for Your CPU

For different processors, adjust block size in `matmul.go`:

```go
const block = 32  // Try: 16, 24, 32, 48, 64
```

- **Intel (32KB L1):** block = 32
- **AMD (32KB L1):** block = 32-48
- **Apple M1/M2 (128KB L1):** block = 64-128

## License

MIT

## Author

Optimized for educational purposes and real-world performance.