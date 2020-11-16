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

type Monad interface {
    Bind(func(interface{}) Monad) Monad
}

func Join(m Monad) Monad {
    return m.Bind(func(x interface{}) Monad {
            y, isOk := x.(Monad)
            if isOk {
                return y
            } else {
                return nil
            }
    })
}

func (m *Option) Bind(f func(interface{}) Monad) Monad {
    if m.IsSome() {
        return f(m.Get())
    } else {
        return None()
    }
}

func (m *Either) Bind(f func(interface{}) Monad) Monad {
    if m.IsRight() {
        return f(m.GetRight())
    } else {
        return Left(m.GetLeft())
    }
}

func (m InterfaceSlice) Bind(f func(interface{}) Monad) Monad {
    ys := make([]interface{}, 0, len(m))
    for _, x := range m {
        m2, isOk := f(x).(InterfaceSlice)
        if isOk {
            for _, y := range m2 {
                ys = append(ys, y)
            }
        }
    }
    return InterfaceSlice(ys)
}
