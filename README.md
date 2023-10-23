# go-jobs-scheduler

Using [HTTPie](https://httpie.io/)

```sh
http http://localhost:8080/job
http http://localhost:8080/job/3
http POST http://localhost:8080/jobs id=3 name=Job3
http PUT http://localhost:8080/job/2 id=2 name=Job2 Status=failed
```
