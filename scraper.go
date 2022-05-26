package main
import 	(
    "bufio"
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "strconv"
   // "github.com/gocolly/colly/v2"
)

type PuppySearch struct {
    Breed string
    Gender string
    ZipCode string
    Radius int
}

func main() {
    validBreeds, err := readFile("breeds.txt")
    if err != nil {
        log.Fatal(err)
    }

    puppySearchList := readInputCsv("input.csv")

    //c := colly.NewCollector()
}

func readInputCsv(fileName string) []PuppySearch {
    f, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    puppySearchList := createPuppySearchList(data)

    return puppySearchList
}

func createPuppySearchList(data [][]string) []PuppySearch {
    var puppySearchList []PuppySearch
    for i, line := range data {
        if i > 0 { // omit header line
            var puppySearch PuppySearch
            for j, field := range line {
                if j == 0 {
                    puppySearch.Breed = field
                } else if j == 1 {
                    puppySearch.Gender = field
                } else if j == 2 {
                    puppySearch.ZipCode = field
                } else if j == 3 {
                    radius, err := strconv.Atoi(field)
                    if err != nil {
                        log.Fatal(err)
                    }
                    puppySearch.Radius = radius
                }
            }
            puppySearchList = append(puppySearchList, puppySearch)
        }
    }
    return puppySearchList
}

func readFile(filePath string) ([]string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }

    defer func() {
        if err = file.Close(); err != nil {
            log.Fatal(err)
        }
    }()

    var result []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        result = append(result, scanner.Text())
    }
    return result, scanner.Err()
}