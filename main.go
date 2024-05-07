/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"

	"github.com/kubespect/server/cmd"
)

func main() {
	log.Println(`
	__ _  _  _  ____  ____  ____  ____  ____  ___  ____ 
	(  / )/ )( \(  _ \(  __)/ ___)(  _ \(  __)/ __)(_  _)
	 )  ( ) \/ ( ) _ ( ) _) \___ \ ) __/ ) _)( (__   )(  
	(__\_)\____/(____/(____)(____/(__)  (____)\___) (__) 
`)

	log.Println("KubeSpect Server is starting...")
	cmd.Start()
}
