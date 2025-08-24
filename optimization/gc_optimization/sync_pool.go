package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Worker represents a worker object that's expensive to create
type Worker struct {
	ID     int
	Buffer []byte
	Data   map[string]int
}

// NewWorker creates a new worker with allocated resources
func NewWorker(id int) *Worker {
	return &Worker{
		ID:     id,
		Buffer: make([]byte, 1024), // 1KB buffer
		Data:   make(map[string]int),
	}
}

// Reset resets the worker for reuse
func (w *Worker) Reset() {
	w.ID = 0
	// Clear buffer without reallocating
	for i := range w.Buffer {
		w.Buffer[i] = 0
	}
	// Clear map without reallocating
	for k := range w.Data {
		delete(w.Data, k)
	}
}

// DoWork simulates some work that modifies the worker
func (w *Worker) DoWork(taskID int) string {
	w.ID = taskID
	w.Data["processed"] = taskID
	copy(w.Buffer, fmt.Sprintf("task_%d_data", taskID))
	return fmt.Sprintf("Completed task %d", taskID)
}

// Global pool for worker objects
var workerPool = sync.Pool{
	New: func() interface{} {
		return NewWorker(0)
	},
}

// demonstrateWithoutPool shows the impact of creating objects without pooling
func demonstrateWithoutPool(tasks int) {
	fmt.Printf("=== Without Object Pooling ===\n")

	var stats1 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&stats1)

	start := time.Now()

	var results []string
	for i := 0; i < tasks; i++ {
		// Create new worker for each task - causes many heap allocations
		worker := NewWorker(i)
		result := worker.DoWork(i)
		results = append(results, result)
		// Worker becomes eligible for GC immediately after use
	}

	duration := time.Since(start)

	var stats2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&stats2)

	fmt.Printf("Tasks processed: %d\n", tasks)
	fmt.Printf("Duration: %v\n", duration)
	fmt.Printf("Heap allocations: %d\n", stats2.Mallocs-stats1.Mallocs)
	fmt.Printf("Heap memory allocated: %d bytes\n", stats2.TotalAlloc-stats1.TotalAlloc)
	fmt.Printf("GC cycles: %d\n", stats2.NumGC-stats1.NumGC)

	// Keep reference to prevent compiler optimization
	_ = results[len(results)-1]
}

// demonstrateWithPool shows the benefits of object pooling
func demonstrateWithPool(tasks int) {
	fmt.Printf("\n=== With Object Pooling (sync.Pool) ===\n")

	var stats1 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&stats1)

	start := time.Now()

	var results []string
	for i := 0; i < tasks; i++ {
		// Get worker from pool - reuses existing objects
		worker := workerPool.Get().(*Worker)
		worker.Reset() // Clean up for reuse

		result := worker.DoWork(i)
		results = append(results, result)

		// Return worker to pool for reuse
		workerPool.Put(worker)
	}

	duration := time.Since(start)

	var stats2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&stats2)

	fmt.Printf("Tasks processed: %d\n", tasks)
	fmt.Printf("Duration: %v\n", duration)
	fmt.Printf("Heap allocations: %d\n", stats2.Mallocs-stats1.Mallocs)
	fmt.Printf("Heap memory allocated: %d bytes\n", stats2.TotalAlloc-stats1.TotalAlloc)
	fmt.Printf("GC cycles: %d\n", stats2.NumGC-stats1.NumGC)

	// Keep reference to prevent compiler optimization
	_ = results[len(results)-1]
}

// Buffer pool example for byte slices
var bufferPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 0, 1024) // Pre-allocate 1KB capacity
	},
}

// demonstrateBufferPool shows pooling for byte slices
func demonstrateBufferPool() {
	fmt.Printf("\n=== Buffer Pooling Example ===\n")

	// Get buffer from pool
	buf := bufferPool.Get().([]byte)
	buf = buf[:0] // Reset length but keep capacity

	// Use buffer
	buf = append(buf, "Hello, pooled buffer!"...)
	fmt.Printf("Buffer content: %s\n", string(buf))
	fmt.Printf("Buffer capacity: %d\n", cap(buf))

	// Return to pool for reuse
	bufferPool.Put(buf)

	fmt.Println("Buffer returned to pool for reuse")
}

func RunPoolDemo() {
	tasks := 10000

	fmt.Println("Object Pooling with sync.Pool Demo")
	fmt.Println("===================================")

	demonstrateWithoutPool(tasks)
	demonstrateWithPool(tasks)
	demonstrateBufferPool()

	fmt.Printf("\n=== Key Benefits of sync.Pool ===\n")
	fmt.Println("1. Reduces heap allocations by reusing objects")
	fmt.Println("2. Decreases GC pressure and frequency")
	fmt.Println("3. Improves performance for frequently allocated objects")
	fmt.Println("4. Automatically manages pool size based on usage")
	fmt.Println("5. Thread-safe and works well with goroutines")

	fmt.Printf("\n=== When to Use sync.Pool ===\n")
	fmt.Println("✓ Frequently allocated temporary objects")
	fmt.Println("✓ Objects with expensive initialization")
	fmt.Println("✓ Buffers, parsers, encoders, network connections")
	fmt.Println("✗ Long-lived objects")
	fmt.Println("✗ Objects that hold references to other objects")
}

// main function for educational demonstration
func main() {
	fmt.Println("=== Object Pooling with sync.Pool Demo ===")
	fmt.Println("This demonstrates the benefits of object pooling vs creating new objects")
	fmt.Println("Shows memory allocation reduction and performance improvement")
	fmt.Println()

	RunPoolDemo()
}
