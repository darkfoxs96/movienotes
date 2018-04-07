package settings

import (
	"fmt"
)

func IsKey(id string) bool {
	row, err := DataBase.Query(`SELECT id FROM keys WHERE id = $1;`, id)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return false
	}
	defer row.Close()
	key := ""
	for row.Next() {
		err = row.Scan(&key)
		if err != nil {
			fmt.Println("ERROR: ", err)
			return false
		}
	}
	if key == "" {
		return false
	} else {
		return true
	}
}

func IsKeyPassword(id, password string) bool {
	row, err := DataBase.Query(`SELECT id FROM keys WHERE id = $1 AND password = $2;`, id, password)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return false
	}
	defer row.Close()
	key := ""
	for row.Next() {
		err = row.Scan(&key)
		if err != nil {
			fmt.Println("ERROR: ", err)
			return false
		}
	}
	if key == "" {
		return false
	} else {
		return true
	}
}

func AddKey(id, password string) error {
	if IsKey(id) {
		return fmt.Errorf("the key is busy")
	}
	_, err := DataBase.Exec(`INSERT INTO keys VALUES($1, $2);`, id, password)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return err
	}
	return nil
}

func GetKeyVersion(id string) int {
	row, err := DataBase.Query(`SELECT version FROM keys WHERE id = $1;`, id)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return -1
	}
	defer row.Close()
	version := -10
	for row.Next() {
		err = row.Scan(&version)
		if err != nil {
			fmt.Println("ERROR: ", err)
			return -1
		}
	}
	if version == -10 {
		return -1
	} else {
		return version
	}
}

func EditKeyVersion(id, act string) {
	row, err := DataBase.Query(`SELECT version FROM keys WHERE id = $1;`, id)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	defer row.Close()
	version := -10
	for row.Next() {
		err = row.Scan(&version)
		if err != nil {
			fmt.Println("ERROR: ", err)
		}
	}
	if version == -10 {
		return
	}
	if act == "remove" {
		row, err := DataBase.Query(`SELECT id FROM notes_movies WHERE key_id = $1;`, id)
		if err != nil {
			fmt.Println("ERROR: ", err)
		}
		defer row.Close()
		col := 0
		for row.Next() {
			col++
			break
		}
		if col == 0 {
			_, err := DataBase.Exec(`UPDATE keys SET version = $1 WHERE id = $2;`, 0, id)
			if err != nil {
				fmt.Println("ERROR: ", err)
			}
		} else {
			if version == 32766 {
				version = 1
			} else {
				version++
			}
			_, err := DataBase.Exec(`UPDATE keys SET version = $1 WHERE id = $2;`, version, id)
			if err != nil {
				fmt.Println("ERROR: ", err)
			}
		}
		return
	}
	if act == "edit,add" {
		if version == 32766 {
			version = 1
		} else {
			version++
		}
		_, err := DataBase.Exec(`UPDATE keys SET version = $1 WHERE id = $2;`, version, id)
		if err != nil {
			fmt.Println("ERROR: ", err)
		}
		return
	}
}
