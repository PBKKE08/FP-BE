package models

import (
	"net/http"

	"github.com/PBKKE08/FP-BE/echo-rest/db"
)

type Category struct {
	Id   int `json:"id"`
	Name string `json:"name"`
}

func FetchAllCategories() (Response, error) {
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