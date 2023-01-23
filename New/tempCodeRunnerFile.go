	go func(){
		for{
			v, ok:= <-ch
			if !ok{
				fmt.Println("Completed Receiving from ch")
				return
			}
			fmt.Printf("Received %v\n", v)
		}
	}()