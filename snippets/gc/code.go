package main

import "fmt"
import "runtime"

func main() {
	/* display current memory allocation on startup */
	var memstats runtime.MemStats
	runtime.ReadMemStats(&memstats)
	fmt.Printf("HeapSys:\t%d\nHeapAlloc:\t%d\n", memstats.HeapSys, memstats.HeapAlloc)

	/* lets hog some ram */
	ram := make([]byte, 1000e6)
	ram[0] = 0
	/* let's disp at end too */
	runtime.ReadMemStats(&memstats)
	fmt.Printf("HeapSys:\t%d\nHeapAlloc:\t%d\n", memstats.HeapSys, memstats.HeapAlloc)

	ram = nil
	runtime.GC() // call the garbage collection now

	/* we have marked ram as being eligible for garbage collection with ram=nil; */
	runtime.ReadMemStats(&memstats)
	fmt.Printf("HeapSys:\t%d\nHeapAlloc:\t%d\n", memstats.HeapSys, memstats.HeapAlloc)

}
