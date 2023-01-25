package main;
type Any interface {};
type Colorǂ interface {};
type Topǂ struct {};
type Topǂ_meta struct {  };
func (thisǂ Topǂ) Opǂ() Any { return Topǂ{};
 };
type Redǂ struct {};
type Redǂ_meta struct {  };
type Blueǂ struct {};
type Blueǂ_meta struct {  };
func (thisǂ Topǂ) doǂ(xǂ Any) Any { return thisǂ;
 };
type Baseǂ interface { g_chanǂ() Any; g1ǂ(p1ǂ Any, p2ǂ Any) Any; g2ǂ(p1ǂ Any, p2ǂ Any) Any };
type Derivedǂ struct {};
type Derivedǂ_meta struct { _type_0 _type_metadata;_type_1 _type_metadata };
func (thisǂ Derivedǂ) g_chanǂ() Any { return Topǂ{}.Opǂ().(Topǂ).Opǂ();
 };
func (thisǂ Derivedǂ) g1ǂ(p1ǂ Any, p2ǂ Any) Any { return thisǂ;
 };
func (thisǂ Derivedǂ) g2ǂ(p1ǂ Any, p2ǂ Any) Any { return thisǂ;
 };
func (thisǂ Topǂ) f0ǂ(xǂ Any, p1ǂ Any, p2ǂ Any) Any { return thisǂ.doǂ(xǂ.(Baseǂ).g1ǂ(p1ǂ, p2ǂ)).(Topǂ).doǂ(xǂ.(Baseǂ).g2ǂ(p1ǂ, p2ǂ)).(Topǂ).doǂ(xǂ.(Baseǂ).g_chanǂ());
 };
func (thisǂ Topǂ) f4ǂ(p1ǂ Any, p2ǂ Any) Any { return thisǂ.f0ǂ(Derivedǂ{}, p1ǂ, p2ǂ);
 };
func (thisǂ Topǂ) f1ǂ(p1ǂ Any, p2ǂ Any) Any { return thisǂ.f2ǂ(p1ǂ, p2ǂ);
 };
func (thisǂ Topǂ) f2ǂ(p1ǂ Any, p2ǂ Any) Any { return thisǂ.f3ǂ(p1ǂ, p2ǂ);
 };
func (thisǂ Topǂ) f3ǂ(p1ǂ Any, p2ǂ Any) Any { return thisǂ.f4ǂ(p1ǂ, p2ǂ);
 };
type spec_metadata_0 struct { _type_0 _type_metadata };
type spec_metadata_2 struct { _type_0 _type_metadata; _type_1 _type_metadata; _type_2 _type_metadata };
type _type_metadata interface { tryCast(x Any) Any; assertEq(x _type_metadata) Any };
type __BidirectionalChannel struct {};
func (this __BidirectionalChannel) assertEq(x Any) Any { return x.(__BidirectionalChannel);
 };
type __sendOnlyChannel struct {};
func (this __sendOnlyChannel) assertEq(x Any) Any { return x.(__sendOnlyChannel);
 };
type __receiveOnlyChannel struct {};
func (this __receiveOnlyChannel) assertEq(x Any) Any { return x.(__receiveOnlyChannel);
 };
type chan_wrapper struct { ch chan Any };
type chan_direction interface { assertEq(x Any) Any };
type chan_metadata struct { dir chan_direction };
func (this chan_metadata) tryCast(x Any) Any {
//_ = x.(chan_wrapper)._type.assertEq(this);
return x
};
func (this chan_metadata) assertEq(x _type_metadata) Any {
//_ = x.(chan_metadata)._type.assertEq(this._type); 
return x.(chan_metadata).dir.assertEq(this.dir)
};
func main() { _ = Topǂ{}.f1ǂ(Redǂ{}, Redǂ{}).(Topǂ).f1ǂ(Redǂ{}, Blueǂ{}).(Topǂ).f1ǂ(Blueǂ{}, Redǂ{}).(Topǂ).f1ǂ(Blueǂ{}, Blueǂ{}) }
