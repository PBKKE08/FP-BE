package models

import (
	"net/http"

	"github.com/PBKKE08/FP-BE/echo-rest/db"
)

type City struct{
	Id int `json:"id"`
	Name string `json:"name"`
}

func ReadAllCities() (Response, error) {
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

func CreateCity(name string) (Response, error){
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

func UpdateCity (id int, name string) (Response, error){
	var res Response
	db := db.GetDb()

	sqlStatement := `UPDATE city SET name = $2 WHERE id = $1 RETURNING id`
	err := db.QueryRow(sqlStatement, id, name).Scan(&id)

	if err != nil{
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int{
		"last_updated_id": id,
	}

	return res, nil
}

func DeleteCity (id int) (Response, error){
	var res Response
	db := db.GetDb()

	sqlStatement := `DELETE FROM public.city WHERE (id = $1) RETURNING id`
	err := db.QueryRow(sqlStatement, id).Scan(&id)

	if err != nil{
		return res, err
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]int{
		"last_deleted_id": id,
	}

	return res, nil
}