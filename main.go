package main

import (
	"flag"
	"fmt"
	"github.com/arbovm/levenshtein"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
)

type vid struct {
	file     os.FileInfo
	distance int
}

type byDistance []vid

func (v byDistance) Len() int           { return len(v) }
func (v byDistance) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v byDistance) Less(i, j int) bool { return v[i].distance < v[j].distance }

// Arguments
var show = flag.Bool("show", false, "")
var okFirst = flag.Bool("okFirst", false, "")

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	flag.Parse()
	// Gather all srt and other files
	var subs []os.FileInfo
	var vids []vid
	for i := 0; i < len(files); i++ {
		if files[i].IsDir() {
			continue
		}
		if filepath.Ext(files[i].Name()) == ".srt" {
			subs = append(subs, files[i])
		} else {
			vids = append(vids, vid{file: files[i]})
		}
	}
	// Loop and srt and compute distance
	for _, sub := range subs {
		for i := 0; i < len(vids); i++ {
			vids[i].distance = levenshtein.Distance(sub.Name(), vids[i].file.Name())
		}
		sort.Sort(byDistance(vids))
		if *okFirst {
			rename(sub, vids[0])
			continue
		}
		if *show {
			fmt.Printf("%s ---> %s\n", sub.Name(), vids[0].file.Name())
			continue
		}
		fmt.Printf("%s :\n\t0 %s\n\t1 %s\n\t2 %s\nChoose:", sub.Name(),
			vids[0].file.Name(), vids[1].file.Name(), vids[2].file.Name())
		var r int
		_, err := fmt.Scanf("%d", &r)
		if err != nil || r > 2 {
			// skip
			continue
		}
		// rename sub with name from number r
		rename(sub, vids[r])
	}
}

func rename(sub os.FileInfo, vid vid) {
	var vidName = vid.file.Name()
	ext := filepath.Ext(vidName)
	var newName = vidName[0:len(vidName)-len(ext)] + ".srt"
	fmt.Printf("Renaming %s -> %s", sub.Name(), newName)
	err := os.Rename(sub.Name(), newName)
	if err != nil {
		panic(err)
	}
}
