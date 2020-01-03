package main

import (
	"training.go/imgproc/filter"
	"training.go/imgproc/task"
)

func main() {
	var f filter.Filter = filter.Grayscale{}
	t := task.NewWaitGrpTask("./imgs", "output", f)

	t.Process()
}
