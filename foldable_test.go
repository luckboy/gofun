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

func TestFoldLeftMethodFoldsNone(t *testing.T) {
    xs := None().FoldLeft(func(x, y interface{}) interface{} {
            return append(InterfaceSliceOrElse(x, InterfaceSlice([]interface{} {})), y)
    }, InterfaceSlice([]interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} {})) {
        t.Errorf("FoldLeft method result is %v; want %v", xs, InterfaceSlice([]interface{} {}))
    }
}

func TestFoldLeftMethodFoldsSome(t *testing.T) {
    xs := Some(1).FoldLeft(func(x, y interface{}) interface{} {
            return append(InterfaceSliceOrElse(x, InterfaceSlice([]interface{} {})), y)
    }, InterfaceSlice([]interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} { 1 })) {
        t.Errorf("FoldLeft method result is %v; want %v", xs, InterfaceSlice([]interface{} { 1 }))
    }
}

func TestFoldRightMethodFoldsNone(t *testing.T) {
    xs := None().FoldRight(func(y, x interface{}) interface{} {
            return append(InterfaceSliceOrElse(x, InterfaceSlice([]interface{} {})), y)
    }, InterfaceSlice([]interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} {})) {
        t.Errorf("FoldRight method result is %v; want %v", xs, InterfaceSlice([]interface{} {}))
    }
}

func TestFoldRightMethodFoldsSome(t *testing.T) {
    xs := Some(1).FoldRight(func(y, x interface{}) interface{} {
            return append(InterfaceSliceOrElse(x, InterfaceSlice([]interface{} {})), y)
    }, InterfaceSlice([]interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} { 1 })) {
        t.Errorf("FoldRight method result is %v; want %v", xs, InterfaceSlice([]interface{} { 1 }))
    }
}

func TestFoldLeftMethodFoldsLeft(t *testing.T) {
    xs := Left("error").FoldLeft(func(x, y interface{}) interface{} {
            return append(InterfaceSliceOrElse(x, InterfaceSlice([]interface{} {})), y)
    }, InterfaceSlice([]interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} {})) {
        t.Errorf("FoldLeft method result is %v; want %v", xs, InterfaceSlice([]interface{} {}))
    }
}

func TestFoldLeftMethodFoldsRight(t *testing.T) {
    xs := Right(2).FoldLeft(func(x, y interface{}) interface{} {
            return append(InterfaceSliceOrElse(x, InterfaceSlice([]interface{} {})), y)
    }, InterfaceSlice([]interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} { 2 })) {
        t.Errorf("FoldLeft method result is %v; want %v", xs, InterfaceSlice([]interface{} { 2 }))
    }
}

func TestFoldRightMethodFoldsLeft(t *testing.T) {
    xs := Left("error").FoldRight(func(y, x interface{}) interface{} {
            return append(InterfaceSliceOrElse(x, InterfaceSlice([]interface{} {})), y)
    }, InterfaceSlice([]interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} {})) {
        t.Errorf("FoldRight method result is %v; want %v", xs, InterfaceSlice([]interface{} {}))
    }
}

func TestFoldRightMethodFoldsRight(t *testing.T) {
    xs := Right(2).FoldRight(func(y, x interface{}) interface{} {
            return append(InterfaceSliceOrElse(x, InterfaceSlice([]interface{} {})), y)
    }, InterfaceSlice([]interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} { 2 })) {
        t.Errorf("FoldRight method result is %v; want %v", xs, InterfaceSlice([]interface{} { 2 }))
    }
}

func TestFoldLeftMethodFoldsEmptyList(t *testing.T) {
    xs := Nil().FoldLeft(func(x, y interface{}) interface{} {
            return append(InterfaceSliceOrElse(x, InterfaceSlice([]interface{} {})), y)
    }, InterfaceSlice([]interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} {})) {
        t.Errorf("FoldLeft method result is %v; want %v", xs, InterfaceSlice([]interface{} {}))
    }
}

