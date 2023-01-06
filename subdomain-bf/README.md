### Usage
```bash
 ./subdomain-bf --host twitter.com -f assets/subs.txt 
```

```
assets/subs.txt

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

`Output`
```json
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


### Improvements
- [ ] Dockerizing the Go Application
- [ ] Implementing Goroutines for better time complexity

