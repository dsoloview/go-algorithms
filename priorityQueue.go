package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* https://stepik.org/lesson/13240/step/8?unit=3426 */

const EXTRACT_MAX = "ExtractMax"
const INSERT = "Insert"

type PriorityQueue []int

func (q PriorityQueue) SiftDown(i int) {
	for {
		left := 2*i + 1
		right := 2*i + 2

		largest := i
		if left < len(q) && q[i] < q[left] {
			largest = left
		}

		if right < len(q) && q[largest] < q[right] {
			largest = right
		}

		if largest == i {
			return
		}

		q[i], q[largest] = q[largest], q[i]
		i = largest
	}
}

func (q PriorityQueue) SiftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if q[parent] >= q[i] {
			return
		}
		q[parent], q[i] = q[i], q[parent]
		i = parent
	}
}

func (q PriorityQueue) Add(element int) PriorityQueue {
	q = append(q, element)
	q.SiftUp(len(q) - 1)

	return q
}

func (q PriorityQueue) BuildHeap() {
	for i := len(q)/2 - 1; i >= 0; i-- {
		q.SiftDown(i)
	}
}

func (q PriorityQueue) ExtractMax() (int, error) {

	if len(q) > 0 {
		result := q[0]
		q[0] = q[len(q)-1]
		q = q[:len(q)-1]
		if len(q) > 0 {
			q.SiftDown(0)
		}
		return result, nil
	}

	return 0, errors.New("queue is empty")
}
func mainPriorityQueue() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	var commands []string

	var queue PriorityQueue
	for i := 0; i < n; i++ {
		sc.Scan()
		commands = append(commands, sc.Text())
	}
	err := executeCommands(commands, queue)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func executeCommands(commands []string, queue PriorityQueue) error {
	for _, command := range commands {
		splitedCommand := strings.Split(command, " ")
		if len(splitedCommand) == 1 && splitedCommand[0] == EXTRACT_MAX {
			max, err := queue.ExtractMax()
			if err != nil {
				return err
			}

			fmt.Println(max)
		}

		if len(splitedCommand) == 2 && splitedCommand[0] == INSERT {
			number, err := strconv.Atoi(splitedCommand[1])
			if err != nil {
				return err
			}
			queue = queue.Add(number)
		}
	}

	return nil
}
