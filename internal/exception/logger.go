package exception

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

func logFile() {
	
	if DEBUG == "false" {
		file, err := os.OpenFile("../../storage/log/app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			color.Red(fmt.Sprintf("Error : %s", err))
		}
		log.SetOutput(file)
		log.Println("Start logger :)")
	}
}
