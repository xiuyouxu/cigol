package main

import (
	"fmt"
	"mathbeta"
	"strconv"
	"time"
)

type MyTask struct {
	Id      string
	Content string
}

func (t MyTask) Run() {
	fmt.Println(t.Id, t.Content)
}

func main() {
	taskNum := []int{200, 300, 400}
	n := len(taskNum)
	tasks := make([][]mathbeta.Task, n)
	for i := 0; i < n; i++ {
		k := taskNum[i]
		tasks[i] = make([]mathbeta.Task, k)
		for j := 0; j < k; j++ {
			tasks[i][j] = &MyTask{Id: "task-" + strconv.Itoa(i) + "-" + strconv.Itoa(j), Content: strconv.FormatInt(time.Now().UnixNano(), 10)}
		}
	}

	mathbeta.Exec(tasks)
}
