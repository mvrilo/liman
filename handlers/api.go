// Copyright 2018 The Liman Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/salihciftci/liman/cmd"
	"github.com/salihciftci/liman/util"
)

var (
	apiKey = util.GeneratePassword(32)
)

func apiStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"ok":     "false",
				"result": "METHOD_NOT_ALLOWED",
			})
		log.Println(r.Method, r.URL.Path)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"ok":     "true",
		"result": nil,
	})

	log.Println(r.Method, r.URL.Path)
}

func apiAuth(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "GET" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"ok":     "false",
				"result": "METHOD_NOT_ALLOWED",
			})
		return fmt.Errorf("METHOD_NOT_ALLOWED")
	}

	params := r.URL.Query()
	key, ok := params["key"]

	if !ok || len(key) < 1 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"ok":     "false",
				"result": "API_KEY_NOT_FOUND",
			})
		return fmt.Errorf("API_KEY_NOT_FOUND")
	}

	if string(key[0]) != apiKey {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(
			map[string]interface{}{
				"ok":     "false",
				"result": "API_KEY_INVALID",
			})
		return fmt.Errorf("API_KEY_INVALID")
	}

	return nil
}

func apiContainer(w http.ResponseWriter, r *http.Request) {
	err := apiAuth(w, r)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	container, err := cmd.ParseContainers()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"ok":     "true",
			"result": container,
		})
	log.Println(r.Method, r.URL.Path)
}

func apiImages(w http.ResponseWriter, r *http.Request) {
	err := apiAuth(w, r)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	images, err := cmd.ParseImages()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"ok":     "true",
			"result": images,
		})

	log.Println(r.Method, r.URL.Path)
}

func apiVolumes(w http.ResponseWriter, r *http.Request) {
	err := apiAuth(w, r)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	volumes, err := cmd.ParseVolumes()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"ok":     "true",
			"result": volumes,
		})

	log.Println(r.Method, r.URL.Path)
}

func apiNetworks(w http.ResponseWriter, r *http.Request) {
	err := apiAuth(w, r)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	networks, err := cmd.ParseNetworks()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"ok":     "true",
			"result": networks,
		})

	log.Println(r.Method, r.URL.Path)
}

func apiStats(w http.ResponseWriter, r *http.Request) {
	err := apiAuth(w, r)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	stats, err := cmd.ParseStats()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"ok":     "true",
			"result": stats,
		})

	log.Println(r.Method, r.URL.Path)
}

func apiLogs(w http.ResponseWriter, r *http.Request) {
	err := apiAuth(w, r)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	container, err := cmd.ParseContainers()
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	logs, err := cmd.ParseLogs(container)
	if err != nil {
		log.Println(r.Method, r.URL.Path, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"ok":     "true",
			"result": logs,
		})

	log.Println(r.Method, r.URL.Path)
}
