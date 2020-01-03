package main

import (
	"fmt"
	"flag"
	"time"
	"training.go/imgproc/filter"
	"training.go/imgproc/task"
)

func main() {
	var srcDir = flag.String("src", "", "Input directory")
	var dstDir = flag.String("dst", "", "Output directory")
	var filterType = flag.String("filter", "grayscale", "grayscale/blur")
	var taskname = flag.String("task", "waitgrp", "waitgrp/channel")
	var poolSize = flag.Int("poolsize", 4, "Number of process")
	flag.Parse()

	var f filter.Filter
	switch *filterType {
	case "grayscale":
		f = filter.Grayscale{}
	case "blur":
		f = filter.Blur{}
	default:
		f = filter.Grayscale{}
	}

	var t task.Tasker
	switch *taskname {
	case "waitgrp":
		t = task.NewWaitGrpTask(*srcDir, *dstDir, f)
	case "channel":
		t = task.NewChanTask(*srcDir, *dstDir, f, *poolSize)
	}

	start := time.Now()
	t.Process()
	elapsed := time.Since(start)

	fmt.Printf("Image processing took %s\n", elapsed)
}
