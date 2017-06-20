package parseJsonOrXml

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
	"task1"
	"task2"
	"task3"
	"task4"
	"task5"
	"task6"
	"task7"
	"validator"
)

type Data struct {
	T1 struct {
		Width, Height uint
		Symbol        string
	}
	T2 struct {
		Env1, Env2 task2.Envelope
	}
	T3 struct {
		SliceOfTriangles []task3.Triangle
	}
	T4 struct {
		Number uint
	}
	T5 struct {
		Min, Max uint
	}
	T6 struct {
		Length, MaxSquare uint
	}
	T7 struct {
		File string
	}
}

func SimpleSpellCheckerJson(file string) (b bool, message string) {
	openingCurlyBraces := strings.Count(file, "{")
	closingCurlyBraces := strings.Count(file, "}")
	quotes := strings.Count(file, "\"")
	openingBrackets := strings.Count(file, "[")
	closingBrackets := strings.Count(file, "]")

	if (openingBrackets+closingBrackets)%2 != 0 {
		b = false
		message += "К-ство открывающих и закрывающих квадратных скобок не равно\n"
	} else if (openingCurlyBraces+closingCurlyBraces)%2 != 0 {
		b = false
		message += "К-ство открывающих и закрывающих фигурных скобок не равно\n"
	} else if quotes%2 != 0 {
		b = false
		message += "К-ство кавычек не парно\n"
	} else {
		b = true
	}
	return b, message
}

func SimpleSpellCheckerXml(file string) (b bool, message string) {
	openingAngleBrackets := strings.Count(file, "<")
	closingAngleBrackets := strings.Count(file, ">")
	slashes := strings.Count(file, "/")

	if openingAngleBrackets != closingAngleBrackets {
		b = false
		message += "К-ство открывающих и закрывающих угловых скобок не равно\n"
	} else if (openingAngleBrackets/2 != slashes) || (closingAngleBrackets/2 != slashes) {

		b = false
		message += "К-ство слэшей не то\n"
	} else {
		b = true
	}
	return b, message
}

func CheckForNames(file string) (b bool, message string) {
	b = true

	var rules []string = []string{
		"T1", "T2", "T3", "T4", "T5", "T6", "T7",
		"Width", "Height", "Symbol",
		"Env1", "Env2", "Side1", "Side2",
		"SliceOfTriangles", "Name", "A", "B", "C",
		"Number",
		"Min", "Max",
		"Length", "MaxSquare",
		"File"}

	for _, v := range rules {
		if strings.Contains(file, v) == false {
			b = false
			message += v + "  "
		}
	}
	if b == false {
		message += " \n\nОтсутствуют либо неправильно написаны перечисленные выше переменные"
	}
	return b, message
}

func GetData(fileName string) (Data, bool, string) {
	var MyData Data
	extension := strings.Split(fileName, ".")

	if extension[len(extension)-1] == "json" {
		var contents []byte
		var err error
		if contents, err = ioutil.ReadFile("data.json"); err != nil {
			return MyData, false, "Не удалось прочитать файл"
		}
		var file string
		for _, v := range contents {
			file += string(v)
		}
		if b, m := validator.ValidateData(contents, extension[len(extension)-1]); b == false {
			return MyData, false, m
		}

		if b, m := SimpleSpellCheckerJson(file); b == false {
			return MyData, false, m
		}

		if b, m := CheckForNames(file); b == false {
			return MyData, false, m
		}
		var unmarshalFails string = "Не удалось распарсить json. " +
			"\nПроверьте данные и их типы." +
			"\nЗначения типа string должны быть заключены в кавычки: \"значение\"" +
			"\nЧисла должны быть без кавычек." +
			"\nЧисло не может быть меньше ноля" +
			"\n\nОбщий формат json'a таков: " +
			"\n{" +
			"\n\t\"ключ\":значение1,значение2\n}" +
			"\nЗначения разделяются запятой." +
			"\nЕсли значение составное, то действует то же правило:" +
			"\n{\n\"ключ1\":\n\t{" +
			"\n\t\"ключ2\":\n\t\tзначение1,значение2\n\t}," +
			"\n\t{\n\t\"ключ3\":\n\t\tзначение1,значение2\n\t}\n......\n}"
		if err = json.Unmarshal(contents, &MyData); err != nil {
			return MyData, false, unmarshalFails
		}

	} else if extension[len(extension)-1] == "xml" {
		var contents []byte
		var err error
		var file string
		if contents, err = ioutil.ReadFile("data.xml"); err != nil {
			return MyData, false, "Не удалось открыть файл"
		}
		for _, v := range contents {
			file += string(v)
		}
		if b, m := validator.ValidateData(contents, extension[len(extension)-1]); b == false {
			return MyData, false, m
		}
		if b, m := SimpleSpellCheckerXml(string(contents)); b == false {
			return MyData, false, m
		}

		if b, m := CheckForNames(string(contents)); b == false {
			return MyData, false, m
		}

		if err = xml.Unmarshal(contents, &MyData); err != nil {
			return MyData, false, "Не удалось распаковать xml"
		}

	} else {
		return MyData, false, "Не удалось открыть файл. \nРасширение файла должно быть формата json или xml."
	}
	return MyData, true, ""
}

func MakeAStructFromJson(fileName string) (Data, bool, string) {
	Data, ok, reason := GetData(fileName)
	return Data, ok, reason
}

func separator(n int) {
	fmt.Printf("\n\n------------->Task #%d<-------------\n\n", n)
}

func Operate(Data Data) {
	separator(1)
	fmt.Println(task1.DoTask1(Data.T1.Width, Data.T1.Height, Data.T1.Symbol))
	separator(2)
	fmt.Println(task2.DoTask2(task2.Envelope{Data.T2.Env1.Side1, Data.T2.Env1.Side2},
		task2.Envelope{Data.T2.Env2.Side1, Data.T2.Env2.Side2}))
	separator(3)
	fmt.Println(task3.DoTask3(Data.T3.SliceOfTriangles))
	separator(4)
	fmt.Println(task4.DoTask4(Data.T4.Number))
	separator(5)
	fmt.Println(task5.DoTask5(Data.T5.Min, Data.T5.Max))
	separator(6)
	fmt.Println(task6.DoTask6(Data.T6.Length, Data.T6.MaxSquare))
	separator(7)
	fmt.Println(task7.DoTask7(Data.T7.File))
}
