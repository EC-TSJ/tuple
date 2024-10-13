package tuple

import (
	"fmt"
	"reflect"
	"std/container/internal/core"
	"strings"
)

type subyacent = comparable

// //////////////////////////////////////////////////////////////////////////////
// tupleV1 is a group of 1 elements. Subyacente
// //////////////////////////////////////////////////////////////////////////////
type TupleV1[A subyacent] struct {
	/** Subyacent */
	A A
}

// //////////////////////////////////////////////////////////////////////////////
// tupleV2 is a group of 2 elements. Subyacente
// //////////////////////////////////////////////////////////////////////////////
type TupleV2[A, B subyacent] struct {
	/** Subyacent */
	A A
	B B
}

// //////////////////////////////////////////////////////////////////////////////
// tupleV3 is a group of 3 elements. Subyacente
// //////////////////////////////////////////////////////////////////////////////
type TupleV3[A, B, C subyacent] struct {
	/** Subyacent */
	A A
	B B
	C C
}

// //////////////////////////////////////////////////////////////////////////////
// tupleV4 is a group of 4 elements. Subyacente
// //////////////////////////////////////////////////////////////////////////////
type TupleV4[A, B, C, D subyacent] struct {
	/** Subyacent */
	A A
	B B
	C C
	D D
}

// //////////////////////////////////////////////////////////////////////////////
// tupleV5 is a group of 5 elements. Subyacente
// //////////////////////////////////////////////////////////////////////////////
type TupleV5[A, B, C, D, E subyacent] struct {
	/** Subyacent */
	A A
	B B
	C C
	D D
	E E
}

// //////////////////////////////////////////////////////////////////////////////
// tupleV6 is a group of 6 elements. Subyacente
// //////////////////////////////////////////////////////////////////////////////
type TupleV6[A, B, C, D, E, F subyacent] struct {
	/** Subyacent */
	A A
	B B
	C C
	D D
	E E
	F F
}

// //////////////////////////////////////////////////////////////////////////////
// tupleV7 is a group of 7 elements. Subyacente
// //////////////////////////////////////////////////////////////////////////////
type TupleV7[A, B, C, D, E, F, G subyacent] struct {
	/** Subyacent */
	A A
	B B
	C C
	D D
	E E
	F F
	G G
}

// //////////////////////////////////////////////////////////////////////////////
// tupleV8 is a group of 8 elements. Subyacente
// //////////////////////////////////////////////////////////////////////////////
type TupleV8[A, B, C, D, E, F, G, H subyacent] struct {
	/** Subyacent */
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
}

// //////////////////////////////////////////////////////////////////////////////
// tupleV9 is a group of 9 elements. Subyacente
// //////////////////////////////////////////////////////////////////////////////
type TupleV9[A, B, C, D, E, F, G, H, I subyacent] struct {
	/** Subyacent */
	A A
	B B
	C C
	D D
	E E
	F F
	G G
	H H
	I I
}

/*******************************************************************************
 * Clase base en sistema tuple. Es para usar las tuples como tal
 *******************************************************************************
 *! @type {T}
 *? @public @class yupl[T]
 */
type tupl[T comparable] struct {
	Class core.ClassType
	_any_ T // tuple subyacente
}

/*******************************************************************************
 * Clase base en sistema tuple. Es para copiar (convertir) tuples. Extends tupl[T]
 *******************************************************************************
 *! @type {T} es el tipo
 *! @type {P} es el parametro
 *? @public @class TuplK[T, K] @extends tupl[T]
 */
type TuplK[T, K comparable] struct {
	/**
	 *! @extends
	 */*tupl[T]   //_any_ T // mantiene el tuple subyacente
	_to_        K // el tuple a convertir
}

// //////////////////////////////////////////////////////////////////////////////
type Tuple = tupl[any]
type T1 = TupleV1[any]
type T2 = TupleV2[any, any]
type T3 = TupleV3[any, any, any]
type T4 = TupleV4[any, any, any, any]
type T5 = TupleV5[any, any, any, any, any]
type T6 = TupleV6[any, any, any, any, any, any]
type T7 = TupleV7[any, any, any, any, any, any, any]
type T8 = TupleV8[any, any, any, any, any, any, any, any]
type T9 = TupleV9[any, any, any, any, any, any, any, any, any]

/////////////////////////////////////////////////////////////////////////////////

/** Obtención del numero de veces que aparece count. Base cero */
func (t *tupl[T]) Count(count any) int { return Count(*(&t._any_), count) }

/** Filtra los datos del array en base a una condición. Base cero */
func (t *tupl[T]) Filter(fn func(...any) []any) []any { return filter(*&t._any_, fn) }

/** Realiza la función forEach para cada elemento del tuple. Base cero */
func (t *tupl[T]) ForEach(fn func(any)) { forEach(*(&t._any_), fn) }

