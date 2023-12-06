package models

import (
	"log"
	"strconv"
)

type (
	Date struct {
		Year  uint16 `json:"year"`
		Month byte   `json:"month"`
		Day   byte   `json:"day"`
	}

	Team struct {
		Title string
		Goals byte
	}

	Match struct {
		Date     Date
		Host     Team
		Guest    Team
		Overtime bool
	}

	ToJSONstruct struct {
		Data    Data    `json:"data"`
		DataTwo DataTwo `json:"data_two"`
	}

	Amounts struct {
		WinsInOvertime  int `json:"wins_in_overtime"`
		WinsInMainTime  int `json:"wins_in_maintime"`
		LosesInOvertime int `json:"loses_in_overtime"`
		LosesInMainTime int `json:"loses_in_maintime"`
		Draw            int `json:"draws"`
		Goals           int `json:"goals"`
		Missed          int `json:"missed"`
	}

	FieldHostGuest struct {
		Home  int
		Guest int
	}

	AmountsTwo struct {
		WinsInOvertime  FieldHostGuest `json:"wins_in_overtime"`
		WinsInMainTime  FieldHostGuest `json:"wins_in_maintime"`
		LosesInOvertime FieldHostGuest `json:"loses_in_overtime"`
		LosesInMainTime FieldHostGuest `json:"loses_in_maintime"`
		Draw            FieldHostGuest `json:"draws"`
		Goals           FieldHostGuest `json:"goals"`
		Missed          FieldHostGuest `json:"missed"`
	}

	Data           map[string]Amounts
	DataTwo        map[string]AmountsTwo
	CommandIndexes map[Team]uint
	CommandTable   [][]int
)

func (i CommandIndexes) ToCommandTable(m []Match) CommandTable {
	tbl := make(CommandTable, len(m))
	for _, mm := range m {
		if len(tbl[i[mm.Host]]) == 0 {
			// log.Println("empty")
			tbl[i[mm.Host]] = make([]int, len(m))
		}
		// if mm.Host.Goals == 0 {
		// 	log.Printf("%s: %d", mm.Host.Title, 0)
		// }
		tbl[i[mm.Host]][i[mm.Guest]] = int(mm.Host.Goals)
	}
	return tbl
}

func (tbl CommandTable) Print() {
	// log.Println("table")
	for _, l := range tbl {
		var t string
		if len(l) < 1 {
			continue
		}
		for _, ll := range l {
			t += strconv.Itoa(ll) + ","
		}
		log.Println(t)
	}
}

func (a *AmountsTwo) ToAmounts() Amounts {
	return Amounts{
		WinsInOvertime:  a.WinsInOvertime.Guest + a.WinsInOvertime.Home,
		WinsInMainTime:  a.WinsInMainTime.Guest + a.WinsInMainTime.Home,
		LosesInOvertime: a.LosesInOvertime.Guest + a.LosesInOvertime.Home,
		LosesInMainTime: a.LosesInMainTime.Guest + a.LosesInMainTime.Home,
		Draw:            a.Draw.Guest + a.Draw.Home,
		Goals:           a.Goals.Guest + a.Goals.Home,
		Missed:          a.Missed.Guest + a.Missed.Home,
	}
}
