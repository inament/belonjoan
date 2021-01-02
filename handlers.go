package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bagus2x/sumux"
	"github.com/bagus2x/sumux/res"
)

// BelonjoanHandler .
type BelonjoanHandler struct {
	bs *BelonjoanStore
}

// InitBelonjoanHandler .
func InitBelonjoanHandler(db *sql.DB, r sumux.Router) {
	bh := BelonjoanHandler{
		NewBelonjoanStore(db),
	}

	r.Group("/api/v1/belonjoan", func(r sumux.Router) {
		r.Get("/", bh.FindAll)
		r.Get("/<itemid>", bh.FindOne)
		r.Post("/", bh.AddItem)
		r.Put("/<itemid>", bh.Update)
		r.Delete("/<itemid>", bh.Delete)
	})
}

// AddItem belonjoan
func (bh BelonjoanHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	var belonjoan Belonjoan
	err := json.NewDecoder(r.Body).Decode(&belonjoan)
	if err != nil {
		res.JSON(w, 400, map[string]string{"error": err.Error()})
		return
	}

	err = bh.bs.Add(belonjoan)
	if err != nil {
		fmt.Println(err)
		res.JSON(w, 400, map[string]string{"error": err.Error()})
		return
	}

	res.JSON(w, 200, map[string]string{"message": "ok"})
}

// FindAll belonjoan
func (bh BelonjoanHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	belonjoans, err := bh.bs.FindAll()
	if err != nil {
		res.JSON(w, 400, map[string]string{"error": err.Error()})
		return
	}
	res.JSON(w, 200, belonjoans)
}

// FindOne belonjoan
func (bh BelonjoanHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	itemidStr, _ := sumux.Param(r, "itemid")
	itemid, err := strconv.Atoi(itemidStr)

	if err != nil {
		res.JSON(w, 400, map[string]string{"error": err.Error()})
		return
	}

	belonjoan, err := bh.bs.FindOne(itemid)
	if err != nil {
		res.JSON(w, 400, map[string]string{"error": err.Error()})
		return
	}

	res.JSON(w, 200, belonjoan)
}

// Update belonjoan
func (bh BelonjoanHandler) Update(w http.ResponseWriter, r *http.Request) {
	itemidStr, _ := sumux.Param(r, "itemid")
	itemid, err := strconv.Atoi(itemidStr)

	if err != nil {
		res.JSON(w, 400, map[string]string{"error": err.Error()})
		return
	}

	var belonjoan Belonjoan
	err = json.NewDecoder(r.Body).Decode(&belonjoan)
	if err != nil {
		res.JSON(w, 400, map[string]string{"error": err.Error()})
		return
	}

	err = bh.bs.UpdateItem(itemid, belonjoan)
	if err != nil {
		res.JSON(w, 400, map[string]string{"error": err.Error()})
		return
	}

	res.JSON(w, 200, map[string]string{"message": "update success"})
}

// Delete belonjoan
func (bh BelonjoanHandler) Delete(w http.ResponseWriter, r *http.Request) {
	itemidStr, _ := sumux.Param(r, "itemid")
	itemid, err := strconv.Atoi(itemidStr)

	if err != nil {
		res.JSON(w, 400, map[string]string{"error": err.Error()})
		return
	}

	err = bh.bs.DeleteItem(itemid)
	if err != nil {
		res.JSON(w, 400, map[string]string{"error": err.Error()})
		return
	}

	res.JSON(w, 200, map[string]string{"message": "delete success"})
}
