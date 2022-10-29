package strmech

import "sync"

//	International System of Units (SI)
//
// 	SI is also referred to as a Metric prefix.
//
//	A metric prefix is a unit prefix that precedes a
//	basic unit of measure to indicate a multiple or
//	submultiple of the unit.
//
//	All metric prefixes used today are decadic (of or
//	relating to the decimal system of counting). Each
//	prefix has a unique symbol that is prepended to
//	any unit symbol.
//
//	The prefix kilo-, for example, may be added to gram to
//	indicate multiplication by one thousand: one kilogram
//	is equal to one thousand grams. The prefix milli-,
//	likewise, may be added to metre to indicate division
//	by one thousand; one millimetre is equal to one
//	thousandth of a metre.
//
//	A metric power is an integer unit affix, written in
//	superscript in formal typography, that follows the
//	basic unit of measure to indicate a multiplicity of
//	the basic unit. In electronic plain text where
//	superscript is not available, the subscript is often
//	omitted, or where confusion is possible, indicated by
//	placing the caret symbol ^ between the base unit and
//	the integer power, thus km2, km2, and km^2 are
//	variously encountered. When no integer affix is
//	supplied, the implied power is 1. When a unit is not
//	mentioned at all, the implied power is 0. Negative
//	powers imply division. With extreme formality, the
//	unit m/s2 can also be rendered m1s-2, but the literal
//	present of the implied integer 1 is considered
//	unconventional in common usage. Often all the units
//	with positive prefixes will be listed first (in some
//	natural order), followed by all the units with
//	negative prefixes (in some natural order); this
//	semi-canonical form is most easily mapped by the mind
//	onto division notation, and makes switching between
//	the two conventions less mentally onerous.
//
//						SI prefixes
//					Prefix	Representations
//
//						Base 	 Base
//		Name	Symbol	1000	  10		Value
//		----	------  -----	-----	-------------
//		yotta	  Y		1000^8	 10^24	1000000000000000000000000
//		zetta	  Z		1000^7	 10^21	1000000000000000000000
//		exa		  E		1000^6	 10^18	1000000000000000000
//		peta 	  P		1000^5	 10^15	1000000000000000
//		tera	  T		1000^4	 10^12	1000000000000
//		giga	  G		1000^3	 10^9	1000000000
//		mega	  M		1000^2	 10^6	1000000
//		kilo	  k		1000^1	 10^3	1000
//						1000^0	 10^0	1
//		milli	  m		1000^−1	 10^−3	0.001
//		micro	  μ		1000^−2	 10^−6	0.000001
//		nano	  n		1000^−3	 10^−9	0.000000001
//		pico	  p		1000^−4	 10^−12	0.000000000001
//		femto	  f		1000^−5	 10^−15	0.000000000000001
//		atto	  a		1000^−6	 10^−18	0.000000000000000001
//		zepto	  z		1000^−7	 10^−21	0.000000000000000000001
//		yocto	  y		1000^−8	 10^−24	0.000000000000000000000001
//
//  Multiplier Notation
//
//     k kilo	1 x 10^3	1,000
//     M mega	1 x 10^6	1,000,000
//     G giga	1 x 10^9	1,000,000,000
//     T tera	1 x 10^12	1,000,000,000,000
//     P peta	1 x 10^15	1,000,000,000,000,000
//     E exa	1 x 10^18	1,000,000,000,000,000,000
//     Z zetta	1 x 10^21	1,000,000,000,000,000,000,000
//     Y yotta	1 x 10^24	1,000,000,000,000,000,000,000,000
//     Base	1 x 10^0	1
//     m milli	1 x 10^-3	0.001
//     μ micro	1 x 10^-6	0.000 001
//     n nano	1 x 10^-9	0.000 000 001
//     p pico	1 x 10^-12	0.000 000 000 001
//     f femto	1 x 10^-15	0.000 000 000 000 001
//     a atto	1 x 10^-18  0.000 000 000 000 000 001
//	   z zepto	1 x 10^-21	0.000 000 000 000 000 000 001
//     y yocto	1 x 10^-24	0.000 000 000 000 000 000 000 001

var lockEngNotationSI = sync.Mutex{}

//	Lock var lockEngNotationSI  before accessing
//	thia maps!

var mEngNotationSISymbols = map[string]string{
	"10^24":  "Y",
	"10^21":  "Z",
	"10^18":  "E",
	"10^15":  "P",
	"10^12":  "T",
	"10^9 ":  "G",
	"10^6 ":  "M",
	"10^3 ":  "k",
	"10^−3 ": "m",
	"10^−6 ": "μ",
	"10^−9 ": "n",
	"10^−12": "p",
	"10^−15": "f",
	"10^−18": "a",
	"10^−21": "z",
	"10^−24": "y",
}

var mEngNotationSINames = map[string]string{
	"10^24":  "yotta",
	"10^21":  "zetta",
	"10^18":  "exa",
	"10^15":  "peta",
	"10^12":  "tera",
	"10^9 ":  "giga",
	"10^6 ":  "mega",
	"10^3 ":  "kilo",
	"10^−3 ": "milli",
	"10^−6 ": "micro",
	"10^−9 ": "nano",
	"10^−12": "pico",
	"10^−15": "femto",
	"10^−18": "atto",
	"10^−21": "zepto",
	"10^−24": "yocto",
}
