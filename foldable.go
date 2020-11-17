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

type Foldable interface {
    FoldLeft(f func(interface{}, interface{}) interface{}, z interface{}) interface{}
    FoldRight(f func(interface{}, interface{}) interface{}, z interface{}) interface{}
}

func FoldableOrElse(x interface{}, y Foldable) Foldable {
    z, isOk := x.(Foldable)
    if isOk {
        return z
    } else {
        return y
    }
}

func FoldableOrElseNil(x interface{}) Foldable {
    z, isOk := x.(Foldable)
    if isOk {
        return z
    } else {
        return nil
    }
}

func Any(f func(interface{}) bool, xs Foldable) bool {
    return BoolOrElse(xs.FoldLeft(func(x, y interface{}) interface{} {
            return BoolOrElse(x, false) || f(y)
    }, false), false)
}

func All(f func(interface{}) bool, xs Foldable) bool {
    return BoolOrElse(xs.FoldLeft(func(x, y interface{}) interface{} {
            return BoolOrElse(x, false) && f(y)
    }, true), false)
}

func Elem(x interface{}, xs Foldable) bool {
    return BoolOrElse(xs.FoldLeft(func(y, z interface{}) interface{} {
            return BoolOrElse(y, false) || z == x
    }, false), false)
}

func Find(f func(interface{}) bool, xs Foldable) *Option {
    return OptionOrElse(xs.FoldLeft(func(x, y interface{}) interface{} {
            o := OptionOrElse(x, None())
            if o.IsSome() {
                return o
            } else {
                if f(y) {
                    return Some(y)
                } else {
                    return None()
                }
            }
    }, None()), None())
}

func NotElem(x interface{}, xs Foldable) bool {
    return !Elem(x, xs)
}

func Length(xs Foldable) int {
    return IntOrElse(xs.FoldLeft(func(x, y interface{}) interface{} {
            return IntOrElse(x, 0) + 1
    }, 0), 0)
}

func (xs *Option) FoldLeft(f func(interface{}, interface{}) interface{}, z interface{}) interface{} {
    if xs.IsSome() {
        return f(z, xs.Get())
    } else {
        return z
    }
}

func (xs *Option) FoldRight(f func(interface{}, interface{}) interface{}, z interface{}) interface{} {
    if xs.IsSome() {
        return f(xs.Get(), z)
    } else {
        return z
    }
}

func (xs *Either) FoldLeft(f func(interface{}, interface{}) interface{}, z interface{}) interface{} {
    if xs.IsRight() {
        return f(z, xs.GetRight())
    } else {
        return z
    }
}

func (xs *Either) FoldRight(f func(interface{}, interface{}) interface{}, z interface{}) interface{} {
    if xs.IsRight() {
        return f(xs.GetRight(), z)
    } else {
        return z
    }
}

func (xs *List) FoldLeft(f func(interface{}, interface{}) interface{}, z interface{}) interface{} {
    y := z
    for l := xs; l.IsCons(); l = l.Tail() {
        y = f(y, l.Head())
    }
    return y
}

func (xs *List) FoldRight(f func(interface{}, interface{}) interface{}, z interface{}) interface{} {
    ys := make([]interface{}, 0, Length(xs))
    for l := xs; l.IsCons(); l = l.Tail() {
        ys = append(ys, l.Head())
    }
    return InterfaceSlice(ys).FoldRight(f, z)
}

func (xs InterfaceSlice) FoldLeft(f func(interface{}, interface{}) interface{}, z interface{}) interface{} {
    y := z
    for _, x := range xs {
        y = f(y, x)
    }
    return y
}

func (xs InterfaceSlice) FoldRight(f func(interface{}, interface{}) interface{}, z interface{}) interface{} {
    y := z
    for i := len(xs) - 1; i >= 0; i-- {
        y = f(xs[i], y)
    }
    return y
}

func (xs InterfacePairMap) FoldLeft(f func(interface{}, interface{}) interface{}, z interface{}) interface{} {
    y := z
    for k, v := range xs {
        y = f(y, NewPair(k, v))
    }
    return y
}

func (xs InterfacePairMap) FoldRight(f func(interface{}, interface{}) interface{}, z interface{}) interface{} {
    ys := make([]interface{}, 0, len(xs))
    for k, v := range xs {
        ys = append(ys, NewPair(k, v))
    }
    return InterfaceSlice(ys).FoldRight(f, z)
}
