**Application type**: Command line interface tool

**Language**: Golang

### Task: Write an HTTP web scanner in Golang that will take hosts (domains, subdomains, IP addresses) from an input file and scan the web pages for the following things:
- [x]  HTTP headers
- [x] Page title
- [x] HTTP headers
- [x] TLS certificates hosts -  Subject Alt Names
- [x] Hosts present in HTML
- [x] Emails addresses
- [x] IP addresses
- [x] S3 buckets
- [x] Social media links (twitter/facebook/instagram)
- [x] Copyrights


**Input format**: Plain text file, each host separated by one line.
**Input type**: domains, subdomains, IP addresses (IPv4/IPv6)

```
1.1.1.1
one.one.one.one
redhuntlabs.com
```

**Output structure**:
```go
{
    Host    string
    Headers map[string]string
    Title string
    Hosts struct {
        Certificate []string
        Headers     []string
        HTML        []string
    }
    Other struct {
        Emails    []string
        Ips       []string
        Buckets   []string
        Social    []string
        Copyright string
    }
}
```
**Output format**: JSON