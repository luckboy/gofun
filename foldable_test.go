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
