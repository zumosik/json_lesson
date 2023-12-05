package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"hockey/models"
	"log"
	"os"
)

func main() {
	var inp, out string
	flag.StringVar(&inp, "input-file-path", "../hockey.json", "path to hockey file (json)")
	flag.StringVar(&out, "output-file-path", "../hockey_result.json", "path to hockey output file (json)")
	flag.Parse()

	matches, err := ReadJSON(inp)
	if err != nil {
		log.Fatalf("cant read json: %v", err)
	}

	dTmp, data2 := AnalyseData(matches)

	j := models.ToJSONstruct{
		Data:    dTmp,
		DataTwo: data2,
	}

	// Creating file
	fout, err := os.Create(out)
	if err != nil {
		log.Fatalf("cant create file (%s): %v", out, err)
	}
	defer fout.Close()
	err = json.NewEncoder(fout).Encode(j)
	if err != nil {
		log.Fatalf("cant encode json to file (%s): %v", out, err)
	}

	// Reading it
	fin, err := os.Open(out)
	if err != nil {
		log.Fatalf("cant open file (%s): %v", out, err)
	}
	defer fin.Close()

	var j2 models.ToJSONstruct
	err = json.NewDecoder(fin).Decode(&j2)
	if err != nil {
		log.Fatalf("cant decoded json from file (%s): %v", out, err)
	}

	// Deleting file
	err = os.Remove(out)
	if err != nil {
		/*
		 !!! this isnt fatal error !!!
		*/
		log.Printf("cant delete file (%s): %v", out, err)
	}

	dTmp = models.Data{}

	data := j2.DataTwo
	tbl := data.ToCommandTable()
	log.Println(tbl)
}

func ReadJSON(fileName string) ([]models.Match, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return []models.Match{}, err
	}
	defer f.Close()

	var res []models.Match
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		var m models.Match
		err := json.Unmarshal(scanner.Bytes(), &m)
		if err != nil {
			return []models.Match{}, err
		}
		res = append(res, m)
	}

	return res, nil

}

// TODO: models.DataTwo

func AnalyseData(matches []models.Match) (models.Data, models.DataTwo) {
	d := make(models.Data)
	d2 := make(models.DataTwo)
	for _, m := range matches {
		h := m.Host
		g := m.Guest

		d2h := d2[h.Title]
		d2g := d2[g.Title]

		d2h.Missed.Home += int(g.Goals)
		d2g.Missed.Guest += int(h.Goals)

		d2h.Goals.Home += int(h.Goals)
		d2g.Goals.Guest += int(g.Goals)

		if h.Goals > g.Goals {
			if m.Overtime == true {
				d2h.WinsInOvertime.Home++
				d2g.LosesInOvertime.Guest++
			} else {
				d2h.WinsInMainTime.Home++
				d2g.LosesInMainTime.Guest++
			}
		} else if h.Goals < g.Goals {
			if m.Overtime == true {
				d2g.WinsInOvertime.Guest++
				d2h.LosesInOvertime.Home++
			} else {
				d2g.WinsInMainTime.Guest++
				d2h.LosesInMainTime.Home++
			}
		} else { // h Goals == g Goals
			d2g.Draw.Guest++
			d2h.Draw.Guest++
		}

		d[h.Title] = d2h.ToAmounts()
		d[g.Title] = d2g.ToAmounts()

		d2[h.Title] = d2h
		d2[g.Title] = d2g
	}

	return d, d2
}
