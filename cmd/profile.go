package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Profiles(w http.ResponseWriter, r *http.Request) {
	profiles := getProfiles()

	out, err := json.Marshal(profiles)

	if err != nil {
        fmt.Println("error encoding people:", err)
    }

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func ProfileId(w http.ResponseWriter, r *http.Request) {
	id := getIdProfile()
	out, err := json.Marshal(id)

	if err != nil {
        fmt.Println("error encoding people:", err)
    }

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func AddProfile(w http.ResponseWriter, r *http.Request) {
	var profile ProfileJson
	var payload jsonResponse

	err := json.NewDecoder(r.Body).Decode(&profile)

	if err != nil{
		fmt.Println("invalid json")
		payload.Error = true
		payload.Message = "invalid json"
		
		out,err := json.MarshalIndent(payload,"","\t")
		if err != nil{
			fmt.Println(err)
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(out)
		return
	}
	
	saveProfile(profile)

	payload.Error = false
	payload.Message = "card saved"
	out,err := json.MarshalIndent(payload,"","\t")
		if err != nil{
			fmt.Println(err)
		}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func EditProfile(w http.ResponseWriter, r *http.Request) {
	var profile ProfileJson
	var payload jsonResponse

	err := json.NewDecoder(r.Body).Decode(&profile)

	if err != nil{
		fmt.Println("invalid json")
		payload.Error = true
		payload.Message = "invalid json"
		
		out,err := json.MarshalIndent(payload,"","\t")
		if err != nil{
			fmt.Println(err)
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(out)
		return
	}
	
	editProfile(profile)

	payload.Error = false
	payload.Message = "profile edited"
	out,err := json.MarshalIndent(payload,"","\t")
		if err != nil{
			fmt.Println(err)
		}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	var profile IdProfile
	var payload jsonResponse

	err := json.NewDecoder(r.Body).Decode(&profile)
	
	if err != nil{
		fmt.Println("invalid json")
		payload.Error = true
		payload.Message = "invalid json"
		
		out,err := json.MarshalIndent(payload,"","\t")
		if err != nil{
			fmt.Println(err)
		}
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(out)
		return
	}
	
	deleteProfile(profile.Id)

	payload.Error = false
	payload.Message = "id card"
	out,err := json.MarshalIndent(payload,"","\t")
		if err != nil{
			fmt.Println(err)
		}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}