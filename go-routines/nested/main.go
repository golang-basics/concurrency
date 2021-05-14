package main

import (
	"fmt"
	"sync"
)

func main() {
	var mainWG sync.WaitGroup
	mainWG.Add(1)
	go func() {
		defer mainWG.Done()
		var g1WG sync.WaitGroup
		g1WG.Add(1)

		go func() {
			defer g1WG.Done()
			var g2WG sync.WaitGroup
			g2WG.Add(1)

			go func() {
				defer g2WG.Done()
				var g3WG sync.WaitGroup
				g3WG.Add(1)

				go func() {
					defer g3WG.Done()
					var g4WG sync.WaitGroup
					g4WG.Add(2)

					go func() {
						fmt.Println("g5 done")
						g4WG.Done()
					}()
					go func() {
						fmt.Println("g6 done")
						g4WG.Done()
					}()

					g4WG.Wait()
					fmt.Println("g4 done")
				}()

				g3WG.Wait()
				fmt.Println("g3 done")
			}()

			g2WG.Wait()
			fmt.Println("g2 done")
		}()

		g1WG.Wait()
		fmt.Println("g1 done")
	}()
	mainWG.Wait()
	fmt.Println("main g done")
}
