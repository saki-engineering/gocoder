package gocoder

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

type city struct {
	name, pref, year, id int
}

type cityList []city

func (a cityList) Len() int      { return len(a) }
func (a cityList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

type byName struct{ cityList }

func (a byName) Less(i, j int) bool { return a.cityList[i].name < a.cityList[j].name }

type byYear struct{ cityList }

func (a byYear) Less(i, j int) bool {
	if a.cityList[i].pref != a.cityList[j].pref {
		return a.cityList[i].pref < a.cityList[j].pref
	}
	return a.cityList[i].year < a.cityList[j].year
}

func main() {
	sc.Split(bufio.ScanWords)

	_, M := scanInt(), scanInt()
	L := make(cityList, M)
	for i := 0; i < M; i++ {
		L[i].name, L[i].pref, L[i].year = i+1, scanInt(), scanInt()
	}

	sort.Sort(byYear{L})
	L[0].id = 1
	for i := 1; i < M; i++ {
		if L[i].pref == L[i-1].pref {
			L[i].id = L[i-1].id + 1
		} else {
			L[i].id = 1
		}
	}

	sort.Sort(byName{L})
	for _, l := range L {
		fmt.Printf("%06d%06d\n", l.pref, l.id)
	}
}
