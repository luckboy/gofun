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

package gofun_test
import (
    "reflect"
    "testing"
    . "gofun"
)

func TestBindMethodBindsNone(t *testing.T) {
    m := None().Bind(func(x interface{}) Monad {
            return OptionUnit(IntOrElse(x, 0) + 1)
    })
    if !reflect.DeepEqual(m, None()) {
        t.Errorf("Bind method result is %v; want %v", m, None())
    }
}

func TestBindMethodBindsSome(t *testing.T) {
    m := Some(2).Bind(func(x interface{}) Monad {
            return OptionUnit(IntOrElse(x, 0) + 1)
    })
    if !reflect.DeepEqual(m, Some(3)) {
        t.Errorf("Bind method result is %v; want %v", m, Some(3))
    }
}

func TestBindMethodBindsSomeForMapMethod(t *testing.T) {
    m := Some(3).Bind(func(x interface{}) Monad {
            m2 := Some(2).Map(func(y interface{}) interface{} {
                    return IntOrElse(x, 0) + IntOrElse(y, 0) + 1
            })
            return MonadOrElse(m2, None())
    })
    if !reflect.DeepEqual(m, Some(6)) {
        t.Errorf("Bind method result is %v; want %v", m, Some(6))
    }
}

func TestBindMethodBindsLeft(t *testing.T) {
    m := Left("error").Bind(func(x interface{}) Monad {
            return EitherUnit(IntOrElse(x, 0) + 1)
    })
    if !reflect.DeepEqual(m, Left("error")) {
        t.Errorf("Bind method result is %v; want %v", m, Left("error"))
    }
}

func TestBindMethodBindsRight(t *testing.T) {
    m := Right(3).Bind(func(x interface{}) Monad {
            return EitherUnit(IntOrElse(x, 0) + 1)
    })
    if !reflect.DeepEqual(m, Right(4)) {
        t.Errorf("Bind method result is %v; want %v", m, Right(4))
    }
}

func TestBindMethodBindsRightForMapMethod(t *testing.T) {
    m := Right(3).Bind(func(x interface{}) Monad {
            m2 := Right(2).Map(func(y interface{}) interface{} {
                    return IntOrElse(x, 0) + IntOrElse(y, 0) + 1
            })
            return MonadOrElse(m2, Left("error"))
    })
    if !reflect.DeepEqual(m, Right(6)) {
        t.Errorf("Bind method result is %v; want %v", m, Right(6))
    }
}

func TestBindMethodBindsEmptyList(t *testing.T) {
    m := Nil().Bind(func(x interface{}) Monad {
            return ListUnit(IntOrElse(x, 0) + 1)
    })
    if !reflect.DeepEqual(m, Nil()) {
        t.Errorf("Bind method result is %v; want %v", m, Nil())
    }
}

func TestBindMethodBindsList(t *testing.T) {
    m := Cons(2, Cons(3, Nil())).Bind(func(x interface{}) Monad {
            return ListUnit(IntOrElse(x, 0) + 1)
    })
    if !reflect.DeepEqual(m, Cons(3, Cons(4, Nil()))) {
        t.Errorf("Bind method result is %v; want %v", m, Cons(3, Cons(4, Nil())))
    }
}

func TestBindMethodBindsListForMapMethod(t *testing.T) {
    m := Cons(2, Cons(3, Nil())).Bind(func(x interface{}) Monad {
            m2 := Cons(4, Cons(5, Nil())).Map(func(y interface{}) interface{} {
                    return IntOrElse(x, 0) + IntOrElse(y, 0) + 1
            })
            return MonadOrElse(m2, Nil())
    })
    if !reflect.DeepEqual(m, Cons(7, Cons(8, Cons(8, Cons(9, Nil()))))) {
        t.Errorf("Bind method result is %v; want %v", m, Cons(7, Cons(8, Cons(8, Cons(9, Nil())))))
    }
}

func TestBindMethodBindsST(t *testing.T) {
    m := ST(func(s interface{}) (interface{}, interface{}) {
            return s, 2
    }).Bind(func(x interface{}) Monad {
            return STUnit(IntOrElse(x, 0) + 1)
    })
    l, isOk := m.(ST)
    if !isOk {
        t.Errorf("Bind method result type isn't ST")
    } else {
        s, x := RunST(l, 1)
        if !reflect.DeepEqual(s, 1) {
            t.Errorf("RunST function first result from Bind method result is %v; want %v", s, 1)
        }
        if !reflect.DeepEqual(x, 3) {
            t.Errorf("RunST function second result from Bind method result is %v; want %v", x, 4)
        }
    }
}

func TestBindMethodBindsSTForMapMethod(t *testing.T) {
    m := ST(func(s interface{}) (interface{}, interface{}) {
            return s, 2
    }).Bind(func(x interface{}) Monad {
            m2 := ST(func(s interface{}) (interface{}, interface{}) {
                    return IntOrElse(s, 0) + 1, 3
            }).Map(func(y interface {}) interface{} {
                return IntOrElse(x, 0) + IntOrElse(y, 0) + 1
            })
            return MonadOrElse(m2, STUnit(1))
    })
    l, isOk := m.(ST)
    if !isOk {
        t.Errorf("Bind method result type isn't ST")
    } else {
        s, x := RunST(l, 1)
        if !reflect.DeepEqual(s, 2) {
            t.Errorf("RunST function first result from Bind method result is %v; want %v", s, 1)
        }
        if !reflect.DeepEqual(x, 6) {
            t.Errorf("RunST function second result from Bind method result is %v; want %v", x, 4)
        }
    }
}

