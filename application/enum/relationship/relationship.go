package relationship

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Relationship int64

const (
	Undefined Relationship = iota // 0 - Indefinido - Gênero não definido
	Parent                        // 1 - Pais
	Child                         // 2 - Filhos
	Spouse                        // 3 - Cônjuge
	Sibling                       // 4 - Irmãos

	name_undefined = "Undefined"
	name_parent    = "Parent"
	name_child     = "Child"
	name_spouse    = "Spouse"
	name_sibling   = "Sibling"
)

var (
	relationship_name = map[int64]string{
		0: name_undefined,
		1: name_parent,
		2: name_child,
		3: name_spouse,
		4: name_sibling,
	}
	relationship_value = map[string]int64{
		name_undefined: 0,
		name_parent:    1,
		name_child:     2,
		name_spouse:    3,
		name_sibling:   4,
	}

	typeAccepts = func() string {
		
		descriptionTypeAccepts := "Types accepts: "
		
		for enum, types := range relationship_name {
			descriptionTypeAccepts += fmt.Sprintf("%d ou %s | ", enum, types)
		}

		return descriptionTypeAccepts
	}
)

func (r Relationship) String() string {
	switch r {
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
	default:
		panic(fmt.Sprintf("type relationship is invalid. %s", typeAccepts()))
	}

}

func (r Relationship) MarshalJSON() ([]byte, error) {

	defer func() {
		errRecover := recover()

		if errRecover != nil {
			panic(fmt.Sprintf("Marshal failed. Type relationship informed: %d. Failed details: %v", r, errRecover))
		}
	}()

	return []byte(fmt.Sprintf(`"%s"`, r.String())), nil
}

func (r *Relationship) UnmarshalJSON(bytes []byte) error {
	value, err := r.tryGetValueFromJSON(bytes)
	if err == nil && !strings.EqualFold(value, "") {

		rtshp, err := r.tryParseValueToRelationship(value)

		if err != nil {
			return err
		}

		*r = Relationship(rtshp)
	}

	return err
}

func (r Relationship) Value() (driver.Value, error) {

	defer func() {
		errRecover := recover()

		if errRecover != nil {
			panic(fmt.Sprintf("Value failed. Type relationship informed: %d. Failed details: %v", r, errRecover))
		}
	}()

	return r.String(), nil
}

func (r *Relationship) Scan(value interface{}) (err error) {

	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = fmt.Errorf("Scan failed for value %v for type relationship. Details: %v", value, errRecover)
		}
	}()

	switch data := value.(type) {
	case []uint8:
		str := string([]byte(data))

		rtshp, err := r.tryParseValueToRelationship(str)

		if err != nil {
			return err
		}

		*r = Relationship(rtshp)

	case int:
		d := int64(data)
		*r = Relationship(d)
		_ = r.String()
	case int32:
		d := int64(data)
		*r = Relationship(d)
		_ = r.String()
	case float32:
		d := int64(data)
		*r = Relationship(d)
		_ = r.String()
	case float64:
		d := int64(data)
		*r = Relationship(d)
		_ = r.String()
	case int64:
		*r = Relationship(data)
		_ = r.String()
	case string:
		rtshp, err := r.tryParseValueToRelationship(data)
		if err != nil {
			panic(err)
		}
		*r = rtshp
	default:
		_ = r.String()
	}

	return nil
}

func (r *Relationship) tryParseValueToRelationship(value string) (rtshp Relationship, err error) {

	toTitle := cases.Title(language.BrazilianPortuguese, cases.NoLower)

	value = toTitle.String(strings.ToLower(value))

	rtshpINT, ok := relationship_value[value]

	if !ok {
		valueINT, err := strconv.Atoi(value)
		if err != nil {
			return Undefined, fmt.Errorf("the %s is incorret to type of relationship", value)
		}

		rtshpStr, ok := relationship_name[int64(valueINT)]
		if !ok {
			return Undefined, fmt.Errorf("the %s not valid type of relationship", value)
		}

		rtshpINT, ok = relationship_value[rtshpStr]
		if !ok {
			return Undefined, fmt.Errorf("the %s is invalid type of relationship", value)
		}
	}

	return Relationship(rtshpINT), nil
}

func (r *Relationship) tryGetValueFromJSON(bytes []byte) (value string, err error) {
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