func TestFoldLeftMethodFoldsList(t *testing.T) {
    xs := Cons(1, Cons(2, Cons(3, Nil()))).FoldLeft(func(x, y interface{}) interface{} {
            return append(InterfaceSliceOrElse(x, InterfaceSlice([]interface{} {})), y)
    }, InterfaceSlice([]interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} { 1, 2, 3 })) {
        t.Errorf("FoldLeft method result is %v; want %v", xs, InterfaceSlice([]interface{} { 1, 2, 3 }))
    }
}

func TestFoldRightMethodFoldsEmptyList(t *testing.T) {
    xs := Nil().FoldRight(func(y, x interface{}) interface{} {
            return append(InterfaceSliceOrElse(x, InterfaceSlice([]interface{} {})), y)
    }, InterfaceSlice([]interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} {})) {
        t.Errorf("FoldRight method result is %v; want %v", xs, InterfaceSlice([]interface{} {}))
    }
}

func TestFoldRightMethodFoldsList(t *testing.T) {
    xs := Cons(1, Cons(2, Cons(3, Nil()))).FoldRight(func(y, x interface{}) interface{} {
            return append(InterfaceSliceOrElse(x, InterfaceSlice([]interface{} {})), y)
    }, InterfaceSlice([]interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} { 3, 2, 1 })) {
        t.Errorf("FoldRight method result is %v; want %v", xs, InterfaceSlice([]interface{} { 3, 2, 1 }))
    }
}

func TestFoldLeftMethodFoldsEmptyIntefaceSlice(t *testing.T) {
    xs := InterfaceSlice([]interface{} {}).FoldLeft(func(x, y interface{}) interface{} {
            return append(InterfaceSliceOrElse(x, InterfaceSlice([]interface{} {})), y)
    }, InterfaceSlice([]interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} {})) {
        t.Errorf("FoldLeft method result is %v; want %v", xs, InterfaceSlice([]interface{} {}))
    }
}

func TestFoldLeftMethodFoldsInterfaceSlice(t *testing.T) {
    xs := InterfaceSlice([]interface{} { 1, 2, 3 }).FoldLeft(func(x, y interface{}) interface{} {
            return append(InterfaceSliceOrElse(x, InterfaceSlice([]interface{} {})), y)
    }, InterfaceSlice([]interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} { 1, 2, 3 })) {
        t.Errorf("FoldLeft method result is %v; want %v", xs, InterfaceSlice([]interface{} { 1, 2, 3 }))
    }
}

func TestFoldRightMethodFoldsEmptyIntefaceSlice(t *testing.T) {
    xs := InterfaceSlice([]interface{} {}).FoldRight(func(y, x interface{}) interface{} {
            return append(InterfaceSliceOrElse(x, InterfaceSlice([]interface{} {})), y)
    }, InterfaceSlice([]interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} {})) {
        t.Errorf("FoldLeft method result is %v; want %v", xs, InterfaceSlice([]interface{} {}))
    }
}

func TestFoldRightMethodFoldsInterfaceSlice(t *testing.T) {
    xs := InterfaceSlice([]interface{} { 1, 2, 3 }).FoldRight(func(y, x interface{}) interface{} {
            return append(InterfaceSliceOrElse(x, InterfaceSlice([]interface{} {})), y)
    }, InterfaceSlice([]interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} { 3, 2, 1 })) {
        t.Errorf("FoldRight method result is %v; want %v", xs, InterfaceSlice([]interface{} { 3, 2, 1 }))
    }
}

func TestFoldLeftMethodFoldsEmptyInterfacePairMap(t *testing.T) {
    xs := InterfacePairMap(map[interface{}]interface{} {}).FoldLeft(func(x, y interface{}) interface{} {
           ys := InterfacePairMapOrElse(x, map[interface{}]interface{} {})
           p := PairOrElse(y, NewPair("", 0))
           ys[p.First] = p.Second
           return ys
    }, InterfacePairMap(map[interface{}]interface{} {}))
    if !reflect.DeepEqual(xs, InterfacePairMap(map[interface{}]interface{} {})) {
        t.Errorf("FoldRight method result is %v; want %v", xs, InterfaceSlice([]interface{} {}))
    }
}

