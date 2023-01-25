# README for [fgg2go](https://github.com/sfzhu93/fgg2go)

---

This repository hosts [fgg2go](https://github.com/sfzhu93/fgg2go). 
It is a fork of [fgg](https://github.com/rhu1/fgg), and implements two homogeneous translations
of generics, via dictionary passing and dynamic type assertion.

It is part of the artifact of:

> Generic Go to Go: Dictionary-Passing, Monomorphisation, and Hybrid
  
> *Stephen Ellis, Shuofei Zhu, Nobuko Yoshida, Linhai Song"  
> https://arxiv.org/abs/2208.06810

The [full artifact](https://zenodo.org/record/7067362#.YzJ-SXbMKUk) is a docker image
that includes a more detailed README and benchmark environment. 

## Installation
The source code is developed before `GO111MODULE` is introduced, and it relies on the `$GOPATH` environment variable.

We suggest the following command for Go 1.17 or newer,

```
export GOPATH=$(go env GOPATH)
cd $GOPATH/src
GO111MODULE=off go get github.com/sfzhu93/fgg2go
cd github.com/sfzhu93/fgg2go
go install .
```

so that you can make a local copy of the source code, because it is easier to find the examples that we provide, and make changes directly to the source code.

## Translating the microbenchmark program
After executing the commands for installation, there should be an executable `$GOPATH/bin/fgg2go` available and the current directory is located at `$GOPATH/github.com/sfzhu93/fgg2go`.

Then, the following command translate the microbenchmark program in FGG into standard Go via dictionary passing, without type metadata for type assertions, following the experimental settings in the paper.
```
~/go/bin/fgg2go -dictpassnoassertion -fgg examples/fgg/microbenchmark.fgg > microbenchmark_dict_pass.go
```
 
Then, the following command translate the microbenchmark program in FGG into standard Go, via erasure.
```
~/go/bin/fgg2go -dynamic -fgg examples/fgg/microbenchmark.fgg > microbenchmark_erasure.go
```

Pre-translated of the two files are provided under `$GOPATH/github.com/sfzhu93/fgg2go/examples/fgg`.
