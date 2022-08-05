package kinship

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKinship_String(t *testing.T) {

	var tpKinship Kinship

	require.Equal(t, name_undefined, tpKinship.String(), "Must be Undefined")

	tpKinship = Parent

	require.Equal(t, name_parent, tpKinship.String(), "Must be Parent")

	tpKinship = Child

	require.Equal(t, name_child, tpKinship.String(), "Must be Child")

	tpKinship = Spouse

	require.Equal(t, name_spouse, tpKinship.String(), "Must be Spouse")

	tpKinship = Sibling

	require.Equal(t, name_sibling, tpKinship.String(), "Must be Sibling")

	tpKinship = Cousins

	require.Equal(t, name_cousin, tpKinship.String(), "Must be Cousins")

	tpKinship = 63

	require.Panics(t, func() {
		_ = tpKinship.String()
	}, "Type kinship don't mapped")
}

func TestKinship_MarshalJSON(t *testing.T) {

	var tpKinship Kinship

	contentJsonExpected := `"Undefined"`

	parsedJson, err := tpKinship.MarshalJSON()

	require.Nil(t, err, "Error must be returned nil when parsed Undefined")

	require.Equal(t, contentJsonExpected, string(parsedJson), "Content parsed to Undefined has been success")

	tpKinship = Parent

	contentJsonExpected = `"Parent"`

	parsedJson, err = tpKinship.MarshalJSON()

	require.Nil(t, err, "Error must be returned nil when parsed Parent")

	require.Equal(t, contentJsonExpected, string(parsedJson), "Content parsed to Parent has been success")

	tpKinship = Child

	contentJsonExpected = `"Child"`

	parsedJson, err = tpKinship.MarshalJSON()

	require.Nil(t, err, "Error must be returned nil when parsed Child")

	require.Equal(t, contentJsonExpected, string(parsedJson), "Content parsed to Child has been success")

	tpKinship = Spouse

	contentJsonExpected = `"Spouse"`

	parsedJson, err = tpKinship.MarshalJSON()

	require.Nil(t, err, "Error must be return nil when parsed Spouse")

	require.Equal(t, contentJsonExpected, string(parsedJson), "Content parsed to Spouse has been success")

	tpKinship = Sibling

	contentJsonExpected = `"Sibling"`

	parsedJson, err = tpKinship.MarshalJSON()

	require.Nil(t, err, "Error must be return nil when parsed Sibling")

	require.Equal(t, contentJsonExpected, string(parsedJson), "Content parsed to Sibling has been success")

	tpKinship = Cousins

	contentJsonExpected = `"Cousins"`

	parsedJson, err = tpKinship.MarshalJSON()

	require.Nil(t, err, "Error must be return nil when parsed Cousins")

	require.Equal(t, contentJsonExpected, string(parsedJson), "Content parsed to Cousins has been success")

	tpKinship = 52

	require.Panics(t, func() {
		parsedJson, err = tpKinship.MarshalJSON()
	}, "Type kinship don't mapped, then will happen panic")

}

func TestKinship_UnmarshalJSON(t *testing.T) {

	var tpKinship Kinship

	possibilities := []string{
		`"Undefined"`,
		`"undefined"`,
		`"UNDEFINED"`,
		`"0"`,
		`0`,
	}

	for _, possibility := range possibilities {

		err := tpKinship.UnmarshalJSON([]byte(possibility))

		typeExpected := Undefined

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpKinship, "Content parsed with value Undefined has been success")
	}

	possibilities = []string{
		`"Parent"`,
		`"parent"`,
		`"PARENT"`,
		`"1"`,
		`1`,
	}

	for _, possibility := range possibilities {

		err := tpKinship.UnmarshalJSON([]byte(possibility))

		typeExpected := Parent

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpKinship, "Content parsed with value Parent has been success")
	}

	possibilities = []string{
		`"Child"`,
		`"child"`,
		`"CHILD"`,
		`"2"`,
		`2`,
	}

	for _, possibility := range possibilities {

		err := tpKinship.UnmarshalJSON([]byte(possibility))

		typeExpected := Child

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpKinship, "Content parsed with value Child has been success")
	}

	possibilities = []string{
		`"Spouse"`,
		`"spouse"`,
		`"SPOUSE"`,
		`"3"`,
		`3`,
	}

	for _, possibility := range possibilities {

		err := tpKinship.UnmarshalJSON([]byte(possibility))

		typeExpected := Spouse

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpKinship, "Content parsed with value Spouse has been success")
	}

	possibilities = []string{
		`"Sibling"`,
		`"sibling"`,
		`"SIBLING"`,
		`"4"`,
		`4`,
	}

	for _, possibility := range possibilities {

		err := tpKinship.UnmarshalJSON([]byte(possibility))

		typeExpected := Sibling

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpKinship, "Content parsed with value Sibling has been success")
	}

	possibilities = []string{
		`"Cousins"`,
		`"cousins"`,
		`"COUSINS"`,
		`"5"`,
		`5`,
	}

	for _, possibility := range possibilities {

		err := tpKinship.UnmarshalJSON([]byte(possibility))

		typeExpected := Cousins

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpKinship, "Content parsed with value Cousins has been success")
	}

	possibilities = []string{
		`"sdklfjsfd"`,
		`"Asdfg"`,
		`"SDFSDF"`,
		`"445"`,
		`4345345`,
	}

	for _, possibility := range possibilities {

		err := tpKinship.UnmarshalJSON([]byte(possibility))

		require.NotNil(t, err, "Content parsed with error, the error must be not nil")

	}

}

