
// Copyright (C) shubaoyu2, SSE@USTC, 2023-2025                                              
                                                                                                
//  FILE NAME             :  menu.go                                                             
//  PRINCIPAL AUTHOR      :  094sby                                                             
//  SUBSYSTEM NAME        :  menu                                                                 
//  MODULE NAME           :  menu                                                                 
//  LANGUAGE              :  GO                                                                    
//  TARGET ENVIRONMENT    :  ANY                                                                 
//  DESCRIPTION           :  This is a menu program                                               
//
// Revision log:

// Created by 094sby

 

package main
 
import "fmt"


func menu() {
	for {
		var cmd string
		fmt.Println("cmd:")
		fmt.Scanln(&cmd)
		if "help" == cmd {
			fmt.Println("this is a help function ")
		} else if "exit" == cmd {
			fmt.Println("exit")
			break
		} else {
			fmt.Println("wrong command")
		}
	}
}

func main() {
	menu()
}