func TestFoldLeftMethodFoldsInterfacePairMap(t *testing.T) {
    xs := InterfacePairMap(map[interface{}]interface{} { "a": 1, "b": 2 }).FoldLeft(func(x, y interface{}) interface{} {
           ys := InterfacePairMapOrElse(x, map[interface{}]interface{} {})
           p := PairOrElse(y, NewPair("", 0))
           ys[p.First] = p.Second
           return ys
    }, InterfacePairMap(map[interface{}]interface{} {}))
    if !reflect.DeepEqual(xs, InterfacePairMap(map[interface{}]interface{} { "a": 1, "b": 2 })) {
        t.Errorf("FoldLeft method result is %v; want %v", xs, InterfacePairMap(map[interface{}]interface{} { "a": 1, "b": 2 }))
    }
}

func TestFoldRightMethodFoldsEmptyInterfacePairMap(t *testing.T) {
    xs := InterfacePairMap(map[interface{}]interface{} {}).FoldRight(func(y, x interface{}) interface{} {
           ys := InterfacePairMapOrElse(x, map[interface{}]interface{} {})
           p := PairOrElse(y, NewPair("", 0))
           ys[p.First] = p.Second
           return ys
    }, InterfacePairMap(map[interface{}]interface{} {}))
    if !reflect.DeepEqual(xs, InterfacePairMap(map[interface{}]interface{} {})) {
        t.Errorf("FoldRight method result is %v; want %v", xs, InterfacePairMap(map[interface{}]interface{} {}))
    }
}

func TestFoldRightMethodFoldsInterfacePairMap(t *testing.T) {
    xs := InterfacePairMap(map[interface{}]interface{} { "a": 1, "b": 2 }).FoldRight(func(y, x interface{}) interface{} {
           ys := InterfacePairMapOrElse(x, map[interface{}]interface{} {})
           p := PairOrElse(y, NewPair("", 0))
           ys[p.First] = p.Second
           return ys
    }, InterfacePairMap(map[interface{}]interface{} {}))
    if !reflect.DeepEqual(xs, InterfacePairMap(map[interface{}]interface{} { "a": 1, "b": 2 })) {
        t.Errorf("FoldRight method result is %v; want %v", xs, InterfacePairMap(map[interface{}]interface{} { "a": 1, "b": 2 }))
    }
}

func TestAllFunctionReturnsFalse(t *testing.T) {
    b := All(func(x interface{}) bool {
            return IntOrElse(x, 0) % 2 == 0
    }, InterfaceSlice([]interface{} { 2, 3, 4 }))
    if b != false {
        t.Errorf("All function result is %v; want %v", b, false)
    }
}

func TestAllFunctionReturnsTrue(t *testing.T) {
    b := All(func(x interface{}) bool {
            return IntOrElse(x, 0) % 2 == 0
    }, InterfaceSlice([]interface{} { 2, 4, 6 }))
    if b != true {
        t.Errorf("All function result is %v; want %v", b, true)
    }
}

func TestAllMFunctionReturnsFalseMonad(t *testing.T) {
    m := AllM(func(x interface{}) Monad {
            return GetST().Bind(func(s interface{}) Monad {
                    return SetST(IntOrElse(s, 0) + 1).Bind(func(r interface{}) Monad {
                            return STUnit((IntOrElse(x, 0) + IntOrElse(s, 0)) % 2 == 0)
                    })
            })
    }, InterfaceSlice([]interface{} { 1, 3, 3 }), STUnit)
    l, isOk := m.(ST)
    if !isOk {
        t.Errorf("AllM function result type isn't ST")
    } else {
        s, x := RunST(l, 0)
        if !reflect.DeepEqual(s, 1) {
            t.Errorf("RunST function first result from AllM function result is %v; want %v", s, 1)
        }
        if !reflect.DeepEqual(x, false) {
            t.Errorf("RunST function second result from AllM function result is %v; want %v", x, false)
        }
    }
}