/** Obtención de un valor determinado. Base cero */
func (t *tupl[T]) Get(pos int) any { return get(*(&t._any_), pos) }

/** Retorna el tuple */
func (t *tupl[T]) GetTuple() T { return t._any_ }

/** Devuelve el primer índice en el que aparece x en el tuple, o -1 si no está en el tupl. Base cero */
func (t *tupl[T]) Index(x any) (int, string) { return Index(*(&t._any_), x) }

/** Nos dice el número de elementos en el tuple. Base cero */
func (t *tupl[T]) Length() int { return Length(*(&t._any_)) }

/** Nos dice el máximo valor del tuple. Base cero */
func (t *tupl[T]) Max() any { return fn[0](*(&t._any_)) }

/** Nos dice el mínimo valor del tuple. Base cero */
func (t *tupl[T]) Min() any { return fn[1](*(&t._any_)) }

/**
 * Crea y empaqueta el tuple
 *? @static @constructor @New
 */
func (*tupl[T]) New(t T) *tupl[T] {
	return &tupl[T]{Class: core.Class.Normal, _any_: t}
}

// /**
//  * Suma dos tuples
//  */
// func (t *tupl[T]) AddTuple(to *Tuple) {
// }

/** Obtiene el puntero al tuple subyacente */
func (t *tupl[T]) Ptr() any { return t._any_ }

/** Convierte el tuple en un Slice. Base cero */
func (t *tupl[T]) ToSlice() []any { return ToSlice(*(&t._any_)) }

/** Obtiene los datos desempaquetados */
func (t *tupl[T]) UnPack() T { return UnPack(t) }

/** Convierte unas tuple a una menor, o mas grande */
func (t *tupl[T]) Convert(to any) any {
	t._any_ = convert(*(&t._any_), to).(T)
	return t._any_
}

/** Convierte unas tuple a una menor, o mas grande */
func (t *TuplK[T, K]) Convert() K {
	return Convert[T, K](*(&t._any_))
}

func (*TuplK[T, K]) New(t T, k K) *TuplK[T, K] {
	return &TuplK[T, K]{tupl: &tupl[T]{core.Class.Normal, t}, _to_: k}
}

/**
 * El tipo en las funciones es el del elemento a devolver, etc. El del
 * elemento en la tuple. Excepto en Convert que deberá ser los tipos que
 * consideran a convertir.
 */

//!+
/*******************************************************************************
 * Nos da un valor determinado. Base cero
 * Debe llamarse desreferenciando el puntero de llamada
 *! @param {any}
 *! @param {int}
 *! @returns {T}
 */
func get(obj any, pos int) any {
	if reflect.ValueOf(obj).NumField() < 0 || pos > reflect.ValueOf(obj).NumField() {
		return any(nil)
	}
	return reflect.ValueOf(obj).Field(pos).Interface()
}

/*******************************************************************************
 * Nos da un valor determinado. Base cero
 * Debe llamarse desreferenciando el puntero de llamada
 *! @param {any}
 *! @param {int}
 *! @returns {T}
 */
func Get[T comparable](obj any, pos int) T {
	return get(obj, pos).(T)
}

//!-
// /**
//  * Nos da los datos enmarcados en su tipo
//  *! @param {any}
//  *! @returns {T}
//  */
// func Typed[T comparable](value any) T {
// 	return value.(T)
// }

/*******************************************************************************
 * Devuelve el número de veces que x aparece en el tuple
 * Debe llamarse desreferenciando el puntero de llamada
 *? <CODE>
 *? ...
 *?   t.Count(*ptr, 25)
 *? ...
 *? </CODE>
 *! @param {any}
 *! @param {T}
 *! @returns {int}
 */
func Count[T comparable](obj any, x T) (count int) {
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		if fmt.Sprintf("%v", v.Field(i)) == fmt.Sprintf("%v", x) {
			count++
		}
	}
	return count
}

/*******************************************************************************
 * Devuelve el primer índice en el que aparece x en el tupl, o -1 si no está en el tuple,
 * Debe llamarse desreferenciando el puntero de llamada,
 *? <CODE>
 *? ...
 *?   t.Index(*ptr, "any")
 *? ...
 *? </CODE>
 *! @param {any}
 *! @param {T}
 *! @returns {int} indice
 *! @return {string} nonbre del campo
 */
func Index[T comparable](obj any, x T) (idx int, field string) {
	idx = -1
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		if fmt.Sprintf("%v", v.Field(i)) == fmt.Sprintf("%v", x) {
			idx++
			return idx, v.Type().Field(i).Name
		} else {
			idx++
		}
	}
	return idx, core.Lit.NullString
}

/*******************************************************************************
 * Nos dice el número de elementos en el tuple
 * Debe llamarse desreferenciando el puntero de llamada
 *? <CODE>
 *? ...
 *?   t.Length(*ptr)
 *? ...
 *? </CODE>
 *! @param {any}
 *! @returns {int}
 */
