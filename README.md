# Seleksi Test Backend Developer PT. Orde Digital Intelektual

## Case Study, E-Commerce
![ERD](/assets/ERD.png "Digital-Kiosk ERD.")

## Technologies
1. Go with Gin
2. PostgreSQL
3. Docker Compose

## Project Directory

```
├── assets                  
├── cmd                      
├── docs                    
├── internal                
│   ├── app
│   │   ├── controller      
│   │   ├── mocks           
│   │   ├── model           
│   │   ├── repository      
│   │   ├── schema          
│   │   └── service         
│   └── pkg
│       ├── config
│       ├── db
│       ├── handler
│       ├── middleware
│       ├── reason
│       └── validator
```

## How To Run App
1. Rename the config file `app.env.sample` to `app.env` and fill the value. Database value (host, port, database) must equal to docker environment variables value
2. Exec `make environment` to start docker container. After the database up, make connection with your database tools such as DBeaver / Navicat / etc.
3. `make server` to run the application.
4. Import `.json` file in `./docs/` to your Postman collection. There are 2 files, one for the collection and other for the collection environment.