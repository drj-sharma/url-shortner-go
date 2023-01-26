URL SHORTNER


## Installation & testing

Use below as package to install glogger in your Go project

```
https://github.com/drj-sharma/url-shortner-go
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
    
