package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// 007、sort.Interface 接口
// go run 007_sorting.go
// 输出：
// 默认排列：
// title       Artist          Album              Year  Length
// -----       ------          -----              ----  ------
// Go          Delilah         From the Roots Up  2012  3m38s
// Go          Moby            Moby               1992  3m37s
// Go Ahead    Alicia Keys     As I Am            2007  4m36s
// Ready 2 Go  Martin Solveig  Smash              2011  4m24s
//
// 按 Artist 正序排列：
// title       Artist          Album              Year  Length
// -----       ------          -----              ----  ------
// Go Ahead    Alicia Keys     As I Am            2007  4m36s
// Go          Delilah         From the Roots Up  2012  3m38s
// Ready 2 Go  Martin Solveig  Smash              2011  4m24s
// Go          Moby            Moby               1992  3m37s
//
// 按 Artist 倒序排列：
// title       Artist          Album              Year  Length
// -----       ------          -----              ----  ------
// Go          Moby            Moby               1992  3m37s
// Ready 2 Go  Martin Solveig  Smash              2011  4m24s
// Go          Delilah         From the Roots Up  2012  3m38s
// Go Ahead    Alicia Keys     As I Am            2007  4m36s
//
// 按 Year 正序排列：
// title       Artist          Album              Year  Length
// -----       ------          -----              ----  ------
// Go          Moby            Moby               1992  3m37s
// Go Ahead    Alicia Keys     As I Am            2007  4m36s
// Ready 2 Go  Martin Solveig  Smash              2011  4m24s
// Go          Delilah         From the Roots Up  2012  3m38s
//
// 按 Title、Year、Length 正序排列：
// title       Artist          Album              Year  Length
// -----       ------          -----              ----  ------
// Go          Moby            Moby               1992  3m37s
// Go          Delilah         From the Roots Up  2012  3m38s
// Go Ahead    Alicia Keys     As I Am            2007  4m36s
// Ready 2 Go  Martin Solveig  Smash              2011  4m24s
//
// false
// [1 1 3 4]
// true
// [4 3 1 1]
// false
func main() {
	fmt.Println("默认排列：")
	printTracks(tracks)
	fmt.Printf("\n按 Artist 正序排列：\n")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)
	fmt.Printf("\n按 Artist 倒序排列：\n")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)
	fmt.Printf("\n按 Year 正序排列：\n")
	sort.Sort(byYear(tracks))
	printTracks(tracks)
	fmt.Printf("\n按 Title、Year、Length 正序排列：\n")
	sort.Sort(byCustom{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)

	fmt.Println()

	values := []int{3, 1, 4, 1}
	fmt.Println(sort.IntsAreSorted(values)) // false
	sort.Ints(values)
	fmt.Println(values)                     // [1 1 3 4]
	fmt.Println(sort.IntsAreSorted(values)) // true
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)                     // [4 3 1 1]
	fmt.Println(sort.IntsAreSorted(values)) // false
}

// Track *
type Track struct {
	Title, Artist, Album string
	Year                 int
	Length               time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // 计算列宽并打印表格
}

// 按照艺术家排序（切片类型）
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// 按照年份排序（切片类型）
type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// 自定义排序（结构体类型）
type byCustom struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x byCustom) Len() int           { return len(x.t) }
func (x byCustom) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x byCustom) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
