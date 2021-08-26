package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var srcFileName string
	var fileCount, fileSize int
	flag.StringVar(&srcFileName, "srcFileName", "gs.log", "待切割的大文件")
	flag.IntVar(&fileCount, "divideCount", 13, "切割文件个数")
	flag.IntVar(&fileSize, "divideSize", 100, "切割文件大小（M）")
	file, err := os.Open(srcFileName)
	fmt.Println("**")
	if err != nil {

		log.Fatal(err)

	}

	defer file.Close()
	dirName := strings.SplitN(srcFileName, ".", 2)[0]
	os.Mkdir(dirName, 0777)

	fsize := fileSize * 1024 * 1024
	for i := 0; i < fileCount; i++ {

		partName := fmt.Sprintf("%s/%d-%s", dirName, i+1, srcFileName)
		buf := make([]byte, fsize)
		file.Read(buf)
		destFile, err := os.OpenFile(partName, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}
		destFile.Write(buf)
		destFile.Close()
	}

}
