package readFile

import ( 
	"bufio"
    "fmt"
    "io"
    "strconv"
    "strings"
)

func ReadInts(fileName io.Reader) ([]int, error) {
    scanner := bufio.NewScanner(fileName)
    scanner.Split(bufio.ScanWords)
    var result []int
    for scanner.Scan() {
        x, err := strconv.Atoi(scanner.Text())
        if err != nil {
            return result, err
        }
        result = append(result, x)
    }
    return result, scanner.Err()
}