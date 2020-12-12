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
import "fmt"

type List struct {
    isCons bool
    head interface{}
    tail *List
}

func ListOrElse(x interface{}, y *List) *List {
    z, isOk := x.(*List)
    if isOk {
        return z
    } else {
        return y
    }
}

func Nil() *List {
    return &List { isCons: false, head: nil, tail: nil }
}

func Cons(head interface{}, tail *List) *List {
    return &List { isCons: true, head: head, tail: tail }
}

func (l *List) IsNil() bool {
    return !l.isCons
}

func (l *List) IsCons() bool {
    return l.isCons
}

func (l *List) Head() interface{} {
    return l.head
}

func (l *List) HeadOption() *Option {
    if l.isCons {
        return Some(l.head)
    } else {
        return None()
    }
}

func (l *List) Tail() *List {
    return l.tail
}

func (l *List) TailOption() *Option {
    if l.isCons {
        return Some(l.tail)
    } else {
        return None()
    }
}

func (l *List) SetTail(tail *List) bool {
    if l.isCons {
        l.tail = tail
        return true
    } else {
        return false
    }
}

func (l *List) String() string {
    s := "List["
    isFirst := true
    for l2 := l; l2.isCons; l2 = l2.tail {
        if !isFirst {
            s += " "
        }
        s += fmt.Sprintf("%v", l2.head)
        isFirst = false
    }
    s += "]"
    return s
}

func (xs *List) Concat(ys *List) *List {
    var zs *List = Nil()
    var prev *List = nil
    for l := xs; l.IsCons(); l = l.Tail() {
        l2 := Cons(l.Head(), Nil())
        if prev != nil {
            prev.SetTail(l2)
        } else {
            zs = l2
        }
        prev = l2
    }
    if prev != nil {
        prev.SetTail(ys)
    } else {
        zs = ys
    }
    return zs
}
