////////////////////////////////////////////////////////////////////////////
// Program: fss
// Purpose: file system summary
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"
	"path/filepath"

	"github.com/mkideal/cli"
	"github.com/spakin/awk"
	"github.com/suntong/set"
)

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var commonExts set.Set = set.NewSet0From([]string{
	"png",
	"jpg",
	"avi",
	"mkv",
	"gif",
	"rar",
	"srt",
	"par2",
	"tgz",
	"html",
	"mp4",
	"flv",
	"htm",
	"zip",
	"sub",
	"rmvb",
	"idx",
	"nfo",
	"m4v",
	"txt",
	"VOB",
	"IFO",
	"BUP",
	"pdf",
	"lst",
	"js",
	"ico",
	"css",
})

////////////////////////////////////////////////////////////////////////////
// Function definitions

func dirsCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	//argv := ctx.Argv().(*dirsT)
	// fmt.Printf("[dirs]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
	return cmdDirs()
}

func cmdDirs() error {
	// https://godoc.org/github.com/spakin/awk
	s := awk.NewScript()

	// == Match & Process
	s.AppendStmt(nil, func(s *awk.Script) {
		// == skip all files
		fn := s.F(0).String()
		// with known extension
		if commonExts.Has(filepath.Ext(fn)) {
			s.Next()
		}
		// or not a dir
		stt, _ := os.Stat(fn)
		if !stt.IsDir() {
			s.Next()
		}
	})

	// 1; # i.e., print all
	s.AppendStmt(nil, nil)
	return s.Run(os.Stdin)
}
