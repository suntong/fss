////////////////////////////////////////////////////////////////////////////
// Program: fss
// Purpose: file system summary
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/mkideal/cli"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

//==========================================================================
// fss

type rootT struct {
	cli.Helper
	Verbose cli.Counter `cli:"v,verbose" usage:"Verbose mode (Multiple -v options increase the verbosity.)"`
}

var root = &cli.Command{
	Name: "fss",
	Desc: "file system summary\nbuilt on " + buildTime,
	Text: "Tool to summarize files and diretories in the file system" +
		"\n\nUsage:\n  fss COMMAND",
	Global: true,
	Argv:   func() interface{} { return new(rootT) },
	Fn:     fss,

	NumArg: cli.AtLeast(1),
}

// Template for main starts here
////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "fss"
//          VERSION   = "0.1.0"
//          buildTime = "2017-07-21"
//  )

//  var rootArgv *rootT

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	cli.SetUsageStyle(cli.ManualStyle) // up-down, for left-right, use NormalStyle
//  	//NOTE: You can set any writer implements io.Writer
//  	// default writer is os.Stdout
//  	if err := cli.Root(root,
//  		cli.Tree(grpsDef),
//  		cli.Tree(dirsDef)).Run(os.Args[1:]); err != nil {
//  		fmt.Fprintln(os.Stderr, err)
//  	}
//  	fmt.Println("")
//  }

// Template for main dispatcher starts here
//==========================================================================
// Main dispatcher

//  func fss(ctx *cli.Context) error {
//  	ctx.JSON(ctx.RootArgv())
//  	ctx.JSON(ctx.Argv())
//  	fmt.Println()

//  	return nil
//  }

// Template for CLI handling starts here

////////////////////////////////////////////////////////////////////////////
// grps

//  func grpsCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*grpsT)
//  	fmt.Printf("[grps]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
//  	return nil
//  }

type grpsT struct {
}

var grpsDef = &cli.Command{
	Name: "grps",
	Desc: "Separate files into groups",
	Text: "List files in groups with total sizes from the given path" +
		"\n\nUsage:\n  fss grps SRC-PATH",
	Argv: func() interface{} { return new(grpsT) },
	Fn:   grpsCLI,

	NumArg:      cli.AtLeast(1),
	CanSubRoute: true,
}

////////////////////////////////////////////////////////////////////////////
// dirs

//  func dirsCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*dirsT)
//  	fmt.Printf("[dirs]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
//  	return nil
//  }

type dirsT struct {
}

var dirsDef = &cli.Command{
	Name: "dirs",
	Desc: "Pick out dirs from the input",
	Text: "Pick out dirs, given a list of files and diretories from stdin" +
		"\n\nUsage:\n  mlocate SRC-PATH | sort | fss dirs",
	Argv: func() interface{} { return new(dirsT) },
	Fn:   dirsCLI,
}
