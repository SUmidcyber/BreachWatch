package network

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func Network_name() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the interface name: ")
	interfaceName, _ := reader.ReadString('\n')
	interfaceName = interfaceName[:len(interfaceName)-1]

	os.Setenv("INTERFACE_NAME", interfaceName)

	cmd := exec.Command("snort", "-i", interfaceName, "-c", "/etc/snort/snort.conf", "-A", "console")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Snort failed to run:", err)
	}
}
