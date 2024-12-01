# golang-red-team
Red Team based Projects made from Golang

Ideas taken from [kurogai](https://github.com/kurogai/100-redteam-projects)

### HTTP Application Scanner

This tool scans HTTP applications for various details such as headers, hosts, and more.

**USAGE**
```bash
./http-app-scan -f assets/hosts.txt
```
`assets/hosts.txt` should contain a list of hosts to scan, for example:
```bash
1.1.1.1
iiitdm.ac.in
regex-generator.olafneumann.org
stackoverflow.com
github.com
```
#### Sample output for `github.com`:
```bash
{
    "Host": "github.com",
    "Headers": {
        "Accept-Ranges": ["bytes"],
        "Cache-Control": ["max-age=0, private, must-revalidate"],
        "Content-Language": ["en-US"],
        "Content-Security-Policy": ["default-src 'none'; ..."]
    },
    "Title": "GitHub: Let’s build from here · GitHub",
    "Hosts": {
        "Certificate": ["github.com", "DigiCert TLS Hybrid ECC SHA384 2020 CA1"],
        "Headers": ["Date", "Content-Type", "Vary"],
        "HTML": [".electronjs.org", "desktop.github.com", "github.community", "services.github.com"]
    },
    "Other": {
        "Emails": null,
        "Ips": ["20.207.73.82"],
        "Buckets": {"github.s3.amazonaws.com": "Private"},
        "Social": ["https://twitter.com/github", "https://www.facebook.com/GitHub", "https://www.linkedin.com/company/github"],
        "Copyright": ["2023 GitHub, Inc."]
    }
}
```
**Developing**

- Finding buckets: The tool checks for public S3 buckets by testing URLs such as http://website-name.s3.amazonaws.com/.
- Copyrights: Extracts copyright information by scraping for specific symbols and words.
- Hosts: Identifies links mentioned on the website.
- Headers: Collects all HTTP headers for the given host.

### Subdomain Brute Force

This tool performs brute force attacks to find subdomains for a given host.

**USAGE**
```bash
./subdomain-bf --host twitter.com -f assets/subs.txt
```
`assets/subs.txt` should contain a list of possible subdomains, for example:
```bash
hello
login
app
careers
admin
dashboard
console
email
box
```

Sample Output:
```bash
{
    "Subdomain": "dashboard.twitter.com",
    "Type": "A",
    "Resource": [
        "104.244.42.67",
        "104.244.42.3",
        "104.244.42.131",
        "104.244.42.195"
    ]
}
```


