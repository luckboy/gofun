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

type Either struct {
    isRight bool
    x interface{}
}

func EitherOrElse(x interface{}, y *Either) *Either {
    z, isOk := x.(*Either)
    if isOk {
        return z
    } else {
        return y
    }
}


func Left(x interface{}) *Either {
    return &Either { isRight: false, x: x }
}

func Right(x interface{}) *Either {
    return &Either { isRight: true, x: x }
}

func (e *Either) IsLeft() bool {
    return !e.isRight
}

func (e *Either) IsRight() bool {
    return e.isRight
}

func (e *Either) GetLeft() interface{} {
    if e.isRight {
        return nil
    } else {
        return e.x
    }
}

func (e *Either) GetRight() interface{} {
    if e.isRight {
        return e.x
    } else {
        return nil
    }
}

func (e *Either) GetLeftOrElse(x func() interface{}) interface{} {
    if e.isRight {
        return x()
    } else {
        return e.x
    }
}

func (e *Either) LeftOrElse(e2 func() *Either) interface{} {
    if e.isRight {
        return e2()
    } else {
        return e
    }
}

func (e *Either) GetRightOrElse(x func() interface{}) interface{} {
    if e.isRight {
        return e.x
    } else {
        return x()
    }
}

func (e *Either) RightOrElse(e2 func() *Either) interface{} {
    if e.isRight {
        return e
    } else {
        return e2()
    }
}
