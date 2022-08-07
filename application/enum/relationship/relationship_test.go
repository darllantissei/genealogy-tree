package relationship

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRelationship_String(t *testing.T) {

	var tpRelationship Relationship

	require.Equal(t, name_undefined, tpRelationship.String(), "Must be Undefined")

	tpRelationship = Parent

	require.Equal(t, name_parent, tpRelationship.String(), "Must be Parent")

	tpRelationship = Child

	require.Equal(t, name_child, tpRelationship.String(), "Must be Child")

	tpRelationship = Spouse

	require.Equal(t, name_spouse, tpRelationship.String(), "Must be Spouse")

	tpRelationship = Sibling

	require.Equal(t, name_sibling, tpRelationship.String(), "Must be Sibling")

	tpRelationship = 5

	require.Panics(t, func() {
		_ = tpRelationship.String()
	}, "Type relationship don't mapped")
}

func TestRelationship_MarshalJSON(t *testing.T) {

	var tpRelationship Relationship

	contentJsonExpected := `"Undefined"`

	parsedJson, err := tpRelationship.MarshalJSON()

	require.Nil(t, err, "Error must be returned nil when parsed Undefined")

	require.Equal(t, contentJsonExpected, string(parsedJson), "Content parsed to Undefined has been success")

	tpRelationship = Parent

	contentJsonExpected = `"Parent"`

	parsedJson, err = tpRelationship.MarshalJSON()

	require.Nil(t, err, "Error must be returned nil when parsed Parent")

	require.Equal(t, contentJsonExpected, string(parsedJson), "Content parsed to Parent has been success")

	tpRelationship = Child

	contentJsonExpected = `"Child"`

	parsedJson, err = tpRelationship.MarshalJSON()

	require.Nil(t, err, "Error must be returned nil when parsed Child")

	require.Equal(t, contentJsonExpected, string(parsedJson), "Content parsed to Child has been success")

	tpRelationship = Spouse

	contentJsonExpected = `"Spouse"`

	parsedJson, err = tpRelationship.MarshalJSON()

	require.Nil(t, err, "Error must be return nil when parsed Spouse")

	require.Equal(t, contentJsonExpected, string(parsedJson), "Content parsed to Spouse has been success")

	tpRelationship = Sibling

	contentJsonExpected = `"Sibling"`

	parsedJson, err = tpRelationship.MarshalJSON()

	require.Nil(t, err, "Error must be return nil when parsed Sibling")

	require.Equal(t, contentJsonExpected, string(parsedJson), "Content parsed to Sibling has been success")

	tpRelationship = 5

	require.Panics(t, func() {
		parsedJson, err = tpRelationship.MarshalJSON()
	}, "Type relationship don't mapped, then will happen panic")

}

func TestRelationship_UnmarshalJSON(t *testing.T) {

	var tpRelationship Relationship

	possibilities := []string{
		`"Undefined"`,
		`"undefined"`,
		`"UNDEFINED"`,
		`"0"`,
		`0`,
	}

	for _, possibility := range possibilities {

		err := tpRelationship.UnmarshalJSON([]byte(possibility))

		typeExpected := Undefined

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpRelationship, "Content parsed with value Undefined has been success")
	}

	possibilities = []string{
		`"Parent"`,
		`"parent"`,
		`"PARENT"`,
		`"1"`,
		`1`,
	}

	for _, possibility := range possibilities {

		err := tpRelationship.UnmarshalJSON([]byte(possibility))

		typeExpected := Parent

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpRelationship, "Content parsed with value Parent has been success")
	}

	possibilities = []string{
		`"Child"`,
		`"child"`,
		`"CHILD"`,
		`"2"`,
		`2`,
	}

	for _, possibility := range possibilities {

		err := tpRelationship.UnmarshalJSON([]byte(possibility))

		typeExpected := Child

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpRelationship, "Content parsed with value Child has been success")
	}

	possibilities = []string{
		`"Spouse"`,
		`"spouse"`,
		`"SPOUSE"`,
		`"3"`,
		`3`,
	}

	for _, possibility := range possibilities {

		err := tpRelationship.UnmarshalJSON([]byte(possibility))

		typeExpected := Spouse

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpRelationship, "Content parsed with value Spouse has been success")
	}

	possibilities = []string{
		`"Sibling"`,
		`"sibling"`,
		`"SIBLING"`,
		`"4"`,
		`4`,
	}

	for _, possibility := range possibilities {

		err := tpRelationship.UnmarshalJSON([]byte(possibility))

		typeExpected := Sibling

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpRelationship, "Content parsed with value Sibling has been success")
	}

	possibilities = []string{
		`"sdklfjsfd"`,
		`"Asdfg"`,
		`"SDFSDF"`,
		`"445"`,
		`4345345`,
	}

	for _, possibility := range possibilities {

		err := tpRelationship.UnmarshalJSON([]byte(possibility))

		require.NotNil(t, err, "Content parsed with error, the error must be not nil")

	}

}

