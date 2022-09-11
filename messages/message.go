package messages

import (
	"encoding/json"
	"log"
	"reflect"
	"strconv"
)

type message struct {
	Int0 int64 `json:"int0"`
	Int1 int64 `json:"int1"`
	Int2 int64 `json:"int2"`
	Int3 int64 `json:"int3"`
	Int4 int64 `json:"int4"`
	Int5 int64 `json:"int5"`

	Double0 float64 `json:"double0"`
	Double1 float64 `json:"double1"`
	Double2 float64 `json:"double2"`
	Double3 float64 `json:"double3"`
	Double4 float64 `json:"double4"`
	Double5 float64 `json:"double5"`

	Str0  string `json:"str0"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
	Str3  string `json:"str3"`
	Str4  string `json:"str4"`
	Str5  string `json:"str5"`
	Str6  string `json:"str6"`
	Str7  string `json:"str7"`
	Str8  string `json:"str8"`
	Str9  string `json:"str9"`
	Str10 string `json:"str10"`
	Str11 string `json:"str11"`
	Str12 string `json:"str12"`
	Str13 string `json:"str13"`
	Str14 string `json:"str14"`
	Str15 string `json:"str15"`
	Str16 string `json:"str16"`
	Str17 string `json:"str17"`
	Str18 string `json:"str18"`
	Str19 string `json:"str19"`
	Str20 string `json:"str20"`
	Str21 string `json:"str21"`
	Str22 string `json:"str22"`
	Str23 string `json:"str23"`
	Str24 string `json:"str24"`
	Str25 string `json:"str25"`
	Str26 string `json:"str26"`
	Str27 string `json:"str27"`
	Str28 string `json:"str28"`
	Str29 string `json:"str29"`
	Str30 string `json:"str30"`
	Str31 string `json:"str31"`
}

func createMessage() *message {
	return &message{
		Int0:    0,
		Int1:    0,
		Int2:    0,
		Int3:    0,
		Int4:    0,
		Int5:    0,
		Double0: 0.0,
		Double1: 0.0,
		Double2: 0.0,
		Double3: 0.0,
		Double4: 0.0,
		Double5: 0.0,
		Str0:    "a",
		Str1:    "b",
		Str2:    "c",
		Str3:    "d",
		Str4:    "e",
		Str5:    "f",
		Str6:    "g",
		Str7:    "h",
		Str8:    "i",
		Str9:    "j",
		Str10:   "k",
		Str11:   "l",
		Str12:   "m",
		Str13:   "n",
		Str14:   "o",
		Str15:   "p",
		Str16:   "q",
		Str17:   "r",
		Str18:   "s",
		Str19:   "t",
		Str20:   "u",
		Str21:   "v",
		Str22:   "w",
		Str23:   "z",
		Str24:   "y",
		Str25:   "z",
		Str26:   "A",
		Str27:   "B",
		Str28:   "C",
		Str29:   "D",
		Str30:   "E",
		Str31:   "F",
	}
}

func (m *message) increaseSize(delta uint) *message {
	// prepare deltas for every string field
	deltaPerField := make([]uint, 32)
	minDelta := delta / 32
	for i, _ := range deltaPerField {
		deltaPerField[i] = minDelta
		delta -= minDelta
	}
	for i := uint(0); i < delta; i++ {
		deltaPerField[i] += 1
	}

	// add symbols to every string field
	reflectValue := reflect.ValueOf(m)
	for i, delta := range deltaPerField {
		reflectField := reflect.Indirect(reflectValue).FieldByName("Str" + strconv.Itoa(i))

		fieldValue := []byte(reflectField.String())
		deltaValue := make([]byte, delta)
		deltaSymbol := fieldValue[len(fieldValue)-1]
		for j, _ := range deltaValue {
			deltaValue[j] = deltaSymbol
		}

		reflectField.SetString(string(fieldValue) + string(deltaValue))
	}

	return m
}

func (m *message) toJson() []byte {
	result, err := json.Marshal(m)
	if err != nil {
		log.Printf("Error marshalling into JSON: %v\n", err)
	}
	return result
}

func (m *message) jsonSize() int {
	return len(m.toJson())
}
