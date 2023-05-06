package models

import (
	"net/http"

	"github.com/PBKKE08/FP-BE/echo-rest/db"
)

type City struct{
	Id int `json:"id"`
	Name string `json:"name"`
}

func FetchAllCities() (Response, error) {
	var obj City
	var arrobj []City
	var res Response

	db := db.GetDb()

	sqlStatement := "SELECT * FROM city"

	rows, err := db.Query(sqlStatement)
	
	if(err != nil){
		return res, err
	}
	defer rows.Close()
	

	for rows.Next() {
		err = rows.Scan(&obj.Id, &obj.Name)
		if (err != nil){
			return res, err
		}

		arrobj = append(arrobj, obj)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = arrobj

	return res, nil
}

func PostCity(name string) (Response, error){
	var res Response
	var id int
	db := db.GetDb()

	sqlStatement := `INSERT INTO public.city (name) VALUES ($1) RETURNING id`
	err := db.QueryRow(sqlStatement, name).Scan(&id)
	if err != nil{
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int{
		"last_inserted_id": id,
	}

	return res, nil
}