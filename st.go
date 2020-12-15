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

// ST represents a state monads.
type ST func(interface{}) (interface{}, interface{})

// STOrElse returns x if x is ST, otherwise y.
func STOrElse(x interface{}, y ST) ST {
    z, isOk := x.(ST)
    if isOk {
        return z
    } else {
        return y
    }
}

// RunST runs the ST monad.
func RunST(st ST, x interface{}) (interface{}, interface{}) {
    return st(x)
}

// GetST returns the ST monad with the state.
func GetST() ST {
    return ST(func(s interface{}) (interface{}, interface{}) {
            return s, s
    })
}

// SetST sets a new state.
func SetST(newS interface{}) ST {
    return ST(func(s interface{}) (interface{}, interface{}) {
            return newS, struct{} {}
    })
}
