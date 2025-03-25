package main

import (
	"bufio"
	"fejsal/filter"
	"fejsal/reader"
	"fmt"
	"strings"
	"sync"
)

func main() {
	// 1. sample data (jsonl, csv, simple text)
	// 2. 라인 별로 나눈다.
	// 3. 라인에서 특정 데이터를 가져올 수 있도록 해야함. csv -> 인덱스로 data 에 접근, json -> key 로 value 에 접근
	//	sampleData := `{"idx":1,"who":"monkey","what":"loves","whom":"banana"},
	//{"idx":2,"who":"dog","what":"eat","whom":"banana"},
	//{"idx":3,"who":"I","what":"drink","whom":"banana smoothie"}`

	sampleData2 := `1,monkey,loves,banana
2,dog,eat,banana
3,I,drink,banana smoothie
`
	strFilter1, _ := filter.NewFilter(filter.OperatorContain, filter.ValueTypeString, "banana")
	strFilter2, _ := filter.NewFilter(filter.OperatorNotEqual, filter.ValueTypeString, "banana smoothie")

	intFilter, _ := filter.NewFilter(filter.OperatorLessThan, filter.ValueTypeNumber, 3)

	strFilter3, _ := filter.NewFilter(filter.OperatorContain, filter.ValueTypeString, "o")

	numWorkers := 3

	channels := make([]chan string, numWorkers)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		channels[i] = make(chan string, 1000) // buffered channel
		csvReader := reader.NewCSVReader()

		ft := filter.FTree{
			Left: &filter.FTree{
				Left: &filter.FTree{
					FilterSet: filter.FSet[string]{
						Filters: []filter.Filter[string]{
							strFilter1,
							strFilter2,
						},
						Condition:  filter.ConditionOr,
						DataGetter: csvReader.StringGetter(2),
					},
				},
				Right: &filter.FTree{
					FilterSet: filter.FSet[string]{
						Filters: []filter.Filter[string]{
							strFilter3,
						},
						Condition:  filter.ConditionAnd,
						DataGetter: csvReader.StringGetter(1),
					},
				},
				Condition: filter.ConditionOr,
			},
			Right: &filter.FTree{
				FilterSet: filter.FSet[int]{
					Filters: []filter.Filter[int]{
						intFilter,
					},
					Condition:  filter.ConditionAnd,
					DataGetter: csvReader.IntGetter(0),
				},
			},
			Condition: filter.ConditionAnd,
		}

		wg.Add(1)

		go func(csvReader *reader.CSVReader, ch <-chan string, wg *sync.WaitGroup, id int) {
			defer wg.Done()

			for line := range ch {
				csvReader.InputStream(strings.NewReader(line))
				if ok := csvReader.LoadNextLine(); ok {
					if ft.Evaluate() {
						fmt.Println(line)
					}
				}
			}
		}(csvReader, channels[i], &wg, i)
	}

	scanner := bufio.NewScanner(strings.NewReader(sampleData2))
	i := 0
	for scanner.Scan() {
		channels[i%numWorkers] <- scanner.Text()
		i++
	}

	for _, ch := range channels {
		close(ch)
	}

	wg.Wait()
	fmt.Println("Done")

}
