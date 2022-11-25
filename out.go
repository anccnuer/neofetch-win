package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var s = `                           ooooo                
                              ooooo             
                      oo       .ooooo.          
               .oooooooooo       oooooo         
             .oooooooooo.         oooooo.       
           .oooooooooo.            oooooo.      
         .oooooooooooo              oooooo      
       ooooooooooooooooo            ooooooo     
         ooooooooooooooooo          .oooooo     
           ooo.   oooooooooo         oooooo     
                    oooooooooo      ooooooo     
                      oooooooooo    ooooooo     
                        oooooooooo ooooooo      
       oo                 ooooooooooooooo.      
     oooooo                 ooooooooooooo       
      ooooooooo              ooooooooooo        
      .ooooooooooooooooooooooooooooooooooo      
   .oooooooooooooooooooooooooooooooooooooo.     
   .ooooo    .oooooooooooooooooo    oooo.       
     .o.            .oooo.            .         
                                                `

func out() {
	cpu := getCpu()
	m1, m2 := getMem()
	user := getHost()
	sys, arch := GetOS()
	d1, d2 := GetDisk()
	ss, m, h, d := getUptime()

	mp := (float64(m2) / float64(m1)) * 100

	info := [6]string{
		user,
		"---------------------------------",
		"OS:" + " " + sys + " " + arch,
		"CPU:" + cpu,
		"Memory:",
		"Disk:",
	}

	asii := strings.Split(s, "\n")
	for i := 0; i < len(asii); i++ {
		color.Set(color.FgRed)
		fmt.Printf("%v", asii[i])
		color.Unset()
		if i == 4 {
			fmt.Printf("Memory:%vMB / %vMB(%2.2f%%)%v", m2, m1, mp, "\n")
		} else if i == 5 {
			fmt.Printf("Disk:%vGB / %vGB(%2.2f%%)%v", d2, d1, (float64(d2)/float64(d1))*100, "\n")
		} else if i == 6 {
			fmt.Printf("Uptime:%v days, %v hours, %v minutes, %v seconds%v", d, h, m, ss, "\n")
		} else if i < 5 {
			fmt.Printf("%s%v", info[i], "\n")
		} else {
			fmt.Println()
		}

	}
}
