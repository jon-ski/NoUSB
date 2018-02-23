![logo](images/logo.png)  
NoUSB allows you to host a directory for another device on the network to download with a single binary.
## Why
- Don't have internet, but do have a local area network.
- No copying.. Only pasting..
- Don't trust someone's usb drive but are on the same lan?
- Potentially quicker.
## How
#### Serve Directory
~~~
nousb
// Custom port (default is 8080)
nousb -p=80
~~~
Now enter the IP:Port on the other device's browser to load up the interface.
#### Download Directory
~~~
// nousb -d -a=ServerIP:Port
nousb -d -a=192.168.1.50:8080
~~~
## Todo
[ ] Stop browser from caching download pages
[ ] Show IP address or a list of IP addresses the client should try in their browser