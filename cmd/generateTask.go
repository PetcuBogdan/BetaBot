package main

import "fmt"

func StartSpecTask(id IdTask) {
	task := getSpecTask(id.Id)
	profile := getSpecProfile(task.Profile)
	card := getSpecCard(task.Card)
	taskStart := ReadyTask{
		Task: task,
		Profile: profile,
		Card: card,
	}
	fmt.Println(taskStart)
	if taskStart.Task.Shop == "Palace" {
		palace(taskStart)
	}else{
		supreme(taskStart)
	}
}