package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

// 019、示例：基于标记的 XML 解码
// go run ../../01、入门/src/008_fetch.go http://www.w3.org/TR/2006/REC-xml11-20060816 | go run 019_xmlselect.go div div h2
// 输出：
// html body div div h2: 1 Introduction
// html body div div h2: 2 Documents
// html body div div h2: 3 Logical Structures
// html body div div h2: 4 Physical Structures
// html body div div h2: 5 Conformance
// html body div div h2: 6 Notation
// html body div div h2: A References
// html body div div h2: B Definitions for Character Normalization
// html body div div h2: C Expansion of Entity and Character References (Non-Normative)
// html body div div h2: D Deterministic Content Models (Non-Normative)
// html body div div h2: E Autodetection of Character Encodings (Non-Normative)
// html body div div h2: F W3C XML Working Group (Non-Normative)
// html body div div h2: G W3C XML Core Working Group (Non-Normative)
// html body div div h2: H Production Notes (Non-Normative)
// html body div div h2: I Suggestions for XML Names (Non-Normative)
func main() {
	dec := xml.NewDecoder(os.Stdin)
	var stack []string
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}
