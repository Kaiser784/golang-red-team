package main

import (
	"fmt"
	"log"
	"os"
    "bufio"
    "net"
    "encoding/json"
    _ "strings"

	"github.com/urfave/cli/v2"
	
)

var host string
var subs string

type info struct {
    Subdomain string
    Type string
    Resource []string
}

func retrieval() {
    file, err := os.Open(subs)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        var ip []string
        sub := scanner.Text()
        cname, _ := net.LookupCNAME(sub+"."+host)
        cname_resource := []string{cname}

        iprecords, _ := net.LookupIP(sub+"."+host)
        for i := 0; i < len(iprecords); i++ {
            ip = append(ip, iprecords[i].String())
        }

        if(cname_resource[0] != "") || (len(ip) > 0){
            resultc := info{sub+"."+host, "CNAME", cname_resource}
            resulta := info{sub+"."+host, "A", ip}
            outc, _ := json.MarshalIndent(resultc, "", "    ")
            outa, _ := json.MarshalIndent(resulta, "", "    ")
            
            fmt.Println(string(outc))
            fmt.Println(string(outa))
        }
    }
}

func main() {

    app := &cli.App{
		Name:     "subdomain-bf",
        Usage: "bruteforce subdomains (A and CNAME records) on a given host",
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:    	"host",
                Usage:   	"Host to bruteforce `URL`(e.g. twitter.com)",
				Destination: &host,
				Required: true,
            },
            &cli.StringFlag{
                Name:    "file",
                Aliases: []string{"f"},
                Usage:   "Load subdomains from `FILE`",
				Destination: &subs,
				Required: true,
            },
        },
        Action: func(cli *cli.Context) error {
            fmt.Println("URL - ", host)
            if _, err := os.Stat(subs); err == nil {
                fmt.Println("Input file:", subs)
            } else {  
                fmt.Printf("Input File NOT Found\n");
            }
            retrieval()

            return nil
        },
        HideHelpCommand: true,
    }


    if err := app.Run(os.Args); err != nil {
        log.Fatal(err)
    }
	
}