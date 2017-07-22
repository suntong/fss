////////////////////////////////////////////////////////////////////////////
// Program: fss
// Purpose: file system summary
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	"github.com/mkideal/cli"
)

func dirsCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*dirsT)
	fmt.Printf("[dirs]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
	return nil
}
