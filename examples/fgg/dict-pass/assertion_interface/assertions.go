package main

type Any interface{}
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

type anyǂ interface{}
type anyǂDictǂ struct{ _type _type_metadata }
type anyǂ_meta struct{}

func (this anyǂ_meta) tryCast(x Any) Any {
	_ = x.(anyǂ)
	return x
}
func (this anyǂ_meta) assertEq(x _type_metadata) Any {
	_ = x.(anyǂ_meta)
	return this
}

type tIǂ interface {
	mǂ(inǂ Any) Any
	spec_m() spec_metadata_1
}
type tIǂDictǂ struct {
	mǂ    func(receiver Any, inǂ Any) Any
	_type _type_metadata
}
type tIǂ_meta struct{}

func (this tIǂ_meta) tryCast(x Any) Any {
	x_ := x.(tIǂ)
	m_actual := x_.spec_m()

	m_actual._type_0.assertEq(intǂ_meta{})
	m_actual._type_1.assertEq(anyǂ_meta{})
	return x
}
func (this tIǂ_meta) assertEq(x _type_metadata) Any {
	_ = x.(tIǂ_meta)
	return this
}

type uSǂ struct{}
type uSǂ_meta struct{}

func (this uSǂ_meta) tryCast(x Any) Any {
	_ = x.(uSǂ_meta)
	return x
}
func (this uSǂ_meta) assertEq(x _type_metadata) Any {
	return this
}
func (thisǂ uSǂ) mǂ(intǂ Any) Any {
	return thisǂ
}
func (thisǂ uSǂ) spec_m() spec_metadata_1 {
	return spec_metadata_1{intǂ_meta{}, anyǂ_meta{}}
}
func (thisǂ boolǂ) fooǂ(xǂ Any) Any {
	return xǂ
}
func (thisǂ boolǂ) spec_foo() spec_metadata_1 {
	return spec_metadata_1{anyǂ_meta{}, anyǂ_meta{}}
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
func main() { _ = tIǂ_meta{}.tryCast(boolǂ{}.fooǂ(uSǂ{})) }
