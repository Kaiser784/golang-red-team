**Application type**: Command line interface tool

**Language**: Golang

### Task: Write a Golang program to bruteforce subdomains (A and CNAME records) on a given host. The program should take words from a subs.txt file. The subdomains that resolve should be logged to console in JSON format and displayed as follows:

**Example**:
```bash
$ ./bruteforcer -host twitter.com -f subdomains.txt
```

**File structure for ***subs.txt*****: Plain text file, a word on each line.
```
hello
login
app
careers
login
admin
dashboard
console
email
box
```

**Output**:
```json
{ "subdomain": "careers.twitter.com", "type": "CNAME", "resource": ["twitter.random-hr-service.com"] }
{ "subdomain": "login.twitter.com", "type": "A", "resource": ["1.1.1.1", "2.2.2.2", "3.3.3.3"] }
```
**Output format**: JSON
