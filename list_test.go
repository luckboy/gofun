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

func TestListConcatMethodConcatenatesEmptyListAndEmptyList(t *testing.T) {
    xs := Nil().Concat(Nil())
        if !reflect.DeepEqual(xs, Nil()) {
        t.Errorf("List.Concat method first result is %v; want %v", xs, Nil())
    }
}

func TestListConcatMethodConcatenatesListAndEmptyList(t *testing.T) {
    xs := Cons(1, Cons(2, Nil())).Concat(Nil())
        if !reflect.DeepEqual(xs, Cons(1, Cons(2, Nil()))) {
        t.Errorf("List.Concat method first result is %v; want %v", xs, Cons(1, Cons(2, Nil())))
    }
}

func TestListConcatMethodConcatenatesEmptyListAndList(t *testing.T) {
    xs := Nil().Concat(Cons(3, Cons(4, Nil())))
        if !reflect.DeepEqual(xs, Cons(3, Cons(4, Nil()))) {
        t.Errorf("List.Concat method first result is %v; want %v", xs, Cons(3, Cons(4, Nil())))
    }
}

func TestListConcatMethodConcatenatesListAndList(t *testing.T) {
    xs := Cons(1, Cons(2, Nil())).Concat(Cons(3, Cons(4, Nil())))
        if !reflect.DeepEqual(xs, Cons(1, Cons(2, Cons(3, Cons(4, Nil()))))) {
        t.Errorf("List.Concat method first result is %v; want %v", xs, Cons(1, Cons(2, Cons(3, Cons(4, Nil())))))
    }
}
