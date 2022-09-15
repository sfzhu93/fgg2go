package main

type Any interface{}
type tauǂ interface{}
type tauǂDictǂ struct{ _type _type_metadata }
type tauǂ_meta struct{}

func (this tauǂ_meta) tryCast(x Any) Any {
	_ = x.(tauǂ)
	return x
}
func (this tauǂ_meta) assertEq(x _type_metadata) Any {
	_ = x.(tauǂ_meta)
	return this
}

type intǂ struct{}
type intǂ_meta struct{}

func (this intǂ_meta) tryCast(x Any) Any {
	return x
}
func (this intǂ_meta) assertEq(x _type_metadata) Any {
	return this
}

type boolǂ struct{}
type boolǂ_meta struct{}

func (this boolǂ_meta) tryCast(x Any) Any {
	return x
}
func (this boolǂ_meta) assertEq(x _type_metadata) Any {
	return this
}

type tsǂ struct {
	arg1   Any
	arg2   Any
	dict_0 tauǂDictǂ
}
type tsǂ_meta struct{ _type_0 _type_metadata }

func (this tsǂ_meta) tryCast(x Any) Any {
	this._type_0.assertEq(x.(tsǂ).dict_0._type)
	return x
}
func (this tsǂ_meta) assertEq(x _type_metadata) Any {
	this._type_0.assertEq(x.(tsǂ_meta)._type_0)
	return this
}

type dummyǂ struct{}
type dummyǂ_meta struct{}

func (this dummyǂ_meta) tryCast(x Any) Any {
	return x
}
func (this dummyǂ_meta) assertEq(x _type_metadata) Any {
	return this
}
func (thisǂ dummyǂ) doitǂ() Any {
	return tsǂ{intǂ{}, boolǂ{}, tauǂDictǂ{boolǂ_meta{}}}
}
func (thisǂ dummyǂ) spec_doit() spec_metadata_0 {
	return spec_metadata_0{tauǂ_meta{}}
}

type spec_metadata_0 struct{ _type_0 _type_metadata }
type spec_metadata_1 struct {
	_type_0 _type_metadata
	_type_1 _type_metadata
}
type _type_metadata interface {
	tryCast(x Any) Any
	assertEq(x _type_metadata) Any
}
type __BidirectionalChannel struct{}

func (this __BidirectionalChannel) assertEq(x Any) Any {
	return x.(__BidirectionalChannel)
}

type __sendOnlyChannel struct{}

func (this __sendOnlyChannel) assertEq(x Any) Any {
	return x.(__sendOnlyChannel)
}

type __receiveOnlyChannel struct{}

func (this __receiveOnlyChannel) assertEq(x Any) Any {
	return x.(__receiveOnlyChannel)
}

type chan_wrapper struct {
	ch    chan Any
	_type _type_metadata
}
type chan_direction interface{ assertEq(x Any) Any }
type chan_metadata struct {
	dir   chan_direction
	_type _type_metadata
}

func (this chan_metadata) tryCast(x Any) Any {
	_ = x.(chan_wrapper)._type.assertEq(this)
	return x
}
func (this chan_metadata) assertEq(x _type_metadata) Any {
	_ = x.(chan_metadata)._type.assertEq(this._type)
	return x.(chan_metadata).dir.assertEq(this.dir)
}
func main() { _ = tsǂ_meta{intǂ_meta{}}.tryCast(dummyǂ{}.doitǂ()) }
