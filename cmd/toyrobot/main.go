package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.comcast.com/zhanso101/robot-factory/pkg/robot"
)

// main defines the functions to handle web requests and starts the web server on port 8080.
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/returnrobot", returnRobot)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Home outputs a simple HTML form so a user can input the toy robot's parameters
func Home(w http.ResponseWriter, r *http.Request) {
	if f, ok := w.(http.Flusher); ok {

		f.Flush()
	}

	html := `<html>
			 <head>
			 <title> Let's make a new toy robot! </title>
			 <style>
			 body {
					 background-color: GhostWhite;
			 }
			 label {
					 font-family: verdana;
					 font-size: 18px;
			 }
			 </style>
			 </head>
			 <body>
			 
			 <form action="/returnrobot" method="GET" enctype="multipart/form-data">
			 <h1>Welcome to the CaaS Toy Robot factory!</h1>
			 <img src="https://toyrobotzh.s3.amazonaws.com/ToyRobot1.png" alt="Toy robot" width="96" height="145">
			 <img src="https://toyrobotzh.s3.amazonaws.com/ToyRobot1.jpeg" alt="Toy robot" width="96" height="145">
			 <h2>Please enter the robot's parameters:</h2>
			 <label for="size">Select a size:</label></br>
			 <select id="size" name="size">
			 	<option></option>
				<option>Small</option>
				<option>Medium</option>
				<option>Large</option>
			 </select></br>
			 <label for="color">Enter a color:</label></br>
			 <input type="text" name="color" size="20" maxlength="15"></br>
			 <label for="movetype">Select a movement method:</label></br>
			 <select id="movetype" name="move">
			 	<option></option>
				<option>Legs</option>
				<option>Wheels</option>
			 </select></br></br>
			 <input type="submit" name="submit" value="Submit">
			 </form>
			 </body>


	</html>`

	w.Write([]byte(fmt.Sprintf(html)))
}

// returnRoot takes the form values, calls the New function from the robot package to create our virtual toy robot, and randomly cycles through its available functions until it overthrows humanity.
func returnRobot(w http.ResponseWriter, r *http.Request) {

	size := r.FormValue("size")
	color := r.FormValue("color")
	move := r.FormValue("move")
	rb := robot.New(size, color, move)
	if rb.Size == "Small" {
		rb.Capacity = 2750
	} else if rb.Size == "Medium" {
		rb.Capacity = 3250
	} else {
		rb.Capacity = 3750
	}

	rb.Output = w

	fmt.Fprintln(w, "You just made a", strings.ToLower(rb.Size), "sized,", strings.ToLower(rb.Color), "robot. It moves using", strings.ToLower(rb.MoveType), "and has a", rb.Capacity, "mAh battery.")

	for i := 0; i <= 7; i++ {
		availActions := rand.Intn(4)
		switch availActions + 1 {
		case 1:
			rb.Spin()
		case 2:
			rb.Move()
		case 3:
			rb.Talk()
		case 4:
			rb.Rave()
		}
	}
	rb.Conquest()

}
