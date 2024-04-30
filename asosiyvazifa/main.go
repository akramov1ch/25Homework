package main

import (
    "context"
    "fmt"
    "io"
    "os"
    "time"
)

func readFile(ctx context.Context, filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    buf := make([]byte, 1024)
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Operatsiya bekor qilindi")
            return ctx.Err()
        default:
            n, err := file.Read(buf)
            if err != nil && err != io.EOF {
                return err
            }
            if n > 0 {
                fmt.Print(string(buf[:n]))
            }
            if err == io.EOF {
                return nil
            }
        }
    }
}

func main() {
    filename := "/home/shakhboz/Downloads/9a147f9f-65d1-4d86-95b2-9a6ca77c1d28/25-dars. Continue_working_with_go_postgresql_driver/asosiyvazifa/file.txt"

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    err := readFile(ctx, filename)
    if err != nil {
        fmt.Println("Xatolik yuz berdi:", err)
    }
}
