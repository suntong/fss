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

////////////////////////////////////////////////////////////////////////////
// grps

func grpsCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*grpsT)
	fmt.Printf("[grps]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
	return nil
}