func TestBindMethodBindsEmptyInterfaceSlice(t *testing.T) {
    m := InterfaceSlice([]interface{} {}).Bind(func(x interface{}) Monad {
            return InterfaceSliceUnit(IntOrElse(x, 0) + 1)
    })
    if !reflect.DeepEqual(m, InterfaceSlice([]interface{} {})) {
        t.Errorf("Bind method result is %v; want %v", m, InterfaceSlice([]interface{} {}))
    }
}

func TestBindMethodBindsInterfaceSlice(t *testing.T) {
    m := InterfaceSlice([]interface{} { 1, 2, 3 }).Bind(func(x interface{}) Monad {
            return InterfaceSliceUnit(IntOrElse(x, 0) + 1)
    })
    if !reflect.DeepEqual(m, InterfaceSlice([]interface{} { 2, 3, 4 })) {
        t.Errorf("Bind method result is %v; want %v", m, InterfaceSlice([]interface{} { 2, 3, 4 }))
    }
}

func TestBindMethodBindsInterfaceSliceForMapMethod(t *testing.T) {
    m := InterfaceSlice([]interface{} { 1, 2, 3 }).Bind(func(x interface{}) Monad {
            m2 := InterfaceSlice([]interface{} { 4, 5 }).Map(func(y interface{}) interface{} {
                    return IntOrElse(x, 0) + IntOrElse(y, 0) + 1
            })
            return MonadOrElse(m2, InterfaceSlice([]interface{} {}))
    })
    if !reflect.DeepEqual(m, InterfaceSlice([]interface{} { 6, 7, 7, 8, 8, 9 })) {
        t.Errorf("Bind method result is %v; want %v", m, InterfaceSlice([]interface{} { 6, 7, 7, 8, 8, 9 }))
    }
}

func TestBindMethodBindsEmptyInterfacePairMap(t *testing.T) {
    m := InterfacePairMap(map[interface{}]interface{} {}).Bind(func(x interface{}) Monad {
            p := PairOrElse(x, NewPair("", 0))
            return InterfacePairMapUnit(NewPair(StringOrElse(p.First, "") + "x", IntOrElse(p.Second, 0) + 1))
    })
    if !reflect.DeepEqual(m, InterfacePairMap(map[interface{}]interface{} {})) {
        t.Errorf("Bind method result is %v; want %v", m, InterfacePairMap(map[interface{}]interface{} {}))
    }
}

func TestBindMethodBindsInterfacePairMap(t *testing.T) {
    m := InterfacePairMap(map[interface{}]interface{} { "a": 1, "b": 2 }).Bind(func(x interface{}) Monad {
            p := PairOrElse(x, NewPair("", 0))
            return InterfacePairMapUnit(NewPair(StringOrElse(p.First, "") + "x", IntOrElse(p.Second, 0) + 1))
    })
    if !reflect.DeepEqual(m, InterfacePairMap(map[interface{}]interface{} { "ax": 2, "bx": 3 })) {
        t.Errorf("Bind method result is %v; want %v", m, InterfacePairMap(map[interface{}]interface{} { "ax": 2, "bx": 3 }))
    }
}

func TestBindMethodBindsInterfacePairMapForMapMethod(t *testing.T) {
    m := InterfacePairMap(map[interface{}]interface{} { "a": 1, "b": 2 }).Bind(func(x interface{}) Monad {
            m2 := InterfacePairMap(map[interface{}]interface{} { "c": 3, "d": 4 }).Map(func(y interface{}) interface{} {
                    p := PairOrElse(x, NewPair("", 0))
                    p2 := PairOrElse(y, NewPair("", 0))
                    return NewPair(StringOrElse(p.First, "") + StringOrElse(p2.First, "") + "x", IntOrElse(p.Second, 0) + IntOrElse(p2.Second, 0) + 1)
            })
            return MonadOrElse(m2, InterfacePairMap(map[interface{}]interface{} {}))
    })
    if !reflect.DeepEqual(m, InterfacePairMap(map[interface{}]interface{} { "acx": 5, "adx": 6, "bcx": 6, "bdx": 7 })) {
        t.Errorf("Bind method result is %v; want %v", m, InterfacePairMap(map[interface{}]interface{} { "acx": 5, "adx": 6, "bcx": 6, "bdx": 7 }))
    }
}

func TestBindMethodBindsInterfacePairFunction(t *testing.T) {
    m := InterfacePairFunction(func(x interface{}) interface{} {
            return IntOrElse(x, 0) + 1
    }).Bind(func(x interface{}) Monad {
            return InterfacePairFunctionUnit(IntOrElse(x, 0) + 2)
    })
    l, isOk := m.(InterfacePairFunction)
    if !isOk {
        t.Errorf("Bind result type isn't InterfacePairFunction")
    } else {
        y := l(4)
        if !reflect.DeepEqual(y, 7) {
            t.Errorf("function result of Bind method result is %v; want %v", y, 7)
        }
    }
}

func TestBindMethodBindsInterfacePairFunctionForMapMethod(t *testing.T) {
    m := InterfacePairFunction(func(x interface{}) interface{} {
            return IntOrElse(x, 0) + 1
    }).Bind(func(x interface{}) Monad {
            m2 := InterfacePairFunction(func(y interface{}) interface{} {
                    return IntOrElse(y, 0) + 2
            }).Map(func(y interface{}) interface{} {
                    return IntOrElse(x, 0) + IntOrElse(y, 0) + 1
            })
            return MonadOrElse(m2, InterfacePairFunction(func(y interface{}) interface{} { return 0 }))
    })
    l, isOk := m.(InterfacePairFunction)
    if !isOk {
        t.Errorf("Bind method result type isn't InterfacePairFunction")
    } else {
        y := l(3)
        if !reflect.DeepEqual(y, 10) {
            t.Errorf("function result of Bind method result is %v; want %v", y, 10)
        }
    }
}
