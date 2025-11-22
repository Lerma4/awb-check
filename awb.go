package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    rng := rand.New(rand.NewSource(time.Now().UnixNano()))
    for {
        fmt.Print("Inserisci le prime 7 cifre dell'AWB, oppure 'r' per generare un casuale (Ctrl+C per uscire): ")
        input, err := reader.ReadString('\n')
        if err != nil {
            return
        }
        input = strings.TrimSpace(input)

        // Comando speciale: genera 7 AWB casuali con CIN
        if strings.EqualFold(input, "r") {
            awb := rng.Intn(10_000_000) // 0..9,999,999
            cin := awb % 7
            fmt.Printf("%07d%d\n", awb, cin)
            continue
        }

        if len(input) != 7 {
            fmt.Println("Errore: devi inserire esattamente 7 cifre.")
            continue
        }

        awbNum, err := strconv.Atoi(input)
        if err != nil {
            fmt.Println("Errore: l'input deve essere un numero.")
            continue
        }

        cin := awbNum % 7
        fmt.Printf("Il CIN (ottava cifra) per l'AWB %s è: %d\n", input, cin)
    }
}
