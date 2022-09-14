package messages

import (
	"encoding/json"
	"log"
	"reflect"
	"strconv"

	flatbuffers "github.com/google/flatbuffers/go"
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

func (m *message) toFbf() []byte {
	builder := flatbuffers.NewBuilder(0)

	str0 := builder.CreateString(m.Str0)
	str1 := builder.CreateString(m.Str1)
	str2 := builder.CreateString(m.Str2)
	str3 := builder.CreateString(m.Str3)
	str4 := builder.CreateString(m.Str4)
	str5 := builder.CreateString(m.Str5)
	str6 := builder.CreateString(m.Str6)
	str7 := builder.CreateString(m.Str7)
	str8 := builder.CreateString(m.Str8)
	str9 := builder.CreateString(m.Str9)
	str10 := builder.CreateString(m.Str10)
	str11 := builder.CreateString(m.Str11)
	str12 := builder.CreateString(m.Str12)
	str13 := builder.CreateString(m.Str13)
	str14 := builder.CreateString(m.Str14)
	str15 := builder.CreateString(m.Str15)
	str16 := builder.CreateString(m.Str16)
	str17 := builder.CreateString(m.Str17)
	str18 := builder.CreateString(m.Str18)
	str19 := builder.CreateString(m.Str19)
	str20 := builder.CreateString(m.Str20)
	str21 := builder.CreateString(m.Str21)
	str22 := builder.CreateString(m.Str22)
	str23 := builder.CreateString(m.Str23)
	str24 := builder.CreateString(m.Str24)
	str25 := builder.CreateString(m.Str25)
	str26 := builder.CreateString(m.Str26)
	str27 := builder.CreateString(m.Str27)
	str28 := builder.CreateString(m.Str28)
	str29 := builder.CreateString(m.Str29)
	str30 := builder.CreateString(m.Str30)
	str31 := builder.CreateString(m.Str31)

	InstrumentStart(builder)

	InstrumentAddInt0(builder, m.Int0)
	InstrumentAddInt1(builder, m.Int1)
	InstrumentAddInt2(builder, m.Int2)
	InstrumentAddInt3(builder, m.Int3)
	InstrumentAddInt4(builder, m.Int4)
	InstrumentAddInt5(builder, m.Int5)

	InstrumentAddDouble0(builder, m.Double0)
	InstrumentAddDouble1(builder, m.Double1)
	InstrumentAddDouble2(builder, m.Double2)
	InstrumentAddDouble3(builder, m.Double3)
	InstrumentAddDouble4(builder, m.Double4)
	InstrumentAddDouble5(builder, m.Double5)

	InstrumentAddStr0(builder, str0)
	InstrumentAddStr1(builder, str1)
	InstrumentAddStr2(builder, str2)
	InstrumentAddStr3(builder, str3)
	InstrumentAddStr4(builder, str4)
	InstrumentAddStr5(builder, str5)
	InstrumentAddStr6(builder, str6)
	InstrumentAddStr7(builder, str7)
	InstrumentAddStr8(builder, str8)
	InstrumentAddStr9(builder, str9)
	InstrumentAddStr10(builder, str10)
	InstrumentAddStr11(builder, str11)
	InstrumentAddStr12(builder, str12)
	InstrumentAddStr13(builder, str13)
	InstrumentAddStr14(builder, str14)
	InstrumentAddStr15(builder, str15)
	InstrumentAddStr16(builder, str16)
	InstrumentAddStr17(builder, str17)
	InstrumentAddStr18(builder, str18)
	InstrumentAddStr19(builder, str19)
	InstrumentAddStr20(builder, str20)
	InstrumentAddStr21(builder, str21)
	InstrumentAddStr22(builder, str22)
	InstrumentAddStr23(builder, str23)
	InstrumentAddStr24(builder, str24)
	InstrumentAddStr25(builder, str25)
	InstrumentAddStr26(builder, str26)
	InstrumentAddStr27(builder, str27)
	InstrumentAddStr28(builder, str28)
	InstrumentAddStr29(builder, str29)
	InstrumentAddStr30(builder, str30)
	InstrumentAddStr31(builder, str31)

	instrument := InstrumentEnd(builder)
	builder.Finish(instrument)

	return builder.FinishedBytes()
}

func (m *message) fbfSize() int {
	return len(m.toFbf())
}
