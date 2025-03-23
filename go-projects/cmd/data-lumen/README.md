Data-Lumen is consists of command line tool which is used to load the csv data into database and provides anAPI to do analysis on the data uploaded.

**Requirements:**

1. Go
2. Database - PostgreSQL
    Used packages 
    - "gorm.io/driver/postgres"
	- "gorm.io/gorm"
    
3. Backend Service - Supabase 
    - Uses postgresql(cloud service which is free to create tables, manages data and authentication support)


**Execution:**

1. ENV **DB_DSN** should be configured. Example: "postgresql://postgres.mqqizrfgukdbygnbmebl:YOUR_PWD@aws-0-ap-south-1.pooler.supabase.com:6543/postgres".


2. Load data
    - cd go-projects/cmd/data-lumen/cli
    - ```./cli load --file-path data-lumen/input.csv``` ---->  If tables already created(Run this as table already there)
    - ```./cli load --file-path data-lumen/input.csv --create-tables``` ----> If tables not exists already.



    ![image1](images/customers.png)


    ![image2](images/items.png)


    ![image3](images/orders.png)


    ![image4](images/products.png)


    ![image5](images/regions.png)




3. API call
    - cd go-projects/cmd/data-lumen/api
    - Run ```./api```
    - Sample Requests
        - ```curl -X GET "http://localhost:8080/revenue?start_date=2023-12-01&end_date=2024-03-01" | jq```
        - ```curl -X GET "http://localhost:8080/revenue_by_product?start_date=2023-12-01&end_date=2024-05-01" | jq```
        - ```curl -X GET "http://localhost:8080/revenue_by_category?start_date=2023-12-01&end_date=2024-05-01" | jq```
        - ```curl -X GET "http://localhost:8080/revenue_by_region?start_date=2023-12-01&end_date=2024-05-01" | jq```


    ![image6](images/total_revenue.png)


    ![image7](images/revenue_by_product.png)


    ![image8](images/revenue_by_category.png)


    ![image9](images/revenue_by_region.png)

