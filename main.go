package main

import (
				"bufio"
				"fmt"
				"log"
				"os"
				"os/exec"
				"strings"
				"time"
)

// AppInfo holds information about an application
type AppInfo struct {
				Name  string
				Path  string
				Alias string
}

// SplitAt splits a string at a delimiter
func SplitAt(delim string) bufio.SplitFunc {
				return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
								if atEOF && len(data) == 0 {
												return 0, nil, nil
								}
								if i := strings.Index(string(data), delim); i >= 0 {
												return i + len(delim), data[0:i], nil
								}
								if atEOF {
												return len(data), data, nil
								}
								return 0, nil, nil
				}
}

func main() {
				t2 := time.Now()

				// Get the name of program file from config
				configFile, err := os.ReadFile("config")
				if err != nil {
								fmt.Println("Error reading config file:", err)
								return
				}

				programFile := strings.ReplaceAll(string(configFile), `"`, "")
				programFile = strings.ReplaceAll(programFile, "\r\n", "")

				// Get DATA from program file
				data, err := os.Open(programFile)
				if err != nil {
								fmt.Println("Error opening program file:", err)
								return
				}
				defer data.Close()

				// Read and process lines
				lineScanner := bufio.NewScanner(data)
				appsMap := make(map[string]AppInfo)

				for lineScanner.Scan() {
								line := strings.ReplaceAll(lineScanner.Text(), `"`, "")
								parts := strings.SplitN(line, ", ", 3)
								if len(parts) != 3 {
												fmt.Println("Invalid line format:", line)
												continue
								}

								appInfo := AppInfo{
												Name:  strings.TrimPrefix(parts[0], "name:"),
												Path:  strings.TrimPrefix(parts[1], "path:"),
												Alias: strings.TrimPrefix(parts[2], "alias:"),
								}

								appsMap[appInfo.Name] = appInfo
								appsMap[appInfo.Alias] = appInfo
				}

				if err := lineScanner.Err(); err != nil {
								fmt.Println("Error reading lines:", err)
								return
				}

				fmt.Printf("Number of applications: %v\n", len(appsMap))

				/* Debug print the apps
				for _, app := range appsMap {
								fmt.Printf("App Name: %v, Path: %v, Alias: %v\n", app.Name, app.Path, app.Alias)
				}
				*/

				// Process command-line arguments
				appsNotFound := []string{}
				for _, argValue := range os.Args[1:] {
								app, found := appsMap[argValue]
								if found {
												//fmt.Printf("Executing: %v\n", argValue)
												cmd := exec.Command(app.Path)
												if err := cmd.Start(); err != nil {
																log.Printf("Error starting command: %v", err)
												}
								} else {
												appsNotFound = append(appsNotFound, argValue)
								}
				}

				if len(appsNotFound) > 0 {
								fmt.Println("Apps not found:", appsNotFound)
				}

				fmt.Printf("For loop took: %v nanoseconds\n", time.Since(t2).Nanoseconds())
}
