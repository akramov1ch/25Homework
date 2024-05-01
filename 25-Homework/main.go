// package main

// import (
//     "context"
//     "fmt"
//     "io"
//     "os"
//     "time"
// )


// func main() {
//     fname := "text.txt"
//     cx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//     defer cancel()

//     err := readFile(cx, fname)
//     if err != nil {
//         fmt.Println("Error: ", err)
//     }
    
// }


// func readFile(cx context.Context, fname string) error {
//     file, err := os.Open(fname)
//     if err != nil {
//         return err
//     }
//     defer file.Close()
//     bf := make([]byte, 2048)
    
    
//     for {
//         select {
//         case <-cx.Done():
            
//             fmt.Println("Operatsion canceled")
//             return cx.Err()
    
//         default:
//             n, err := file.Read(bf)
            
//             if err != nil && err != io.EOF {
//                 return err
//             }
            
//             if n > 0 {
//                 fmt.Print(string(bf[:n]))
//             }
            
//             if err == io.EOF {
//                 return nil
//             }
//         }
//     }
// }
