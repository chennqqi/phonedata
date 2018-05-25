package main

import (
	"fmt"
	"os"

	"github.com/chennqqi/phonedata"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("请输入手机号")
		return
	}
	pd, err := phonedata.Parse("../phone.dat")
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	pr, err := pd.Find(os.Args[1])
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Print(pr)
}
