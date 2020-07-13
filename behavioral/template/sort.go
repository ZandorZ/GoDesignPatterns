package main

import (
	"fmt"
	"sort"
)

type MyList []int

func (m MyList) Len() int {
	return len(m)
}

func (m MyList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MyList) Less(i, j int) bool {
	return m[i] < m[j]
}

func main() {
	var myList MyList = []int{61, 3, 2, 6, 12, 11}
	fmt.Println(myList)
	sort.Sort(myList)
	fmt.Println(myList)
}
