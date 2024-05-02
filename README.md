1. Install dependencies

- Install Go (follow the instructions given)
https://go.dev/doc/install

- Install PostgreSQL (follow the instructions given)
https://www.postgresql.org/download/

Additional resources and guides:
https://www.w3schools.com/postgresql/postgresql_install.php

2. Download the source code / clone the repository from GitHub

3. Head into commandprompt or bash in the directory and type refreshenv command in case Go environment is not properly set up for the project. 

4. Check if Go is properly set up with typing go version

5. Run these commands

go build -o cloudfyp
go mod init 
go mod tidy 
go get ./..

6. Try running by running the cmmand go run . 
-Dont worry if showed error. as long as it started to listen at port 9000.
If failed continue first,

7. Copy cloudbridge_dump.sql in C drive or any drive as long as its visible and easy to access. Something like this -> C:\cloudbridge_dump.sql

8. Head into where PostgreSQL is installed. directory will look something like this C:\Program Files\PostgreSQL\16\bin   <- head into this directory 

note that 16 is the version of the postgresql

9. open up command prompt in that directory.
 type and run 
 
 psql -U postgres cloudbridge <C:\cloudbridge_dump.sql

<{this is where your cloudbridge_dump.sql file is located. make sure its visible and accessible}

enter password for your postgres

check by running query select * from "user";

will show records otherwise something went wrong

Additional resources and guides:

follow this guide if failed to retrieve database from dump file:

https://www.youtube.com/watch?v=26_Bpdf2aH8

for accessing database
user=postgres password=admin@123 dbname=cloudbridge sslmode=disable

