package main

import (
	"os"
	"fmt"
	"github.com/samparsky/task/db"
	"path/filepath"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/samparsky/task/cmd"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error){
	if(err != nil){
		fmt.Println(err.Error())
		os.Exit(1)
	}
}