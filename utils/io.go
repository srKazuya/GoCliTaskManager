package utils

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func GetTaskId(reader *bufio.Reader) (int, error) {
	fmt.Print("Введите ID задачи")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	taskId, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("введите корректное число")
	}
	return taskId, nil
}
