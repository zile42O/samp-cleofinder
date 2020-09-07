package main
import (
	"fmt"
	"os"
	"path/filepath"
	"log"
	"github.com/fatih/color"
	"bufio"
	"strings"
	"time"
	"os/exec"
	"regexp"
)

func main() {
	//Start program
	color.Green("THE                       ><<<<")
	color.Green("      ><<    ><< ><     ><<    ><<")
	color.Green("    > ><<   ><     ><<><<        ><<")
	color.Green("   >< ><<        ><<  ><<        ><<")
	color.Green(" ><<  ><<      ><<    ><<        ><<")
	color.Green("><<<< >< ><< ><<        ><<     ><<")
	color.Green("      ><<   ><<<<<<<<     ><<<<    PRODUCTIONS")
	fmt.Println("\n")
	color.Blue("Author:")
	color.Yellow("Aleksandar Zivkovic (Zile42O)")
	color.Blue("Software:")
	color.Yellow("Cleo/ASI Finder (GTA:SA - SAMP)")
	color.Blue("Version:")
	color.Yellow("1.1")
	fmt.Println("\n")
	//Sleep 5 sec
	color.Blue("Please wait...")
	time.Sleep(5 * time.Second)
	//Clear console
	os.Stdout.Write([]byte{0x1B, 0x5B, 0x33, 0x3B, 0x4A, 0x1B, 0x5B, 0x48, 0x1B, 0x5B, 0x32, 0x4A})
	//Show Rules
	color.Red(" ____  _   _ _     _____ ____")  
	color.Red("|  _ \\| | | | |   | ____/ ___|") 
	color.Red("| |_) | | | | |   |  _| \\___ \\") 
	color.Red("|  _ <| |_| | |___| |___ ___) |")
	color.Red("|_| \\_\\\\___/|_____|_____|____/")

	fmt.Println("\n")
	color.Yellow("1. This software will scroll through your folders to find unwanted files.")
	color.Yellow("2. This software will not collect information about you.")
	color.Yellow("3. This software is secure, and does not have any illegal processes in it.")
	color.Yellow("4. It is undesirable to interrupt the program while scanning files.")
	//Show Confirmation
	color.Yellow("\nThe scan will start if you accept the software rules.\n")
	question := ShowConfirmationOption("Do you agree with the stated rules of this software?")
	if question {
		StartScan()	
	} else {
		color.Red("You have declined the policy, thus interrupting the software process, please wait..")
		time.Sleep(3 * time.Second)
		return
	}
}
func StartScan() {
	FoundedFiles := 0
	ScannedFiles := 0
	var found_files []string
	//Find current proccess path and volume name
	/*path, err := os.Getwd()
	if err != nil {
	    log.Println(err)
	}*/
	//Start searching for file samp.exe filepath.VolumeName(path)
	color.Blue("Searching for 'samp.exe', please wait...")
	//output, err := exec.Command("where", "/R", filepath.VolumeName(path), "\\", "samp.exe").Output()
	output, err := exec.Command("cmd.exe", "/c", "dir \\*samp.exe /s /b").Output()
	if err != nil {
		color.Red("An error occurred, 'samp.exe' may not exist in this disk volume | %s", err)
		color.Yellow("The program must be in disk volume")
		time.Sleep(3 * time.Second)
		return
	}
	///
	start := time.Now()
	os.Stdout.Write([]byte{0x1B, 0x5B, 0x33, 0x3B, 0x4A, 0x1B, 0x5B, 0x48, 0x1B, 0x5B, 0x32, 0x4A})
	var myReg = regexp.MustCompile(`(?P<filepath>(?P<root>[/]?)(?P<rest_of_the_path>.+))`)
	color.Green("Scan running successfully, please wait ..")
	for i, match := range myReg.FindAllString(string(output), -1) {
		if i != -1 {
			group := myReg.FindStringSubmatch(match)
			directory := filepath.Dir(group[0])	
			color.Blue("Scan directory is %s", directory)
			time.Sleep(2 * time.Second)
			exts := []string{".cs", ".cleo", ".asi"}
			//Listing
			err = filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
				//Checking files
				if stringInSlice(filepath.Ext(path), exts) {
					found_files = append(found_files, path)
					FoundedFiles++
				} else {
					color.Blue("Scanning: %s", path)
					ScannedFiles++
				} 
				return nil
			})
			if err != nil {
				panic(err)
			}
			for _, file := range found_files {		
				color.Red("\n> %s\n", file)
			}
		}	
	}
	//Total	
	if FoundedFiles > 0 {
		color.Red("\nTotal found files: %v\n", FoundedFiles)
	} else {
		color.Green("\nNo files found\n")
	}
	color.Blue("\nTotal scanned files: %v\n", ScannedFiles)
	//End work
	//Time
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("\n----------------------------------------------------")
	color.Blue("Scanning took:")
	color.Yellow("%s", elapsed)
	time.Sleep(120 * time.Second)
}

func stringInSlice(v string, ss []string) bool {
    for _, s := range ss {
        if s == v {
            return true
        }
    }
    return false
}
func ShowConfirmationOption(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	for {
		color.Blue("%s [y/n]: ", s)

		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		response = strings.ToLower(strings.TrimSpace(response))
		if response == "y" || response == "yes" {
			return true
		} else if response == "n" || response == "no" {
			return false
		}
	}
}
