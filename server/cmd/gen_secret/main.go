package main

import (
    crand "crypto/rand"
    "fmt"
)

func secureRandomStr(b int) string {
    k := make([]byte, b)
    if _, err := crand.Read(k); err != nil {
        panic(err)
    }
    return fmt.Sprintf("%x", k)
}

func main() {
    fmt.Println(secureRandomStr(16))
}
