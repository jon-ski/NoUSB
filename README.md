![logo](images/logo.png)  
Serve a directory on your lan
## How
#### Serve Directory
~~~
cd /folder/you/want/to/host/
nousb
// Custom port (default is 8080)
nousb -p=80
~~~
Enter the IP:Port on the other device's browser to load up the web ui.
#### Download Directory
~~~
cd /folder/you/want/to/download/to
// nousb -d -a=ServerIP:Port
nousb -d -a=192.168.1.50:8080
~~~