package main

import (
	"fmt"
	"runtime"
)

// Point represents a simple 2D point
type Point struct {
	X, Y float64
}

// largeStruct represents a struct that's likely to escape to heap due to size
type LargeStruct struct {
	Data [1024]int // Large array - likely to escape
	ID   int
}

// stackAllocation demonstrates values that stay on the stack
func stackAllocation() *Point {
	// This will likely stay on stack (escape analysis will show)
	p := Point{X: 1.0, Y: 2.0}
	return &p // Even though we return a pointer, escape analysis may optimize this
}

// heapAllocation demonstrates values that escape to heap
func heapAllocation() *Point {
	// This escapes to heap because we store it in a slice
	points := make([]*Point, 1)
	p := &Point{X: 1.0, Y: 2.0}
	points[0] = p // Storing in slice causes escape
	return points[0]
}

// interfaceEscape demonstrates how interfaces cause escapes
func interfaceEscape() interface{} {
	// Any value assigned to interface{} escapes to heap
	p := Point{X: 1.0, Y: 2.0}
	return p // Escapes because of interface conversion
}

// sliceEscape demonstrates how slice operations cause escapes
func sliceEscape() []*Point {
	var points []*Point
	for i := 0; i < 10; i++ {
		// Each &Point{} allocation escapes because it's stored in slice
		p := &Point{X: float64(i), Y: float64(i * 2)}
		points = append(points, p)
	}
	return points
}

// noEscape demonstrates a pattern that doesn't escape
func noEscape() Point {
	// Returns by value - no escape
	return Point{X: 1.0, Y: 2.0}
}

// largeStructEscape shows how large structs escape
func largeStructEscape() *LargeStruct {
	// Large structs often escape to heap regardless of usage
	large := LargeStruct{ID: 1}
	return &large
}

// demonstrateEscapePatterns shows different escape scenarios
func demonstrateEscapePatterns() {
	fmt.Println("=== Escape Analysis Patterns ===")

	var stats1, stats2 runtime.MemStats

	// Test stack allocation
	runtime.GC()
	runtime.ReadMemStats(&stats1)
	for i := 0; i < 1000; i++ {
		_ = stackAllocation()
	}
	runtime.ReadMemStats(&stats2)
	fmt.Printf("Stack allocation pattern - Heap allocs: %d\n", stats2.Mallocs-stats1.Mallocs)

	// Test heap allocation
	runtime.GC()
	runtime.ReadMemStats(&stats1)
	for i := 0; i < 1000; i++ {
		_ = heapAllocation()
	}
	runtime.ReadMemStats(&stats2)
	fmt.Printf("Heap allocation pattern - Heap allocs: %d\n", stats2.Mallocs-stats1.Mallocs)

	// Test interface escape
	runtime.GC()
	runtime.ReadMemStats(&stats1)
	for i := 0; i < 1000; i++ {
		_ = interfaceEscape()
	}
	runtime.ReadMemStats(&stats2)
	fmt.Printf("Interface escape pattern - Heap allocs: %d\n", stats2.Mallocs-stats1.Mallocs)

	// Test no escape
	runtime.GC()
	runtime.ReadMemStats(&stats1)
	for i := 0; i < 1000; i++ {
		_ = noEscape()
	}
	runtime.ReadMemStats(&stats2)
	fmt.Printf("No escape pattern - Heap allocs: %d\n", stats2.Mallocs-stats1.Mallocs)
}

// RunEscapeAnalysisDemo demonstrates escape analysis concepts
func RunEscapeAnalysisDemo() {
	fmt.Println("Escape Analysis Demo")
	fmt.Println("====================")

	demonstrateEscapePatterns()

	fmt.Printf("\n=== Common Escape Scenarios ===\n")
	fmt.Println("1. Storing pointer in slice/map -> ESCAPES")
	fmt.Println("2. Returning pointer to local variable -> MAY ESCAPE")
	fmt.Println("3. Assignment to interface{} -> ESCAPES")
	fmt.Println("4. Large structs (>64KB typically) -> ESCAPES")
	fmt.Println("5. Calling unknown functions with pointer -> MAY ESCAPE")
	fmt.Println("6. Goroutine closure capturing variables -> MAY ESCAPE")

	fmt.Printf("\n=== How to Check Escape Analysis ===\n")
	fmt.Println("Run: go build -gcflags=-m")
	fmt.Println("Look for messages like:")
	fmt.Println("  - 'moved to heap' (escapes)")
	fmt.Println("  - 'does not escape' (stays on stack)")
	fmt.Println("  - 'leaking param' (parameter escapes)")

	fmt.Printf("\n=== Optimization Tips ===\n")
	fmt.Println("1. Return values by copy when possible")
	fmt.Println("2. Use fixed-size arrays instead of slices for small data")
	fmt.Println("3. Avoid interface{} when you know the concrete type")
	fmt.Println("4. Pre-allocate slices with known capacity")
	fmt.Println("5. Use sync.Pool for frequently allocated objects")
}

// main function for educational demonstration
func main() {
	fmt.Println("=== Escape Analysis Demo ===")
	fmt.Println("This demonstrates when Go values escape to heap vs stay on stack")
	fmt.Println("Run with: go build -gcflags=-m escape_analysis.go")
	fmt.Println("to see detailed escape analysis")
	fmt.Println()

	RunEscapeAnalysisDemo()
}
