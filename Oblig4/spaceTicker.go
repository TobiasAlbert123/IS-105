package main

import "time"
import "fmt"

func main() {

	start := time.Date(2017, time.December, 17, 0, 0, 0, 0, time.UTC)
	ticker:= time.Tick(time.Second*5)
	time.Sleep(time.Millisecond*50)
	newStart := time.Date(2017, time.December, 17, 0, 0, 0, 0, time.UTC)
	newTicker := time.Tick(time.Second*5)
	time.Sleep(time.Millisecond*50)
	newStart1 := time.Date(2017, time.December, 17, 0, 0, 0, 0, time.UTC)
	newTicker1:= time.Tick(time.Second*5)
	time.Sleep(time.Millisecond*50)
	newStart2 := time.Date(2018, time.March, 21, 0, 0, 0, 0, time.UTC)
	newTicker2 := time.Tick(time.Second*5)
	time.Sleep(time.Millisecond*50)
	newStart3 := time.Date(2018, time.March, 21, 0, 0, 0, 0, time.UTC)
	newTicker3 := time.Tick(time.Second*5)
	time.Sleep(time.Millisecond*50)
	newStart4 := time.Date(2018, time.March, 21, 0, 0, 0, 0, time.UTC)
	newTicker4 := time.Tick(time.Second*5)
	time.Sleep(time.Millisecond*50)




	//go func() {
		for{
			select {
			case <- ticker:
				t := time.Now()
				elapsed := t.Sub(start)
				fmt.Printf("Scott Tingle has been in space %.0f days\n", elapsed.Hours()/24)

			case <- newTicker:
				t := time.Now()
				elapsed := t.Sub(newStart)
				fmt.Printf("Anton Skhaplerov has been in space %.0f days\n", elapsed.Hours()/24)

			case <- newTicker1:
				t := time.Now()
				elapsed := t.Sub(newStart1)
				fmt.Printf("Norishige Kanai has been in space %.0f days\n", elapsed.Hours()/24)

			case <- newTicker2:
				t := time.Now()
				elapsed := t.Sub(newStart2)
				fmt.Printf("Andrew Feustel has been in space %.0f days\n", elapsed.Hours()/24)

			case <- newTicker3:
				t := time.Now()
				elapsed := t.Sub(newStart3)
				fmt.Printf("Richard Arnold has been in space %.0f days\n", elapsed.Hours()/24)

			case <- newTicker4:
				t := time.Now()
				elapsed := t.Sub(newStart4)
				fmt.Printf("Oleg Martemyev has been in space %.0f days\n", elapsed.Hours()/24)

			}


		}
	//}()

	select {}

}
