package main

import (
  "time"
  "fmt"
  "os"
  "os/exec"
)

func main() {
  fmt.Println("Wifi connection being monitored.....")
  for {
  time.Sleep(20 * time.Minute)
  wifiControl("off")
  time.Sleep(3 * time.Second)
  wifiControl("on")
}
}


func wifiControl(state string) {
  cmdName := "networksetup"
  cmdArgs := []string{"-setairportpower", "en1", state}
  if _, err := exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running : ", err)
		os.Exit(1)
	}
}
