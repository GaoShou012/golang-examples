package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type Persons []*Person

func (s Persons) Len() int           { return len(s) }
func (s Persons) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Persons) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	//p := []Person{Person{"Lily", 20}, Person{"Bing", 18}, Person{"Tom", 23}, Person{"Vivy", 16}, Person{"John", 18}}

	var p []*Person
	p = append(p, &Person{"Lily", 20})
	p = append(p, &Person{"Bing", 18})
	p = append(p, &Person{"Tom", 23})
	p = append(p, &Person{"Vivy", 16})
	p = append(p, &Person{"John", 18})
	//sort.Sort(sort.Reverse(Persons(p)))
	sort.Sort(Persons(p))
	for _, row := range p {
		fmt.Println(row)
	}
	//sort.Sort(sort.Reverse(Persons(p))) //sort.Reverse 生成递减序列
	//fmt.Println(p)
}
