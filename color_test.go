package color

import (
	"testing"
)

func TestAllStrings(t *testing.T) {
	t.Logf("%s", GetRed("Red"))
	t.Logf("%s", GetGreen("Green"))
	t.Logf("%s", GetYellow("Yellow"))
	t.Logf("%s", GetBlue("Blue"))
	t.Logf("%s", GetPurple("Purple"))
	t.Logf("%s", GetCyan("Cyan"))
	t.Logf("%s", GetWhite("White"))
	t.Logf("%s", GetGray("Gray"))
	t.Logf("%s", GetLightRed("LightRed"))
	t.Logf("%s", GetLightGreen("LightGreen"))
	t.Logf("%s", GetLightYellow("LightYellow"))
	t.Logf("%s", GetLightBlue("LightBlue"))
	t.Logf("%s", GetLightPurple("LightPurple"))
	t.Logf("%s", GetLightCyan("LightCyan"))
	t.Logf("%s", GetLightWhite("LightWhite"))
	t.Logf("%s", GetReset("Reset"))
	t.Logf("%s", GetBold("Bold"))
	t.Logf("%s", GetDim("Dim"))
	t.Logf("%s", GetItalic("Italic"))
	t.Logf("%s", GetUnderline("Underline"))
	t.Logf("%s", GetReverse("Reverse"))
}
