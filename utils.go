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

func BoolOrElse(x interface{}, y bool) bool {
    z, isOk := x.(bool)
    if isOk {
        return z
    } else {
        return y
    }
}

func ByteOrElse(x interface{}, y byte) byte {
    z, isOk := x.(byte)
    if isOk {
        return z
    } else {
        return y
    }
}

func Complex64OrElse(x interface{}, y complex64) complex64 {
    z, isOk := x.(complex64)
    if isOk {
        return z
    } else {
        return y
    }
}

func Complex128OrElse(x interface{}, y complex128) complex128 {
    z, isOk := x.(complex128)
    if isOk {
        return z
    } else {
        return y
    }
}

func ErrorOrElse(x interface{}, y error) error {
    z, isOk := x.(error)
    if isOk {
        return z
    } else {
        return y
    }
}

func Float32OrElse(x interface{}, y float32) float32 {
    z, isOk := x.(float32)
    if isOk {
        return z
    } else {
        return y
    }
}

func Float64OrElse(x interface{}, y float64) float64 {
    z, isOk := x.(float64)
    if isOk {
        return z
    } else {
        return y
    }
}

func IntOrElse(x interface{}, y int) int {
    z, isOk := x.(int)
    if isOk {
        return z
    } else {
        return y
    }
}

func Int8OrElse(x interface{}, y int8) int8 {
    z, isOk := x.(int8)
    if isOk {
        return z
    } else {
        return y
    }
}

func Int16OrElse(x interface{}, y int16) int16 {
    z, isOk := x.(int16)
    if isOk {
        return z
    } else {
        return y
    }
}

func Int32OrElse(x interface{}, y int32) int32 {
    z, isOk := x.(int32)
    if isOk {
        return z
    } else {
        return y
    }
}

func Int64OrElse(x interface{}, y int64) int64 {
    z, isOk := x.(int64)
    if isOk {
        return z
    } else {
        return y
    }
}

func RuneOrElse(x interface{}, y rune) rune {
    z, isOk := x.(rune)
    if isOk {
        return z
    } else {
        return y
    }
}

func StringOrElse(x interface{}, y string) string {
    z, isOk := x.(string)
    if isOk {
        return z
    } else {
        return y
    }
}

func UintOrElse(x interface{}, y uint) uint {
    z, isOk := x.(uint)
    if isOk {
        return z
    } else {
        return y
    }
}

func Uint8OrElse(x interface{}, y uint8) uint8 {
    z, isOk := x.(uint8)
    if isOk {
        return z
    } else {
        return y
    }
}

func Uint16OrElse(x interface{}, y uint16) uint16 {
    z, isOk := x.(uint16)
    if isOk {
        return z
    } else {
        return y
    }
}

func Uint32OrElse(x interface{}, y uint32) uint32 {
    z, isOk := x.(uint32)
    if isOk {
        return z
    } else {
        return y
    }
}

func Uint64OrElse(x interface{}, y uint64) uint64 {
    z, isOk := x.(uint64)
    if isOk {
        return z
    } else {
        return y
    }
}

func UintptrOrElse(x interface{}, y uintptr) uintptr {
    z, isOk := x.(uintptr)
    if isOk {
        return z
    } else {
        return y
    }
}
