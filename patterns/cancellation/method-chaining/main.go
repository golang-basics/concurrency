package main

import (
	"fmt"

	"concurrency/patterns/cancellation"
)

func main() {
	p := cancellation.NewIntPipeline(1, 2, 3)
	defer p.Done()
	//for i := range p.Inc().Sq().Done().Dec().Res() {
	for i := range p.Inc().Sq().Dec().Res() {
		fmt.Println(i)
	}
}
