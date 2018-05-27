package quran

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeEmptyString(t *testing.T) {
	input := ""
	expected := []string{}
	actual := quranizeTest.Encode(input)
	assert.Equal(t, expected, actual)
}

func TestEncodeNonAlquran(t *testing.T) {
	input := "alfan nur fauzan"
	expected := []string{}
	actual := quranizeTest.Encode(input)
	assert.Equal(t, expected, actual)
}

func TestEncodeAlquran(t *testing.T) {
	testCases := map[string][]string{
		"tajri":                 []string{"تجري"},
		"alhamdulillah":         []string{"الحمد لله"},
		"bismillah":             []string{"بسم الله", "بشماله"},
		"wa'tasimu":             []string{"واعتصموا"},
		"wa'tasimu bihablillah": []string{"واعتصموا بحبل الله"},
		"shummun bukmun":        []string{"صم وبكم", "صم بكم", "الصم البكم"},
		"kahfi":                 []string{"الكهف"},
		"wabasyiris sobirin":    []string{"وبشر الصابرين"},
		"bissobri wassolah":     []string{"بالصبر والصلاة"},
		"ya aiyuhalladzina":     []string{"يا أيها الذين"},

		"bismillah hirrohman nirrohim":                                                []string{"بسم الله الرحمن الرحيم"},
		"alhamdu lillahi robbil 'alamin":                                              []string{"الحمد لله رب العالمين"},
		"arrohma nirrohim":                                                            []string{"الرحمن الرحيم"},
		"maaliki yau middin":                                                          []string{"مالك يوم الدين"},
		"iyya kanakbudu waiyya kanastain":                                             []string{"إياك نعبد وإياك نستعين"},
		"ihdinash shirothol mustaqim":                                                 []string{"اهدنا الصراط المستقيم"},
		"shirotholladzina an'am ta'alaihim ghoiril maghdzu bi'alaihim waladh dhollin": []string{"صراط الذين أنعمت عليهم غير المغضوب عليهم ولا الضالين"},
	}
	for input, expected := range testCases {
		actual := quranizeTest.Encode(input)
		assert.Equal(t, expected, actual)
	}
}

func TestLocateEmptyString(t *testing.T) {
	input := ""
	expected := zeroLocs
	actual := quranizeTest.Locate(input)
	assert.Equal(t, expected, actual)
}

func TestLocateNonAlquran(t *testing.T) {
	input := "alfan"
	expected := zeroLocs
	actual := quranizeTest.Locate(input)
	assert.Equal(t, expected, actual)
}

func TestLocateAlquran(t *testing.T) {
	input := "بسم الله الرحمن الرحيم"
	expected := []Location{Location{1, 1, 0}, Location{27, 30, 4}}
	actual := quranizeTest.Locate(input)
	assert.Equal(t, expected, actual)
}

func TestLocateAlquranBeforeBuildIndex(t *testing.T) {
	root := quranizeTest.root
	defer func() { quranizeTest.root = root }()
	quranizeTest.root = nil
	input := "بسم الله الرحمن الرحيم"
	expected := zeroLocs
	actual := quranizeTest.Locate(input)
	assert.Equal(t, expected, actual)
}