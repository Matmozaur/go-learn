package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

// GCStats holds garbage collection statistics
type GCStats struct {
	NumGC       uint32
	PauseTotal  time.Duration
	LastPause   time.Duration
	HeapSize    uint64
	HeapObjects uint64
	HeapAlloc   uint64
	TotalAlloc  uint64
}

// getGCStats captures current GC statistics
func getGCStats() GCStats {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)

	return GCStats{
		NumGC:       stats.NumGC,
		PauseTotal:  time.Duration(stats.PauseTotalNs),
		LastPause:   time.Duration(stats.PauseNs[(stats.NumGC+255)%256]),
		HeapSize:    stats.HeapSys,
		HeapObjects: stats.HeapObjects,
		HeapAlloc:   stats.HeapAlloc,
		TotalAlloc:  stats.TotalAlloc,
	}
}

// printGCStats prints formatted GC statistics
func printGCStats(label string, stats GCStats) {
	fmt.Printf("=== %s ===\n", label)
	fmt.Printf("GC cycles: %d\n", stats.NumGC)
	fmt.Printf("Total GC pause time: %v\n", stats.PauseTotal)
	fmt.Printf("Last GC pause: %v\n", stats.LastPause)
	fmt.Printf("Heap size: %d bytes (%.2f MB)\n", stats.HeapSize, float64(stats.HeapSize)/1024/1024)
	fmt.Printf("Heap objects: %d\n", stats.HeapObjects)
	fmt.Printf("Heap allocated: %d bytes (%.2f MB)\n", stats.HeapAlloc, float64(stats.HeapAlloc)/1024/1024)
	fmt.Printf("Total allocated: %d bytes (%.2f MB)\n", stats.TotalAlloc, float64(stats.TotalAlloc)/1024/1024)
	fmt.Println()
}

// simulateMemoryPressure creates memory pressure to trigger GC
func simulateMemoryPressure() {
	fmt.Println("Creating memory pressure...")

	// Create lots of allocations to trigger GC
	var data [][]byte
	for i := 0; i < 1000; i++ {
		// Allocate 1MB chunks
		chunk := make([]byte, 1024*1024)
		data = append(data, chunk)

		// Trigger GC manually every 100 iterations
		if i%100 == 0 {
			runtime.GC()
		}
	}

	// Keep reference to prevent optimization
	_ = data[len(data)-1]
}

// demonstrateGCTuning shows how to tune GC behavior
func demonstrateGCTuning() {
	fmt.Println("=== GC Tuning Demonstration ===")

	// Show default GC settings
	fmt.Printf("Default GOGC: %d%%\n", debug.SetGCPercent(-1))
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(0))

	// Set more aggressive GC (collects more frequently)
	oldGOGC := debug.SetGCPercent(50) // Collect when heap grows 50% (default is 100%)
	fmt.Printf("Set GOGC to 50%% (more frequent collections)\n")

	stats1 := getGCStats()
	simulateMemoryPressure()
	stats2 := getGCStats()

	fmt.Printf("With GOGC=50%%:\n")
	fmt.Printf("GC cycles: %d\n", stats2.NumGC-stats1.NumGC)
	fmt.Printf("Total pause time: %v\n", stats2.PauseTotal-stats1.PauseTotal)

	// Reset to default
	debug.SetGCPercent(oldGOGC)
	fmt.Printf("Reset GOGC to %d%%\n", oldGOGC)
}

// memoryLeakSimulation demonstrates a memory leak pattern
func memoryLeakSimulation() {
	fmt.Println("=== Memory Leak Simulation ===")

	// Simulate a memory leak - growing slice that's never cleaned
	var leakyData [][]byte

	stats1 := getGCStats()

	for i := 0; i < 500; i++ {
		// Keep adding data without removing old data
		data := make([]byte, 10*1024) // 10KB chunks
		leakyData = append(leakyData, data)

		if i%100 == 0 {
			currentStats := getGCStats()
			fmt.Printf("Iteration %d - Heap: %.2f MB, Objects: %d\n",
				i,
				float64(currentStats.HeapAlloc)/1024/1024,
				currentStats.HeapObjects)
		}
	}

	stats2 := getGCStats()

	fmt.Printf("Memory leak results:\n")
	fmt.Printf("Heap growth: %.2f MB\n", float64(stats2.HeapAlloc-stats1.HeapAlloc)/1024/1024)
	fmt.Printf("Object growth: %d\n", stats2.HeapObjects-stats1.HeapObjects)

	// Keep reference to maintain the "leak"
	_ = leakyData[len(leakyData)-1]
}

// RunGCDemo demonstrates garbage collection monitoring and tuning
func RunGCDemo() {
	fmt.Println("Garbage Collection Monitoring Demo")
	fmt.Println("===================================")

	// Initial stats
	initialStats := getGCStats()
	printGCStats("Initial State", initialStats)

	// Force a GC cycle
	runtime.GC()
	afterGCStats := getGCStats()
	printGCStats("After Manual GC", afterGCStats)

	// Demonstrate GC tuning
	demonstrateGCTuning()

	// Show memory leak pattern
	memoryLeakSimulation()

	// Final stats
	finalStats := getGCStats()
	printGCStats("Final State", finalStats)

	fmt.Printf("=== GC Monitoring Tips ===\n")
	fmt.Println("1. Monitor GC frequency and pause times in production")
	fmt.Println("2. Use runtime.ReadMemStats() for detailed statistics")
	fmt.Println("3. Set GOGC environment variable to tune GC frequency")
	fmt.Println("4. Lower GOGC = more frequent GC, less memory usage")
	fmt.Println("5. Higher GOGC = less frequent GC, more memory usage")
	fmt.Println("6. Use pprof for detailed memory profiling")

	fmt.Printf("\n=== Environment Variables ===\n")
	fmt.Println("GOGC=100 (default) - GC when heap doubles")
	fmt.Println("GOGC=50 - GC when heap grows 50%")
	fmt.Println("GOGC=200 - GC when heap triples")
	fmt.Println("GOGC=off - Disable automatic GC (not recommended)")

	fmt.Printf("\n=== Profiling Commands ===\n")
	fmt.Println("go tool pprof http://localhost:6060/debug/pprof/heap")
	fmt.Println("go tool pprof http://localhost:6060/debug/pprof/allocs")
	fmt.Println("GODEBUG=gctrace=1 go run main.go")
}

// main function for educational demonstration
func main() {
	fmt.Println("=== Garbage Collection Monitoring Demo ===")
	fmt.Println("This demonstrates GC monitoring, tuning, and memory leak detection")
	fmt.Println("Run with: GODEBUG=gctrace=1 go run gc_monitoring.go")
	fmt.Println("to see GC trace output")
	fmt.Println()

	RunGCDemo()
}
