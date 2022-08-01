package gender

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Gender int64

const (
	Undefined Gender = iota // 0 - Indefinido - Gênero não definido
	Masculine               // 1 - Masculino
	Female                  // 2 - Feminino

	name_undefined = "Undefined"
	name_masculine = "Masculine"
	name_female    = "Female"
)

var (
	gender_name = map[int64]string{
		0: name_undefined,
		1: name_masculine,
		2: name_female,
	}
	gender_value = map[string]int64{
		name_undefined: 0,
		name_masculine: 1,
		name_female:    2,
	}

	typeAccepts = func() string {

		descriptionTypeAccepts := "Types accepts: "

		for enum, types := range gender_name {
			descriptionTypeAccepts += fmt.Sprintf("%d ou %s | ", enum, types)
		}

		return descriptionTypeAccepts
	}
)

func (g Gender) String() string {
	switch g {
	case Undefined:
		return name_undefined
	case Masculine:
		return name_masculine
	case Female:
		return name_female
	default:
		panic(fmt.Sprintf("type gender is invalid. %s", typeAccepts()))
	}
}

func (g Gender) MarshalJSON() ([]byte, error) {

	defer func() {
		errRecover := recover()

		if errRecover != nil {
			panic(fmt.Sprintf("Marshal failed. Type gender informed: %d. Failed details: %v", g, errRecover))
		}
	}()

	return []byte(fmt.Sprintf(`"%s"`, g.String())), nil
}

func (g *Gender) UnmarshalJSON(bytes []byte) error {
	value, err := g.tryGetValueFromJSON(bytes)
	if err == nil && !strings.EqualFold(value, "") {

		gndr, err := g.tryParseValueToGender(value)

		if err != nil {
			return err
		}

		*g = Gender(gndr)
	}

	return err
}

func (g Gender) Value() (driver.Value, error) {

	defer func() {
		errRecover := recover()

		if errRecover != nil {
			panic(fmt.Sprintf("Value failed. Type gender informed: %d. Failed details: %v", g, errRecover))
		}
	}()

	return g.String(), nil
}

func (g *Gender) Scan(value interface{}) (err error) {

	defer func() {
		if errRecover := recover(); errRecover != nil {
			err = fmt.Errorf("Scan failed for value %v for type gender. Details: %v", value, errRecover)
		}
	}()

	switch data := value.(type) {
	case []uint8:
		str := string([]byte(data))

		gndr, err := g.tryParseValueToGender(str)

		if err != nil {
			return err
		}

		*g = Gender(gndr)

	case int:
		d := int64(data)
		*g = Gender(d)
		_ = g.String()
	case int32:
		d := int64(data)
		*g = Gender(d)
		_ = g.String()
	case float32:
		d := int64(data)
		*g = Gender(d)
		_ = g.String()
	case float64:
		d := int64(data)
		*g = Gender(d)
		_ = g.String()
	case int64:
		*g = Gender(data)
		_ = g.String()
	case string:
		gndr, err := g.tryParseValueToGender(data)
		if err != nil {
			panic(err)
		}
		*g = gndr
	default:
		_ = g.String()
	}

	return nil
}

func (g *Gender) tryParseValueToGender(value string) (gndr Gender, err error) {

	toTitle := cases.Title(language.BrazilianPortuguese, cases.NoLower)

	value = toTitle.String(strings.ToLower(value))

	gndrINT, ok := gender_value[value]

	if !ok {
		valueINT, err := strconv.Atoi(value)
		if err != nil {
			return Undefined, fmt.Errorf("the %s is incorret to type of gender", value)
		}

		gndrStr, ok := gender_name[int64(valueINT)]
		if !ok {
			return Undefined, fmt.Errorf("the %s not valid type of gender", value)
		}

		gndrINT, ok = gender_value[gndrStr]
		if !ok {
			return Undefined, fmt.Errorf("the %s is invalid type of gender", value)
		}
	}

	return Gender(gndrINT), nil
}

func (g *Gender) tryGetValueFromJSON(bytes []byte) (value string, err error) {
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
