package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	out, err := getDiffLink(os.Args[1], os.Args[2])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(out)
}

// Takes in two url links and returns a url with the diff between the 2 links
func getDiffLink(url1 string, url2 string) (string, error) {
	url1 = strings.ReplaceAll(url1, "?", "/")
	url2 = strings.ReplaceAll(url2, "?", "/")
	location1 := -1
	arr1 := strings.Split(url1, "/")
	for i, v := range arr1 {
		if v == "commit" {
			location1 = i + 1
		}
	}
	if location1 == -1 {
		return "", fmt.Errorf("url1 bad url")
	}
	location2 := -1
	arr2 := strings.Split(url2, "/")
	for i, v := range arr2 {
		if v == "commit" {
			location2 = i + 1
		}
	}
	if location2 == -1 {
		return "", fmt.Errorf("url2 bad url")
	}
	outUrlBase := arr1[1 : location1-1]
	outUrl := strings.Join(outUrlBase, "/")
	return fmt.Sprintf("https:/%s/branchCompare?baseVersion=GC%s&targetVersion=GC%s", outUrl, arr1[location1], arr2[location2]), nil
}