func TestAllMFunctionReturnsTrueMonad(t *testing.T) {
    m := AllM(func(x interface{}) Monad {
            return GetST().Bind(func(s interface{}) Monad {
                    return SetST(IntOrElse(s, 0) + 1).Bind(func(r interface{}) Monad {
                            return STUnit((IntOrElse(x, 0) + IntOrElse(s, 0)) % 2 == 0)
                    })
            })
    }, InterfaceSlice([]interface{} { 2, 3, 4 }), STUnit)
    l, isOk := m.(ST)
    if !isOk {
        t.Errorf("AllM function result type isn't ST")
    } else {
        s, x := RunST(l, 0)
        if !reflect.DeepEqual(s, 3) {
            t.Errorf("RunST function first result from AllM function result is %v; want %v", s, 3)
        }
        if !reflect.DeepEqual(x, true) {
            t.Errorf("RunST function second result from AllM function result is %v; want %v", x, true)
        }
    }
}

func TestAnyFunctionReturnsFalse(t *testing.T) {
    b := Any(func(x interface{}) bool {
            return IntOrElse(x, 0) % 2 == 0
    }, InterfaceSlice([]interface{} { 1, 3, 5 }))
    if b != false {
        t.Errorf("Any function result is %v; want %v", b, false)
    }
}

func TestAnyFunctionReturnsTrue(t *testing.T) {
    b := Any(func(x interface{}) bool {
            return IntOrElse(x, 0) % 2 == 0
    }, InterfaceSlice([]interface{} { 1, 2, 3 }))
    if b != true {
        t.Errorf("Any function result is %v; want %v", b, true)
    }
}

func TestAnyMFunctionReturnsFalseMonad(t *testing.T) {
    m := AnyM(func(x interface{}) Monad {
            return GetST().Bind(func(s interface{}) Monad {
                    return SetST(IntOrElse(s, 0) + 1).Bind(func(r interface{}) Monad {
                            return STUnit((IntOrElse(x, 0) + IntOrElse(s, 0)) % 2 == 0)
                    })
            })
    }, InterfaceSlice([]interface{} { 1, 2, 3 }), STUnit)
    l, isOk := m.(ST)
    if !isOk {
        t.Errorf("AnyM function result type isn't ST")
    } else {
        s, x := RunST(l, 0)
        if !reflect.DeepEqual(s, 3) {
            t.Errorf("RunST function first result from AnyM function result is %v; want %v", s, 3)
        }
        if !reflect.DeepEqual(x, false) {
            t.Errorf("RunST function second result from AnyM function result is %v; want %v", x, false)
        }
    }
}

func TestAnyMFunctionReturnsTrueMonad(t *testing.T) {
    m := AnyM(func(x interface{}) Monad {
            return GetST().Bind(func(s interface{}) Monad {
                    return SetST(IntOrElse(s, 0) + 1).Bind(func(r interface{}) Monad {
                            return STUnit((IntOrElse(x, 0) + IntOrElse(s, 0)) % 2 == 0)
                    })
            })
    }, InterfaceSlice([]interface{} { 1, 3, 3 }), STUnit)
    l, isOk := m.(ST)
    if !isOk {
        t.Errorf("AnyM function result type isn't ST")
    } else {
        s, x := RunST(l, 0)
        if !reflect.DeepEqual(s, 2) {
            t.Errorf("RunST function first result from AnyM function result is %v; want %v", s, 2)
        }
        if !reflect.DeepEqual(x, true) {
            t.Errorf("RunST function second result from AnyM function result is %v; want %v", x, true)
        }
    }
}

