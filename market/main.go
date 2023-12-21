package main

import (
	"fmt"
	"os"
)

func main() {
    argsWithoutProg := os.Args[1:]

    if len(argsWithoutProg) == 0 {
        fmt.Println("Usage: my-program <command>")
        fmt.Println("Available commands:")
        fmt.Println("- server")
        fmt.Println("- sync-models")
        fmt.Println("- seed-category")
        fmt.Println("- sync-views")
        fmt.Println("- update-deposits")
        fmt.Println("- user <username> <action> <role>")
        fmt.Println("- index")
        fmt.Println("- import-metro")
        fmt.Println("- staff-stats")
        return
    }

    switch argsWithoutProg[0] {
        case "server":
            runServer()
        case "sync-models":
            syncModels()
        case "seed-category":
            seedCategories()
        case "sync-views":
            syncDatabaseViews()
        // case "update-deposits":
        //     updateDeposits()
        case "user":
            if len(argsWithoutProg) < 4 {
                fmt.Println("Usage: my-program user <username> <action> <role>")
                return
            }
            username, action, role := argsWithoutProg[1], argsWithoutProg[2], argsWithoutProg[3]
            manageRole(username, action, role)
        // case "index":
        //     indexItems()
        // case "import-metro":
        //     importMetroStations()
        // case "staff-stats":
        //     staffStats()
        default:
            fmt.Println("Unknown command:", argsWithoutProg[0])
    }
}

