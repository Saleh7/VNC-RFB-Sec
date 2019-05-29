package main

import (
    "bufio"
    "fmt"
    "net"
    "log"
    "os/exec"
    "time"
    "os"
    "github.com/bclicn/color"
)

const (
  v0 = "RFB 000.000\n"  // UltraVNC repeater
  v1 = "RFB 003.003\n"
  v2 = "RFB 003.006\n"  // UltraVNC
  v3 = "RFB 003.889\n"  // Apple Remote Desktop
  v4 = "RFB 003.007\n"
  v5 = "RFB 003.008\n"
  v6 = "RFB 004.000\n"  // Intel AMT KVM
  v7 = "RFB 004.001\n"  // RealVNC 4.6
  TCP_PORT = "5900"  // TCP PORT
  TIMEOUT = 2  // timeout x second
)

func system(cmd string, arg ...string) string{
    out, err := exec.Command(cmd, arg...).Output()
    if err != nil {
        log.Fatal(err)
    }
    return string(out)
}
func checkSecurityType(ip string, port string) bool{
  conn, err := net.DialTimeout("tcp", ip + ":" + port, TIMEOUT * time.Second)
  if err != nil {
    return false
  }
  connbuf := bufio.NewReader(conn)
  str, err := connbuf.ReadString('\n')
  if len(str)>0 {
    if v5 == str {
      // vncsnapshot https://github.com/shamun/vncsnapshot
      // system("timeout","60","vncsnapshot","-allowblank","-vncQuality","6","-quality","60",ip,""+ip+".jpg")
      fmt.Println(color.Green("No Authentication: ") + color.LightYellow(ip))
      return true
    }
  }
  if err!= nil {
      return false
  }
  return false
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print(color.BYellow("Please enter full file path to list of IPs.") + color.LightRed("\nExample: ") + "/home/user/Documents/list.txt\nFull path: ")
        scanner.Scan()
        file_path := scanner.Text()
        if len(file_path) != 0 {
            fmt.Print("\nRunning ...\n")
            file, err := os.Open(file_path)
            if err != nil {
                fmt.Print(color.LightRed("\nPlease enter full file path\n"))
                break
            }
            defer file.Close()
            scanner := bufio.NewScanner(file)
            for scanner.Scan() {
                checkSecurityType(scanner.Text(),TCP_PORT)
            }
            if err := scanner.Err(); err != nil {
                log.Fatal(err)
                break
            }
        } else {
            break
        }
        fmt.Println("\nProgram complete, enjoy sifting through your results :-)")
        break
    }
}
