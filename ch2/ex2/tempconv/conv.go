package tempconv

func FToY(f Foot) Yard {
	return Yard(f / 3)
}

func FToM(f Foot) Metre {
	return Metre(f * 0.3048)
}

func MToF(m Metre) Foot {
	return Foot(m / 0.3048)
}

func MToY(m Metre) Yard {
	return Yard(m * 1.0936)
}

func YToF(y Yard) Foot {
	return Foot(y * 3)
}

func YToM(y Yard) Metre {
	return Metre(y * 0.9144)
}
