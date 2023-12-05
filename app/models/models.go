package models

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

	Data         map[string]Amounts
	DataTwo      map[string]AmountsTwo
	CommandTable map[AmountsTwo]uint
)

func (d DataTwo) ToCommandTable() CommandTable {
	t := make(CommandTable)
	var i uint
	for _, a := range d {
		t[a] = i
	}
	return t
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
