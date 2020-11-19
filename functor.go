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

type Functor interface {
    Map(func(interface{}) interface{}) Functor
}

func FunctorOrElse(x interface{}, y Functor) Functor {
    z, isOk := x.(Functor)
    if isOk {
        return z
    } else {
        return y
    }
}

func (xs *Option) Map(f func(interface{}) interface{}) Functor {
    if xs.IsSome() {
        return Some(f(xs.Get()))
    } else {
        return None()
    }
}

func (xs *Either) Map(f func(interface{}) interface{}) Functor {
    if xs.IsRight() {
        return Right(f(xs.GetRight()))
    } else {
        return Left(xs.GetLeft())
    }
}

func (xs *List) Map(f func(interface{}) interface{}) Functor {
    var ys *List = Nil()
    var prev *List = nil
    for l := xs; l.IsCons(); l = l.Tail() {
        l2 := Cons(f(l.Head()), Nil())
        if prev != nil {
            prev.SetTail(l2)
        } else {
            ys = l2
        }
        prev = l2
    }
    return ys
}

func (xs InterfaceSlice) Map(f func(interface{}) interface{}) Functor {
    ys := make([]interface{}, 0, len(xs))
    for _, x := range xs {
        ys = append(ys, f(x))
    }
    return InterfaceSlice(ys)
}

func (xs InterfacePairMap) Map(f func(interface{}) interface{}) Functor {
    ys := make(map[interface{}]interface{}, len(xs))
    for k, v := range xs {
        p, isOk := f(NewPair(k, v)).(*Pair)
        if isOk {
            ys[p.First] = p.Second
        }
    }
    return InterfacePairMap(ys)
}

func (xs InterfacePairFunction) Map(f func(interface{}) interface{}) Functor {
    return InterfacePairFunction(func(x interface {}) interface{} {
        return f(xs(x))
    })
}
