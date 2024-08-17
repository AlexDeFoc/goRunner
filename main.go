package main

import (
    "encoding/json"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
)

type Config struct {
    Advanced_Status bool `json:"Advanced mode"`

    Shell string `json:"shell"`
}

type App struct {
    Path string `json:"path"`
    Desc string `json:"desc"`
}

type Alias struct {
    App_Name string `json:"app"`

    Status string `json:"status"`

    Command string `json:"command"`

    Prefix string `json:"prefix"`
    Connect string `json:"connect"`
    Fallback string `json:"fallback"`

    Desc string `json:"desc"`
}

func main () {
    /// Exit if no names are given
    if len(os.Args) < 2 {
        fmt.Printf("Please provide name\n")
        os.Exit(1)
    }

    /// Setup Config
    settings := setupConfig()


    /// Init the lists
    var Alias_List map[string]Alias
    var App_List map[string]App
    Load_List_From_File("alias.json", &Alias_List)
    Load_List_From_File("app.json", &App_List)

    if (os.Args[1]) == "list" {
        List_Available_Aliases(&App_List, &Alias_List)
        os.Exit(0)
    }

    if settings.Advanced_Status == false {
        Run_Multi(&App_List)
        os.Exit(0)
    }

    termName := Get_Terminal_Name()

    // Check if termName exists in list
    if !Check_Name_Exists(termName, &Alias_List) && !Check_Name_Exists(termName, &App_List) {
        fmt.Printf("Please provide valid name\nTry 'list' to list the names\n\n")
        Print(*termName)
        os.Exit(1)
    } else if !Check_Name_Exists(termName, &Alias_List) {
        Run_Multi(&App_List)
        os.Exit(0)
    }

    // Get status from aliase
    status := Get_Status_From_Alias(termName, &Alias_List)
    switch status {
    case "command":
        Run_Command(termName, &settings.Shell, &Alias_List)
    case "search":
        Run_Search(termName, &App_List, &Alias_List)
    case "pass":
        Run_Pass(termName, &App_List, &Alias_List)
    case "call":
        Run_Call(termName, &settings, &App_List, &Alias_List)
    default:
        Print("Status not found")
    }
}


func Run_Multi (list *map[string]App) {
    names := os.Args[1:]
    var cmd *exec.Cmd
    var err error

    for _, appName := range names {
        cmd = exec.Command((*list)[appName].Path)
        cmd.Dir = filepath.Dir((*list)[appName].Path)
        err = cmd.Start()
        BlowUp(err)
    }
}




















func Run_Call (name *string, settings *Config, App_List *map[string]App, Alias_List *map[string]Alias) {
    //Print("Calling an alias\n")

    alias_Name := (*Alias_List)[*name].App_Name

    if entry, ok := (*Alias_List)[alias_Name]; ok {
        entry.Command = (*Alias_List)[*name].Command
        entry.Prefix= (*Alias_List)[*name].Prefix
        entry.Connect = (*Alias_List)[*name].Connect
        entry.Fallback = (*Alias_List)[*name].Fallback

        (*Alias_List)[alias_Name] = entry
    }

    switch (*Alias_List)[alias_Name].Status {
    case "command":
        Run_Command(&alias_Name, &settings.Shell, Alias_List)
    case "search":
        Run_Search(&alias_Name, App_List, Alias_List)
    case "pass":
        Run_Pass(&alias_Name, App_List, Alias_List)
    case "call":
        Run_Call(&alias_Name, settings, App_List, Alias_List)
    default:
        Print("Status not found")
    }

}










































func Run_Pass (name *string, App_List *map[string]App, Alias_List *map[string]Alias) {
    //Print("Passing to an app\n")
    
    item := (*Alias_List)[*name]

    // Check for overrides
    if item.Prefix == "" {
        item.Prefix = (*Alias_List)[item.App_Name].Prefix
    }

    path := (*App_List)[item.App_Name].Path
    command := strings.Split(item.Command, " ")

    cmd := exec.Command(path, command...)
    cmd.Dir = filepath.Dir(path)

    err := cmd.Start()
    BlowUp(err)
}




