func TestDeepElementFunctionFindsElement(t *testing.T) {
    b := DeepElement(Some(2), InterfaceSlice([]interface{} { Some(1), Some(2), Some(3) }))
    if b != true {
        t.Errorf("DeepElement function result is %v; want %v", b, true)
    }
}

func TestDeepElementFunctionDoesNotFindElement(t *testing.T) {
    b := DeepElement(Some(4), InterfaceSlice([]interface{} { Some(1), Some(2), Some(3) }))
    if b != false {
        t.Errorf("DeepElement function result is %v; want %v", b, false)
    }
}

func TestElementFunctionFindsElement(t *testing.T) {
    b := Element(2, InterfaceSlice([]interface{} { 1, 2, 3 }))
    if b != true {
        t.Errorf("Element function result is %v; want %v", b, true)
    }
}

func TestElementFunctionDoesNotFindElement(t *testing.T) {
    b := Element(4, InterfaceSlice([]interface{} { 1, 2, 3 }))
    if b != false {
        t.Errorf("Element function result is %v; want %v", b, false)
    }
}

func TestFilterFunctionFilters(t *testing.T) {
    xs := Filter(func(x interface{}) bool {
            return IntOrElse(x, 0) % 2 == 0
    }, InterfaceSlice([]interface{} { 1, 2, 3, 4 }))
    if !reflect.DeepEqual(xs, Cons(2, Cons(4, Nil()))) {
        t.Errorf("Filter function result is %v; want %v", xs, Cons(2, Cons(4, Nil())))
    }
}

func TestFilterMFunctionFilters(t *testing.T) {
    m := FilterM(func(x interface{}) Monad {
            return GetST().Bind(func(s interface{}) Monad {
                    return SetST(IntOrElse(s, 0) + 1).Bind(func(r interface{}) Monad {
                            return STUnit((IntOrElse(x, 0) + IntOrElse(s, 0)) % 2 == 0)
                    })
            })
    }, InterfaceSlice([]interface{} { 1, 3, 5, 7 }), STUnit)
    l, isOk := m.(ST)
    if !isOk {
        t.Errorf("FilterM function result type isn't ST")
    } else {
        s, x := RunST(l, 0)
        if !reflect.DeepEqual(s, 4) {
            t.Errorf("RunST function first result from FilterM function result is %v; want %v", s, 4)
        }
        if !reflect.DeepEqual(x, Cons(3, Cons(7, Nil()))) {
            t.Errorf("RunST function second result from FilterM function result is %v; want %v", x, Cons(3, Cons(7, Nil())))
        }
    }
}

func TestFilterSliceFunctionFilters(t *testing.T) {
    xs := FilterSlice(func(x interface{}) bool {
            return IntOrElse(x, 0) % 2 == 0
    }, InterfaceSlice([]interface{} { 1, 2, 3, 4 }))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} { 2, 4 })) {
        t.Errorf("FilterSlice function result is %v; want %v", xs, InterfaceSlice([]interface{} { 2, 4 }))
    }
}

func TestFilterSliceMFunctionFilters(t *testing.T) {
    m := FilterSliceM(func(x interface{}) Monad {
            return GetST().Bind(func(s interface{}) Monad {
                    return SetST(IntOrElse(s, 0) + 1).Bind(func(r interface{}) Monad {
                            return STUnit((IntOrElse(x, 0) + IntOrElse(s, 0)) % 2 == 0)
                    })
            })
    }, InterfaceSlice([]interface{} { 1, 3, 5, 7 }), STUnit)
    l, isOk := m.(ST)
    if !isOk {
        t.Errorf("FilterSliceM function result type isn't ST")
    } else {
        s, x := RunST(l, 0)
        if !reflect.DeepEqual(s, 4) {
            t.Errorf("RunST function first result from FilterSliceM function result is %v; want %v", s, 4)
        }
        if !reflect.DeepEqual(x, InterfaceSlice([]interface{} { 3, 7 })) {
            t.Errorf("RunST function second result from FilterSliceM function result is %v; want %v", x, InterfaceSlice([]interface{} { 3, 7 }))
        }
    }
}

