// +build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/99designs/gqlgen/cmd"
)

func main() {
	// ctx := context.Background()
	dir := "./api"
	apath, err := filepath.Abs(dir)
	if err != nil {
		log.Fatalf("could not get absolute path %s", err)
	}
	grdir, err := filepath.Abs("../../internal/graphql/resolver")
	if err != nil {
		log.Fatalf("could not get absolute path for genresolver %s", err)
	}
	grfile := filepath.Join(grdir, "resolver.go")
	// if generated resolver file already exists, remove it
	if _, err := os.Stat(grfile); err == nil {
		os.Remove(grfile)
		fmt.Printf("successfully removed genresolver file %s\n", grfile)
	}
	// create temporary directory for genresolver
	if _, err := os.Stat(grdir); os.IsNotExist(err) {
		// 0777 denotes read, write, & execute for owner, group and others
		os.Mkdir(grdir, 0777)
	}
	// create temporary directory for schema if it doesn't exist
	if _, err := os.Stat(apath); os.IsNotExist(err) {
		os.Mkdir(apath, 0777)
	}
	// defer cleanup(apath)
	// client := github.NewClient(nil)
	// _, files, _, err := client.Repositories.GetContents(ctx, "dictyBase", "graphql-schema", "/", nil)
	// if err != nil {
	// 	log.Fatalf("error in getting graphql schema", err)
	// }
	// for _, n := range files {
	// 	file, _, _, err := client.Repositories.GetContents(ctx, "dictyBase", "graphql-schema", n.GetName(), nil)
	// 	if err != nil {
	// 		log.Fatalf("error in getting individual schema file", err)
	// 	}
	// 	// need to decode file contents
	// 	d, err := base64.StdEncoding.DecodeString(*file.Content)
	// 	if err != nil {
	// 		log.Fatalf("error decoding github file contents", err)
	// 	}
	// 	fp := filepath.Join(apath, file.GetName())
	// 	err = ioutil.WriteFile(fp, d, 0777)
	// 	if err != nil {
	// 		log.Fatalf("error writing file", err)
	// 	}
	// 	fmt.Printf("successfully wrote file %s\n", fp)
	// }
	cmd.Execute()
}

func cleanup(path string) {
	s := strings.Split(path, "/")
	tmpfile := filepath.Join(path, "query.graphql")
	if _, err := os.Stat(tmpfile); err == nil {
		if s[len(s)-2] == "graphql-server" {
			err = os.RemoveAll(path)
			if err != nil {
				fmt.Printf("unable to remove temporary schema directory %s", err)
				os.Exit(1)
			}
			fmt.Printf("successfully removed temporary schema directory %s\n", path)
		} else {
			fmt.Printf("folder path does not match")
			os.Exit(1)
		}
	} else {
		fmt.Printf("unable to remove temporary schema directory")
		os.Exit(1)
	}
}
