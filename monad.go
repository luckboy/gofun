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
    Functor
    Bind(func(interface{}) Monad) Monad
}

func MonadOrElse(x interface{}, y Monad) Monad {
    z, isOk := x.(Monad)
    if isOk {
        return z
    } else {
        return y
    }
}

func MonadOrElseNil(x interface{}) Monad {
    z, isOk := x.(Monad)
    if isOk {
        return z
    } else {
        return nil
    }
}

func IfM(cond Monad, ifTrue, ifFalse func() Monad) Monad {
    return cond.Bind(func(x interface{}) Monad {
            if BoolOrElse(x, false) {
                return ifTrue()
            } else {
                return ifFalse()
            }
    });
}

func Join(m Monad) Monad {
    return m.Bind(func(x interface{}) Monad {
            return MonadOrElseNil(x)
    })
}

func UntilM(m Monad, cond func() Monad, u func(interface{}) Monad) Monad {
    return m.Bind(func(x interface{}) Monad {
            return cond().Bind(func(y interface{}) Monad {
                    if !BoolOrElse(x, false) {
                        return UntilM(m, cond, u)
                    } else {
                        return u(struct{} {})
                    }
            })
    })
}

func WhileM(cond Monad, body func() Monad, u func(interface{}) Monad) Monad {
    return cond.Bind(func(x interface{}) Monad {
            if BoolOrElse(x, false) {
                return body().Bind(func(y interface{}) Monad {
                        return WhileM(cond, body, u)
                })
            } else {
                return u(struct{}{})
            }
    })
}

func (m *Option) Bind(f func(interface{}) Monad) Monad {
    if m.IsSome() {
        m2 := f(m.Get())
        if m2 != nil {
            return m2
        } else {
            return None()
        }
    } else {
        return None()
    }
}

func OptionUnit(x interface{}) Monad {
    return Some(x)
}

func (m *Either) Bind(f func(interface{}) Monad) Monad {
    if m.IsRight() {
        m2 := f(m.GetRight())
        if m2 != nil {
            return m2
        } else {
            return Left(nil)
        }
    } else {
        return Left(m.GetLeft())
    }
}

func EitherUnit(x interface{}) Monad {
    return Right(x)
}

func (m *List) Bind(f func(interface{}) Monad) Monad {
    var ys *List = Nil()
    var prev *List = nil
    for l := m; l.IsCons(); l = l.Tail() {
        zs, isOk := f(l.Head()).(*List)
        if isOk {
            if zs != nil {
                for l2 := zs; l.IsCons(); l2 = l2.Tail() {
                    l3 := Cons(l2.Head(), Nil())
                    if prev != nil {
                        prev.SetTail(l3)
                    } else {
                        ys = l3
                    }
                    prev = l3
                }
            }
        }
    }
    return ys
}

func ListUnit(x interface{}) Monad {
    return Cons(x, Nil())
}

func (m InterfaceSlice) Bind(f func(interface{}) Monad) Monad {
    ys := make([]interface{}, 0, len(m))
    for _, x := range m {
        m2, isOk := f(x).(InterfaceSlice)
        if isOk {
            if m2 != nil {
                for _, y := range m2 {
                    ys = append(ys, y)
                }
            }
        }
    }
    return InterfaceSlice(ys)
}

func InterfaceSliceUnit(x interface{}) Monad {
    return InterfaceSlice([]interface{} { x })
}

func (m InterfacePairMap) Bind(f func(interface{}) Monad) Monad {
    ys := make(map[interface{}]interface{}, len(m))
    for k, v := range m {
        m2, isOk := f(NewPair(k, v)).(InterfacePairMap)
        if isOk {
            if m2 != nil {
                for k2, v2 := range m2 {
                    ys[k2] = v2
                }
            }
        }
    }
    return InterfacePairMap(ys)
}

func InterfacePairMapUnit(x interface{}) Monad {
    p, isOk := x.(*Pair)
    if isOk {
        return InterfacePairMap(map[interface{}]interface{} { p.First : p.Second })
    } else {
        return nil
    }
}

func (m InterfacePairFunction) Bind(f func(interface{}) Monad) Monad {
    return InterfacePairFunction(func(x interface{}) interface{} {
            g, isOk := f(m(x)).(InterfacePairFunction)
            if isOk {
                if g != nil {
                    return g(x)
                } else {
                    return x
                }
            } else {
                return x
            }
    })
}

func InterfacePairFunctionUnit(x interface{}) Monad {
    return InterfacePairFunction(func(y interface{}) interface{} {
            return x
    })
}
