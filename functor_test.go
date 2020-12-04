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

func TestMapMethodMapsNone(t *testing.T) {
    xs := None().Map(func(x interface{}) interface{} {
            return IntOrElse(x, 0) + 2
    })
    if !reflect.DeepEqual(xs, None()) {
        t.Errorf("Map method result is %v; want %v", xs, None())
    }
}

func TestMapMethodMapsSome(t *testing.T) {
    xs := Some(1).Map(func(x interface{}) interface{} {
            return IntOrElse(x, 0) + 1
    })
    if !reflect.DeepEqual(xs, Some(2)) {
        t.Errorf("Map method result is %v; want %v", xs, Some(2))
    }
}

func TestMapMethodMapsLeft(t *testing.T) {
    xs := Left("error").Map(func(x interface{}) interface{} {
            return IntOrElse(x, 0) + 2
    })
    if !reflect.DeepEqual(xs, Left("error")) {
        t.Errorf("Map method result is %v; want %v", xs, Left("error"))
    }
}

func TestMapMethodMapsRight(t *testing.T) {
    xs := Right(1).Map(func(x interface{}) interface{} {
            return IntOrElse(x, 0) + 1
    })
    if !reflect.DeepEqual(xs, Right(2)) {
        t.Errorf("Map method result is %v; want %v", xs, Right(2))
    }
}

func TestMapMethodMapsEmptyList(t *testing.T) {
    xs := Nil().Map(func(x interface{}) interface{} {
            return IntOrElse(x, 0) + 1
    })
    if !reflect.DeepEqual(xs, Nil()) {
        t.Errorf("Map method result is %v; want %v", xs, Nil())
    }
}

func TestMapMethodMapsList(t *testing.T) {
    xs := Cons(1, Cons(2, Cons(3, Nil()))).Map(func(x interface{}) interface{} {
            return IntOrElse(x, 0) + 3
    })
    if !reflect.DeepEqual(xs, Cons(4, Cons(5, Cons(6, Nil())))) {
        t.Errorf("Map method result is %v; want %v", xs, Cons(4, Cons(5, Cons(6, Nil()))))
    }
}

func TestMapMethodMapsST(t *testing.T) {
    xs := ST(func(s interface{}) (interface{}, interface{}) {
            return s, 2
    }).Map(func(x interface{}) interface{} {
            return IntOrElse(x, 0) + 2
    })
    ys, isOk := xs.(ST)
    if !isOk {
        t.Errorf("Map result type isn't ST")
    } else {
        s, x := RunST(ys, 1)
        if !reflect.DeepEqual(s, 1) {
            t.Errorf("RunST function first result from Map method result is %v; want %v", s, 1)
        }
        if !reflect.DeepEqual(x, 4) {
            t.Errorf("RunST function second result from Map method result is %v; want %v", x, 4)
        }
    }
}

func TestMapMethodMapsEmptyInterfaceSlice(t *testing.T) {
    xs := InterfaceSlice([]interface{} {}).Map(func(x interface{}) interface{} {
            return IntOrElse(x, 0) + 1
    })
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} {})) {
        t.Errorf("Map method result is %v; want %v", xs, InterfaceSlice([]interface{} {}))
    }
}

func TestMapMethodMapsInterfaceSlice(t *testing.T) {
    xs := InterfaceSlice([]interface{} { 1, 2, 3 }).Map(func(x interface{}) interface{} {
            return IntOrElse(x, 0) + 2
    })
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} { 3, 4, 5 })) {
        t.Errorf("Map method result is %v; want %v", xs, InterfaceSlice([]interface{} { 3, 4, 5 }))
    }
}

func TestMapMethodMapsEmptyInterfacePairMap(t *testing.T) {
    xs := InterfacePairMap(map[interface{}]interface{} {}).Map(func(x interface{}) interface{} {
            p := PairOrElse(x, NewPair("", 0))
            return NewPair(StringOrElse(p.First, "") + "x", IntOrElse(p.Second, 0) + 1)
    })
    if !reflect.DeepEqual(xs, InterfacePairMap(map[interface{}]interface{} {})) {
        t.Errorf("Map method result is %v; want %v", xs, InterfacePairMap(map[interface{}]interface{} {}))
    }
}

func TestMapMethodMapsInterfacePairMap(t *testing.T) {
    xs := InterfacePairMap(map[interface{}]interface{} { "a": 1, "b": 2 }).Map(func(x interface{}) interface{} {
            p := PairOrElse(x, NewPair("", 0))
            return NewPair(StringOrElse(p.First, "") + "x", IntOrElse(p.Second, 0) + 2)
    })
    if !reflect.DeepEqual(xs, InterfacePairMap(map[interface{}]interface{} { "ax": 3, "bx": 4 })) {
        t.Errorf("Map method result is %v; want %v", xs, InterfacePairMap(map[interface{}]interface{} { "ax": 3, "bx": 4 }))
    }
}

func TestMapMethodMapsInterfacePairFunction(t *testing.T) {
    xs := InterfacePairFunction(func(x interface{}) interface{} {
            return IntOrElse(x, 0) + 1
    }).Map(func(x interface{}) interface{} {
            return IntOrElse(x, 0) + 2
    })
    ys, isOk := xs.(InterfacePairFunction)
    if !isOk {
        t.Errorf("Map result type isn't InterfacePairFunction")
    } else {
        y := ys(4)
        if !reflect.DeepEqual(y, 7) {
            t.Errorf("function result of Map method result is %v; want %v", y, 7)
        }
    }
}
