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

	Data    map[string]Amounts
	DataTwo map[string]AmountsTwo
)
