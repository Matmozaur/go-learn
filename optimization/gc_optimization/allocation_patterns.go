package main

import (
	"fmt"
	"runtime"
	"time"
)

// BadStruct represents a poorly optimized struct with heap allocations
type BadStruct struct {
	ID       int
	Name     string
	Data     []byte
	Metadata map[string]interface{}
}

// GoodStruct represents a better optimized struct for stack allocation
type GoodStruct struct {
	ID   int
	Name string
	Data [64]byte // Fixed size array stays on stack
	Flag bool
}

// demonstrateBadAllocation shows the problematic pattern that causes heap thrashing
func demonstrateBadAllocation(iterations int) {
	fmt.Printf("=== Bad Allocation Pattern (Heap Thrashing) ===\n")

	var stats1 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&stats1)

	start := time.Now()

	// This is the problematic pattern - creating new structs in tight loop
	var results []*BadStruct
	for i := 0; i < iterations; i++ {
		// Each struct allocation likely escapes to heap
		badStruct := &BadStruct{
			ID:       i,
			Name:     fmt.Sprintf("item_%d", i),    // String formatting allocates
			Data:     make([]byte, 100),            // Slice allocation
			Metadata: make(map[string]interface{}), // Map allocation
		}
		badStruct.Metadata["index"] = i
		results = append(results, badStruct) // Slice growth causes more allocations
	}

	duration := time.Since(start)

	var stats2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&stats2)

	fmt.Printf("Iterations: %d\n", iterations)
	fmt.Printf("Duration: %v\n", duration)
	fmt.Printf("Heap allocations: %d\n", stats2.Mallocs-stats1.Mallocs)
	fmt.Printf("Heap memory allocated: %d bytes\n", stats2.TotalAlloc-stats1.TotalAlloc)
	fmt.Printf("Final heap size: %d bytes\n", stats2.HeapAlloc)
	fmt.Printf("GC cycles: %d\n", stats2.NumGC-stats1.NumGC)

	// Keep reference to prevent compiler optimization
	_ = results[len(results)-1]
}

// demonstrateGoodAllocation shows optimized pattern that prefers stack allocation
func demonstrateGoodAllocation(iterations int) {
	fmt.Printf("\n=== Good Allocation Pattern (Stack Preferred) ===\n")

	var stats1 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&stats1)

	start := time.Now()

	// Pre-allocate slice to avoid growth
	results := make([]GoodStruct, 0, iterations)

	for i := 0; i < iterations; i++ {
		// Create struct by value (not pointer) - stays on stack if possible
		goodStruct := GoodStruct{
			ID:   i,
			Name: "fixed_name", // Constant string - no allocation
			Flag: i%2 == 0,
		}

		// Copy some data without additional allocation
		copy(goodStruct.Data[:], "some_fixed_data")

		// Append by value, not pointer
		results = append(results, goodStruct)
	}

	duration := time.Since(start)

	var stats2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&stats2)

	fmt.Printf("Iterations: %d\n", iterations)
	fmt.Printf("Duration: %v\n", duration)
	fmt.Printf("Heap allocations: %d\n", stats2.Mallocs-stats1.Mallocs)
	fmt.Printf("Heap memory allocated: %d bytes\n", stats2.TotalAlloc-stats1.TotalAlloc)
	fmt.Printf("Final heap size: %d bytes\n", stats2.HeapAlloc)
	fmt.Printf("GC cycles: %d\n", stats2.NumGC-stats1.NumGC)

	// Keep reference to prevent compiler optimization
	_ = results[len(results)-1]
}

func RunAllocationDemo() {
	iterations := 100000

	fmt.Println("Memory Allocation Patterns Demo")
	fmt.Println("===============================")

	demonstrateBadAllocation(iterations)
	demonstrateGoodAllocation(iterations)

	fmt.Printf("\n=== Key Takeaways ===\n")
	fmt.Println("1. Pointer allocations in loops often escape to heap")
	fmt.Println("2. String formatting and dynamic data structures allocate")
	fmt.Println("3. Value types with fixed sizes prefer stack allocation")
	fmt.Println("4. Pre-allocating slices reduces growth-related allocations")
	fmt.Println("\nRun with: go build -gcflags=-m to see escape analysis")
}

// main function for educational demonstration
func main() {
	fmt.Println("=== Allocation Patterns Demo ===")
	fmt.Println("This demonstrates the difference between good and bad allocation patterns")
	fmt.Println("Run with: go build -gcflags=-m allocation_patterns.go")
	fmt.Println("to see escape analysis output")
	fmt.Println()

	RunAllocationDemo()
}
