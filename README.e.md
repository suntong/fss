
# {{.Name}}

{{render "license/shields" . "License" "MIT"}}
{{template "badge/godoc" .}}
{{template "badge/goreport" .}}
{{template "badge/travis" .}}

## {{toc 5}}

## {{.Name}} - file system summary

The `fss` is a tool to summarize files and diretories in the file system.


# Usage

### $ {{exec "fss" | color "sh"}}

### $ {{shell "fss grps" | color "sh"}}

### $ {{shell "fss dirs" | color "sh"}}


# Examples

### $ {{shell "fss grps ." | color "sh"}}

### $ {{shell "ls -ld `ls -1d /etc/* | fss dirs`" | color "sh"}}


# Installation

```
go get github.com/suntong/fss
```

All patches welcome.


