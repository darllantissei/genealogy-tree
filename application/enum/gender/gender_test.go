package gender

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGender_String(t *testing.T) {

	var tpGender Gender

	require.Equal(t, name_undefined, tpGender.String(), "Must be Undefined")

	tpGender = Masculine

	require.Equal(t, name_masculine, tpGender.String(), "Must be Masculine")

	tpGender = Female

	require.Equal(t, name_female, tpGender.String(), "Must be Female")

	tpGender = 3

	require.Panics(t, func() {
		_ = tpGender.String()
	}, "Type gender don't mapped")
}

func TestGender_MarshalJSON(t *testing.T) {

	var tpGender Gender

	contentJsonExpected := `"Undefined"`

	parsedJson, err := tpGender.MarshalJSON()

	require.Nil(t, err, "Error must be returned nil when parsed Undefined")

	require.Equal(t, contentJsonExpected, string(parsedJson), "Content parsed to Undefined has been success")

	tpGender = Masculine

	contentJsonExpected = `"Masculine"`

	parsedJson, err = tpGender.MarshalJSON()

	require.Nil(t, err, "Error must be returned nil when parsed Masculine")

	require.Equal(t, contentJsonExpected, string(parsedJson), "Content parsed to Masculine has been success")

	tpGender = Female

	contentJsonExpected = `"Female"`

	parsedJson, err = tpGender.MarshalJSON()

	require.Nil(t, err, "Error must be returned nil when parsed Female")

	require.Equal(t, contentJsonExpected, string(parsedJson), "Content parsed to Female has been success")

	tpGender = 3

	require.Panics(t, func() {
		parsedJson, err = tpGender.MarshalJSON()
	}, "Type gender don't mapped, then will happen panic")

}

func TestGender_UnmarshalJSON(t *testing.T) {

	var tpGender Gender

	possibilities := []string{
		`"Undefined"`,
		`"undefined"`,
		`"UNDEFINED"`,
		`"0"`,
		`0`,
	}

	for _, possibility := range possibilities {

		err := tpGender.UnmarshalJSON([]byte(possibility))

		typeExpected := Undefined

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpGender, "Content parsed with value Undefined has been success")
	}

	possibilities = []string{
		`"Masculine"`,
		`"masculine"`,
		`"MASCULINE"`,
		`"1"`,
		`1`,
	}

	for _, possibility := range possibilities {

		err := tpGender.UnmarshalJSON([]byte(possibility))

		typeExpected := Masculine

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpGender, "Content parsed with value Masculine has been success")
	}

	possibilities = []string{
		`"Female"`,
		`"female"`,
		`"FEMALE"`,
		`"2"`,
		`2`,
	}

	for _, possibility := range possibilities {

		err := tpGender.UnmarshalJSON([]byte(possibility))

		typeExpected := Female

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpGender, "Content parsed with value Female has been success")
	}

	possibilities = []string{
		`"sdklfjsfd"`,
		`"Asdfg"`,
		`"SDFSDF"`,
		`"445"`,
		`4345345`,
	}

	for _, possibility := range possibilities {

		err := tpGender.UnmarshalJSON([]byte(possibility))

		require.NotNil(t, err, "Content parsed with error, the error must be not nil")

	}

}

func TestGender_Value(t *testing.T) {
	var tpGender Gender

	contentValueExpected := `Undefined`

	parsedValue, err := tpGender.Value()

	require.Nil(t, err, "Error must be returned nil when parsed Undefined")

	require.Equal(t, contentValueExpected, parsedValue, "Content parsed to Undefined has been success")

	tpGender = Masculine

	contentValueExpected = `Masculine`

	parsedValue, err = tpGender.Value()

	require.Nil(t, err, "Error must be returned nil when parsed Masculine")

	require.Equal(t, contentValueExpected, parsedValue, "Content parsed to Masculine has been success")

	tpGender = Female

	contentValueExpected = `Female`

	parsedValue, err = tpGender.Value()

	require.Nil(t, err, "Error must be returned nil when parsed Female")

	require.Equal(t, contentValueExpected, parsedValue, "Content parsed to Masculine has been success")

	tpGender = 3

	require.Panics(t, func() {
		parsedValue, err = tpGender.Value()
	}, "Type gender don't mapped, then will happen panic")
}

