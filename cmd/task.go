package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Tasks(w http.ResponseWriter, r *http.Request) {
	tasks := getTasks()

	out, err := json.Marshal(tasks)

	if err != nil {
        fmt.Println("error encoding people:", err)
    }

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func TaskId(w http.ResponseWriter, r *http.Request) {
	id := getIdTask()

	out, err := json.Marshal(id)

	if err != nil {
        fmt.Println("error encoding people:", err)
    }

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func AddTask(w http.ResponseWriter, r *http.Request) {
	var task TaskJson
	var payload jsonResponse

	err := json.NewDecoder(r.Body).Decode(&task)
	
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
	
	saveTask(task)

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

func EditTask(w http.ResponseWriter, r *http.Request) {
 	var task TaskJson
 	var payload jsonResponse

 	err := json.NewDecoder(r.Body).Decode(&task)
	
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
	
	editTask(task)

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

 func DeleteTask(w http.ResponseWriter, r *http.Request) {
 	var task IdTask
	var payload jsonResponse

	err := json.NewDecoder(r.Body).Decode(&task)
	
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
	
	deleteTask(task.Id)

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


func StartTask(w http.ResponseWriter, r *http.Request) {
	var id IdTask
	var payload jsonResponse

	err := json.NewDecoder(r.Body).Decode(&id)

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
	
	StartSpecTask(id)

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

func StopTask(w http.ResponseWriter, r *http.Request) {
	var task IdTask
	var payload jsonResponse
	err := json.NewDecoder(r.Body).Decode(&task)
	
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
	
	//deleteTask(task.Id)

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