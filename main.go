package main

import (
	"fmt"
	"os/exec"
	"io/ioutil"
	"bytes"
	"time"
	"strconv"
	"math/rand"
)

func exec_shell(arg ...string) (string, error){
	cmd := exec.Command(arg[0], arg[1:]...)

	var out bytes.Buffer
    cmd.Stdout = &out

    err := cmd.Run()
    return out.String(), err
}

func writeFile() {
	now := time.Now().Format("2006-01-02T15:04:05.999999-07:00")
	fmt.Println(now)
	d1 := []byte(now + "\n")
	err := ioutil.WriteFile("./hacker.txt", d1, 0644)
	if err != nil {
		fmt.Print("write file error", err)
	}
}

func generateTimeStr(num int) string {
	// example: --date=\"Thu Aug 9 00:00:00 2018 "
	dur := strconv.Itoa(num * -24) + "h"
	hisTime, _ := time.ParseDuration(dur)
	// fmt.Println(dur, hisTime)
	timeBefore := time.Now().Add(hisTime)
	// fmt.Println(timeBefore.Format("Mon Jan 2 15:04:05 2006"))
	timeStamp := "--date=\"" + timeBefore.Format("Mon Jan 2 15:04:05 2006") + "\""
	// fmt.Println(timeStamp)
	return timeStamp
}

func add() {
	s := []string {"git", "add", "hacker.txt"}
	res, err := exec_shell(s...)
	if err != nil {
		fmt.Print("add fail:", err, res)
	}
}

func commit(timeStr string) {
	str := time.Now().Format("Mon Jan 2 15:04:05 2006")
	s := []string {"git", "commit", "-m", "\"" + str + "\"", 
	""}
	s[4] = timeStr
	res, err := exec_shell(s...)
	if err != nil {
		fmt.Print("commit fail:", err, res)
	}
}

func push() {
	s := []string {"git", "push"}
	res, err := exec_shell(s...)
	if err != nil {
		fmt.Print("push fail:", err, res)
	}
}

func gitProcess(i int) {
	writeFile()
	add()
	s := generateTimeStr(i) 
	commit(s)
}

func main()  {
	for i := 0; i <= 365; i++ {
		if rand.Intn(100) > 70 {
			gitProcess(i)
			if rand.Intn(100) > 42 {
				for t := 0; t < rand.Intn(10); t ++ {
					gitProcess(i)
				}
			}
		}
	}
	push()	
}