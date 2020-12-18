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

// Option represents optional values.
type Option struct {
    isSome bool
    x interface{}
}

// OptionOrElse returns x if x is Option pointer, otherwise y.
func OptionOrElse(x interface{}, y *Option) *Option {
    z, isOk := x.(*Option)
    if isOk {
        return z
    } else {
        return y
    }
}

// None creates an Option without a value.
func None() *Option {
    return &Option { isSome: false, x: nil } 
}

// Some creates an Option with a value.
func Some(x interface{}) *Option {
    return &Option { isSome: true, x: x }
}

// IsNone returns true if o doesn't contain the value, otherwise false.
func (o *Option) IsNone() bool {
    return !o.isSome
}

// IsNone returns true if o contains the value, otherwise false.
func (o *Option) IsSome() bool {
    return o.isSome
}

// Get returns the value.
func (o *Option) Get() interface{} {
    return o.x
}

// GetOrElse returns the value if o contains the value, otherwise x().
func (o *Option) GetOrElse(x func() interface{}) interface{} {
    if o.isSome {
        return o.x
    } else {
        return x()
    }
}

// OrElse returns o if o contains the value, otherwise o2().
func (o *Option) OrElse(o2 func() *Option) *Option {
    if o.isSome {
        return o
    } else {
        return o2()
    }
}

func (o *Option) String() string {
    if o.isSome {
        return fmt.Sprintf("Some[%v]", o.x)
    } else {
        return "None"
    }
}
