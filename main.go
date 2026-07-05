package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "math"
    "strings"
)

func main() {
    gameBody := make(map[int]map[int]string, 9)
    for i := range 3 {
        gameBody[i] = make(map[int]string, 3)
    }
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println(
`

 \\\\\\\\\\\\\\\\\\\\\\\\\\
 \\\\\\\\|         |\\\\\\\
 \\\\\\\\|TicTacGo!|\\\\\\\
 \\\\\\\\|         |\\\\\\\
 \\\\\\\\\\\\\\\\\\\\\\\\\\
`)

    isX := true
    count := 0
    currentChar := "X"
    cache := make(map[int][]int)
    initErr := drawFrame(gameBody, true, cache, &count)

    if initErr != nil {
            fmt.Println("frames failed to generate...Closing")
            os.Exit(1)
    }

    for {
        //fmt.Println("currrent cached data:", cache)
        if isX {
            fmt.Printf("\nX's turn> ")
        } else {
            fmt.Printf("\nO's turn> ")
        }
        scanner.Scan()

        if len( strings.Fields(scanner.Text())) == 0 {
            fmt.Println("Please enter a non empty value")
            continue
        }
        text := strings.Fields(scanner.Text())[0]
        pos, _ := strconv.Atoi(text)

        if (text == "return" || text == "back") && count > 1 {
            rtnErr := drawFrame(gameBody, false, cache, &count)
            if rtnErr != nil {
                fmt.Println("invalid return, exiting program")
                os.Exit(1)
            }
            if currentChar == "X" {
            currentChar = "O"
            } else { currentChar = "X" }
            isX = !isX
            continue
        }

        if pos > 9 || pos < 1 {
            fmt.Printf("\ninvalid index, please enter again\n")
            continue
        }

        for {
            row, col := int(math.Floor((float64(pos) - 1.0) / 3.0)), (pos - 1) % 3
            if _, ok := gameBody[row][col]; ok {
                fmt.Println("invalid position, please enter a different place")
                if isX {
                    fmt.Printf("X's turn> ")
                } else {
                    fmt.Printf("O's turn> ")
                }
                scanner.Scan()
                text = scanner.Text()
                pos, _ = strconv.Atoi(text)
                if pos > 9 || pos < 1 {
                    continue
                }
            } else {
                gameBody[row][col] = currentChar
                break
            }
        }
        err := drawFrame(gameBody, true, cache, &count)
        if err != nil {
            fmt.Println("frames failed to generate...Closing")
            os.Exit(1)
        }
        if currentChar == "X" {
            currentChar = "O"
        } else { currentChar = "X" }

        isX = !isX
        status := gameStatus(gameBody, count - 1)
        caching(cache, pos, count)
        if status != 0 {
            if status == 1 {
               fmt.Println(`
\\\\\\\\\\\\\\\\\\\\\\\\\\\\
\\\\\\\\|           |\\\\\\\
\\\\\\\\|   X WON!  |\\\\\\\
\\\\\\\\|           |\\\\\\\
\\\\\\\\\\\\\\\\\\\\\\\\\\\\
`)
            } else if status == -1 {
               fmt.Println(`
\\\\\\\\\\\\\\\\\\\\\\\\\\\\
\\\\\\\\|           |\\\\\\\
\\\\\\\\|   O WON!  |\\\\\\\
\\\\\\\\|           |\\\\\\\
\\\\\\\\\\\\\\\\\\\\\\\\\\\\
`)
            } else if status == 2 {
               fmt.Println(`
\\\\\\\\\\\\\\\\\\\\\\\\\\\\
\\\\\\\\|           |\\\\\\\
\\\\\\\\|   DRAW!   |\\\\\\\
\\\\\\\\|           |\\\\\\\
\\\\\\\\\\\\\\\\\\\\\\\\\\\\
`)
            }
            os.Exit(1)
        }
    }
}
