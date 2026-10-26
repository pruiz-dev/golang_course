package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	entry := "EVENT_LOG: status=200, action=process_payment, code=OK\n"

	startBuildWithPlus := time.Now()
	sBuildWithPlus := BuildWithPlus(40000, entry)
	durationBuildWithPlus := time.Since(startBuildWithPlus)

	startBuildWithBuilder := time.Now()
	sBuildWithBuilder := BuildWithBuilder(40000, entry)
	durationBuildWithBuilder := time.Since(startBuildWithBuilder)

	startBuildWithBuilderPrealloc := time.Now()
	sBuildWithBuilderPrealloc := BuildWithBuilderPrealloc(40000, entry)
	durationBuildWithBuilderPrealloc := time.Since(startBuildWithBuilderPrealloc)

	formattedPrint(
		"BuildWithPlus",
		len(sBuildWithPlus),
		durationBuildWithPlus,
	)

	formattedPrint(
		"BuildWithBuilder",
		len(sBuildWithBuilder),
		durationBuildWithBuilder,
	)

	formattedPrint(
		"BuildWithBuilderPrealloc",
		len(sBuildWithBuilderPrealloc),
		durationBuildWithBuilderPrealloc,
	)

	fmt.Printf(
		"All strings are equal: %v\n\n",
		sBuildWithBuilder == sBuildWithBuilderPrealloc &&
			sBuildWithBuilderPrealloc == sBuildWithPlus,
	)

	divisor := durationBuildWithBuilderPrealloc
	if divisor == 0 {
		divisor = time.Nanosecond // Защита: считаем, что заняло минимум 1 нс
	}

	multiplicity := durationBuildWithPlus / divisor
	fmt.Printf("string plus method is >= %d times slower\n\n", multiplicity)

}

func formattedPrint(method string, size int, duration time.Duration) {
	fmt.Printf(
		"Method: %v\n"+
			"Size in bytes: %d\n"+
			"Execution time: %v\n\n",
		method, size, duration,
	)
}

func BuildWithPlus(n int, entry string) string {
	var s string
	for range n {
		s += entry
	}

	return s
}

func BuildWithBuilder(n int, entry string) string {
	var sb strings.Builder
	for range n {
		sb.WriteString(entry)
	}

	return sb.String()
}

func BuildWithBuilderPrealloc(n int, entry string) string {
	var sb strings.Builder
	size := n * len(entry)
	sb.Grow(size)
	for range n {
		sb.WriteString(entry)
	}

	return sb.String()
}
