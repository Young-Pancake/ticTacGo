package main
import (
    "fmt"
    "math"
)

func drawFrame(body map[int]map[int]string, drawMode bool, cache map[int][]int, count *int) error {
    template := []rune(`

         |       |
     1   |   2   |   3
  _______|_______|_______
         |       |
     4   |   5   |   6
  _______|_______|_______
         |       |
     7   |   8   |   9
         |       |
`)

	placeHolders := [][]int{[]int{26, 34, 42},
                            []int{94, 102, 110},
                            []int{162, 170, 178}}
    for keyRow, rows := range body {
        for keyCol, val := range rows {
                template[placeHolders[keyRow][keyCol]] = []rune(val)[0]
        }
    }
    if drawMode {
        *count++
    } else {
        template[placeHolders[rune(cache[*count - 1][0])][rune(cache[*count - 1][1])]] = ' '

        delete(body[cache[*count - 1][0]], cache[*count - 1][1])
        delete(cache, *count - 1)
        *count--
    }
	fmt.Println(string(template))
    return nil
}

func gameStatus(body map[int]map[int]string, count int) int {

    switch true {
    case body[0][0] == "X" && body[0][1] == "X" && body[0][2] == "X":
        return 1
    case body[1][0] == "X" && body[1][1] == "X" && body[1][2] == "X":
        return 1
    case body[2][0] == "X" && body[2][1] == "X" && body[2][2] == "X":
        return 1
    case body[0][0] == "X" && body[1][0] == "X" && body[2][0] == "X":
        return 1
    case body[0][1] == "X" && body[1][1] == "X" && body[2][1] == "X":
        return 1
    case body[0][2] == "X" && body[1][2] == "X" && body[2][2] == "X":
        return 1
    case body[0][0] == "X" && body[1][1] == "X" && body[2][2] == "X":
        return 1
    case body[0][2] == "X" && body[1][1] == "X" && body[2][0] == "X":
        return 1


    case body[0][0] == "O" && body[0][1] == "O" && body[0][2] == "O":
        return -1
    case body[1][0] == "O" && body[1][1] == "O" && body[1][2] == "O":
        return -1
    case body[2][0] == "O" && body[2][1] == "O" && body[2][2] == "O":
        return -1
    case body[0][0] == "O" && body[1][0] == "O" && body[2][0] == "O":
        return -1
    case body[0][1] == "O" && body[1][1] == "O" && body[2][1] == "O":
        return -1
    case body[0][2] == "O" && body[1][2] == "O" && body[2][2] == "O":
        return -1
    case body[0][0] == "O" && body[1][1] == "O" && body[2][2] == "O":
        return -1
    case body[0][2] == "O" && body[1][1] == "O" && body[2][0] == "O":
        return -1
    case count == 9:
        return 2
    default:
        return 0
    }
}

func caching(cache map[int][]int, pos, count int) {
    cache[count - 1] = []int{int(math.Floor(float64(pos - 1) / 3.0)), (pos - 1) % 3}
}
