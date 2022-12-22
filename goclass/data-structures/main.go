package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	words := make(map[string]int)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		words[scan.Text()]++
	}

	type kv struct {
		key string
		val int
	}

	var ss []kv

	for k, v := range words {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].val > ss[j].val
	})

	fmt.Println("==== Top 3 ====")
	for _, s := range ss[:3] {
		fmt.Println(s.key, "appears", s.val, "times")
	}

	fmt.Println("\n==== Total ====")
	fmt.Println(len(words), "unique words")
}
