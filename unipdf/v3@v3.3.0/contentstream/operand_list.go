/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package contentstream

// isValidOperand returns true if `operand` is a valid PDF contentstream operand.
func isValidOperand(operand string) bool {
	_, has := validOperands[operand]
	return has
}

// List of valid PDF contentstream operands in a map format for quick lookups.
// (Table A.1 PDF32000_2008).
var validOperands = map[string]struct{}{
	"b":   struct{}{},
	"B":   struct{}{},
	"b*":  struct{}{},
	"B*":  struct{}{},
	"BDC": struct{}{},
	"BI":  struct{}{},
	"BMC": struct{}{},
	"BT":  struct{}{},
	"BX":  struct{}{},
	"c":   struct{}{},
	"cm":  struct{}{},
	"CS":  struct{}{},
	"cs":  struct{}{},
	"d":   struct{}{},
	"d0":  struct{}{},
	"d1":  struct{}{},
	"Do":  struct{}{},
	"DP":  struct{}{},
	"EI":  struct{}{},
	"EMC": struct{}{},
	"ET":  struct{}{},
	"EX":  struct{}{},
	"f":   struct{}{},
	"F":   struct{}{},
	"f*":  struct{}{},
	"G":   struct{}{},
	"g":   struct{}{},
	"gs":  struct{}{},
	"h":   struct{}{},
	"i":   struct{}{},
	"ID":  struct{}{},
	"j":   struct{}{},
	"J":   struct{}{},
	"K":   struct{}{},
	"k":   struct{}{},
	"l":   struct{}{},
	"m":   struct{}{},
	"M":   struct{}{},
	"MP":  struct{}{},
	"n":   struct{}{},
	"q":   struct{}{},
	"Q":   struct{}{},
	"re":  struct{}{},
	"RG":  struct{}{},
	"rg":  struct{}{},
	"ri":  struct{}{},
	"s":   struct{}{},
	"S":   struct{}{},
	"SC":  struct{}{},
	"sc":  struct{}{},
	"SCN": struct{}{},
	"scn": struct{}{},
	"sh":  struct{}{},
	"T*":  struct{}{},
	"Tc":  struct{}{},
	"Td":  struct{}{},
	"TD":  struct{}{},
	"Tf":  struct{}{},
	"Tj":  struct{}{},
	"TJ":  struct{}{},
	"TL":  struct{}{},
	"Tm":  struct{}{},
	"Tr":  struct{}{},
	"Ts":  struct{}{},
	"Tw":  struct{}{},
	"Tz":  struct{}{},
	"v":   struct{}{},
	"w":   struct{}{},
	"W":   struct{}{},
	"W*":  struct{}{},
	"y":   struct{}{},
	`'`:   struct{}{},
	`''`:  struct{}{},
}
