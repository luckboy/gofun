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

func TestUnzipMethodUnzipsNone(t *testing.T) {
    xs, ys := None().Unzip(None())
    if !reflect.DeepEqual(xs, None()) {
        t.Errorf("Unzip method first result is %v; want %v", xs, None())
    }
    if !reflect.DeepEqual(ys, None()) {
        t.Errorf("Unzip method second result is %v; want %v", ys, None())
    }
}

func TestUnzipMethodUnzipsSome(t *testing.T) {
    xs, ys := Some(NewPair(1, 2)).Unzip(None())
    if !reflect.DeepEqual(xs, Some(1)) {
        t.Errorf("Unzip method first result is %v; want %v", xs, Some(1))
    }
    if !reflect.DeepEqual(ys, Some(2)) {
        t.Errorf("Unzip method second result is %v; want %v", ys, Some(2))
    }
}

func TestUnzipMethodUnzipsLeft(t *testing.T) {
    xs, ys := Left("error1").Unzip(Left("error2"))
    if !reflect.DeepEqual(xs, Left("error1")) {
        t.Errorf("Unzip method first result is %v; want %v", xs, Left("error1"))
    }
    if !reflect.DeepEqual(ys, Left("error1")) {
        t.Errorf("Unzip method second result is %v; want %v", ys, Left("error1"))
    }
}

func TestUnzipMethodUnzipsRight(t *testing.T) {
    xs, ys := Right(NewPair(1, 2)).Unzip(Left("error2"))
    if !reflect.DeepEqual(xs, Right(1)) {
        t.Errorf("Unzip method first result is %v; want %v", xs, Right(1))
    }
    if !reflect.DeepEqual(ys, Right(2)) {
        t.Errorf("Unzip method second result is %v; want %v", ys, Right(2))
    }
}

func TestUnzipMethodUnzipsEmptyList(t *testing.T) {
    xs, ys := Nil().Unzip(Nil())
    if !reflect.DeepEqual(xs, Nil()) {
        t.Errorf("Unzip method first result is %v; want %v", xs, Nil())
    }
    if !reflect.DeepEqual(ys, Nil()) {
        t.Errorf("Unzip method second result is %v; want %v", ys, Nil())
    }
}

func TestUnzipMethodUnzipsList(t *testing.T) {
    xs, ys := Cons(NewPair(1, 4), Cons(NewPair(2, 5), Cons(NewPair(3, 6), Nil()))).Unzip(Nil())
    if !reflect.DeepEqual(xs, Cons(1, Cons(2, Cons(3, Nil())))) {
        t.Errorf("Unzip method first result is %v; want %v", xs, Cons(1, Cons(2, Cons(3, Nil()))))
    }
    if !reflect.DeepEqual(ys, Cons(4, Cons(5, Cons(6, Nil())))) {
        t.Errorf("Unzip method second result is %v; want %v", ys, Cons(4, Cons(5, Cons(6, Nil()))))
    }
}

func TestUnzipMethodUnzipsEmptyInterfaceSlice(t *testing.T) {
    xs, ys := InterfaceSlice([]interface{} {}).Unzip(InterfaceSlice([]interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} {})) {
        t.Errorf("Unzip method first result is %v; want %v", xs, InterfaceSlice([]interface{} {}))
    }
    if !reflect.DeepEqual(ys, InterfaceSlice([]interface{} {})) {
        t.Errorf("Unzip method second result is %v; want %v", ys, InterfaceSlice([]interface{} {}))
    }
}

func TestUnzipMethodUnzipsInterfaceSlice(t *testing.T) {
    xs, ys := InterfaceSlice([]interface{} { NewPair(1, 4), NewPair(2, 5), NewPair(3, 6) }).Unzip(InterfaceSlice([]interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([]interface{} { 1, 2, 3 })) {
        t.Errorf("Unzip method first result is %v; want %v", xs, InterfaceSlice([]interface{} { 1, 2, 3 }))
    }
    if !reflect.DeepEqual(ys, InterfaceSlice([]interface{} { 4, 5, 6 })) {
        t.Errorf("Unzip method second result is %v; want %v", ys, InterfaceSlice([]interface{} { 4, 5, 6 }))
    }
}

func TestUnzipMethodUnzipsInterfacePairFunction(t *testing.T) {
    xs, ys := InterfacePairFunction(func(x interface{}) interface{} {
            return NewPair(IntOrElse(x, 0) + 1, IntOrElse(x, 0) + 2)
    }).Unzip(InterfacePairFunction(func(x interface{}) interface{} {
            return NewPair(x, x)
    }))
    xs2, isOk := xs.(InterfacePairFunction)
    if !isOk {
        t.Errorf("Unzip method first result type isn't InterfacePairFunction")
    } else {
        x := xs2(1)
        if !reflect.DeepEqual(x, 2) {
            t.Errorf("function result of Unzip method first result is %v; want %v", x, 2)
        }
    }
    ys2, isOk := ys.(InterfacePairFunction)
    if !isOk {
        t.Errorf("Unzip method second result type isn't InterfacePairFunction")
    } else {
        y := ys2(2)
        if !reflect.DeepEqual(y, 4) {
            t.Errorf("function result of Unzip method second result is %v; want %v", y, 4)
        }
    }
}
