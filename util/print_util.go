package util

import "fmt"

func PrintInTextPlain(resultMap map[string]string) {
	fmt.Println()
	for k, v := range resultMap {
		fmt.Printf("%s %s\n", k, v)
	}
}

func PrintInMarkdownFormat(resultMap map[string]string) {
	fmt.Println()
	for path, url := range resultMap {
		name := GetFileName(path)
		fmt.Printf("![%s](%s)\n", name, url)
	}
}
