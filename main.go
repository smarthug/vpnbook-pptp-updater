package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// app := "ls"

	// arg0 := "-e"
	// arg1 := "Hello world"
	// arg2 := "\n\tfrom"
	// arg3 := "golang"

	// cmd := exec.Command(app, arg0, arg1, arg2, arg3)
	cmd := exec.Command("cmd", "/C", "start ms-settings:network-vpn")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))

	exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://www.vpnbook.com/").Start()

	//string fileName = Environment.GetEnvironmentVariable("USERPROFILE")+"\\AppData\\Roaming\\Microsoft\\Network\\Connections\\PBK\\rasphone.pbk";

	//https://docs.microsoft.com/en-us/powershell/module/vpnclient/add-vpnconnection?view=win10-ps
	//https://github.com/ugurturhal/VPN_Creator_CSharp/blob/master/Codes
	//https://stackoverflow.com/questions/318233/how-can-i-programmatically-create-a-windows-vpn-connection

}