func TestGender_Scan(t *testing.T) {
	var tpGender Gender

	possibilities := []interface{}{
		`Undefined`,
		`undefined`,
		`UNDEFINED`,
		`0`,
		[]uint8(`Undefined`),
		[]uint8(`undefined`),
		[]uint8(`UNDEFINED`),
		[]uint8(`0`),
		int(0),
		int32(0),
		int64(0),
		float32(0),
		float64(0),
		Gender(0),
	}

	for _, possibility := range possibilities {

		err := tpGender.Scan(possibility)

		typeExpected := Undefined

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpGender, "Content parsed with value Undefined has been success")
	}

	possibilities = []interface{}{
		`Masculine`,
		`masculine`,
		`MASCULINE`,
		`1`,
		[]uint8(`Masculine`),
		[]uint8(`masculine`),
		[]uint8(`MASCULINE`),
		[]uint8(`1`),
		int(1),
		int32(1),
		int64(1),
		float32(1),
		float64(1),
		Gender(1),
	}

	for _, possibility := range possibilities {

		err := tpGender.Scan(possibility)

		typeExpected := Masculine

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpGender, "Content parsed with value Masculine has been success")
	}

	possibilities = []interface{}{
		`Female`,
		`female`,
		`FEMALE`,
		`2`,
		[]uint8(`Female`),
		[]uint8(`female`),
		[]uint8(`FEMALE`),
		[]uint8(`2`),
		int(2),
		int32(2),
		int64(2),
		float32(2),
		float64(2),
		Gender(2),
	}

	for _, possibility := range possibilities {

		err := tpGender.Scan(possibility)

		typeExpected := Female

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpGender, "Content parsed with value Female has been success")
	}

	possibilities = []interface{}{
		`Adfsf`,
		`sdfsdf`,
		`SFDDSFDS`,
		`645`,
		[]uint8(`Adfsf`),
		[]uint8(`sdfsdf`),
		[]uint8(`SFDDSFDS`),
		[]uint8(`756`),
		int(8657),
		int32(8657),
		int64(8657),
		float32(8657),
		float64(8657),
		Gender(8657),
	}

	for _, possibility := range possibilities {

		err := tpGender.Scan(possibility)

		require.NotNil(t, err, "Content parsed with error, the error must be not nil")

	}
}

func TestGender_MarshalXML(t *testing.T) {
	possibilities := []interface{}{
		0,
		`Undefined`,
		`undefined`,
		`UNDEFINED`,
		`0`,
		1,
		`Masculine`,
		`masculine`,
		`MASCULINE`,
		`1`,
		2,
		`Female`,
		`female`,
		`FEMALE`,
		`2`,
	}

	var tpGender Gender

	for _, possibility := range possibilities {

		valueStsApp, err := tpGender.tryParseValueToGender(fmt.Sprint(possibility))

		require.Nil(t, err, "Error must be nil when parse possibility to gender")

		expected := []byte(fmt.Sprintf(`<gender>%s</gender>`, valueStsApp))

		contentOUT := new(bytes.Buffer)

		err = valueStsApp.MarshalXML(xml.NewEncoder(contentOUT), xml.StartElement{Name: xml.Name{Space: "", Local: "gender"}})

		require.Nil(t, err, "Error must be returned nil when parsed gender")

		require.Equal(t, expected, contentOUT.Bytes(), "The parse XML must be equal expected")

	}
}

func TestGender_UnmarshalXML(t *testing.T) {
	possibilities := []interface{}{
		0,
		`Undefined`,
		`undefined`,
		`UNDEFINED`,
		`0`,
		1,
		`Masculine`,
		`masculine`,
		`MASCULINE`,
		`1`,
		2,
		`Female`,
		`female`,
		`FEMALE`,
		`2`,
	}

	var tpGender Gender

	for _, possibility := range possibilities {

		expectGender, err := tpGender.tryParseValueToGender(fmt.Sprint(possibility))

		require.Nil(t, err, "Error must be nil")

		contentIN := []byte(`<Gender>` + fmt.Sprint(possibility) + `</Gender>`)

		err = xml.Unmarshal(contentIN, &tpGender)

		require.Nil(t, err, "Error must be returned nil when parsed gender")

		require.Equal(t, expectGender, tpGender, "The gender parsed must be equal expected")
	}

	contentIN := []byte(`<Gender></Gender>`)

	err := xml.Unmarshal(contentIN, &tpGender)

	require.NotNil(t, err, "Error must be returned when parsed gender empty")
}
