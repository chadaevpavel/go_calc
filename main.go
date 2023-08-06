package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	str      [3]string
	operands [2]int
	operator string
	is_roman bool
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
	romans := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VII",
		9:  "IX",
		10: "X",
	}
	operators := []string{"+", "-", "*", "/"}
	var err1, err2 error
	var found, found1, found2 bool

	data.operands[0], err1 = strconv.Atoi(data.str[0])
	data.operands[1], err2 = strconv.Atoi(data.str[2])
	if err1 != nil || err2 != nil {
		for i, literal := range romans {
			data.str[0] = strings.ToUpper(data.str[0])
			data.str[2] = strings.ToUpper(data.str[2])
			if data.str[0] == literal {
				data.operands[0] = i
				found1 = true
			}
			if data.str[2] == literal {
				data.operands[1] = i
				found2 = true
			}
		}
		if found1 && found2 {
			data.is_roman = true
		} else if !(err1 != nil && err2 != nil) && (found1 || found2) {
			return errors.New("Ошибка: нельзя смешивать арабские и римские цифры")
		} else {
			return errors.New("Ошибка: неверный формат операндов")
		}
	}

	for _, op := range operators {
		if data.str[1] == op {
			found = true
		}
	}

	if !found {
		return errors.New("Ошибка: неверный оператор")
	}
	return nil
}
