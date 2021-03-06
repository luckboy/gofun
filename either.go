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

// Either represents one of two values.
type Either struct {
    isRight bool
    x interface{}
}

// EitherOrElse returns x if x is Either pointer, otherwise y.
func EitherOrElse(x interface{}, y *Either) *Either {
    z, isOk := x.(*Either)
    if isOk {
        return z
    } else {
        return y
    }
}

// Left creates an Either with a left value. 
func Left(x interface{}) *Either {
    return &Either { isRight: false, x: x }
}

// Right creates an Either with a right value. 
func Right(x interface{}) *Either {
    return &Either { isRight: true, x: x }
}

// IsLeft returns true if e contains the left value, otherwise false.
func (e *Either) IsLeft() bool {
    return !e.isRight
}

// IsRight returns true if e contains the right value, otherwise false.
func (e *Either) IsRight() bool {
    return e.isRight
}

// GetLeft returns the left value.
func (e *Either) GetLeft() interface{} {
    if e.isRight {
        return nil
    } else {
        return e.x
    }
}

// GetRight returns the right value.
func (e *Either) GetRight() interface{} {
    if e.isRight {
        return e.x
    } else {
        return nil
    }
}

// GetLeftOrElse returns the left value if e contains the left value, otherwise x().
func (e *Either) GetLeftOrElse(x func() interface{}) interface{} {
    if e.isRight {
        return x()
    } else {
        return e.x
    }
}

// LeftOrElse returns e if e contains the left value, otherwise e2().
func (e *Either) LeftOrElse(e2 func() *Either) interface{} {
    if e.isRight {
        return e2()
    } else {
        return e
    }
}

// GetRightOrElse returns the right value if e contains the right value, otherwise x().
func (e *Either) GetRightOrElse(x func() interface{}) interface{} {
    if e.isRight {
        return e.x
    } else {
        return x()
    }
}

// RightOrElse returns e if e contains the right value, otherwise e2().
func (e *Either) RightOrElse(e2 func() *Either) interface{} {
    if e.isRight {
        return e
    } else {
        return e2()
    }
}

func (e *Either) String() string {
    if e.isRight {
        return fmt.Sprintf("Right[%v]", e.x)
    } else {
        return fmt.Sprintf("Left[%v]", e.x)
    }
}
