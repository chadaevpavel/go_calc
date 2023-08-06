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
	operator int
	is_roman bool
	result   int
}

func main() {
	var data Data

	err := reader(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = parser(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = calculator(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = printer(&data)
	if err != nil {
		fmt.Println(err)
		return
	}

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
	roman_operands := map[int]string{
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
	operators := map[int]string{
		1: "+",
		2: "-",
		3: "*",
		4: "/"}
	var err1, err2 error
	var found1, found2 bool

	data.operands[0], err1 = strconv.Atoi(data.str[0])
	data.operands[1], err2 = strconv.Atoi(data.str[2])
	if err1 != nil || err2 != nil {
		for i, literal := range roman_operands {
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
	} else {
		if data.operands[0] < 1 || data.operands[0] > 10 || data.operands[1] < 1 || data.operands[1] > 10 {
			return errors.New("Ошибка: значение оператора выходит за допустимый диапазон")
		}
	}

	for i, op := range operators {
		if data.str[1] == op {
			data.operator = i
		}
	}
	if data.operator == 0 {
		return errors.New("Ошибка: неверный оператор")
	}

	return nil
}

func calculator(data *Data) error {
	switch data.operator {
	case 1:
		data.result = data.operands[0] + data.operands[1]
	case 2:
		data.result = data.operands[0] - data.operands[1]
	case 3:
		data.result = data.operands[0] * data.operands[1]
	case 4:
		data.result = data.operands[0] / data.operands[1]
	default:
		return errors.New("Ошибка: неверный оператор")
	}
	if data.is_roman && data.result < 1 {
		return errors.New("Ошибка: результат выходит за допустимые пределы")
	}
	return nil
}

func printer(data *Data) error {
	if data.is_roman {
		fmt.Println(int_to_roman(data.result))
	} else {
		fmt.Println(data.result)
	}
	return nil
}

func int_to_roman(n int) string {
	var roman_numbers = map[int]string{
		1:   "I",
		4:   "IV",
		5:   "V",
		9:   "IX",
		10:  "X",
		40:  "XL",
		50:  "L",
		90:  "XC",
		100: "C",
	}
	var res string
	for n > 0 {
		nearest := 1
		for i := range roman_numbers {
			if i <= n && i > nearest {
				nearest = i
			}
		}
		res += roman_numbers[nearest]
		n -= nearest
	}
	return res
}
