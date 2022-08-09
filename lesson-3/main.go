package main

import "fmt"

func main() {
	scanWithSscan()
	scanWithSscanf()
	scanWithSscanln()
	scanWithScan()

}

func scanWithSscan() {
	var d1, d2 int
	var s1 string
	fmt.Printf("Before scanning: %d,%d,%s\n", d1, d2, s1)
	if _, err := fmt.Sscan("5 7\n9", &d1, &d2, &s1); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("After scanning with Sscan: %d, %d, %s\n", d1, d2, s1)
}

func scanWithSscanf() {
	var d1, d2 int
	var s1 string
	fmt.Printf("Before scanning: %d, %d, %s\n", d1, d2, s1)
	if _, err := fmt.Sscanf("5, 7 and 9 are my values", "%d, %d and %s", &d1, &d2, &s1); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("After scanning with Sscanf: %d, %d, %s\n", d1, d2, s1)
}

func scanWithSscanln() {
	var d1, d2 int
	var s1 string
	fmt.Printf("Before scanning: %d, %d, %s\n", d1, d2, s1)
	if _, err := fmt.Sscanln("5 7 9\n", &d1, &d2, &s1); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("After scanning with Sscanln: %d, %d, %s\n", d1, d2, s1)
}

func scanWithScan() {
	var s1, s2, s3 string
	fmt.Print("You inputs please: ")
	if _, err := fmt.Scan(&s1, &s2, &s3); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("After scanning with Scan: %s, %s, %s\n", s1, s2, s3)
}
