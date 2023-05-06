package models

import (
	"net/http"

	"github.com/PBKKE08/FP-BE/echo-rest/db"
)

type Category struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func ReadAllCategories() (Response, error) {
	var obj Category
	var arrobj []Category
	var res Response

	db := db.GetDb()

	sqlStatement := "SELECT * FROM category"

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

func CreateCategory(name string) (Response, error){
	var res Response
	var id int
	db := db.GetDb()

	sqlStatement := `INSERT INTO public.category (name) VALUES ($1) RETURNING id`
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

func UpdateCategory (id int, name string) (Response, error){
	var res Response
	db := db.GetDb()

	sqlStatement := `UPDATE category SET name = $2 WHERE id = $1 RETURNING id`
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

func DeleteCategory (id int) (Response, error){
	var res Response
	db := db.GetDb()

	sqlStatement := `DELETE FROM public.category WHERE (id = $1) RETURNING id`
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