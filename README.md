# Validators Library
Some useful data validators:
- validate email address for correct syntax
- check email domain is a disposable service or not
- check any IPv4 for blacklists
- cache data lists in memory for speed

## Disposable Email Provider
```golang
// return true
IsDisposableEmailProvider("disposableaddress.com")

// return false
IsDisposableEmailProvider("gmail.com")
```

## Suspicious IPv4
```golang
// return true
IsSuspiciousIPv4("141.98.10.125")

// return false
IsSuspiciousIPv4("1.1.1.1")
```
