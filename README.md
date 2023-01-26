URL SHORTNER


## Installation & testing

Use below as package to install glogger in your Go project

```
https://github.com/drj-sharma/url-shortner-go
```

NOTE: Use your own MySQL config, update below values accordingly
```
Username = "root"
Password = "12345678"
Hostname = "127.0.0.1:3306"
DBname   = "urlshortner"
```

```
Migration support has already been done in this project (whenever server starts), that means you don't have to create table. 
It will automatically create the table using Migration.
You only have to create database and update that database in Config Object -> /config/config.go
```


CURL to get shortner url against given url
provide url (full path url with http) to the url (param) It will return shortner url (e.g. localhost:8080/510e9896-19d) 
that will always redirect to original url (provided one)
```
curl --location --request \
GET 'localhost:8080/get-short-url?url=http://www.google.com' \
--data-raw ''
```

```
for example, 
we got localhost:8080/510e9896-19d against http://www.google.com, new url will redirect to http://www.google.com
```

---
&nbsp;

Easiest way to test this library by using Docker

```bash
docker build -t app .

docker run app
```

OR

Install make tool (https://askubuntu.com/questions/161104/how-do-i-install-make)

then,
``` bash
make docker-build

make docker-run
    
