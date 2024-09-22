package network

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Interface() {
	file, err := os.OpenFile("interface.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Learn the interface [Y/n]: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if strings.EqualFold(input, "Y") {
		fmt.Print("Enter command (Windows [ipconfig] | Linux [ifconfig]): ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		var cmd *exec.Cmd
		if command == "ifconfig" {
			cmd = exec.Command("ifconfig")
		} else if command == "ipconfig" {
			cmd = exec.Command("ipconfig")
		} else {
			fmt.Println("Unknown command...")
			logger.Println("Unknown command entered.")
			return
		}

		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Failed to execute command:", err)
			logger.Println("Failed to execute command:", err)
			return
		}

		logger.Println("Command executed successfully:", command)
		logger.Println("Output:\n", string(output))

		fmt.Println("Command executed successfully. Check the log file for details.")
	} else {
		fmt.Println("Interface learning not initiated.")
		logger.Println("Interface learning not initiated.")
	}
}
