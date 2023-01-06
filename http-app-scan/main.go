package main

import (
	"bufio"
	"crypto/tls"
	"encoding/json"
	_ "encoding/xml"
	"fmt"
	_ "io"
	"io/ioutil"
	_ "io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
    "sort"

	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli/v2"
	_ "golang.org/x/net/html"
)

var hosts string
var media = regexp.MustCompile(`instagram|facebook|linkedin|reddit|twitter`)
var domain = regexp.MustCompile(`.*([^\.]+)(com|net|org|info|coop|int|co\.uk|org\.uk|ac\.uk|uk|__and so on__)$`)
var pattern =  regexp.MustCompile(`(<\/?[a-zA-A]+?[^>]*\/?>)*`)

type Hosts struct {
    Certificate []string
    Headers     []string
    HTML        []string
}

type other struct {
    Emails    []string
    Ips       []string
    Buckets   map[string]string
    Social    []string
    Copyright []string
}

type info struct {
    Host    string
    Headers map[string][]string
    Title string
    Hosts Hosts
    Other other
}

func contains(elems []string, v string) bool {
    for _, s := range elems {
        if v == s {
            return true
        }
    }
    return false
}

func removeDuplicateStr(strSlice []string) []string {
    allKeys := make(map[string]bool)
    list := []string{}
    for _, item := range strSlice {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}

func getLinks(sub string) ([]string, []string){
    response, _ := http.Get("https://"+sub)
    defer response.Body.Close()
    doc, _ := goquery.NewDocumentFromReader(response.Body)

    var social []string
    var hosts []string

    f := func(i int, s *goquery.Selection) bool {
        link, _ := s.Attr("href")
        return strings.HasPrefix(link, "https")
    }

    doc.Find("body a").FilterFunction(f).Each(func(_ int, tag *goquery.Selection) {
        link, _ := tag.Attr("href")
        if media.MatchString(link) {
            social = append(social, link)
        }else { 
            url, _ := url.Parse(link)
            host := strings.TrimPrefix(url.Hostname(), "www")
            if !(contains(hosts, host)){
                hosts = append(hosts, host)
            }    
        }
    })
    return social, hosts
}

func getMails(sub string) []string{
    response, _ := http.Get("https://"+sub)
    defer response.Body.Close()
    doc, _ := goquery.NewDocumentFromReader(response.Body)

    var mails []string

    f := func(i int, s *goquery.Selection) bool {
        link, _ := s.Attr("href")
        return strings.HasPrefix(link, "mailto:")
    }

    doc.Find("body a").FilterFunction(f).Each(func(_ int, tag *goquery.Selection) {
        link, _ := tag.Attr("href")
        split := strings.Split(link, ":")
        mails = append(mails, split[1])
    })
    return mails
}

func getIP(sub string) []string{
    var ip []string
    iprecords, _ := net.LookupIP(sub)
    for i := 0; i < len(iprecords); i++ {
        ip = append(ip, iprecords[i].String())
    }
    ip = removeDuplicateStr(ip)
    return ip
}

func getCert(sub string) []string{
    var certs []string
    conf := &tls.Config{
        InsecureSkipVerify: true,
    }
    conn, _ := tls.Dial("tcp", sub+":443", conf)
    defer conn.Close()
    tls := conn.ConnectionState().PeerCertificates

    for _, cert := range tls {
        certs = append(certs, cert.Subject.CommonName)
    }
    return certs
}

func getS3(sub string) map[string]string{
    bucket := make(map[string]string)
    var hostname string = sub
    aws_s3 := ".s3.amazonaws.com"
    ext := []string{"www.", ".com", ".org",".ac", ".co", ".in", ".gov", ".edu"}
    
    for i:=0; i < len(ext); i++ {
        hostname = strings.Replace(hostname, ext[i], "" ,-1)
    }
    res, _ := http.Get("http://"+hostname+aws_s3)
    content, _ := ioutil.ReadAll(res.Body)

    if strings.Contains(string(content), "NoSuchBucket") {
        bucket[hostname+aws_s3] = "Does not exist" 
    } else if strings.Contains(string(content), "AccessDenied"){
        bucket[hostname+aws_s3] = "Private" 
    }else {
        bucket[hostname+aws_s3] = "Public" 
    }
    
    return bucket
}

func removeTag(line string) string{
    groups := pattern.FindAllString(line, -1)
	// should replace long string first
	sort.Slice(groups, func(i, j int) bool {
		return len(groups[i]) > len(groups[j])
	})
	for _, group := range groups {
		if strings.TrimSpace(group) != "" {
			line = strings.ReplaceAll(line, group, "")
		}
	}
    line = strings.Replace(line, "&copy;", "", -1)
    line = strings.Replace(line, "&#169", "", -1)
	return line
}

func getCopyrights(sub string) []string{
    var lines []string
    copy := []string{"Copyright ", "© ", "&copy; ", "&#169; "}
    var copyrights []string

    response, _ := http.Get("https://"+sub)
    scanner := bufio.NewScanner(response.Body)
    for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

    for _, line := range lines {
        for i:=0; i < len(copy); i++  {
            if strings.Contains(line, copy[i]){
                line = removeTag(line)
                copyrights = append(copyrights, line)
            }
        }
	}
    copyrights = removeDuplicateStr(copyrights)
    return copyrights
}

func retrieval() {
    file, err := os.Open(hosts)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        sub := scanner.Text()
        response, _ := http.Get("https://"+sub)
        doc, _ := goquery.NewDocumentFromReader(response.Body)

        title := doc.Find("title").Text()
        headers := response.Header
        social, hosts := getLinks(sub)
        mails := getMails(sub)
        ip := getIP(sub)
        keys := make([]string, 0 , len(headers))
        for k := range headers {
            keys = append(keys, k)
        }
        certs := getCert(sub)
        buckets := make(map[string]string)
        buckets = getS3(sub)
        copy := getCopyrights(sub)

        resultc := &info{
            Host: sub, 
            Headers: headers, 
            Title: title,
            Hosts: Hosts{
                Certificate: certs,
                Headers: keys,
                HTML:hosts,
            },
            Other: other{
                Emails: mails,   
                Ips: ip,
                Buckets: buckets,
                Social: social,
                Copyright: copy,
            },
        }
        outc, _ := json.MarshalIndent(resultc, "", "    ")
    
        fmt.Println(string(outc))
    }
}

func main() {

    app := &cli.App{
		Name:  "http-app-scan",
        Usage: "scans the web pages for interesting and useful things",
        Flags: []cli.Flag{
            &cli.StringFlag{
                Name:    "file",
                Aliases: []string{"f"},
                Usage:   "Load subdomains from `FILE`",
				Destination: &hosts,
				Required: true,
            },
        },
        Action: func(cli *cli.Context) error {
            if _, err := os.Stat(hosts); err == nil {
                fmt.Println("Input file:", hosts)
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