func TestRelationship_Value(t *testing.T) {
	var tpRelationship Relationship

	contentValueExpected := `Undefined`

	parsedValue, err := tpRelationship.Value()

	require.Nil(t, err, "Error must be returned nil when parsed Undefined")

	require.Equal(t, contentValueExpected, parsedValue, "Content parsed to Undefined has been success")

	tpRelationship = Parent

	contentValueExpected = `Parent`

	parsedValue, err = tpRelationship.Value()

	require.Nil(t, err, "Error must be returned nil when parsed Parent")

	require.Equal(t, contentValueExpected, parsedValue, "Content parsed to Parent has been success")

	tpRelationship = Child

	contentValueExpected = `Child`

	parsedValue, err = tpRelationship.Value()

	require.Nil(t, err, "Error must be returned nil when parsed Child")

	require.Equal(t, contentValueExpected, parsedValue, "Content parsed to Child has been success")

	tpRelationship = Spouse

	contentValueExpected = `Spouse`

	parsedValue, err = tpRelationship.Value()

	require.Nil(t, err, "Error must be returned nul when parsed Spouse")

	require.Equal(t, contentValueExpected, parsedValue, "Content parsed to Spouse has been success")

	tpRelationship = Sibling

	contentValueExpected = `Sibling`

	parsedValue, err = tpRelationship.Value()

	require.Nil(t, err, "Error must be returned nul when parsed Sibling")

	require.Equal(t, contentValueExpected, parsedValue, "Content parsed to Sibling has been success")

	tpRelationship = 5

	require.Panics(t, func() {
		parsedValue, err = tpRelationship.Value()
	}, "Type relationship don't mapped, then will happen panic")
}

