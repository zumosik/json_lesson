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

	dTmp := AnalyseData(matches)

	// Creating file
	fout, err := os.Create(out)
	if err != nil {
		log.Fatalf("cant create file (%s): %v", out, err)
	}
	defer fout.Close()
	err = json.NewEncoder(fout).Encode(dTmp)
	if err != nil {
		log.Fatalf("cant encode json to file (%s): %v", out, err)
	}

	// Reading it
	fin, err := os.Open(out)
	if err != nil {
		log.Fatalf("cant open file (%s): %v", out, err)
	}
	defer fin.Close()

	var data models.Data
	err = json.NewDecoder(fin).Decode(&data)
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

	log.Println(data)
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

func AnalyseData(matches []models.Match) models.Data {
	d := make(models.Data)
	for _, m := range matches {
		h := m.Host
		g := m.Guest
		dh := d[h.Title]
		dg := d[g.Title]

		dh.Missed += int(g.Goals)
		dg.Missed += int(h.Goals)

		dh.Goals += int(h.Goals)
		dg.Goals += int(g.Goals)

		if h.Goals > g.Goals {
			if m.Overtime == true {
				dh.WinsInOvertime++
				dg.LosesInOvertime++
			} else {
				dh.WinsInMainTime++
				dg.LosesInMainTime++
			}
		} else if h.Goals < g.Goals {
			if m.Overtime == true {
				dg.WinsInOvertime++
				dh.LosesInOvertime++
			} else {
				dg.WinsInMainTime++
				dh.LosesInMainTime++
			}
		} else { // h Goals == g Goals
			dg.Draw++
			dh.Draw++
		}

		d[h.Title] = dh
		d[g.Title] = dg

	}

	return d
}
