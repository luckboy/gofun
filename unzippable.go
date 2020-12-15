/*
 * Copyright (c) 2020 Łukasz Szpakowski
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */

package gofun

// Unzippable is the interface for unzipping.
type Unzippable interface {
    // Unzip is zipping inverse. Fail must be a failure Zippable.
    Unzip(fail Zippable) (Zippable, Zippable)
}

func UnzippableOrElse(x interface{}, y Unzippable) Unzippable {
    z, isOk := x.(Unzippable)
    if isOk {
        return z
    } else {
        return y
    }
}

func (xs *Option) Unzip(fail Zippable) (Zippable, Zippable) {
    if xs.IsSome() {
        p, isOk := xs.Get().(*Pair)
        if isOk {
            return Some(p.First), Some(p.Second)
        } else {
            return None(), None()
        }
    } else {
        return None(), None()
    }
}

func (xs *Either) Unzip(fail Zippable) (Zippable, Zippable) {
    if xs.IsRight() {
        p, isOk := xs.GetRight().(*Pair)
        if isOk {
            return Right(p.First), Right(p.Second)
        } else {
            return fail, fail
        }
    } else {
        return Left(xs.GetLeft()), Left(xs.GetLeft())
    }
}

func (xs *List) Unzip(fail Zippable) (Zippable, Zippable) {
    var ys *List = Nil()
    var prev1 *List = nil
    var zs *List = Nil()
    var prev2 *List = nil
    for l := xs; l.IsCons(); l = l.Tail() {
        p, isOk := l.Head().(*Pair)
        if isOk {
            l2 := Cons(p.First, Nil())
            l3 := Cons(p.Second, Nil())
            if prev1 != nil {
                prev1.SetTail(l2)
            } else {
                ys = l2
            }
            if prev2 != nil {
                prev2.SetTail(l3)
            } else {
                zs = l3
            }
            prev1 = l2
            prev2 = l3
        }
    }
    return ys, zs
}

func (xs InterfaceSlice) Unzip(fail Zippable) (Zippable, Zippable) {
    ys := make([]interface{}, 0, len(xs))
    zs := make([]interface{}, 0, len(xs))
    for _, x := range xs {
        p, isOk := x.(*Pair)
        if isOk {
            ys = append(ys, p.First)
            zs = append(zs, p.Second)
        }
    }
    return InterfaceSlice(ys), InterfaceSlice(zs)
}

func (xs InterfacePairFunction) Unzip(fail Zippable) (Zippable, Zippable) {
    f := InterfacePairFunction(func(x interface{}) interface{} {
            p, isOk := xs(x).(*Pair)
            if isOk {
                return p.First
            } else {
                return x
            }
    })
    g := InterfacePairFunction(func(x interface{}) interface{} {
            p, isOk := xs(x).(*Pair)
            if isOk {
                return p.Second
            } else {
                return x
            }
    })
    return f, g
}
