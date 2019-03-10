package main

import (
	"github.com/Equanox/gotron"
)

func main() {

	//Create a new window
	window,err := gotron.New("webapp")
	if err != nil {
		panic(err)
	}


	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 900
	window.WindowOptions.Title = "First GUI!"

    done, err := window.Start()
    if err != nil {
        panic(err)
	}
	
	
	//Open dev tools must be used after window.Start
	//window.OpenDevTools()

	//Wait for the application to close 
	 <-done
}