package tictacgo

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestMapDBCanBeBuilt(t *testing.T) {
	board := NewBoard()
	var emptyBoardData = make(map[string]float64)
	for i := 0; i < 9; i++ {
		emptyBoardData[fmt.Sprint(i)] = 0.0
	}
	data := make(map[string]map[string]float64)
	data[board.String()] = emptyBoardData
	db := MapDB{source: "some/where", data: data}

	if db.data == nil {
		t.Errorf("Data did not load into DB")
	}
}

func TestMapDBCanLoadJson(t *testing.T) {
	tests := []struct {
		file       string
		shouldFail bool
		expected   map[string]float64
	}{
		{
			file:       "test/state_values.json",
			shouldFail: false,
			expected: map[string]float64{
				"4": 1.0,
				"5": 0.0,
				"6": -1.0,
				"7": 0.0,
				"8": 1.0,
			},
		},
		{
			file:       "wrong/file.json",
			shouldFail: true,
			expected:   make(map[string]float64),
		},
	}

	for _, test := range tests {
		db := NewMapDB(test.file)
		err := db.read()
		if err != nil {
			if !test.shouldFail {
				t.Errorf("File not found!")
			}
		} else {
			board := Board{state: [9]int{1, 2, 1, 2, 0, 0, 0, 0, 0}}
			value := db.get(board)
			if !reflect.DeepEqual(value, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, value)
			}
		}

	}
}

func TestMapDbExists(t *testing.T) {
	db := NewMapDB("some/path")
	board := NewBoard()
	db.update(board, map[string]float64{"4": 0.0, "5": 1.0})

	if !db.exists(board) {
		t.Error("Could not find a board that should exist")
	}

	board.makeMove(5)
	if db.exists(board) {
		t.Error("Returned true for a board that doesn't exist")
	}

}

func TestMapDbCommit(t *testing.T) {
	db := NewMapDB("test/state_values.json")
	err := db.read()
	if err != nil {
		t.Error("Could not read JSON file")
		return
	}

	db.source = "test/new_state_values.json"
	board := NewBoard()
	db.update(board, map[string]float64{"4": 0.0, "5": 1.0})
	db.commit()

	anotherDb := NewMapDB(db.source)
	err = anotherDb.read()
	if err != nil {
		t.Error("New JSON file failed to be read")
		return
	}

	if !reflect.DeepEqual(db, anotherDb) {
		t.Error("Inconsistency when reading saved DB")
	}

	os.Remove(db.source)
}