func Length(obj any) int {
	return reflect.ValueOf(obj).NumField()
}

// !+ __switch
var fn []func(any) any = []func(any) any{
	max,
	min,
}

/*******************************************************************************
 */
func __switch[T comparable](obj any, fn func(any) any) (t T) {
	switch v := fn(obj).(type) {
	case
		string,
		bool,
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float64, float32,
		complex64, complex128:
		t = (v).(T)
	}
	return t
}

//!+
/*******************************************************************************
 * Nos dice el máximo valor del tuple
 * Debe llamarse desreferenciando el puntero de llamada
 *? <CODE>
 *? ...
 *?   t.Max(*ptr)
 *? ...
 *? </CODE>
 *! @param {any}
 *! @returns {T}
 */
func max(obj any) (gh any) {
	gh = strings.Repeat(" ", 7)
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i).Interface()
		t := fmt.Sprintf("%v", f) // t es string, int, bool, etc
		if t == "{}" || t == "{<nil>}" {
			continue
		}
		if t > fmt.Sprintf("%v", gh) {
			gh = f
		}
	}
	return gh
}

/*******************************************************************************
 * - Nos dice el máximo valor del tuple
 * - Debe llamarse desreferenciando el puntero de llamada
 *? <CODE>
 *? ...
 *?   t.Max(*ptr)
 *? ...
 *? </CODE>
 *! @param {any}
 *! @returns {T}
 */
func Max[T comparable](obj any) (t T) {
	return __switch[T](obj, fn[0])
}

//!-
//!+
/*******************************************************************************
 * Nos dice el mínimo valor del tuple
 * Debe llamarse desreferenciando el puntero de llamada
 *? <CODE>
 *? ...
 *?   t.Min(*ptr)
 *? ...
 *? </CODE>
 *! @param {any}
 *! @returns {T}
 */
func min(obj any) (gh any) {
	gh = strings.Repeat("~", 7)
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i).Interface()
		t := fmt.Sprintf("%v", f) // t es string, int, bool, etc
		if t == "{}" || t == "{<nil>}" {
			continue
		}
		if t < fmt.Sprintf("%v", gh) {
			gh = f
		}
	}
	return gh
}

/*******************************************************************************
 * Nos dice el mínimo valor del tuple
 * Debe llamarse desreferenciando el puntero de llamada
 *? <CODE>
 *? ...
 *?   t.Min(*ptr)
 *? ...
 *? </CODE>
 *! @param {any}
 *! @returns {T}
 */
func Min[T comparable](obj any) (t T) {
	return __switch[T](obj, fn[1])
}

//!-
//!-    __switch
/*******************************************************************************
 * Nos da un slice del tuple
 * Debe llamarse desreferenciando el puntero de llamada
 *! @param{any}
 *! @return{[]interface}
 */
func ToSlice(obj any) (arr []any) {
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		arr = append(arr, v.Field(i).Interface()) //.(T)
	}
	return arr
}

// /**
//  * Nos da un slice del tuple
//  * Debe llamarse desreferenciando el puntero de llamada
//  *! @param{any}
//  *! @return{[]interface}
//  */
// func ToSlice[T comparable](obj any) (arr []T) {
// 	v := reflect.ValueOf(obj)
// 	for i := 0; i < v.NumField(); i++ {
// 		arr = append(arr, v.Field(i).Interface().(T))
// 	}
// 	return arr
// }

/*******************************************************************************
 * Convierte a otra tuple, más grande o pequeña.
 *? <CODE>
 *?   ...
 *?		convert(TupleV2[string, int]{"pp", 123}, TupleV3[string, int, float64]{"qq", 51, 6.46})
 *?   ...
 *? </CODE>
 *! @param {any} la tupla original
 *! @param {any} la tupla convertida
 *! @return {any}
 */
func convert(orig any, to any) any {
	var (
		max int
	)
	vorig := reflect.ValueOf(orig)
	_to_ := reflect.ValueOf(to)
	if vorig.Kind() == reflect.Ptr {
		vorig = vorig.Elem()
	}
	// saca el mas pequeño
	if vorig.NumField() > _to_.NumField() {
		max = _to_.NumField() // es una conversión a otra inferior
	} else { // saca el mayor
		max = vorig.NumField() // es una conversión a otra superior
	}
	// lo corre
	for i := 0; i < max; i++ {
		// vv is the interface{}
		vv := reflect.ValueOf(&to).Elem()
		// Allocate a temporary variable with type of the struct.
		//    vv.Elem() is the vale contained in the interface.
		tmp := reflect.New(vv.Elem().Type()).Elem()
		// Copy the struct value contained in interface to
		// the temporary variable.
		tmp.Set(vv.Elem())
		// Set the field.
		tmp.Field(i).Set(_to_.Field(i))
		// Set the interface to the modified struct value.
		vv.Set(tmp)
	}
	return to
}

