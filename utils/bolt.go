package utils

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

func OpenDB() *bolt.DB {
	db, err := bolt.Open("./savedata/database.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CloseDB(db *bolt.DB) {
	db.Close()
}

func SaveCharacterToDB(k []byte, v []byte) {
	db := OpenDB()
	defer CloseDB(db)

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("Characters"))
		return b.Put(k, v)
	})
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteCharacterFromDB(k []byte) {
	db := OpenDB()
	defer CloseDB(db)

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("Characters"))
		err := b.Delete(k)
		if err != nil {
			log.Fatal(err)
		}
		return nil

	})
	if err != nil {
		log.Fatal(err)
	}
}

func GetAllCharactersFromDB() [][]byte {
	var charList [][]byte

	db := OpenDB()
	defer CloseDB(db)

	err := db.View(func(tx *bolt.Tx) error {
		c := tx.Bucket([]byte("DB")).Bucket([]byte("Characters")).Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			vc := make([]byte, len(v))
			copy(vc, v)
			charList = append(charList, vc)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return charList
}

func SetupDB() {
	if !DirectoryCheck(DBDirectory) {
		MakeDirectory(DBDirectory)
	}

	db := OpenDB()
	defer CloseDB(db)

	err := db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}

		_, err = root.CreateBucketIfNotExists([]byte("Characters"))
		if err != nil {
			return fmt.Errorf("could not create character bucket: %v", err)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