func TestFindFunctionFindsElement(t *testing.T) {
    o := Find(func(x interface{}) bool {
            return IntOrElse(x, 0) % 2 == 0
    }, InterfaceSlice([]interface{} { 1, 2, 3 }))
    if !reflect.DeepEqual(o, Some(2)) {
        t.Errorf("Find function result is %v; want %v", o, Some(2))
    }
}

func TestFindFunctionDoesNotFindElement(t *testing.T) {
    o := Find(func(x interface{}) bool {
            return IntOrElse(x, 0) % 2 == 0
    }, InterfaceSlice([]interface{} { 1, 3, 5 }))
    if !reflect.DeepEqual(o, None()) {
        t.Errorf("Find function result is %v; want %v", o, None())
    }
}

func TestFindMFunctionFindsElement(t *testing.T) {
    m := FindM(func(x interface{}) Monad {
            return GetST().Bind(func(s interface{}) Monad {
                    return SetST(IntOrElse(s, 0) + 1).Bind(func(r interface{}) Monad {
                            return STUnit((IntOrElse(x, 0) + IntOrElse(s, 0)) % 2 == 0)
                    })
            })
    }, InterfaceSlice([]interface{} { 1, 3, 5 }), STUnit)
    l, isOk := m.(ST)
    if !isOk {
        t.Errorf("FindM function result type isn't ST")
    } else {
        s, x := RunST(l, 0)
        if !reflect.DeepEqual(s, 2) {
            t.Errorf("RunST function first result from FindM function result is %v; want %v", s, 2)
        }
        if !reflect.DeepEqual(x, Some(3)) {
            t.Errorf("RunST function second result from FindM function result is %v; want %v", x, Some(3))
        }
    }
}

func TestFindMFunctionDoesNotFindElement(t *testing.T) {
    m := FindM(func(x interface{}) Monad {
            return GetST().Bind(func(s interface{}) Monad {
                    return SetST(IntOrElse(s, 0) + 1).Bind(func(r interface{}) Monad {
                            return STUnit((IntOrElse(x, 0) + IntOrElse(s, 0)) % 2 == 0)
                    })
            })
    }, InterfaceSlice([]interface{} { 1, 2, 3 }), STUnit)
    l, isOk := m.(ST)
    if !isOk {
        t.Errorf("FindM function result type isn't ST")
    } else {
        s, x := RunST(l, 0)
        if !reflect.DeepEqual(s, 3) {
            t.Errorf("RunST function first result from FindM function result is %v; want %v", s, 3)
        }
        if !reflect.DeepEqual(x, None()) {
            t.Errorf("RunST function second result from FindM function result is %v; want %v", x, None())
        }
    }
}

func TestFoldLeftMFunctionFolds(t *testing.T) {
    m := FoldLeftM(func(x, y interface{}) Monad {
            return GetST().Bind(func(s interface{}) Monad {
                    return SetST(IntOrElse(s, 0) + 1).Bind(func(r interface{}) Monad {
                            return STUnit(append(InterfaceSliceOrElse(x, InterfaceSlice([]interface{} {})), IntOrElse(y, 0) + IntOrElse(s, 0)))
                    })
            })
    }, InterfaceSlice([]interface{} {}), InterfaceSlice([]interface{} { 1, 2, 3 }), STUnit)
    l, isOk := m.(ST)
    if !isOk {
        t.Errorf("FoldLeftM function result type isn't ST")
    } else {
        s, x := RunST(l, 0)
        if !reflect.DeepEqual(s, 3) {
            t.Errorf("RunST function first result from FoldLeftM function result is %v; want %v", s, 3)
        }
        if !reflect.DeepEqual(x, InterfaceSlice([]interface{} { 1, 3, 5 })) {
            t.Errorf("RunST function second result from FoldLeftM function result is %v; want %v", x, InterfaceSlice([]interface{} { 1, 3, 5 }))
        }
    }
}

