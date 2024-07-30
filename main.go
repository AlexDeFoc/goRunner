package main

import (
				"bufio"
				//"fmt" //DEBUG
				"log"
				"os"
				"os/exec"
				"strings"
)

//TIP:
//OS.ARGS len =>
//1: no args
//2: one arg
//3: two arg

func main(){
				//CHECK if User Inputed at least one argument
				if len(os.Args) > 1 {
								run()
				}
}

func run() {
				rawConfig, err := os.ReadFile("config")

				//Error Handling
				if err != nil {
								log.Panic("Programs file not found")
				}

				config := string(rawConfig)
				config = strings.TrimSpace(config)

				rawFile, err := os.Open(config)
				defer rawFile.Close()

				//Error Handling
				if err != nil {
								log.Panic("Programs file not found")
				}

				lineScanner := bufio.NewScanner(rawFile)
				lineData := []string {}

				for lineScanner.Scan() {
								lineData = append(lineData, lineScanner.Text())
				}

				appScanner := make([]bufio.Scanner, len(lineData))
				for i := range lineData{
								appScanner[i] = *bufio.NewScanner(strings.NewReader(lineData[i]))
				}

				fieldData := make([][]string, len(appScanner))

				//appScanner = bufio.NewScanner(strings.NewReader(lineData[0]))
				for i:= range appScanner{
								appScanner[i].Split(bufio.ScanWords)
								for appScanner[i].Scan(){
												//fmt.Printf("%v\n", appScanner[i].Text())
												fieldData[i] = append(fieldData[i], appScanner[i].Text())
								}
				}

				//Clean Fields for proper use
				for i := range len(fieldData) {
								for j := range len(fieldData[i]) {

												fieldData[i][j] = strings.TrimSuffix(fieldData[i][j], ",")
												if strings.Contains(fieldData[i][j], "name"){
																fieldData[i][j] = strings.ReplaceAll(fieldData[i][j], "name:", "")
												}else if strings.Contains(fieldData[i][j], "path"){
																fieldData[i][j] = strings.ReplaceAll(fieldData[i][j], "path:", "")
																fieldData[i][j] = strings.ReplaceAll(fieldData[i][j], "\\", "//")
												}else if strings.Contains(fieldData[i][j], "alias"){
																fieldData[i][j] = strings.ReplaceAll(fieldData[i][j], "alias:", "")
												}
												fieldData[i][j] = strings.Trim(fieldData[i][j], `\"`)
												//fmt.Printf("%v\n", fieldData[i][j]) //DEBUG

								}
				}



				//Get user flags
				var panicStatus bool
				for _, arg := range os.Args[1:] {
								for n := range len(fieldData) {
												if arg == fieldData[n][2] || arg == fieldData[n][0]{
																cmd := exec.Command(fieldData[n][1])

																err := cmd.Start()

																if err != nil {
																				log.Printf("Couldn't start: %v at path: %v\n", fieldData[n][0], fieldData[n][1])
																				panicStatus = true
																}
																//fmt.Printf("Going to run: %v\n", fieldData[n][1]) //DEBUG
												}
								}
				}

				if panicStatus == true {
								log.Panic()
				}
}
