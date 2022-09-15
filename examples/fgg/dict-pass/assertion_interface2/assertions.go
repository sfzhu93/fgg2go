package main

type Any interface{}
type Anyǂ interface{}
type AnyǂDictǂ struct{ _type _type_metadata }
type Anyǂ_meta struct{}

func (this Anyǂ_meta) tryCast(x Any) Any {
	_ = x.(Anyǂ)
	return x
}
func (this Anyǂ_meta) assertEq(x _type_metadata) Any {
	_ = x.(Anyǂ_meta)
	return this
}

type tiǂ interface {
	mǂ(inǂ Any) Any
	spec_m() spec_metadata_1
}
type tiǂDictǂ struct {
	mǂ    func(receiver Any, inǂ Any) Any
	_type _type_metadata
}
type tiǂ_meta struct{ _type_0 _type_metadata }

func (this tiǂ_meta) tryCast(x Any) Any {
	x_ := x.(tiǂ)
	m_actual := x_.spec_m()

	m_actual._type_0.assertEq(this._type_0)
	m_actual._type_1.assertEq(this._type_0)
	return x
}
func (this tiǂ_meta) assertEq(x _type_metadata) Any {
	x_ := x.(tiǂ_meta)
	this._type_0.assertEq(x_._type_0)
	return this
}

type intǂ struct{}
type intǂ_meta struct{}

func (this intǂ_meta) tryCast(x Any) Any {
	_ = x.(intǂ_meta)
	return x
}
func (this intǂ_meta) assertEq(x _type_metadata) Any {
	return this
}

type boolǂ struct{}
type boolǂ_meta struct{}

func (this boolǂ_meta) tryCast(x Any) Any {
	_ = x.(boolǂ_meta)
	return x
}
func (this boolǂ_meta) assertEq(x _type_metadata) Any {
	return this
}

type usǂ struct{}
type usǂ_meta struct{}

func (this usǂ_meta) tryCast(x Any) Any {
	_ = x.(usǂ_meta)
	return x
}
func (this usǂ_meta) assertEq(x _type_metadata) Any {
	return this
}
func (thisǂ usǂ) mǂ(inǂ Any) Any {
	return inǂ
}
func (thisǂ usǂ) spec_m() spec_metadata_1 {
	return spec_metadata_1{intǂ_meta{}, intǂ_meta{}}
}
func (thisǂ boolǂ) fooǂ(dict_0 AnyǂDictǂ, inǂ Any) Any {
	return dict_0._type.tryCast(inǂ)
}
func (thisǂ boolǂ) spec_foo() spec_metadata_2 {
	return spec_metadata_2{Anyǂ_meta{}, tiǂ_meta{intǂ_meta{}}, param_index_0{}}
}

type spec_metadata_2 struct {
	_type_0 _type_metadata
	_type_1 _type_metadata
	_type_2 _type_metadata
}
type spec_metadata_1 struct {
	_type_0 _type_metadata
	_type_1 _type_metadata
}
type spec_metadata_0 struct{ _type_0 _type_metadata }
type param_index_0 struct{}

func (this param_index_0) tryCast(x Any) Any             { _ = x.(param_index_0); return x }
func (this param_index_0) assertEq(x _type_metadata) Any { _ = x.(param_index_0); return x }

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
func main() { _ = boolǂ{}.fooǂ(AnyǂDictǂ{tiǂ_meta{boolǂ_meta{}}}, usǂ{}) }
