package tuple

import (
	"ec-tsj/core"
	"ec-tsj/event"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type (
	// Dato
	T = core.T
	// Tuple
	Element struct {
		D T
	}
	// Slice de Tuples
	Tuple []Element

	// interface
	ITuple interface {
		Insert(int, ...T)
		Add(...T)
		Remove(int) T
		Clone() ITuple
		ToSlice(...int) []Element
		Set(int, T)
		Get(int) T
		String() string
		Reverse()
		Pop(pop) T
		Count(T) int
		Index(T) int
		Sort(order)
		IsEmpty() bool
		Size() int
		Max() T
		Min() T
	}
)

func init() {
	Events = event.EventEmitter("Element", false)
}

var (
	Events event.IEventEmitter
)

func (t Element) String() string {
	return fmt.Sprintf("%v", t.D)
}

// Crea un Objeto Tuple
// @param {...T} t
// return {*Tuple}
func NewTuple(items ...T) ITuple {
	st := &Tuple{}
	s := len(items)

	if s != 0 {
		st.Add(items...)
	}

	return st
}

// Añade uno o más elementos al Tuple
// Puede ser un Tuple, slice de interfaces ó elementos sueltos
// @param {...T}
func (this *Tuple) Add(items ...T) {
	if v, ok := items[0].([]T); ok {
		for _, f := range v {
			*this = append(*this, Element{f})
			Events.Emit("Add", f)
		}
	} else if v, ok := items[0].([]Element); ok {
		for _, f := range v {
			*this = append(*this, f)
			Events.Emit("Add", f)
		}
	} else {
		for _, v := range items {
			*this = append(*this, Element{v})
			Events.Emit("Add", v)
		}
	}
}

// !+

func (this *Tuple) _parms(i int) int {
	if i < 0 {
		i = 0
	}
	if i >= this.Size() {
		i = this.Size()
	}

	return i
}

// Inserta uno o más elementos al Tuple, empezando en un elemento dado
// Puede ser un Tuple, slice de interfaces ó elementos sueltos
// @param {int}
// @param {...T} t
func (this *Tuple) Insert(start int, items ...T) {
	euler := *this
	start = this._parms(start)
	rhs := euler[start:]
	if v, ok := items[0].([]T); ok {
		for _, f := range v {
			euler = append(euler, Element{f})
			Events.Emit("Add", f)
		}
	} else if v, ok := items[0].([]Element); ok {
		for _, f := range v {
			euler = append(euler, f)
			Events.Emit("Add", f)
		}
	} else {
		for _, v := range items {
			euler = append(euler, Element{v})
			Events.Emit("Add", v)
		}
	}
	euler = append(euler, rhs...)
	*this = euler
}

// Borra un elemento dado del Tuple
// @param {int} t
// return {T}
func (this *Tuple) Remove(item int) T {
	euler := *this
	item = this._parms(item)
	aser := euler[item]
	euler = append(euler[:item], euler[item+1:]...)
	*this = euler
	Events.Emit("Remove", item)

	return aser.D
}

//!-

// Hace una copia del Tuple subyacente
// return {*Tuple}
func (this *Tuple) Clone() ITuple {
	t := NewTuple()
	v := t.(*Tuple)
	*v = append(*v, *this...)

	return v
}

// Retorna un slice de los datos del Tuple. Base 0.
// @param {en blanco | int,int} first, last ó en blanco
// return {[]tuple}
func (this *Tuple) ToSlice(parms ...int) []Element {
	copy := this.Clone()
	fcopy := copy.(*Tuple)

	parms = append(parms, -1, -1)
	parms[0] = map[bool]int{true: parms[0], false: 0}[parms[0] != -1]
	parms[1] = map[bool]int{true: parms[1], false: len(*fcopy)}[parms[1] != -1]

	euler := *fcopy
	euler = euler[parms[0]:parms[1]]

	return euler
}

// Pone el dato n en el Tuple. Base 0
// @param {int}
// @param {T} t
func (this Tuple) Set(idx int, item T) {
	if idx < 0 {
		idx = 0
	} else if idx >= this.Size() {
		this = append(this, Element{item}) // Añade otro item
		return
	}
	this[idx] = Element{item} // modifica un item
}

// Obtiene el dato n del Tuple. Base 0
// @param {int} t
// return {T}
func (this Tuple) Get(idx int) T {
	if idx < 0 {
		idx = 0
	} else if idx >= this.Size() {
		idx = this.Size() - 1 // idx es el máximo
	}
	return this[idx].D
}

// Retorna representación en cadena del Tuple
// return {string}
func (this *Tuple) String() string {
	items := make([]string, 0, len(*this))
	for _, val := range *this {
		if fmt.Sprintf("%T", val) == "tuple.Element" {
			items = append(items, fmt.Sprintf("%v", val.D))
		} else {
			items = append(items, fmt.Sprintf("%v", val))
		}
	}

	return fmt.Sprintf("(%s)", strings.Join(items, ", "))
}

// Reversa el Tuple subyacente
func (this Tuple) Reverse() {
	for i, j := 0, this.Size()-1; i < j; i, j = i+1, j-1 {
		this[i], this[j] = this[j], this[i]
	}
}

// Pop el item más a la izquierda/derecha del Tuple y lo retorna
// @param {pop} t
// return {T}
func (this *Tuple) Pop(enum pop) (ret T) {
	euler := *this
	if enum == POP_LEFT {
		ret = euler[0]
		*this = euler[1:] // recorta el primero
	} else if enum == POP_RIGHT {
		ret = euler[this.Size()-1]
		*this = euler[:this.Size()-1] // recorta el ultimo
	}

	return
}

// Cuenta el número de veces que aparece un objeto en el Tuple
// Típicode Python
// @param {T} t
// return {int}
func (this *Tuple) Count(value T) (count int) {
	for _, m := range *this {
		if m == value {
			count++
		}
	}

	return
}

// Nos dice el sitio, más bajo, en que aparece. Típico de Python
// @param {T} t
// return {int}
func (this *Tuple) Index(value T) int {
	var index int = -1

	for n, v := range *this {
		if v == value {
			index = n
			break
		}
	}

	return index
}

// IsEmpty retorna true si el Tuple está vacío
// return {bool}
func (this *Tuple) IsEmpty() bool {
	return this.Size() == 0
}

// Size retorna el tamaño del Tuple
// return {int}
func (this *Tuple) Size() int {
	return len(*this)
}

// Devuelve el elemento máximo de un tuplePtr
// @param {Tuple}
// return {T}
func (this *Tuple) Max() (max T) {
	var maxA, maxB, maxC = 0, nullString, 0.0

	for _, m := range *this {
		switch g := m.D.(type) {
		case string:
			if g > maxB {
				max, maxB = g, g
				maxA, _ = strconv.Atoi(g)
				maxC, _ = strconv.ParseFloat(g, 64)
			}
		case int:
			if g > maxA {
				max, maxA = g, g
				maxB = strconv.Itoa(g)
				maxC = float64(g)
			}
		case float64:
			if g > maxC {
				max, maxC = g, g
				maxA = int(g)
				maxB = fmt.Sprintf("%f", g)
			}
		}
	}

	return
}

// Devuelve el elemento mínimo de un tuplePtr
// @param {Tuple} t
// return {T}
func (this *Tuple) Min() (min T) {
	var minA, minB, minC = math.MaxInt64, "�", math.Inf(1)

	for _, m := range *this {
		switch g := m.D.(type) {
		case string:
			if g < minB {
				min, minB = g, g
				minA, _ = strconv.Atoi(g)
				minC, _ = strconv.ParseFloat(g, 64)
			}
		case int:
			if g < minA {
				min, minA = g, g
				minB = strconv.Itoa(g)
				minC = float64(g)
			}
		case float64:
			if g < minC {
				min, minC = g, g
				minA = int(g)
				minB = fmt.Sprintf("%f", g)
			}
		}
	}

	return
}

// Clasifica los elementos por orden
// @param {order} t
func (this *Tuple) Sort(enum order) {
	__enum = enum
	sort.Sort(data(*this))
}
