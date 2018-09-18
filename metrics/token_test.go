package metrics

import (
	"testing"
)

/*
	It is necessary to test:

	Numbers:
	- 1 
	- 2 
	- >2

	The length of the term and doc (number of words):
	- 1 
	- 2 
	- >2

	Lang.:
	- One (rus || eng)
	- Mix (array || string)

	Case:
	- Upper
	- Lower
	- Mix

	Others:
	- Hyphen
	- Contractions (including .*'s)
	- Mixing letters
	- Double (all)
*/

var tokenTests = [][]string{
	[]string{"НЕЙТРИНО"}, // Numbers(1), Words(1), Language(One(rus)), Case(Upper)
	[]string{"таким", "образом"}, // Numbers(2), Words(1), Language(One(rus)), Case(Lower)
	[]string{"в его квартире", "номер тридцать пять", "в вентиляции в", "уборной в газетной", "бумаге четыреста долларов"}, // Numbers(>2), Words(>2), Language(One(rus)), Case(Lower)
	[]string{"IN", "FACT"}, // Numbers(2), Words(1), Language(One(eng)), Case(Upper)
	[]string{"ЖУРНАЛЕ CONCORDIA"}, // Numbers(1), Words(2), Language(One(mix)), Case(Upper)
	[]string{"А фрекен", "Бок и", "дядя Юлиус", "отправились в", "гостиную чтобы", "как обычно", "выпить кофе", "с глазу", "на глаз", "Дядя Юлиус", "и фрекен", "Бок обычно", "сидели на", "маленьком диванчике", "который прекрасно", "было видно", "в глазок"}, // Numbers(>2), Words(2), Language(One(rus)), Case(MIX)
	[]string{"нужно", "было", "добежать", "до", "ближайшего", "телефона-автомата", "и", "сообщить", "в", "бюро", "иностранцев", "о", "том", "что", "вот", "мол", "приезжий", "из-за", "границы", "консультант"},  // Numbers(>2), Words(1), Language(One(rus)), Case(Lower), Hyphen
	[]string{"his watch-face", "appeared before", "Rimsky's eyes"}, // Numbers(>2), Words(2), Language(One(eng)), Case(Mix), Hyphen, Contractions(.*'s)
	[]string{"The contraction aren't is used", "in standard English to mean", "am not in questions as", "in I'm right aren't I"}, // Numbers(>2), Words(>2), Language(One(eng)), Case(Mix), Contractions
	[]string{"gentlemаn"}, // Numbers(1), Words(1), Language(One(eng)), Case(Lower), Mixing letters("а")
	[]string{"А", "сeйчас"}, // Numbers(2), Words(1), Language(One(rus)), Case(Mix), Mixing letters("e")
	[]string{"Карлсон сказал Малыш", "Карлсон сказал Малыш", "Карлсон сказал Малыш"}, // Numbers(>2), Words(>2), Language(One(rus)), Case(Mix), Double (all)
}

