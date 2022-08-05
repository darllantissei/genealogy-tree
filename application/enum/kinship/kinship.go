package kinship

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Kinship int64

const (
	Undefined Kinship = iota // 0 - Indefinido - Parentesco não definido
	Parent                   // 1 - Pais
	Child                    // 2 - Filhos
	Spouse                   // 3 - Cônjuge
	Sibling                  // 4 - Irmãos
	Cousins                  // 5 - Primos

	name_undefined = "Undefined"
	name_parent    = "Parent"
	name_child     = "Child"
	name_spouse    = "Spouse"
	name_sibling   = "Sibling"
	name_cousin    = "Cousins"
)

var (
	kinship_name = map[int64]string{
		0: name_undefined,
		1: name_parent,
		2: name_child,
		3: name_spouse,
		4: name_sibling,
		5: name_cousin,
	}
	kinship_value = map[string]int64{
		name_undefined: 0,
		name_parent:    1,
		name_child:     2,
		name_spouse:    3,
		name_sibling:   4,
		name_cousin:    5,
	}

	typeAccepts = func() string {

		descriptionTypeAccepts := "Types accepts: "

		for enum, types := range kinship_name {
			descriptionTypeAccepts += fmt.Sprintf("%d ou %s | ", enum, types)
		}

		return descriptionTypeAccepts
	}
)

func (k Kinship) String() string {
	switch k {
	case Undefined:
		return name_undefined
	case Parent:
		return name_parent
	case Child:
		return name_child
	case Spouse:
		return name_spouse
	case Sibling:
		return name_sibling
	case Cousins:
		return name_cousin
	default:
		panic(fmt.Sprintf("type kinship is invalid. %s", typeAccepts()))
	}

}

func (k Kinship) MarshalJSON() ([]byte, error) {

	defer func() {
		errRecover := recover()

		if errRecover != nil {
			panic(fmt.Sprintf("Marshal failed. Type kinship informed: %d. Failed details: %v", k, errRecover))
		}
	}()

	return []byte(fmt.Sprintf(`"%s"`, k.String())), nil
}

func (k *Kinship) UnmarshalJSON(bytes []byte) error {
	value, err := k.tryGetValueFromJSON(bytes)
	if err == nil && !strings.EqualFold(value, "") {

		knshp, err := k.tryParseValueToKinship(value)

		if err != nil {
			return err
		}

		*k = Kinship(knshp)
	}

	return err
}

func (k Kinship) Value() (driver.Value, error) {

	defer func() {
		errRecover := recover()

		if errRecover != nil {
			panic(fmt.Sprintf("Value failed. Type kinship informed: %d. Failed details: %v", k, errRecover))
		}
	}()

	return k.String(), nil
}

func (k *Kinship) Scan(value interface{}) (err error) {

	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = fmt.Errorf("Scan failed for value %v for type kinship. Details: %v", value, errRecover)
		}
	}()

	switch data := value.(type) {
	case []uint8:
		str := string([]byte(data))

		knshp, err := k.tryParseValueToKinship(str)

		if err != nil {
			return err
		}

		*k = Kinship(knshp)

	case int:
		d := int64(data)
		*k = Kinship(d)
		_ = k.String()
	case int32:
		d := int64(data)
		*k = Kinship(d)
		_ = k.String()
	case float32:
		d := int64(data)
		*k = Kinship(d)
		_ = k.String()
	case float64:
		d := int64(data)
		*k = Kinship(d)
		_ = k.String()
	case int64:
		*k = Kinship(data)
		_ = k.String()
	case string:
		knshp, err := k.tryParseValueToKinship(data)
		if err != nil {
			panic(err)
		}
		*k = knshp
	default:
		_ = k.String()
	}

	return nil
}

func (k *Kinship) tryParseValueToKinship(value string) (knshp Kinship, err error) {

	toTitle := cases.Title(language.BrazilianPortuguese, cases.NoLower)

	value = toTitle.String(strings.ToLower(value))

	knshpINT, ok := kinship_value[value]

	if !ok {
		valueINT, err := strconv.Atoi(value)
		if err != nil {
			return Undefined, fmt.Errorf("the %s is incorret to type of kinship", value)
		}

		knshpStr, ok := kinship_name[int64(valueINT)]
		if !ok {
			return Undefined, fmt.Errorf("the %s not valid type of kinship", value)
		}

		knshpINT, ok = kinship_value[knshpStr]
		if !ok {
			return Undefined, fmt.Errorf("the %s is invalid type of kinship", value)
		}
	}

	return Kinship(knshpINT), nil
}

func (k *Kinship) tryGetValueFromJSON(bytes []byte) (value string, err error) {
	value, err = strconv.Unquote(string(bytes))

	if err != nil {

		valueINT := int(Undefined)

		valueINT, err = strconv.Atoi(string(bytes))

		if err != nil {
			return
		}

		value = fmt.Sprintf("%d", valueINT)
		err = nil
	}

	return
}