func TestFoldRightMFunctionFolds(t *testing.T) {
    m := FoldRightM(func(y, x interface{}) Monad {
            return GetST().Bind(func(s interface{}) Monad {
                    return SetST(IntOrElse(s, 0) + 1).Bind(func(r interface{}) Monad {
                            return STUnit(append(InterfaceSliceOrElse(x, InterfaceSlice([]interface{} {})), IntOrElse(y, 0) + IntOrElse(s, 0)))
                    })
            })
    }, InterfaceSlice([]interface{} {}), InterfaceSlice([]interface{} { 1, 2, 3 }), STUnit)
    l, isOk := m.(ST)
    if !isOk {
        t.Errorf("FoldRightM function result type isn't ST")
    } else {
        s, x := RunST(l, 0)
        if !reflect.DeepEqual(s, 3) {
            t.Errorf("RunST function first result from FoldRightM function result is %v; want %v", s, 3)
        }
        if !reflect.DeepEqual(x, InterfaceSlice([]interface{} { 3, 3, 3 })) {
            t.Errorf("RunST function second result from FoldRightM function result is %v; want %v", x, InterfaceSlice([]interface{} { 3, 3, 3 }))
        }
    }
}

func TestLengthFunctionCalculatesLength(t *testing.T) {
    x := Length(InterfaceSlice([]interface{} { 1, 2, 3 }))
    if x != 3 {
        t.Errorf("Length function result is %v; want %v", x, 3)
    }
}

func TestNotDeepElementFunctionFindsElement(t *testing.T) {
    b := NotDeepElement(Some(2), InterfaceSlice([]interface{} { Some(1), Some(2), Some(3) }))
    if b != false {
        t.Errorf("NotDeepElement function result is %v; want %v", b, false)
    }
}

func TestNotDeepElementFunctionDoesNotFindElement(t *testing.T) {
    b := NotDeepElement(Some(4), InterfaceSlice([]interface{} { Some(1), Some(2), Some(3) }))
    if b != true {
        t.Errorf("NotDeepElement function result is %v; want %v", b, true)
    }
}

func TestNotElementFunctionFindsElement(t *testing.T) {
    b := NotElement(2, InterfaceSlice([]interface{} { 1, 2, 3 }))
    if b != false {
        t.Errorf("NotElement function result is %v; want %v", b, false)
    }
}

func TestNotElementFunctionDoesNotFindElement(t *testing.T) {
    b := NotElement(4, InterfaceSlice([]interface{} { 1, 2, 3 }))
    if b != true {
        t.Errorf("NotElement function result is %v; want %v", b, true)
    }
}

func TestNullFunctionReturnsFalse(t *testing.T) {
    b := Null(InterfaceSlice([]interface{} { 1, 2, 3 }))
    if b != false {
        t.Errorf("Null function result is %v; want %v", b, false)
    }
}

func TestNullFunctionReturnsTrue(t *testing.T) {
    b := Null(InterfaceSlice([]interface{} {}))
    if b != true {
        t.Errorf("Null function result is %v; want %v", b, true)
    }
}

func TestToListFunctionReturnsList(t *testing.T) {
    xs := ToList(InterfaceSlice([]interface{} { 1, 2, 3 }))
    if !reflect.DeepEqual(xs, Cons(1, Cons(2, Cons(3, Nil())))) {
        t.Errorf("ToList function result is %v; want %v", xs, Cons(1, Cons(2, Cons(3, Nil()))))
    }
}

func TestToSliceFunctionReturnsSlice(t *testing.T) {
    xs := ToSlice(InterfaceSlice([]interface{} { 1, 2, 3 }))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} { 1, 2, 3 })) {
        t.Errorf("ToSlice function result is %v; want %v", xs, InterfaceSlice([]interface{} { 1, 2, 3 }))
    }
}
