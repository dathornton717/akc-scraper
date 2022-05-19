package main
import 	(
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
    f, err := os.Open("input.csv")
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

    fmt.Printf("%+v\n", puppySearchList)

    //c := colly.NewCollector()
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