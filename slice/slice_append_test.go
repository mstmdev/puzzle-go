package slice

import (
	"fmt"
	"testing"
)

func TestPrintFirstCase(t *testing.T) {
	printFirstCase(0)  // print p3
	printFirstCase(1)  // print p7
	printFirstCase(2)  // print p3
	printFirstCase(3)  // print p7
	printFirstCase(4)  // print p7
	printFirstCase(5)  // print p7
	printFirstCase(6)  // print p3
	printFirstCase(7)  // print p7
	printFirstCase(8)  // print p7
	printFirstCase(9)  // print p7
	printFirstCase(10) // print p7
}

type Person struct {
	Name string
}

func printFirstCase(nilItemCount int) {
	var users []*Person
	users = append(users, &Person{Name: "p1"})
	users = append(users, &Person{Name: "p2"})
	for i := 0; i < nilItemCount; i++ {
		users = append(users, nil)
	}

	fmt.Printf("nilItemCount=%d cap=%d len=%d. ", nilItemCount, cap(users), len(users))
	if len(users) < cap(users) {
		fmt.Printf("It has sufficient capacity, the destination is resliced to accommodate the new elements.\n")
	} else {
		fmt.Printf("It has not sufficient capacity, a new underlying array will be allocated.\n")
	}

	testCases := []struct {
		Users []*Person
	}{
		{append(users, &Person{Name: "p3"})},
		{append(users, &Person{Name: "p4"})},
		{append(users, &Person{Name: "p5"})},
		{append(users, &Person{Name: "p6"})},
		{append(users, &Person{Name: "p7"})},
	}

	fmt.Printf("append [%d] nil item, the first case's last username is [%s]\n", nilItemCount, testCases[0].Users[len(testCases[0].Users)-1].Name)
}
