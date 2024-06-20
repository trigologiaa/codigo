// Package mapadebits proporciona una implementación de un mapa de bits sobre un entero de 32 bits.
// Expone la estructura MapaDeBits y sus métodos para manipular un mapa de bits.
package bitmap

import "errors"

const (
	// MapaDeBitsTamaño define el tamaño del mapa de bits.
	MapaDeBitsTamaño uint8 = 32
)

// MapaDeBits implementa un mapa de bits sobre un entero de 32 bits.
type MapaDeBits struct {
	bits uint32
}

// NuevoMapaDeBits crea un mapa de bits y lo inicializa con todos los bits apagados.
//
// Uso:
//
//	mapadebits := mapadebits.NuevoMapaDeBits() // Crea un nuevo mapa de bits.
func NuevoMapaDeBits() *MapaDeBits {
	return &MapaDeBits{bits: 0b0}
}

// Encender enciende el bit de la posición indicada.
//
// Uso:
//
//	mapadebits.Encender(5) // Enciende el bit en la posición 5.
//
// Parámetros:
//   - `posicion`: la posición del bit a encender.
//
// Retorna:
//   - `error` si la posición es inválida.
func (mapadebits *MapaDeBits) Encender(posicion uint8) error {
	if estaFueraDeRango(posicion) {
		return errors.New("posición no válida")
	}
	mapadebits.bits |= 0b1 << posicion
	return nil
}

// Apagar apaga el bit de la posición indicada.
//
// Uso:
//
//	mapadebits.Apagar(5) // Apaga el bit en la posición 5.
//
// Parámetros:
//   - `posicion`: la posición del bit a apagar.
//
// Retorna:
//   - `error` si la posición es inválida.
func (mapadebits *MapaDeBits) Apagar(posicion uint8) error {
	if estaFueraDeRango(posicion) {
		return errors.New("posición no válida")
	}
	mapadebits.bits &= ^(0b1 << posicion)
	return nil
}

// EstaEncendido obtiene el estado del bit en la posición indicada.
//
// Uso:
//
//	on, _ := mapadebits.EstaEncendido(5) // Verifica si el bit en la posición 5 está encendido.
//
// Parámetros:
//   - `posicion`: la posición del bit a verificar.
//
// Retorna:
//   - `true` si el bit está encendido; `false` si está apagado.
func (mapadebits *MapaDeBits) EstaEncendido(posicion uint8) (bool, error) {
	if estaFueraDeRango(posicion) {
		return false, errors.New("posición no válida")
	}
	return mapadebits.bits&(1<<posicion) != 0b0, nil
}

// ObtenerMapa obtiene la representación interna del mapa de bits
//
// Uso:
//
//	bits := mapadebits.ObtenerMapa() // Obtiene la representación interna del mapa de bits.
//
// Retorna:
//   - La representación interna del mapa de bits como un uint32.
func (mapadebits *MapaDeBits) ObtenerMapa() uint32 {
	return mapadebits.bits
}

// estaFueraDeRango verifica si la posición está fuera de rango.
//
// Uso:
//
//	if estaFueraDeRango(32) {
//		fmt.Println("Posición no válida")
//	}
//
// Retorna:
//   - `true` si la posición está fuera de rango; `false` en caso contrario.
func estaFueraDeRango(posicion uint8) bool {
	return posicion >= MapaDeBitsTamaño
}