package validator

import (
	"fmt"
	parse "parseJsonOrXml"
	"task1"
	"task2"
	"task3"
	"task4"
	"task5"
	"task6"
	"task7"
	"encoding/json"
	"io/ioutil"
	"encoding/xml"
	"strings"
)

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func Task1Rules(d parse.Data) (ok bool, reason string) {
	ok = true
	return
}
func Task2Rules(d parse.DataChecking) (ok bool, reason string) {
	ok = true
	return
}
func Task3Rules(d parse.DataChecking) (ok bool, reason string) {
	ok = true
	return
}
func Task4Rules(d parse.DataChecking) (ok bool, reason string) {
	ok = true
	return
}
func Task5Rules(d parse.DataChecking) (ok bool, reason string) {
	ok = true
	return
}
func Task6Rules(d parse.DataChecking) (ok bool, reason string) {
	ok = true
	return
}
func Task7Rules(d parse.DataChecking) (ok bool, reason string) {
	ok = true
	return
}

func separator(n int) {
	fmt.Printf("\n\n------------->Task #%d<-------------\n\n", n)
}
func Operate(Data parse.Data){
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

func validateData(file string) (ok bool, reason string) {
	
}

func RiseAndShine(fileName string) (parse.Data, bool, string) {
	Data, ok, reason := parse.GetData(fileName)
	return Data, ok, reason
}
