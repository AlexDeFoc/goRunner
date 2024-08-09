package main

import (
    "encoding/json"
    "flag"
    "log"
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
    "strings"
    "sync"
)

//log.Printf("Checkpoint met")
type App struct {
    Path string `json:"path"`
    Desc string `json:"desc,omitempty"`
}

func main() {
    //set jsonDATA to aliasMap
    aliasMap := jsonExtractor()


    //set all FLAGS + name one
    name := flag.String("n", "", "")
    prefix := flag.String("p", "", "")
    connector := flag.String("c", "", "")
    singleArg := flag.String("a", "", "")


    //parse flags
    flag.Parse()


    // Get multiArg after parsing
    multiArg := flag.Args()


    // Early exit when nothing is provided
    if *name == "" && len(multiArg) == 0 {
        return
    }


    //set var command to commandComposer (prefix, connector, singleArgument)
    command := commandComposer(*prefix, *connector, *singleArg, multiArg)


    //filter for loop for name if matches any alias in the aliasMap
    if *name != "" {
        if app, found := aliasMap[*name]; found {
            appLauncher(app.Path, command)
        } else {
            log.Printf("%v not found\n", *name)
        }
    } else {
        // Set go routines to use nr of cores
        runtime.GOMAXPROCS(runtime.NumCPU() * 70/100)

        // Handle multiple app launches
        var wg sync.WaitGroup

        for _, app := range multiArg {
            if appEntry, found := aliasMap[app]; found{
                wg.Add(1)
                go func(appPath string){
                    defer wg.Done()
                    appLauncher(appPath, command)
                }(appEntry.Path)
            } else {
                log.Printf("%v not found\n", app)
            }
        }

        wg.Wait()
    }
}


func appLauncher(path string, command []string) {
    appToRun := exec.Command(path, command...)
    appToRun.Dir = filepath.Dir(path)
    if err := appToRun.Start(); err != nil {
        log.Fatalf("Error opening app: %v", err)
    }
}

func commandComposer(prefix, connector, singleArg string, multiArgs []string) []string{
    var command []string

    if prefix != "" && connector != "" {
        combinedArgs := strings.Join(multiArgs, connector)
        command = append(command, prefix + combinedArgs)
    } else if singleArg != "" {
        command = strings.Fields(singleArg)
    }

    return command
}

func jsonExtractor() map[string]App{
    //init jsonExtractor
    file, err := os.Open(filepath.Join(getScriptDir(), "apps.json"))
    if err != nil {
        fileT, err2 := os.Open("apps.json")
        file = fileT
        if err2 != nil {
            log.Panicf("Couldn't find file apps.json")
        }
    }

    defer file.Close()

    // Create a json decoder for the file
    decoder := json.NewDecoder(file)

    // Create a map to hold the decoded data
    aliasMap := make(map[string]App)

    // Decode the JSON data into the map
    err = decoder.Decode(&aliasMap)
    if err != nil {
        log.Fatalf("Error decoding JSON: %v", err)
    }

    return aliasMap
}


func getScriptDir() string {
    script, _ := os.Executable()
    return filepath.Dir(script)
}
