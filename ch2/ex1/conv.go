package tempconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToK(c Celsius) Kelvinscale {
	return Kelvinscale(c + 273.15)
}

func KToC(k Kelvinscale) Celsius {
	return Celsius(k - 273.15)
}
