package cache

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Item struct {
	Value       []byte    `json:"value"`
	ExpiredTime time.Time `json:"expiredTime"`
}

type Items map[string]Item

const FileName = "db.json"

var Storage Items

func InitCache() {
	Storage = Items{}
	Storage.Save()

	fmt.Println("Cache init success!!")
}

func Get(k string) *Item {
	Storage = *(Read())
	v, ok := Storage[k]
	if !ok {
		return nil
	}

	if time.Now().Compare(v.ExpiredTime) != -1 {
		Remove(k)
		return nil
	}

	return &v
}

func Set(k string, v []byte, t time.Duration) {
	i := Item{
		Value:       v,
		ExpiredTime: time.Now().Add(t),
	}

	Storage[k] = i
	Storage.Save()
}

func Remove(k string) {
	delete(Storage, k)
	Storage.Save()
}

func RemoveAll() {
	Storage = Items{}
	Storage.Save()
}

func (i *Items) Save() {
	b, err := json.Marshal(i)
	if err != nil {
		fmt.Println("can not convert items to json", err)
		return
	}

	err = os.WriteFile(FileName, b, os.ModePerm)
	if err != nil {
		fmt.Println("can not save to db", err)
		return
	}
}

func Read() (items *Items) {
	b, err := os.ReadFile(FileName)
	if err != nil {
		fmt.Println("can not read db", err)
		return nil
	}

	items = &Items{}
	err = json.Unmarshal(b, items)
	if err != nil {
		fmt.Println("can not json to items", err)
		return nil
	}

	return items
}
