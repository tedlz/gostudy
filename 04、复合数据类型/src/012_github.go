package main

import (
	"fmt"
	"gostudy/04、复合数据类型/files"
	"log"
	"os"
)

// 012、JSON github
// go run 012_github.go repo:golang/go is:open json decoder
// 输出：
// 45 issues:
// #36225     dsnet encoding/json: the Decoder.Decode API lends itself to m
// #33416   bserdar encoding/json: This CL adds Decoder.InternKeys
// #34647 babolivie encoding/json: fix byte counter increments when using d
// #5901        rsc proposal: encoding/json: allow override type marshaling
// #29035    jaswdr proposal: encoding/json: add error var to compare  the
// #34543  maxatome encoding/json: Unmarshal & json.(*Decoder).Token report
// #32779       rsc proposal: encoding/json: memoize strings during decode?
// #28923     mvdan encoding/json: speed up the decoding scanner
// #11046     kurin encoding/json: Decoder internally buffers full input
// #12001 lukescott encoding/json: Marshaler/Unmarshaler not stream friendl
// #34564  mdempsky go/internal/gcimporter: single source of truth for deco
// #31789  mgritter encoding/json: provide a way to limit recursion depth
// #31701    lr1980 encoding/json: second decode after error impossible
// #30301     zelch encoding/xml: option to treat unknown fields as an erro
// #33854     Qhesz encoding/json: unmarshal option to treat omitted fields
// #26946    deuill encoding/json: clarify what happens when unmarshaling i
// #36353     dsnet proposal: encoding/gob: allow override type marshaling
// #16212 josharian encoding/json: do all reflect work before decoding
// #33835     Qhesz encoding/json: unmarshalling null into non-nullable gol
// #6647    btracey x/tools/cmd/godoc: display type kind of each named type
// #27179  lavalamp encoding/json: no way to preserve the order of map keys
// #28143    arp242 proposal: encoding/json: add "readonly" tag
// #22752  buyology proposal: encoding/json: add access to the underlying d
// #14750 cyberphon encoding/json: parser ignores the case of member names
// #28189     adnsv encoding/json: confusing errors when unmarshaling custo
// #22480     blixt proposal: encoding/json: add omitnil option
// #33714    flimzy proposal: encoding/json: Opt-in for true streaming supp
// #7872  extempora encoding/json: Encoder internally buffers full output
// #19636 josselin- encoding/base64: decoding is slow
// #30701 LouAdrien encoding/json: ignore tag "-" not working on embedded s
func main() {
	result, err := files.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
