package settings

import (
	"fmt"

	"movienotes/models"
)

func GetNotes(keyId string) (note []*models.Note) {
	if !IsKey(keyId) {
		return
	}
	row, err := DataBase.Query(`SELECT * FROM notes_movies WHERE key_id = $1;`, keyId)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	defer row.Close()
	id := 0
	name := ""
	notestr := ""
	rating := 1
	keyIdNote := ""
	for row.Next() {
		err = row.Scan(&id, &name, &notestr, &rating, &keyIdNote)
		if err != nil {
			fmt.Println("ERROR: ", err)
			return
		}
		note = append(note, &models.Note{
			Id:     id,
			Name:   name,
			Note:   notestr,
			Rating: rating,
			KeyId:  keyIdNote,
		})
	}
	return
}

func GetNote(id int) (note models.Note) {
	row, err := DataBase.Query(`SELECT * FROM notes_movies WHERE id = $1;`, id)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	defer row.Close()
	idNote := 0
	name := ""
	notestr := ""
	rating := 1
	keyIdNote := ""
	for row.Next() {
		err = row.Scan(&idNote, &name, &notestr, &rating, &keyIdNote)
		if err != nil {
			fmt.Println("ERROR: ", err)
			return
		}
		note = models.Note{
			Id:     idNote,
			Name:   name,
			Note:   notestr,
			Rating: rating,
			KeyId:  keyIdNote,
		}
	}
	return
}

func IsKeyNote(key string, idNote int) bool {
	row, err := DataBase.Query(`SELECT id FROM notes_movies WHERE id = $1 AND key_id = $2;`, idNote, key)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return false
	}
	defer row.Close()
	id := -1
	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			fmt.Println("ERROR: ", err)
			return false
		}
	}
	if id == -1 {
		return false
	} else {
		return true
	}
}

func DeleteNote(key string, idNote int) error {
	if !IsKeyNote(key, idNote) {
		return fmt.Errorf("Not your key")
	}
	_, err := DataBase.Exec(`DELETE FROM notes_movies WHERE id = $1 AND key_id = $2;`, idNote, key)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return err
	}
	return nil
}

func EditNote(note models.Note) error {
	if !IsKeyNote(note.KeyId, note.Id) {
		return fmt.Errorf("Not your key")
	}
	_, err := DataBase.Exec(`UPDATE notes_movies SET name = $1, note = $2, rating = $3 WHERE id = $4;`, note.Name, note.Note, note.Rating, note.Id)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return err
	}
	return nil
}

func AddNote(note models.AddNote) (err error, id int) {
	if !IsKey(note.KeyId) {
		return fmt.Errorf("Not your key"), -1
	}
	row, err := DataBase.Query(`INSERT INTO notes_movies (name, note, rating, key_id) 
	VALUES ($1, $2, $3, $4) 
	RETURNING id;`, note.Name, note.Note, note.Rating, note.KeyId)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return err, -1
	}
	defer row.Close()
	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			fmt.Println("ERROR: ", err)
			return err, -1
		}
	}
	return
}
