# Wallet-Service
Simple wallet service written in clean Go Rest API architecture covering dependency injections, decoupled systems, interfaces, db associations, mocking example, etc following SOLID principles.  
  
## Build 
Step 1. Clone the repository.    
  
Step 2. Run `docker-compose up` [It'll create 4 containers : App Server, DB Server, Cache Redis Server & Adminer for pg client].  
  
Step 3. (Optional) For swagger documentation,  
  
        A. Run `make swagger` [It'll fetch and generate swagger.yaml file]. 
          
        B. Run `make server-swagger` [It'll redirect to your browser rendering swagger apis on localhost:51673 ]. 
          
          
  
## Technology Stack
1. Golang (Mux Router, Gorm ORM, Go Modules, Testify)   
2. Mysql (RDBMS)   
3. Redis (Caching)  
4. Docker   
5. Swagger (go-swagger)    
6. Cron & distributed locking (robfig, redis)   

  