func Run_Search (name *string, App_List *map[string]App, Alias_List *map[string]Alias) {
    //Print("Running a search\n")
    
    item := (*Alias_List)[*name]

    // Check for overrides
    if item.Prefix == "" {
        item.Prefix = (*Alias_List)[item.App_Name].Prefix
    }
    if item.Connect == "" {
        item.Connect = (*Alias_List)[item.App_Name].Connect
    }

    path := (*App_List)[item.App_Name].Path
    prefix := item.Prefix
    connect := item.Connect
    arguments := strings.Join(os.Args[2:], connect)
    command := prefix + arguments

    fallback := item.Fallback

    var cmd *exec.Cmd
    if len(arguments) != 0 {
        cmd = exec.Command(path, command)
    } else {
        cmd = exec.Command(path, fallback)
    }

    cmd.Dir = filepath.Dir(path)

    err := cmd.Start()
    BlowUp(err)
}




func Run_Command (name *string, shell *string, Alias_List *map[string]Alias) {
    //Print("Running a command\n")

    item := (*Alias_List)[*name]

    // Check for overrides
    if item.Command == "" {
        item.Command = (*Alias_List)[item.App_Name].Command
    }

    commandPrefix := Make_Cmd_Prefix(*shell)
    command := item.Command

    cmd := exec.Command(*shell, commandPrefix, command)

    err := cmd.Start()
    BlowUp(err)
}

func Make_Cmd_Prefix (shell string) string{
    switch shell {

    case "cmd": 
        return "/C"

    case "pwsh":
        return "-Command"

    case "bash":
        return "-c"

    default:
        fmt.Print("\nError running command. Please add the shell you prefer in the config\n\n")
        os.Exit(1)
        return ""
    }
}
































func Get_Status_From_Alias (aliasName *string, list *map[string]Alias) string {
    return (*list)[*aliasName].Status
}

func Check_Name_Exists [T any] (name *string, list *map[string]T) bool {
    if _, found := (*list)[*name]; found {
        return true
    }

    return false
}

func Load_List_From_File [T any] (fileName string, list *map[string]T) {
    file, err := os.Open(Find_File_Path(fileName))
    defer file.Close()
    BlowUp(err)

    decoder := json.NewDecoder(file)

    err = decoder.Decode(list)
    BlowUp(err)
}

func List_Available_Aliases (App_List  *map[string]App, Alias_List *map[string]Alias) {
    var Alias_Names []string
    var App_Names []string

    for name := range *Alias_List {
        Alias_Names = append(Alias_Names, name)
    }
    for name := range *App_List {
        App_Names = append(App_Names, name)
    }

    Print("-------- App & Alias Names --------\n")

    Print("--- App Names ---\n")
    for _, name := range App_Names {
        fmt.Printf("    -> %v", name)
        fmt.Printf(":  %v\n", (*App_List)[name].Desc)
    }

    Print("\n\n--- Alias Names ---\n")
    for _, name := range Alias_Names {
        fmt.Printf("    -> %v", name)
        fmt.Printf(":  %v\n", (*Alias_List)[name].Desc)
    }

    Print("\n")
}

func Find_File_Path (fileName string) string{
    exe, err := os.Executable()
    BlowUp(err)

    path := filepath.Join(filepath.Dir(exe), fileName)

    return path
}

func BlowUp (err error) {
    if err != nil {
        fmt.Printf("Error: %v", err)
    }
}

func Print (stuff any) {
    fmt.Println(stuff)
}






















func Get_Terminal_Name () *string {
    return &os.Args[1]
}

func Match_Name [T any] (name *string, list *map[string]T) bool {
    if _, found := (*list)[*name]; found {
        return true
    }
    return false
}




















func setupConfig() Config {
    var configData Config

    configFile, err := os.Open(filepath.Join(getScriptDir(), "config.json"))
    if err != nil {
        fmt.Printf("\nError opening config file:\n%v\n\n", err)
        os.Exit(1)
    }
    defer configFile.Close()

    decoder := json.NewDecoder(configFile)

    err = decoder.Decode(&configData)

    if err != nil {
        fmt.Printf("\nError decoding config file:\n%v\n\n", err)
        os.Exit(1)
    }

    return configData
}

func getScriptDir() string {
    script, _ := os.Executable()
    return filepath.Dir(script)
}
