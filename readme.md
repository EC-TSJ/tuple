# **Core**

[![Home](https://godoc.org/github.com/gookit/event?status.svg)](file:///D:/EC-TSJ/Documents/CODE/SOURCE/Go/pkg/lib/cli)
[![Build Status](https://travis-ci.org/gookit/event.svg?branch=master)](https://travis-ci.org/)
[![Coverage Status](https://coveralls.io/repos/github/gookit/event/badge.svg?branch=master)](https://coveralls.io/github/)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/event)](https://goreportcard.com/report/github.com/)

> **[EN README](README.md)**

Es una librería para gestionar tuples.

## GoDoc

- [godoc for github](https://godoc.org/github.com/)

## Funciones Principales

Los tipos principales ( y de carácter subyacente ) son:

- _**TupleV1[A]**_
- _**TupleV2[A, B]**_
- _**TupleV3[A, B, C]**_
- _**TupleV4[A, B, C, D]**_
- _**TupleV5[A, B, C, D, E]**_
- _**TupleV6[A, B, C; D, E, F]**_
- _**TupleV7[A, B, C, D, E, F, G]**_
- _**TupleV8[A, B, C, D, E, F, G, I]**_
- _**TupleV9[A, B, C, D, E, F, G, I, J]**_

Tipos Principales y de carácter genérico:

_**TupleConv[A, B]**_ con sus métodos:

- _`New(T, P) *TupleC[T, P]`_
- _`Convert() P`_

_**Tuple[T]**_, El tipo genérico con sus métodos:

- _`Get(pos int) any`_
- _`Count(count any) int`_
- _`Filter(fn func(...any) []any) []any`_
- _`ForEach(fn func(any))`_
- _`Index(x any) (int, string)`_
- _`Length() int`_
- _`Max() any`_
- _`Min() any`_
- _`ToSlice() []any`_
- _`UnPack() T`_
- _`Ptr() any`_

Y funciones:

- _`Get[T any](obj any, pos int) T`_
- _`Count[T any](obj any, x T) (count int)`_
- _`Index[T any](obj any, x T) (idx int, field string)`_
- _`Length(obj any) int`_
- _`Max[T any](obj any) (t T)`_
- _`Min[T any](obj any) (t T)`_
- _`ToSlice(obj any) (arr []any)`_
- _`Convert[TOrig any, TConv any](orig TOrig) (to TConv)`_
- _`Filter[T any](obj any, fn func(...any) []T) []T`_
- _`ForEach(obj any, fn func(any))`_
- _`UnPack[T any](p *Tuple[T]) T`_

## Ejemplo

```go
```

## Ejemplos

```go
func main() {
	a := (&Tuple[TupleV6[string, int, float64, string, int, int]]{}).New(TupleV6[string, int, float64, string, int, int]{"joder", 25, 25.64, "value", 62, 12})
	//a := &Tuple[TupleV6[string, int, float64, string, int, int]]{TupleV6[string, int, float64, string, int, int]{"joder", 25, 25.64, "value", 62, 12}}
	//flip := UnPack(a)
	zzz1 := Convert[TupleV3[int, string, float64], TupleV1[int]](TupleV3[int, string, float64]{25, "jeje", 56.12})
	zzz2 := Convert[TupleV3[int, string, float64], TupleV6[int, string, float64, int, float64, string]](TupleV3[int, string, float64]{25, "jeje", 56.12})
	zzz2.F = "caballo"
	zar := (&TupleConv[TupleV3[int, string, float64], TupleV1[int]]{}).New(TupleV3[int, string, float64]{13, "XO", 61.10}, TupleV1[int]{})
	//zar := TupleV[TupleV3[int, string, float64], TupleV1[int]]{TupleV3[int, string, float64]{13, "XO", 61.10}, TupleV1[int]{}}
	zzzf := zar.Convert()
	zare := TupleConv[TupleV3[int, string, float64], TupleV7[int, string, float64, string, int, int, float64]]{TupleV3[int, string, float64]{13, "XO", 61.10}, TupleV7[int, string, float64, string, int, int, float64]{}}
	zzzg := zare.Convert()
	zflip := a.UnPack()
	fmt.Println(zflip.C, zflip.D, zzz1, zzz2, zzzf, zzzg)
	fmt.Println(a.ToSlice())
	value := a.Get(2)
	zz1 := a.Length()
	fmt.Println(value)
	blb := a.Count(25.64)
	println(blb, zz1)
	b := &Tuple[TupleV3[int, string, int]]{TupleV3[int, string, int]{15, "tocomocho", 61}, nil}
	zz := b.Length()
	zx := b.Max()
	zz10 := Max[string](b.Ptr())
	zd := b.Min()
	zz11 := Min[int](b.Ptr())
	z := a.Length()
	z1 := a.Count(25)
	z2, z3 := a.Index(25.64)
	xfg := reflect.ValueOf(z2).Kind()
	z2, z3 = a.Index("value")
	zal := a.Max()
	zas := a.Min()
	fmt.Println(zz10, zz11, xfg, a, zas, zal, z2, z3, z1, z, zz, zx, zd)
}
```

## Notas

<!-- - [gookit/ini](https://github.com/gookit/ini) INI配置读取管理，支持多文件加载，数据覆盖合并, 解析ENV变量, 解析变量引用
-->

## LICENSE

**[MIT](LICENSE)**
