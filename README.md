## Golang restfull API with gin web framework, postgresql, nginx and docker-compose

### Installation


1. Clone this repo
   ```git clone git@github.com:Witaly3/golang-rest.git```
2. Enter repo folder
3. ``` sudo docker-compose up -d```
4. API:
 ```
GET    /api/                        
GET    /api/todos                  
POST   /api/todo                  
GET    /api/todo/:todoId           
PUT    /api/todo/:todoId             
DELETE /api/todo/:todoId 
```

5. ```sudo docker-compose down```
