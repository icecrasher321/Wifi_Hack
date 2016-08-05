package main

import (
  "time"
  "fmt"
  "os"
  "os/exec"
)

func main() {
  var timeInt int
  fmt.Println("Enter the time interval (whole number - in minutes >= 1) between each wifi reset")
  fmt.Scanln(&timeInt)
  fmt.Println("Wifi connection being monitored.....")
  for {
  time.Sleep(time.Duration(timeInt) * time.Minute)
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
