package prisma

import (
	"bytes"

	"github.com/ianlopshire/go-fixedwidth"
	"golang.org/x/text/encoding/charmap"
)

//Prisma represent prisma
type Prisma struct {
	charmap *charmap.Charmap
}

//DefaultPrisma prisma with default char encoding
func DefaultPrisma() *Prisma {
	return &Prisma{charmap.CodePage866}
}

//Decode decode fixed width message
func (p *Prisma) Decode(data []byte) (*Message, error) {
	var d *fixedwidth.Decoder

	if p.charmap != nil {
		cd := p.charmap.NewDecoder()
		b, err := cd.Bytes(data)
		if err != nil {
			return nil, err
		}
		d = fixedwidth.NewDecoder(bytes.NewReader(b))
		d.SetUseCodepointIndices(true)
	} else {
		d = fixedwidth.NewDecoder(bytes.NewReader(data))
	}

	var m Message

	//err := fixedwidth.Unmarshal(buf, &m)
	if err := d.Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

//Ecode ecode fixed width message
func (p *Prisma) Ecode(m *Message) ([]byte, error) {
	return fixedwidth.Marshal(m)
	/*
		data, err := fixedwidth.Marshal(m)
		if err != nil || p.charmap == nil {
			return data, fmt.Errorf("Prisma.Marshal %s", err.Error())
		}
		// convert char encoding
		e := p.charmap.NewEncoder()
		d, err := e.Bytes(data)
		if err != nil {
			fmt.Println(string(data))
			err = fmt.Errorf("Prisma.Encode %s", err.Error())
		}
		return d, err
	*/
}

var cm *charmap.Charmap = charmap.CodePage866

//Pstring string to encode to cp866
type Pstring string

//MarshalText implements TextMarshaler
func (s Pstring) MarshalText() ([]byte, error) {
	return cm.NewEncoder().Bytes([]byte(s))
}
func (s Pstring) String() string {
	return string(s)
}

/*
func (s cstring) UnmarshalText() error {
	d := cm.NewDecoder()
	b, err := d.Bytes([]byte(s))
	if err != nil {
		return err
	}
	s = cstring(b)
	return nil
}
*/

//Message base prisma message
type Message struct {
	Prefix     string  `fixed:"1,3"`     //3 array[0..2] of char Символьный Префикс объекта (ККМ)
	Number     int     `fixed:"4,9"`     //6  array[0..5] of char Целый числовой Номер ККМ
	Mode       int     `fixed:"10,13"`   //4 array[0..3] of char Целый числовой Код события
	CassirItem string  `fixed:"14,33"`   //20 array[0..19] of char Символьный Код кассира (таб.		номер)
	Cassir     Pstring `fixed:"34,63"`   //30 array[0..29] of char Символьный Имя кассира
	CKNumber   string  `fixed:"64,73"`   //10 array[0..9] of char Целый числовой Номер чека
	Count      int     `fixed:"74,76"`   //3 array[0..2] of char Целый числовой Номер позиции в		чеке
	BarCode    string  `fixed:"77,89"`   //13 array[0..12] of char Символьный Штрих-код		товара
	GoodsItem  Pstring `fixed:"90,119"`  //30 array[0..29] of char Символьный Код товара
	GoodsName  Pstring `fixed:"120,149"` //30 array[0..29] of char Символьный Наименование		товара
	GoodsPrice float64 `fixed:"150,164"` //15 array[0..14] of char Числовой Цена товара
	GoodsQuant float64 `fixed:"165,179"` //15 array[0..14] of char Числовой Количество		товара
	GoodsSum   float64 `fixed:"180,194"` //15 array[0..14] of char Числовой Сумма по товарной позиции
	Sum        float64 `fixed:"195,209"` //15 array[0..14] of char Числовой Сумма по чеку
	CardType   string  `fixed:"210,212"` //3 array[0..2] of char Символьный Тип карты(дисконтная,кредитная)
	CardNumber Pstring `fixed:"213,232"` //20 array[0..19] of char Символьный Номер карты
	DiscStr    float64 `fixed:"233,247"` //15 array[0..14] of char Числовой Скидка по строкечека
	DiscSum    float64 `fixed:"248,262"` //15 array[0..14] of char Числовой Скидка по чеку
	Day        string  `fixed:"263,264"` //2 array[0..1] of char Числовой День
	Month      string  `fixed:"265,266"` //2 array[0..1] of char Числовой Месяц
	Year       string  `fixed:"267,268"` //2 array[0..1] of char Числовой Год
	Sec100     string  `fixed:"269,271"` //3 array[0..2] of char Числовой Милисекунды
	Sec        string  `fixed:"272,273"` //2 array[0..1] of char Числовой Секунды
	Min        string  `fixed:"274,275"` //2 array[0..1] of char Числовой Минуты
	Hour       string  `fixed:"276,277"` //2 array[0..1] of char Числовой Часы
}