/*******************************************************************************
 * Convierte a otra tuple, más grande o pequeña.
 *? <CODE>
 *?   ...
 *?		Convert[tupleV3[int, string, float64], tupleV1[int]](v3)
 *?   ...
 *? </CODE>
 *! @type {TOrig} la tupla original, instanciada.
 *! @type {TConv} la tupla convertida.
 *! @param {TOrig} la tupla original
 *! @return {TConv}
 */
func Convert[TOrig, TConv comparable](orig TOrig) (to TConv) {
	var (
		max int
	)
	v := reflect.ValueOf(orig)
	// saca el mas pequeño
	if v.NumField() > reflect.ValueOf(to).NumField() {
		max = reflect.ValueOf(to).NumField() // es una conversión a otra inferior
	} else { // saca el mayor
		max = v.NumField() // es una conversión a otra superior
	}
	// lo corre
	for i := 0; i < max; i++ {
		reflect.ValueOf(&to).Elem().Field(i).Set(v.Field(i))
	}
	return to
}

//!+
/*******************************************************************************
 * Filtra los datos del array en base a una condición
 * Debe llamarse desreferenciando el puntero de llamada
 *! @param {func(...any) []T}
 *! @return {[]T}
 */
func Filter[T comparable](obj any, fn func(...any) []any) (arr []any) {
	arr = ToSlice(obj)
	return fn(arr...)
}

/*******************************************************************************
 * Filtra los datos del array en base a una condición
 * Debe llamarse desreferenciando el puntero de llamada
 *! @param {func(...any) []T}
 *! @return {[]T}
 */
func filter(obj any, fn func(...any) []any) (ret []any) {
	ret = ToSlice(obj)
	return fn(ret...)
}

//!-
/*******************************************************************************
 * Realiza la función forEach para cada elemento del tuple
 * Debe llamarse desreferenciando el puntero de llamada
 *! @param {func(any)}
 */
func ForEach[T comparable](obj any, fn func(T)) {
	t := ToSlice(obj)
	for _, v := range t {
		fn(v.(T))
	}
}

/*******************************************************************************
 * Realiza la función forEach para cada elemento del tuple
 * Debe llamarse desreferenciando el puntero de llamada
 *! @param {func(any)}
 */
func forEach(obj any, fn func(any)) {
	t := ToSlice(obj)
	for _, v := range t {
		fn(v)
	}
}

/*******************************************************************************
 * Desempaqueta los contenidos de una tuple
 *?  <CODE>
 *?    ...
 *?    flip := UnPack[tupleV6[string, int, float64, string, int, int]](a)
 *?    flip.D
 *?    flip.B
 *?    ...
 *?  </CODE>
 */
func UnPack[T comparable](p *tupl[T]) T {
	//return any(*(&p._any_)).(T)
	return *(&p._any_)
}

// func main() {
// 	type T2 = TupleV2[string, int]
// 	vz := Tuple{TupleV6[string, int, int, float64, string, float64]{"jesus", 61, 71, 23.32, "pelo", 62.32}}
// 	fdf := vz.Min()
// 	faz := vz.Count(23.32)
// 	vs := ToSlice(TupleV6[string, int, int, float64, string, float64]{"jesus", 61, 71, 23.32, "pelo", 62.32})[2]
// 	f0 := Tuple{T2{"pepe", 51}}
// 	//f0.AddTuple(&Tuple{TupleV3[string, int, int]{"jesus", 13, 42}})
// 	fz := Tuple{TupleV2[string, int]{"polla", 51}}
// 	zsx := Convert[TupleV2[string, int], TupleV3[string, int, float64]](TupleV2[string, int]{"pepe", 42})
// 	zsd := fz.Convert(TupleV3[string, int, float64]{"pepe", 82, 0.00})
// 	fmt.Println(fz, zsd, zsx, vs, vz, fdf, faz)
// 	zzz := f0.GetTuple()
// 	vol := fz.GetTuple()
// 	vola := vol.(TupleV3[string, int, float64])
// 	zs := fz.GetTuple().(TupleV3[string, int, float64])
// 	zi := fz.UnPack().(TupleV3[string, int, float64])

// 	f := TupleV2[string, int]{"polla", 51}
// 	z := &tupl[TupleV2[string, int]]{f}
// 	zep := &tupl[TupleV2[string, int]]{TupleV2[string, int]{"polla", 51}}
// 	a := (&tupl[TupleV2[string, int]]{}).New(f)
// 	zz := UnPack[TupleV2[string, int]](z)
// 	fmt.Println(z, a, zz, f0, fz, zi, zep, zs, vol, vola, zzz)
// }
