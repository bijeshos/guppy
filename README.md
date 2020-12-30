# Guppy
![build](https://github.com/bijeshos/guppy/workflows/go-build/badge.svg)
![Go Report Card](https://goreportcard.com/badge/github.com/bijeshos/guppy)

**Guppy** provides a collection of utility functions. 
At the moment, **Guppy** offers following utilities:

|Area|Function|Purpose|
|---|---|---|
|Array|IsPresent|Checks whether an element is present in the array|
|Array|Find|Finds index of the element|
|Directory|ReadFiles|Read Files|
|Directory|ReadDirs|Read directories|
|Directory|MkDirAll|Create directory|
|Directory|IsExist|Check whether the directory exists|
|Directory|IsSame|Checks whether two directory paths are same|
|File|CreateFile|Create File|
|File|CopyFile|Copy file|
|File|MoveFile|Move File|
|File|isSameMetadata|Checks whether two files' metadata is same or not|
|File|ReadFileContent|Reads content of a file as string array|


## Installation

If you do not have Go installed yet, you can find installation instructions
[here](https://golang.org/doc/install).

To pull the most recent version of **Guppy**, use `go get`.

```
go get github.com/bijeshos/guppy
```

Then import respective packages into your project.

To avoid any potential conflicts, it is suggested to use custom package name as follows:

|Full Package Name|Suggested Custom name|Stands for|
|---|---|---|
|github.com/bijeshos/guppy/arrayutil|gau| Guppy Array Util|
|github.com/bijeshos/guppy/fileutil|gfu| Guppy File Util|
|github.com/bijeshos/guppy/dirutil|gdu| Guppy Directory Util|
|github.com/bijeshos/guppy/stringutil|gsu| Guppy String Util|

So, the package import statements would look like the following:

```go
import gau "github.com/bijeshos/guppy/arrayutil"
import gfu "github.com/bijeshos/guppy/fileutil"
import gdu "github.com/bijeshos/guppy/dirutil"
import gsu "github.com/bijeshos/guppy/stringutil"
```

## Usage
Usage details are available as part of the godoc. 

In order to view godoc, follow the below steps:
- Refer documentation by running 
```
godoc -http=:6060
```
- Open the following URL in a browser:
http://localhost:6060/pkg/github.com/bijeshos/guppy/
  
- Traverse to respective sub packages