func TestKinship_Value(t *testing.T) {
	var tpKinship Kinship

	contentValueExpected := `Undefined`

	parsedValue, err := tpKinship.Value()

	require.Nil(t, err, "Error must be returned nil when parsed Undefined")

	require.Equal(t, contentValueExpected, parsedValue, "Content parsed to Undefined has been success")

	tpKinship = Parent

	contentValueExpected = `Parent`

	parsedValue, err = tpKinship.Value()

	require.Nil(t, err, "Error must be returned nil when parsed Parent")

	require.Equal(t, contentValueExpected, parsedValue, "Content parsed to Parent has been success")

	tpKinship = Child

	contentValueExpected = `Child`

	parsedValue, err = tpKinship.Value()

	require.Nil(t, err, "Error must be returned nil when parsed Child")

	require.Equal(t, contentValueExpected, parsedValue, "Content parsed to Child has been success")

	tpKinship = Spouse

	contentValueExpected = `Spouse`

	parsedValue, err = tpKinship.Value()

	require.Nil(t, err, "Error must be returned nul when parsed Spouse")

	require.Equal(t, contentValueExpected, parsedValue, "Content parsed to Spouse has been success")

	tpKinship = Sibling

	contentValueExpected = `Sibling`

	parsedValue, err = tpKinship.Value()

	require.Nil(t, err, "Error must be returned nul when parsed Sibling")

	require.Equal(t, contentValueExpected, parsedValue, "Content parsed to Sibling has been success")

	tpKinship = Cousins

	contentValueExpected = `Cousins`

	parsedValue, err = tpKinship.Value()

	require.Nil(t, err, "Error must be returned nul when parsed Cousins")

	require.Equal(t, contentValueExpected, parsedValue, "Content parsed to Cousins has been success")

	tpKinship = 53

	require.Panics(t, func() {
		parsedValue, err = tpKinship.Value()
	}, "Type kinship don't mapped, then will happen panic")
}

func TestKinship_Scan(t *testing.T) {
	var tpKinship Kinship

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
		Kinship(0),
	}

	for _, possibility := range possibilities {

		err := tpKinship.Scan(possibility)

		typeExpected := Undefined

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpKinship, "Content parsed with value Undefined has been success")
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
		Kinship(1),
	}

	for _, possibility := range possibilities {

		err := tpKinship.Scan(possibility)

		typeExpected := Parent

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpKinship, "Content parsed with value Parent has been success")
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
		Kinship(2),
	}

	for _, possibility := range possibilities {

		err := tpKinship.Scan(possibility)

		typeExpected := Child

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpKinship, "Content parsed with value Child has been success")
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
		Kinship(3),
	}

	for _, possibility := range possibilities {

		err := tpKinship.Scan(possibility)

		typeExpected := Spouse

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpKinship, "Content parsed with value Spouse has been success")
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
		Kinship(4),
	}

	for _, possibility := range possibilities {

		err := tpKinship.Scan(possibility)

		typeExpected := Sibling

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpKinship, "Content parsed with value Sibling has been success")
	}

	possibilities = []interface{}{
		`Cousins`,
		`cousins`,
		`COUSINS`,
		`5`,
		[]uint8(`Cousins`),
		[]uint8(`cousins`),
		[]uint8(`COUSINS`),
		[]uint8(`5`),
		int(5),
		int32(5),
		int64(5),
		float32(5),
		float64(5),
		Kinship(5),
	}

	for _, possibility := range possibilities {

		err := tpKinship.Scan(possibility)

		typeExpected := Cousins

		require.Nil(t, err, "Content parsed with success, the error must be nil")

		require.Equal(t, typeExpected, tpKinship, "Content parsed with value Cousins has been success")
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
		Kinship(8657),
	}

	for _, possibility := range possibilities {

		err := tpKinship.Scan(possibility)

		require.NotNil(t, err, "Content parsed with error, the error must be not nil")

	}
}
