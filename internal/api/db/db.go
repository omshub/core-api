package db

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"omshub/core-api/internal/api/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	if !db.Migrator().HasTable("users") {
		err = db.AutoMigrate(&models.Program{}, &models.Semester{}, &models.Course{}, &models.Specialization{}, &models.User{}, &models.Review{})
		user := models.User{ID: 1, UserName: "OMSCentral"}
		db.Create(&user)

		res, _ := http.Get("https://omshub-data.s3.amazonaws.com/data/omscentral_programs.json")
		file, _ := ioutil.ReadAll(res.Body)
		var programs []models.Program
		_ = json.Unmarshal([]byte(file), &programs)
		db.Create(&programs)

		res, _ = http.Get("https://omshub-data.s3.amazonaws.com/data/omscentral_semesters.json")
		file, _ = ioutil.ReadAll(res.Body)
		var semesters []models.Semester
		_ = json.Unmarshal([]byte(file), &semesters)
		db.Create(&semesters)

		res, _ = http.Get("https://omshub-data.s3.amazonaws.com/data/omscentral_courses.json")
		file, _ = ioutil.ReadAll(res.Body)
		var courses []models.Course
		_ = json.Unmarshal([]byte(file), &courses)
		db.Create(&courses)

		res, _ = http.Get("https://omshub-data.s3.amazonaws.com/data/omscentral_specializations.json")
		file, _ = ioutil.ReadAll(res.Body)
		var specializations []models.Specialization
		_ = json.Unmarshal([]byte(file), &specializations)
		db.Create(&specializations)

		res, _ = http.Get("https://omshub-data.s3.amazonaws.com/data/omscentral_reviews.json")
		file, _ = ioutil.ReadAll(res.Body)
		var reviews []models.Review
		_ = json.Unmarshal([]byte(file), &reviews)
		db.CreateInBatches(reviews, 200)
	} else {
		err = db.AutoMigrate(&models.Program{}, &models.Semester{}, &models.Course{}, &models.Specialization{}, &models.User{}, &models.Review{})
	}
	// err = db.AutoMigrate(&models.Program{}, &models.Semester{}, &models.Course{}, &models.Specialization{}, &models.User{}, &models.Review{})

	// return migration error
	return db, err
}
