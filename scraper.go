package main
import 	(
    "bufio"
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "strings"
    "github.com/gocolly/colly/v2"
)


func main() {
    validBreeds, err := readFile("breeds.txt")
    if err != nil {
        log.Fatal(err)
    }

    puppySearchList := readInputCsv("input.csv")

    for i, puppySearch := range puppySearchList {
        err = puppySearch.validate()

        if err != nil {
            fmt.Printf("skipping due to validation error on line %d: %s\n", i + 1, err.Error())
        }
        if !validBreeds[puppySearch.Breed] {
            fmt.Printf("%s is not a valid breed, skipping\n", puppySearch.Breed)
            continue
        }

        c := colly.NewCollector(
            colly.AllowedDomains("marketplace.akc.org"),
        )
        c.OnRequest(func(r *colly.Request) {
            fmt.Printf("Visiting %s\n", r.URL.String())
        })
        c.OnHTML("a[href]", func(e *colly.HTMLElement) {
            link := e.Attr("href")
            // Print link
            fmt.Printf("Link found: %q -> %s\n", e.Text, link)
        })

        url := "https://marketplace.akc.org/puppies/"
        url = url + strings.ReplaceAll(puppySearch.Breed, " ", "-")
        url = url + puppySearch.createRequestParams(1)
        c.Visit(url)
    }
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
                    puppySearch.Radius = field
                }
            }
            puppySearchList = append(puppySearchList, puppySearch)
        }
    }
    return puppySearchList
}

func readFile(filePath string) (map[string] bool, error) {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }

    defer func() {
        if err = file.Close(); err != nil {
            log.Fatal(err)
        }
    }()

    result := make(map[string] bool)
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        result[scanner.Text()] = true
    }
    return result, scanner.Err()
}