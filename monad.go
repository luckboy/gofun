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

// Package gofun provides functions, types, and constructions from functional languages.
package gofun

// Monad is the interface for monads.
type Monad interface {
    Functor
    // Bind binds Monad and a function that returns Monad. Binding allows to
    // create data pipes.
    Bind(func(interface{}) Monad) Monad
}

// MonadOrElse returns x if x is Monad, otherwise y.
func MonadOrElse(x interface{}, y Monad) Monad {
    z, isOk := x.(Monad)
    if isOk {
        return z
    } else {
        return y
    }
}

// IfM returns ifTrue() Monad if cond is Monad with true, otherwise ifFalse() Monad. 
func IfM(cond Monad, ifTrue, ifFalse func() Monad) Monad {
    return cond.Bind(func(x interface{}) Monad {
            if BoolOrElse(x, false) {
                return ifTrue()
            } else {
                return ifFalse()
            }
    });
}

// Join joins Monad. Fail must be a failure Monad. 
func Join(m, fail Monad) Monad {
    return m.Bind(func(x interface{}) Monad {
            return MonadOrElse(x, fail)
    })
}

// UntilM is a loop of until type for monads. Unit must be the unit function for specified monad.
func UntilM(m Monad, cond func() Monad, unit func(interface{}) Monad) Monad {
    return m.Bind(func(x interface{}) Monad {
            return cond().Bind(func(y interface{}) Monad {
                    if !BoolOrElse(y, false) {
                        return UntilM(m, cond, unit)
                    } else {
                        return unit(struct{} {})
                    }
            })
    })
}

// WhileM is a loop of while type for monads. Unit must be the unit function for specified monad.
func WhileM(cond Monad, body func() Monad, unit func(interface{}) Monad) Monad {
    return cond.Bind(func(x interface{}) Monad {
            if BoolOrElse(x, false) {
                return body().Bind(func(y interface{}) Monad {
                        return WhileM(cond, body, unit)
                })
            } else {
                return unit(struct{} {})
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

// OptionUnit is an unit function for Option.
func OptionUnit(x interface{}) Monad {
    return Some(x)
}

func (m *Either) Bind(f func(interface{}) Monad) Monad {
    if m.IsRight() {
        return f(m.GetRight())
    } else {
        return Left(m.GetLeft())
    }
}

// EitherUnit is an unit function for Either.
func EitherUnit(x interface{}) Monad {
    return Right(x)
}

func (m *List) Bind(f func(interface{}) Monad) Monad {
    var ys *List = Nil()
    var prev *List = nil
    for l := m; l.IsCons(); l = l.Tail() {
        zs, isOk := f(l.Head()).(*List)
        if isOk {
            for l2 := zs; l2.IsCons(); l2 = l2.Tail() {
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
    return ys
}

// ListUnit is an unit function for List.
func ListUnit(x interface{}) Monad {
    return Cons(x, Nil())
}

func (m ST) Bind(f func(interface{}) Monad) Monad {
    return ST(func(s interface{}) (interface{}, interface{}) {
            s2, x := m(s)
            m2, isOk := f(x).(ST)
            if isOk {
                return m2(s2)
            } else {
                return s2, x
            }
    })
}

// STUnit is an unit function for ST.
func STUnit(x interface{}) Monad {
    return ST(func(s interface{}) (interface{}, interface{}) {
            return s, x
    })
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

// InterfaceSliceUnit is an unit function for InterfaceSlice.
func InterfaceSliceUnit(x interface{}) Monad {
    return InterfaceSlice([]interface{} { x })
}

func (m InterfacePairMap) Bind(f func(interface{}) Monad) Monad {
    ys := make(map[interface{}]interface{}, len(m))
    for k, v := range m {
        m2, isOk := f(NewPair(k, v)).(InterfacePairMap)
        if isOk {
            for k2, v2 := range m2 {
                ys[k2] = v2
            }
        }
    }
    return InterfacePairMap(ys)
}

// InterfacePairMapUnit is an unit function for InterfacePairMap.
func InterfacePairMapUnit(x interface{}) Monad {
    p, isOk := x.(*Pair)
    if isOk {
        return InterfacePairMap(map[interface{}]interface{} { p.First : p.Second })
    } else {
        return InterfacePairMap(map[interface{}]interface{} {})
    }
}

func (m InterfacePairFunction) Bind(f func(interface{}) Monad) Monad {
    return InterfacePairFunction(func(x interface{}) interface{} {
            g, isOk := f(m(x)).(InterfacePairFunction)
            if isOk {
                return g(x)
            } else {
                return x
            }
    })
}

// InterfacePairFunctionUnit is an unit function for InterfacePairFunction.
func InterfacePairFunctionUnit(x interface{}) Monad {
    return InterfacePairFunction(func(y interface{}) interface{} {
            return x
    })
}
