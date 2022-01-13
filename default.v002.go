/************************************/
/*  (%T% %S%), %J% <$B$> <$1.00$>   */
/*  (%W% 30-09-1991 )               */
/*  (%X%            )               */
/*  (%M%            )               */
/*  <$  $>                          */
/************************************/
package tuple

/*
type MinMax interface {
	Len(Item) int
	Min(MinMax) Item
	Max(MinMax) Item
}
*/

var nullString string

// Devuelve el número de elementos de un tuple
//
// @param {STuple} t
// return {int}
func Len(t Tuple) int {
	return t.Size()
}

// Devuelve el elemento máximo de un tuplePtr
//
// @param {STuple}
// return {Item}
func Max(t Tuple) (max T) {
	return t.Max()
}

// Devuelve el elemento mínimo de un tuplePtr
//
// @param {STuple} t
// return {Item}
func Min(t Tuple) (min T) {
	return t.Min()
}