func TestRelationship_Scan(t *testing.T) {
	var tpRelationship Relationship

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
		Relationship(0),
	}

	for _, possibility := range possibilities {

		err := tpRelationship.Scan(possibility)

		typeExpected := Undefined

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpRelationship, "Content parsed with value Undefined has been success")
	}

	possibilities = []interface{}{
		`Parent`,
		`parent`,
		`PARENT`,
		`1`,
		[]uint8(`Parent`),
		[]uint8(`parent`),
		[]uint8(`parent`),
		[]uint8(`1`),
		int(1),
		int32(1),
		int64(1),
		float32(1),
		float64(1),
		Relationship(1),
	}

	for _, possibility := range possibilities {

		err := tpRelationship.Scan(possibility)

		typeExpected := Parent

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpRelationship, "Content parsed with value Parent has been success")
	}

	possibilities = []interface{}{
		`Child`,
		`child`,
		`CHILD`,
		`2`,
		[]uint8(`Child`),
		[]uint8(`child`),
		[]uint8(`CHILD`),
		[]uint8(`2`),
		int(2),
		int32(2),
		int64(2),
		float32(2),
		float64(2),
		Relationship(2),
	}

	for _, possibility := range possibilities {

		err := tpRelationship.Scan(possibility)

		typeExpected := Child

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpRelationship, "Content parsed with value Child has been success")
	}

	possibilities = []interface{}{
		`Spouse`,
		`spouse`,
		`SPOUSE`,
		`3`,
		[]uint8(`Spouse`),
		[]uint8(`spouse`),
		[]uint8(`spouse`),
		[]uint8(`3`),
		int(3),
		int32(3),
		int64(3),
		float32(3),
		float64(3),
		Relationship(3),
	}

	for _, possibility := range possibilities {

		err := tpRelationship.Scan(possibility)

		typeExpected := Spouse

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpRelationship, "Content parsed with value Spouse has been success")
	}

	possibilities = []interface{}{
		`Sibling`,
		`sibling`,
		`SIBLING`,
		`4`,
		[]uint8(`Sibling`),
		[]uint8(`sibling`),
		[]uint8(`SIBLING`),
		[]uint8(`4`),
		int(4),
		int32(4),
		int64(4),
		float32(4),
		float64(4),
		Relationship(4),
	}

	for _, possibility := range possibilities {

		err := tpRelationship.Scan(possibility)

		typeExpected := Sibling

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpRelationship, "Content parsed with value Sibling has been success")
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
		Relationship(8657),
	}

	for _, possibility := range possibilities {

		err := tpRelationship.Scan(possibility)

		require.NotNil(t, err, "Content parsed with error, the error must be not nil")

	}
}

func TestRelationship_MarshalXML(t *testing.T) {
	possibilities := []interface{}{
		0,
		`Undefined`,
		`undefined`,
		`UNDEFINED`,
		"0",
		1,
		`Parent`,
		`parent`,
		`PARENT`,
		"1",
		2,
		`Child`,
		`child`,
		`CHILD`,
		`2`,
		3,
		`Spouse`,
		`spouse`,
		`SPOUSE`,
		`3`,
		4,
		`Sibling`,
		`sibling`,
		`SIBLING`,
		`4`,
	}

	var tpRelationship Relationship

	for _, possibility := range possibilities {

		valueStsApp, err := tpRelationship.tryParseValueToRelationship(fmt.Sprint(possibility))

		require.Nil(t, err, "Error must be nil when parse possibility to relationship")

		expected := []byte(fmt.Sprintf(`<relationship>%s</relationship>`, valueStsApp))

		contentOUT := new(bytes.Buffer)

		err = valueStsApp.MarshalXML(xml.NewEncoder(contentOUT), xml.StartElement{Name: xml.Name{Space: "", Local: "relationship"}})

		require.Nil(t, err, "Error must be returned nil when parsed relationship")

		require.Equal(t, expected, contentOUT.Bytes(), "The parse XML must be equal expected")

	}
}

func TestRelationship_UnmarshalXML(t *testing.T) {
	possibilities := []interface{}{
		0,
		`Undefined`,
		`undefined`,
		`UNDEFINED`,
		"0",
		1,
		`Parent`,
		`parent`,
		`PARENT`,
		"1",
		2,
		`Child`,
		`child`,
		`CHILD`,
		`2`,
		3,
		`Spouse`,
		`spouse`,
		`SPOUSE`,
		`3`,
		4,
		`Sibling`,
		`sibling`,
		`SIBLING`,
		`4`,
	}

	var tpRelationship Relationship

	for _, possibility := range possibilities {

		expectRelationship, err := tpRelationship.tryParseValueToRelationship(fmt.Sprint(possibility))

		require.Nil(t, err, "Error must be nil")

		contentIN := []byte(`<Relationship>` + fmt.Sprint(possibility) + `</Relationship>`)

		err = xml.Unmarshal(contentIN, &tpRelationship)

		require.Nil(t, err, "Error must be returned nil when parsed relationship")

		require.Equal(t, expectRelationship, tpRelationship, "The relationship parsed must be equal expected")
	}

	contentIN := []byte(`<Relationship></Relationship>`)

	err := xml.Unmarshal(contentIN, &tpRelationship)

	require.NotNil(t, err, "Error must be returned when parsed relationship empty")

}
