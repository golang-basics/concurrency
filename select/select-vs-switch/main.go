package main

import "fmt"

func main() {
	howSwitchWorks()
	howSelectWorks()
}

func howSwitchWorks() {
	// switch falls through every case till one case is true
	// the rest of the cases become unreachable
	switch {
	case boolean(false, "1"):
		fmt.Println("case 1")
	case boolean(false, "2"):
		fmt.Println("case 2")
	case boolean(true, "3"):
		fmt.Println("case 3")
	case boolean(true, "4"):
		// this case is unreachable
		fmt.Println("case 4")
	}
}

func boolean(b bool, label string) bool {
	fmt.Println(label)
	return b
}

func howSelectWorks() {
	// select does not fall through the cases
	// instead it checks all of them simultaneously
	// and randomly picks the one case who's most ready
	select {
	case <-channel("1"):
		fmt.Println("channel 1")
	case <-channel("2"):
		fmt.Println("channel 2")
	case <-channel("3"):
		fmt.Println("channel 2")
	case <-channel("4"):
		fmt.Println("channel 4")
	}
}

func channel(label string) chan struct{} {
	fmt.Println(label)
	out := make(chan struct{})
	close(out)
	return out
}
