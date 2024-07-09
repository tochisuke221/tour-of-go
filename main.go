package main

import (
	"flag"
	"log"
	"path/filepath"
	"io/fs"
	"strings"

	"tour-of-go/converter"
)

func main() {
	var srcFormat, dstFormat string

	flag.StringVar(&srcFormat, "src", "jpg", "変換前のフォーマット")
	flag.StringVar(&dstFormat, "dst", "png", "変換後のフォーマット")

	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatal("ディレクトリが指定されていません")
	}


	rootDir := flag.Arg(0)

	err := filepath.Walk(rootDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), "."+srcFormat) {
			newPath := strings.TrimSuffix(path, "."+srcFormat) + "." + dstFormat
			if err := converter.Convert(path, newPath, srcFormat, dstFormat); err != nil {
				log.Printf("failed to convert %s: %v", path, err)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("error walking the path %q: %v\n", rootDir, err)
	}
}