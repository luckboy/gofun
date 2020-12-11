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

type Zippable interface {
    Zip(ys Zippable, fail Unzippable) Unzippable
}

func (xs *Option) Zip(ys Zippable, fail Unzippable) Unzippable {
    ys2, isOk := ys.(*Option)
    if isOk {
        if xs.IsSome() && ys2.IsSome() {
            return Some(NewPair(xs.Get(), ys2.Get()))
        } else {
            return None()
        }
    } else {
        return None()
    }
}

func (xs *Either) Zip(ys Zippable, fail Unzippable) Unzippable {
    ys2, isOk := ys.(*Either)
    if isOk {
        if xs.IsRight() && ys2.IsRight() {
            return Right(NewPair(xs.GetRight(), ys2.GetRight()))
        } else {
            if xs.IsLeft() {
                return xs
            } else {
                return ys2
            }
        }
    } else {
        return fail
    }
}

func (xs *List) Zip(ys Zippable, fail Unzippable) Unzippable {
    var zs *List = Nil()
    var prev *List = nil
    ys2, isOk := ys.(*List)
    if isOk {
        for l1, l2 := xs, ys2; l1.IsCons() && l2.IsCons(); l1, l2 = l1.Tail(), l2.Tail() {
            l3 := Cons(NewPair(l1.Head(), l2.Head()), Nil())
            if prev != nil {
                prev.SetTail(l3)
            } else {
                zs = l3
            }
            prev = l3
        }
    }
    return zs
}

func (xs InterfaceSlice) Zip(ys Zippable, fail Unzippable) Unzippable {
    ys2, isOk := ys.(InterfaceSlice)
    if isOk {
        var length int
        if len(xs) < len(ys2) {
            length = len(xs)
        } else {
            length = len(ys2)
        }
        zs := make([]interface{}, 0, length)
        for i := 0; i < len(xs) && i < len(ys2); i++ {
            zs = append(zs, NewPair(xs[i], ys2[i]))
        }
        return InterfaceSlice(zs)
    } else {
        return InterfaceSlice([]interface{} {})
    }
}

func (xs InterfacePairFunction) Zip(ys Zippable, fail Unzippable) Unzippable {
    ys2, isOk := ys.(InterfacePairFunction)
    if isOk {
        return InterfacePairFunction(func(x interface{}) interface{} {
                return NewPair(xs(x), ys2(x))
        })
    } else {
        return fail
    }
}