func TestTF(t *testing.T) {
	actual, err := TF("", tokenTests[0])
	if err == nil {
		t.Errorf("cannot calculate TF: %v", err)
	}
	actual, err = TF("слово", []string{})
	if err == nil {
		t.Errorf("cannot calculate TF: %v", err)
	}

	expected := []map[string]float64{
		map[string]float64{"НЕЙТРИНО": 1.0, "нейтрино": 0.0, "НЕЙТРИНО ЧАСТИЦА": 0.0, "ЭЛЕМЕНТАРНАЯ ЧАСТИЦА": 0.0, "НEЙТРИНО": 0.0, " НЕЙТРИНО": 0.0},
		map[string]float64{"таким": 0.5, "образом": 0.5, "внесенные таким образом": 0.0, "Таким": 0.0, "ОБРАЗОМ": 0.0, "тaким": 0.0, "вскоре": 0.0, "образом ": 0.0},
		map[string]float64{"в его квартире": 0.2, "номер тридцать пять": 0.2, "в вентиляции в": 0.2, "уборной в газетной": 0.2, "бумаге четыреста долларов": 0.2, "БУМАГЕ ЧЕТЫРЕСТА ДОЛЛАРОВ": 0.0, "в eго квартире": 0.0, "в  вентиляции в ": 0.0},
		map[string]float64{"IN": 0.5, "FACT": 0.5, "IN FACT REVOLUTIONISED": 0.0, "In": 0.0, "fact": 0.0, "fаct": 0.0, "капиталист": 0.0, "IN ": 0.0},
		map[string]float64{"ЖУРНАЛЕ CONCORDIA": 1.0, "журнале Concordia": 0.0, "CONCORDIA": 0.0, "ЖУРНАЛЕ CONCORDIА": 0.0, "CONCORDIA ЖУРНАЛЕ": 0.0, "ЖУРНАЛЕ  CONCORDIA": 0.0},
		map[string]float64{"А фрекен": 0.058823529411764705, "Бок и": 0.058823529411764705, "дядя Юлиус": 0.058823529411764705, "отправились в": 0.058823529411764705, "гостиную чтобы": 0.058823529411764705, "как обычно": 0.058823529411764705, "выпить кофе": 0.058823529411764705, "с глазу": 0.058823529411764705, "на глаз": 0.058823529411764705, "Дядя Юлиус": 0.058823529411764705, "и фрекен": 0.058823529411764705, "Бок обычно": 0.058823529411764705, "сидели на": 0.058823529411764705, "маленьком диванчике": 0.058823529411764705, "который прекрасно": 0.058823529411764705, "было видно": 0.058823529411764705, "в глазок": 0.058823529411764705, "ОТПРАВИЛИСЬ В": 0.0, "сидели нa": 0.0, "в глазок ": 0.0},
		map[string]float64{"нужно":0.05, "было":0.05, "добежать":0.05, "до":0.05, "ближайшего":0.05, "телефона-автомата":0.05, "и":0.05, "сообщить":0.05, "в":0.05, "бюро":0.05, "иностранцев":0.05, "о":0.05, "том":0.05, "что":0.05, "вот":0.05, "мол":0.05, "приезжий":0.05, "из-за":0.05, "границы":0.05, "консультант":0.05, "НУЖНО": 0.0, "телефона автомата": 0.0, "из за": 0.0, "вот ": 0.0},
		map[string]float64{"his watch-face": 0.3333333333333333, "appeared before": 0.3333333333333333, "Rimsky's eyes": 0.3333333333333333, "HIS WATCH-FACE": 0.0, "his watch face": 0.0, "Rimskys eyes": 0.0, "appeared  before": 0.0},
		map[string]float64{"The contraction aren't is used": 0.25, "in standard English to mean": 0.25, "am not in questions as": 0.25, "in I'm right aren't I": 0.25, "THE CONTRACTION AREN'T IS USED": 0.0, "The contraction are not is used": 0.0, "in I am right aren't I": 0.0, "in I'm right are not I": 0.0, "in I am right are not I": 0.0, "in  I'm right  aren't I": 0.0},
		map[string]float64{"gentlemаn": 1.0, " if a gentlemаn walks into my rooms smelling of iodoform": 0.0, "Gentlemаn": 0.0, "gentleman": 0.0, "gentlemаn ": 0.0},
		map[string]float64{"А": 0.5, "сeйчас": 0.5, "сейчас": 0.0, "A": 0.0, "сeйчас ": 0.0, "Wagner": 0.0},
		map[string]float64{"Карлсон сказал Малыш": 1.0, "КАРЛСОН СКАЗАЛ МАЛЫШ": 0.0, "Кaрлсон сказал Малыш": 0.0, "Карлсон сказал  Малыш": 0.0},
	}

	for i, tokens := range tokenTests {
		for token, tf := range expected[i] {
			actual, _ = TF(token, tokens)
			if actual != tf {
				t.Errorf("cannot calculate TF:\ntoken:\t%s\ntokens:\t%#v\ngot:\t%#v\nwant:\t%#v\n\n", token, tokens, actual, tf)
			}
		}
	}
}	





















