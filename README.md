# VNC-RFB-Sec
 The script begins by prompting for user input for the location of their IP list. The function checkSecurityType(TCP_IP) takes an IP address, connects to port 5900 over TCP/IP, negotiates the version/security handshake and sets a flag if unauthenticated access is possible. If this flag is set, vncsnapshot is invoked in the main section of the script. The script repeats until all IPs in the list with no password have been documented with screenshots. 

------------
# Usage

Python

```python
python vnc.py
OR
python3 vnc.py
# Then input for file path | IP list
```
------------
GO

```go
go run vnc.go // Then input for file path | IP list
```
#  Remote Frame Buffer (RFB) protocol

To truly understand vnc.py/vnc.go, you have to understand how the RFB protocol performs the version handshake and the security handshake.
Full documentation on the RFB protocol can be found [here](https://tools.ietf.org/rfc/rfc6143.txt).

# vncsnapshot

https://github.com/shamun/vncsnapshot

------------

#### [more details](https://grumpy-sec.blogspot.com/2017/02/scanning-entire-internet.html)
