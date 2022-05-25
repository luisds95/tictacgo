package tictacgo

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Database interface {
	read()
	get(board Board) map[string]float64
	update(board Board, values map[string]float64)
	commit()
	exists(board Board) bool
}

type MapDB struct {
	data   map[string]map[string]float64
	source string
}

func NewMapDB(source string) MapDB {
	return MapDB{source: source, data: make(map[string]map[string]float64)}
}

func (db *MapDB) read() error {
	jsonFile, err := os.Open(db.source)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data map[string]map[string]float64
	json.Unmarshal([]byte(byteValue), &data)

	db.data = data

	return nil
}

func (db *MapDB) get(board Board) map[string]float64 {

	return db.data[board.String()]
}

func (db *MapDB) update(board Board, values map[string]float64) {
	db.data[board.String()] = values
}

func (db *MapDB) commit() {
	file, _ := json.Marshal(db.data)
	_ = os.WriteFile(db.source, file, 0644)
}

func (db *MapDB) exists(board Board) bool {
	_, ok := db.data[board.String()]
	return ok
}
