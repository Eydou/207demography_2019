//
// EPITECH PROJECT, 2020
// 207demography_2019
// File description:
// error
//

package functions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

//description ------- xy = sumTotal : x = sumDate : y = sumCountry : X = sumPowDate

type dataParse struct {
	country []string
	data    [][]int
	dates   []int
	rootM1  float64
	rootM2  float64
}

const (
	//Million of population
	Million = 1000000
)

//Print country of args in ascending
func (st *dataParse) PrintCountry() {
	fmt.Printf("Country: ")
	for i := 0; i != len(st.country); i++ {
		if i+1 != len(st.country) {
			fmt.Printf("%s, ", st.country[i])
		} else {
			fmt.Printf("%s\n", st.country[i])
		}
	}
}

//Fit2 algo and printing
func (st *dataParse) Fit2(sumCountry int, sumDate int, sumPowCountry float64, sumTotal int,
	dataCountry []int) {
	a, b, predict := float64(0), float64(0), float64(0)

	aBis := float64((len(st.dates)))
	aC := (float64(sumCountry) * float64(sumCountry))
	a = (float64(sumDate)*sumPowCountry - float64(sumCountry)*float64(sumTotal)) /
		(aBis*sumPowCountry - aC)

	b = (float64(len(st.dates))*float64(sumTotal) - float64(sumCountry)*float64(sumDate)) /
		(float64(len(st.dates))*sumPowCountry - aC)
	predict = ((2050 - a) / b) / Million
	if a < 0 {
		fmt.Printf("Fit2\n\tX = %.2f Y - %.2f\n", math.Abs(b)*Million, math.Abs(a))
	} else {
		fmt.Printf("Fit2\n\tX = %.2f Y + %.2f\n", math.Abs(b)*Million, math.Abs(a))
	}
	for i := 0; i != len(st.dates); i++ {
		st.rootM2 += math.Pow(
			((float64(st.dates[i])-a)/b)-float64(dataCountry[i]), 2) /
			float64(len(st.dates))
	}
	st.rootM2 = math.Sqrt(st.rootM2) / Million
	fmt.Printf("\tRoot-mean-square deviation: %.2f\n", st.rootM2)
	fmt.Printf("\tPopulation in 2050: %.2f\n", predict)
}

//Fit1 algo and printing
func (st *dataParse) Fit1(sumDate int, sumCountry int, sumPowDate int, sumTotal int,
	dataCountry []int) {
	a, b, predict := float64(0), float64(0), float64(0)

	a = (float64(sumCountry)*float64(sumPowDate) - float64(sumDate)*float64(sumTotal)) /
		(float64((len(st.dates))*sumPowDate - (sumDate * sumDate)))

	b = (float64(len(st.dates))*float64(sumTotal) - float64(sumDate)*float64(sumCountry)) /
		(float64(len(st.dates))*float64(sumPowDate) - (float64(sumDate) * float64(sumDate)))
	predict = a/Million + (2050 * b / Million)
	if a < 0 {
		fmt.Printf("Fit1\n\tY = %.2f X - %.2f\n", math.Abs(b)/Million, math.Abs(a/Million))
	} else {
		fmt.Printf("Fit1\n\tY = %.2f X + %.2f\n", math.Abs(b)/Million, math.Abs(a)/Million)
	}
	for i := 0; i != len(st.dates); i++ {
		st.rootM1 += math.Pow(((float64(st.dates[i])*b+a)-float64(dataCountry[i])), 2) /
			float64(len(st.dates))
	}
	st.rootM1 = math.Sqrt(st.rootM1) / Million
	fmt.Printf("\tRoot-mean-square deviation: %.2f\n", st.rootM1)
	fmt.Printf("\tPopulation in 2050: %.2f\n", predict)
}

//Start algorithm of demography
func (st *dataParse) StartAlgoDemo() {
	// ----------------- get the sums ----------------
	sumCountry, sumDate, sumPowDate, sumTotal := 0, 0, 0, 0
	sumPowCountry := float64(0)
	dataCountry := make([]int, len(st.dates))

	for _, num := range st.dates {
		sumDate += num
	}
	for _, num := range st.dates {
		sumPowDate += (num * num)
	}
	for i := 0; i != len(st.data); i++ {
		for j := 0; j != len(st.dates); j++ {
			dataCountry[j] += st.data[i][j]
		}
	}
	for _, num := range dataCountry {
		sumCountry += num
	}
	for i := 0; i != len(st.dates); i++ {
		sumTotal += (st.dates[i] * dataCountry[i])
	}
	for i := 0; i != len(dataCountry); i++ {
		sumPowCountry += (float64(dataCountry[i]) * float64(dataCountry[i]))
	}
	// ----------------- Print Algo ----------------
	st.PrintCountry()
	st.Fit1(sumDate, sumCountry, sumPowDate, sumTotal, dataCountry)
	st.Fit2(sumCountry, sumDate, sumPowCountry, sumTotal, dataCountry)
	fmt.Printf("Correlation: %.4f\n", st.rootM1/st.rootM2)
}

func (st *dataParse) AddData(data string, check int) {
	populations := make([]int, 0)
	recupData := strings.Split(data, "\n")
	splitData := strings.Split(recupData[check], ";")

	for i := 2; i != len(splitData); i++ {
		intData, err := strconv.Atoi(splitData[i])
		if err != nil {
			fmt.Println("Parsing Error", err)
			os.Exit(84)
		}
		populations = append(populations, intData)
	}
	st.data = append(st.data, populations)
}

//Start the loop of the project
func (st *dataParse) LoopDemography(data string, args []string, reader *bufio.Reader,
	err error) {
	var line string
	check := 0
	for ; ; check++ {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		split := strings.Split(line, ";")
		if split[0] == "Country Name" {
			recupDate := strings.Split(data, "\n")
			_Date := strings.Split(recupDate[0], ";")
			for i := 2; i != len(_Date); i++ {
				intDate, _ := strconv.Atoi(_Date[i])
				st.dates = append(st.dates, intDate)
			}
		}
		for i := 1; i != len(args); i++ {
			if args[i] == split[1] {
				st.country = append(st.country, split[0])
				st.AddData(data, check)
				break
			}
		}
	}
	st.StartAlgoDemo()
}

//MathParse args
func MathParse(data string, args []string) int {
	st := dataParse{}
	files, err := os.Open("207demography_data.csv")
	defer files.Close()
	if err != nil {
		fmt.Println("File reading error", err)
		return 84
	}
	reader := bufio.NewReader(files)
	st.LoopDemography(data, args, reader, err)
	if len(st.country) == 0 || len(st.country) != len(args)-1 {
		fmt.Println("Error: invalid country in parameter")
		return 84
	}
	sort.Sort(sort.StringSlice(st.country))
	return 0
}
