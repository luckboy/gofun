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

func TestZipMethodZipsNoneAndNone(t *testing.T) {
    xs := None().Zip(None(), None())
    if !reflect.DeepEqual(xs, None()) {
        t.Errorf("Zip method result is %v; want %v", xs, None())
    }
}

func TestZipMethodZipsSomeAndNone(t *testing.T) {
    xs := Some(1).Zip(None(), None())
    if !reflect.DeepEqual(xs, None()) {
        t.Errorf("Zip method result is %v; want %v", xs, None())
    }
}

func TestZipMethodZipsNoneAndSome(t *testing.T) {
    xs := None().Zip(Some(2), None())
    if !reflect.DeepEqual(xs, None()) {
        t.Errorf("Zip method result is %v; want %v", xs, None())
    }
}

func TestZipMethodZipsSomeAndSome(t *testing.T) {
    xs := Some(1).Zip(Some(2), None())
    if !reflect.DeepEqual(xs, Some(NewPair(1, 2))) {
        t.Errorf("Zip method result is %v; want %v", xs, Some(NewPair(1, 2)))
    }
}

func TestZipMethodZipsLeftAndLeft(t *testing.T) {
    xs := Left("error1").Zip(Left("error2"), Left("error3"))
    if !reflect.DeepEqual(xs, Left("error1")) {
        t.Errorf("Zip method result is %v; want %v", xs, Left("error1"))
    }
}

func TestZipMethodZipsRightAndLeft(t *testing.T) {
    xs := Right(1).Zip(Left("error2"), Left("error3"))
    if !reflect.DeepEqual(xs, Left("error2")) {
        t.Errorf("Zip method result is %v; want %v", xs, Left("error2"))
    }
}

func TestZipMethodZipsLeftAndRight(t *testing.T) {
    xs := Left("error1").Zip(Right(2), Left("error3"))
    if !reflect.DeepEqual(xs, Left("error1")) {
        t.Errorf("Zip method result is %v; want %v", xs, Left("error1"))
    }
}

func TestZipMethodZipsRightAndRight(t *testing.T) {
    xs := Right(1).Zip(Right(2), Left("error3"))
    if !reflect.DeepEqual(xs,Right(NewPair(1, 2))) {
        t.Errorf("Zip method result is %v; want %v", xs, Right(NewPair(1, 2)))
    }
}

func TestZipMethodZipsEmptyListAndEmptyList(t *testing.T) {
    xs := Nil().Zip(Nil(), Nil())
    if !reflect.DeepEqual(xs, Nil()) {
        t.Errorf("Zip method result is %v; want %v", xs, Nil())
    }
}

func TestZipMethodZipsListAndEmptyList(t *testing.T) {
    xs := Cons(1, Cons(2, Cons(3, Nil()))).Zip(Nil(), Nil())
    if !reflect.DeepEqual(xs, Nil()) {
        t.Errorf("Zip method result is %v; want %v", xs, Nil())
    }
}

func TestZipMethodZipsEmptyListAndList(t *testing.T) {
    xs := Nil().Zip(Cons(4, Cons(5, Cons(6, Nil()))), Nil())
    if !reflect.DeepEqual(xs, Nil()) {
        t.Errorf("Zip method result is %v; want %v", xs, Nil())
    }
}

func TestZipMethodZipsListAndList(t *testing.T) {
    xs := Cons(1, Cons(2, Cons(3, Nil()))).Zip(Cons(4, Cons(5, Cons(6, Nil()))), Nil())
    if !reflect.DeepEqual(xs, Cons(NewPair(1, 4), Cons(NewPair(2, 5), Cons(NewPair(3, 6), Nil())))) {
        t.Errorf("Zip method result is %v; want %v", xs, Cons(NewPair(1, 4), Cons(NewPair(2, 5), Cons(NewPair(3, 6), Nil()))))
    }
}

func TestZipMethodZipsEmptyInterfaceSliceAndEmptyInterfaceSlice(t *testing.T) {
    xs := InterfaceSlice([] interface{} {}).Zip(InterfaceSlice([] interface{} {}), InterfaceSlice([] interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([] interface{} {})) {
        t.Errorf("Zip method result is %v; want %v", xs, InterfaceSlice([] interface{} {}))
    }
}

func TestZipMethodZipsInterfaceSliceAndEmptyInterfaceSlice(t *testing.T) {
    xs := InterfaceSlice([] interface{} { 1, 2, 3 }).Zip(InterfaceSlice([] interface{} {}), InterfaceSlice([] interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([] interface{} {})) {
        t.Errorf("Zip method result is %v; want %v", xs, InterfaceSlice([] interface{} {}))
    }
}

func TestZipMethodZipsEmptyInterfaceSliceAndInterfaceSlice(t *testing.T) {
    xs := InterfaceSlice([] interface{} {}).Zip(InterfaceSlice([] interface{} { 4, 5, 6 }), InterfaceSlice([] interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([] interface{} {})) {
        t.Errorf("Zip method result is %v; want %v", xs, InterfaceSlice([] interface{} {}))
    }
}

func TestZipMethodZipsInterfaceSliceAndInterfaceSlice(t *testing.T) {
    xs := InterfaceSlice([] interface{} { 1, 2, 3 }).Zip(InterfaceSlice([] interface{} { 4, 5, 6 }), InterfaceSlice([] interface{} {}))
    if !reflect.DeepEqual(xs, InterfaceSlice([] interface{} { NewPair(1, 4), NewPair(2, 5), NewPair(3, 6) })) {
        t.Errorf("Zip method result is %v; want %v", xs, InterfaceSlice([] interface{} { NewPair(1, 4), NewPair(2, 5), NewPair(3, 6) }))
    }
}

func TestZipMethodZipsInterfacePairFunctionAndInterfacePairFunction(t *testing.T) {
    xs := InterfacePairFunction(func(x interface{}) interface{} {
            return IntOrElse(x, 0) + 1
    }).Zip(InterfacePairFunction(func(x interface{}) interface{} {
            return IntOrElse(x, 0) + 2
    }), InterfacePairFunction(func(x interface{}) interface{} {
            return x
    }))
    ys, isOk := xs.(InterfacePairFunction)
    if !isOk {
        t.Errorf("Zip method result type isn't InterfacePairFunction")
    } else {
        y := ys(1)
        if !reflect.DeepEqual(y, NewPair(2, 3)) {
            t.Errorf("function result of Zip method result is %v; want %v", y, NewPair(2, 3))
        }
    }
}
