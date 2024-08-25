package main

import (
    "github.com/joho/godotenv"
    "os/exec"
    "fmt"
    "os"
)

func main() {
    if err := godotenv.Load(); err != nil {
        panic(err)
    }

    cmd := exec.Command(
        "tern",
        "migrate",
        "--migrations",
        "./internal/store/pgstore/migrations",
        "--config",
        "./internal/store/pgstore/migrations/tern.conf",
    )

    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Run(); err != nil {
        fmt.Printf("Command failed with error: %v\n", err)
        os.Exit(1)
    }
}
