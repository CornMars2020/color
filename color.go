package color

func GetRed(s string) string {
	return "\033[31m" + s + "\033[0m"
}

func GetGreen(s string) string {
	return "\033[32m" + s + "\033[0m"
}

func GetYellow(s string) string {
	return "\033[33m" + s + "\033[0m"
}

func GetBlue(s string) string {
	return "\033[34m" + s + "\033[0m"
}

func GetPurple(s string) string {
	return "\033[35m" + s + "\033[0m"
}

func GetCyan(s string) string {
	return "\033[36m" + s + "\033[0m"
}

func GetWhite(s string) string {
	return "\033[37m" + s + "\033[0m"
}

func GetGray(s string) string {
	return "\033[90m" + s + "\033[0m"
}

func GetLightRed(s string) string {
	return "\033[91m" + s + "\033[0m"
}

func GetLightGreen(s string) string {
	return "\033[92m" + s + "\033[0m"
}

func GetLightYellow(s string) string {
	return "\033[93m" + s + "\033[0m"
}

func GetLightBlue(s string) string {
	return "\033[94m" + s + "\033[0m"
}

func GetLightPurple(s string) string {
	return "\033[95m" + s + "\033[0m"
}

func GetLightCyan(s string) string {
	return "\033[96m" + s + "\033[0m"
}

func GetLightWhite(s string) string {
	return "\033[97m" + s + "\033[0m"
}

func GetReset(s string) string {
	return "\033[0m" + s + "\033[0m"
}

func GetBold(s string) string {
	return "\033[1m" + s + "\033[0m"
}

func GetDim(s string) string {
	return "\033[2m" + s + "\033[0m"
}

func GetItalic(s string) string {
	return "\033[3m" + s + "\033[0m"
}

func GetUnderline(s string) string {
	return "\033[4m" + s + "\033[0m"
}

func GetReverse(s string) string {
	return "\033[7m" + s + "\033[0m"
}
