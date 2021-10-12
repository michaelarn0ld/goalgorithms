package moduleone

func BaseToBase(str string, base, newBase int) string {
	return DecToBase(BaseToDec(str, base), newBase)
}
