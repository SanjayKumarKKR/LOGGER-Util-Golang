# LOGGER-Util-Golang
logger-util-golang is a Go package that helps to view to view the logs in a webpage rather than a command line. It instantly makes the logging easy as the logs are served under a port. 

# Usage
First, get the code from this repo 

``go get github.com/SanjayKumarKKR/LOGGER-Util-Golang``

Then import it to your project.

``import logger "github.com/SanjayKumarKKR/LOGGER-Util-Golang" ``


## Example Usage

Presumably your application already uses the default log package. To switch, you'll want your code to look like the following:


```package main

import (
	"fmt"
	"strconv"

	logger "github.com/SanjayKumarKKR/LOGGER-Util-Golang"
)

func main() {
	l := logger.Logger{}
	l.Log("Started the project")
	fmt.Println("sanjay")

	l.Log("Implemented the logic")
	for i := 0; i < 10; i++ {
		l.Log("Implemented the logic" + strconv.Itoa(i))
	}
	l.Log("Project Completed Successfully")

	l.Show()
}
```
