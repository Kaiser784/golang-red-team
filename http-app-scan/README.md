### Usage

```bash
./http-app-scan -f assets/hosts.txt
```

```
assets/hosts.txt

redhuntlabs.com
1.1.1.1
iiitdm.ac.in
regex-generator.olafneumann.org
stackoverflow.com
github.com
```

`Output` for `github.com`  (not the whole actual output)
```json
{
    "Host": "github.com",
    "Headers": {
        "Accept-Ranges": [
            "bytes"
        ],
        "Cache-Control": [
            "max-age=0, private, must-revalidate"
        ],
        "Content-Language": [
            "en-US"
        ],
        "Content-Security-Policy": [
            "default-src 'none'; ....."
        ],
		".............."
    },
    "Title": "GitHub: Let’s build from here · GitHubPythonJavaScriptGo",
    "Hosts": {
        "Certificate": [
            "github.com",
            "DigiCert TLS Hybrid ECC SHA384 2020 CA1"
        ],
        "Headers": [
            "Date",
            "Content-Type",
            "Vary",
            "....."
        ],
        "HTML": [
            ".electronjs.org",
            "desktop.github.com",
            "github.community",
            "services.github.com",
            "......."
        ]
    },
    "Other": {
        "Emails": null,
        "Ips": [
            "20.207.73.82"
        ],
        "Buckets": {
            "github.s3.amazonaws.com": "Private"
        },
        "Social": [
            "https://twitter.com/github",
            "https://www.facebook.com/GitHub",
            "https://www.linkedin.com/company/github"
        ],
        "Copyright": [
            "           2023 GitHub, Inc."
        ]
    }
}
```

### Developing

**Finding buckets**
For this I didn't understand if I had to scrape the website for leaked configs or bucket names, I stuck to the basic checking of `http://website-name.s3.amazonaws.com/` to check if there exists a public bucket and output if there exists a bucket and if it does if it's public or private.

More functionality can be added here as an improement to check for leaked configs

**Copyrights**
Getting copyright information from the website was a bit trickier as all websites do not follow a standard tag for copyrights so I had to scrape for specific symbols and words and clean the line of tags and then get the copyright information

**Hosts**
I understood hosts as any links mentioned in the website, if this meant something else please do specify so I can update 

**Headers**
In the instructions Headers was mentioned twice one under Hosts, as of now I've printed all the headers under one and only the keys for the one under the hosts. But the one under the hosts meant to coleect headers for the hosts which were found in the site, it can be updated

### Improvements
- [ ] Dockerizing the Go Application
- [ ] Implementing Goroutines for better time complexity