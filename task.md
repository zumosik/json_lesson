Хоккейный чемпионат - результаты игр. Данные приводятся в трёх файлах: текстовый файл `hockey.dat`, csv-файл `hockey.csv` и json-файл `hockey.json`. В dat-файле и csv-файле хранится одна и та же информация - по одной строке строке на игру, а в json файле хранится слайс []TMatch, тип TMatch специфицирован ниже.

++Формат строки текстового файла:++
	- месяц/число/год 
	- название_команды_хозяев:голы_команды_хозяев
	- название_команды_гостей:голы_команды_гостей
	- если в игре было дополнительное время, то *, иначе ничего

Все четыре части хранятся в одной строке и разделены запятыми. Если в игре не было дополнительного времени, то запятая после `голов_команды_гостей` не ставится. Две строки-примера поясняют формат:

	10/1/1999,Ventspils Masti:4,Ventspils Trubas:6
	10/1/1999,Pociema Peles:1,Blomu Bites:1,*

++Формат строки csv-файла++:

`месяц, число, год, название_команды_хозяев, голы_команды_хозяев, название_команды_гостей, голы_команды_гостей, возможно * - признак_дополнительного времени`

Все поля хранятся в одной строке и разделены запятыми. Если в игре не было дополнительного времени, то запятая после `голов_команды_гостей` всё равно ставится ставится. Две строки-примера поясняют формат:

	10,1,1999,Ventspils Masti,4,Ventspils Trubas,6,
	10,1,1999,Pociema Peles,1,Blomu Bites,1,*

++Формат json-файла ++ определяется типом []tMatch: 

```go
type  ( 
	tDate struct  {
		Year uint16 `json:"year"` 
		Month byte `json:"month"`
		Day byte `json:"day"` 
	}

	tTeam struct  {
		Title string
		Goals byte
	}

	tMatch struct  {
		Date tDate
		Host tTeam `json:"Host team"`
		Guest tTeam `json:"Guest team"`
		Overtime bool
	}
)
```

Примеры фрагментов json-файла:

	{"Date":{"year":1999,"month":10,"day":1},"Host team":{"Title":"Ventspils Masti","Goals":4},"Guest team":{"Title":"Ventspils Trubas","Goals":6},"Overtime":false}

	{"Date":{"year":1999,"month":10,"day":1},"Host team":{"Title":"Pociema Peles","Goals":1},"Guest team":{"Title":"Blomu Bites","Goals":1},"Overtime":true}

**Возможные вопросы, на которые надо дать ответ:**

A)	Сколько команд участвовало в чемпионате?
	Cik komandas piedalījās čempionātā?

B)	Какие три команды набрали в чемпионате больше всего очков? Сколько?
	Kuras trīs komandas čempionātā izcīnīja visvairāk punktus? Cik?

C) Какие три команды, играя дома, набрали в чемпионате больше всего очков? Сколько?
Kuras trīs komandas, spēlējot mājās, izcīnīja visvairāk punktus? Cik?

D) Какие три команды, играя в гостях, набрали в чемпионате больше всего очков? Сколько?
Kuras trīs komandas, spēlējot viesos, izcīnīja visvairāk punktus? Cik?

E) Какая команда, играя дома, провела больше всего "сухих" игр (не пропустила ни одного гола)? Сколько?
Kura komanda, spēlējot mājās, aizvadījusi visvairāk sauso spēļu? Cik?

F) В каких двух играх разница забитых и пропущенных голов была наибольшей? Какой? Когда состоялись эти игры?
Kurās divās spēlēs gūto un zaudēto vārtu starpība bija vislielākā? Kāda? Kad šīs spēles notika?

G) Какая команда забила больше всего голов по воскресеньям? Сколько? 
Kura no komandām svētdienās guvusi visvairāk vārtus? Cik?

H) Какая команда набрала больше всего очков по 13-м числам? Сколько?
Kura komanda 13.datumos kopā guvusi visvairāk punktus? Cik?

I) Какая команда и против какой забила больше всего голов одной команде? Сколько?
Kura komanda, spēlējot pret kuru, čempionāta laikā guvusi visvairāk vārtus? Cik?

J) Какая команда сделала больше всего ничьих, играя в гостях? Сколько? 
Kura komanda, spēlējot viesos, visvairāk reizes nospēlējusi neizšķirti? Cik?

K) Какая команда, играя дома, по понедельникам выигшрала больше всего игр в дополнитьельное время? Сколько?
Kura komanda pirmdienās, spēlējot mājās, papildlaikā izcīnījusi visvairāk uzvaru? Cik?

L) Какой команде принадлежит самая длинная серия игр без проигрышей?
Kurai komandai čempionāta laikā izdevies nezaudēt visvairāk pēc kārtas nospēlētās spēlēs?

