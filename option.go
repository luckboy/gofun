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

type Option struct {
    isSome bool
    x interface{}
}

func OptionOrElse(x interface{}, y *Option) *Option {
    z, isOk := x.(*Option)
    if isOk {
        return z
    } else {
        return y
    }
}

func None() *Option {
    return &Option { isSome: false, x: nil } 
}

func Some(x interface{}) *Option {
    return &Option { isSome: true, x: x }
}

func (o *Option) IsNone() bool {
    return !o.isSome
}

func (o *Option) IsSome() bool {
    return o.isSome
}

func (o *Option) Get() interface{} {
    return o.x
}

func (o *Option) GetOrElse(x interface{}) interface{} {
    if o.isSome {
        return o.x
    } else {
        return x
    }
}

func (o *Option) OrElse(o2 *Option) *Option {
    if o.isSome {
        return o
    } else {
        return o2
    }
}
