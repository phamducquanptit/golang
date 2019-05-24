• Setup require of project: 
    1. Databases: MySQL
    2. Framework for API: Gin - https://github.com/gin-gonic/gin
    3. Swag: using document for API
       - Link: https://github.com/swaggo/swag
       - Run command: $ go get -u github.com/swaggo/swag/cmd/swag
    4. Makefile 
    5. Config env database at: conf/config.yaml 
    6. Run query init table in database at file: databases/queryInitDB.sql
    7. Enable GO111MODULE: $ export GO111MODULE=on
• Run project: 
    1. Run project by command: $ make
    - NOTE: can instead command $ make by command: $ swag init & go run main.go in case can't run command $ make 
    2. See document API: http://localhost:3001/api/v1/docs/index.html
    3. Attached File: rest-api-postman.zip - using import postman
