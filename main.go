package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Data struct {
	str      [3]string
	operands [2]int
	operator int // 1 - сложение, 2 - вычитание, 3 - умножение, 4 - деление
}

func main() {
	var data Data

	err := reader(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)

	err = parser(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)

}

func reader(data *Data) error {
	rdr := bufio.NewReader(os.Stdin)
	str, err := rdr.ReadString('\n')
	if err != nil {
		return errors.New("Ошибка чтения из консоли")
	}
	str = strings.TrimSpace(str)
	substr := strings.Split(str, " ")
	if len(substr) != 3 {
		return errors.New("Ошибка: неверное количество аргументов")
	}
	for i, s := range substr {
		data.str[i] = s
	}
	return nil
}

func parser(data *Data) error {

	//
	return nil
}
