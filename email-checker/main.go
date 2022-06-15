package main

import(
	"fmt"
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func checkDomain(domain string){

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err :=net.LookupMX(domain)

	if err!=nil{
		log.Printf("Errror: %v\n", err)
	}
	if len(mxRecords) > 0{
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err !=nil{
		log.Printf("Errror:%v\n", err)
	}

	for _, record := range txtRecords{
		if strings.HasPrefix(record, "v=spfi"){
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil{
		log.Printf("Error%v\n", err)
	}

	for _, record :=range dmarcRecords{
		if strings.HasPrefix(record, "v=DMARC1"){
			hasDMARC = false
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v,    %v,    %v,    %v,    %v,    %v    ", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}


func main(){

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")

	for scanner.Scan(){
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err!=nil{
		log.Fatal("Error: could not read from inpit: %v\n", err)
	}
}

