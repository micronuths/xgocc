package main

import (
	"fmt"
	"log"
	"xgocc"
)

func main() {
	s2t, err := xgocc.New("t2s")
	if err != nil {
		log.Fatal(err)
	}
	in := `自然语言处理是人工智能领域中的一个重要方向。`
	out, err := s2t.Convert(in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n%s\n", in, out)
	//自然语言处理是人工智能领域中的一个重要方向。
	//自然語言處理是人工智能領域中的一個重要方向。
}
