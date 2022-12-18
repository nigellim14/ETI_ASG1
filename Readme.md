5.1.3.1 Design consideration of your microservices  
For my Microservices I have planned to build an console application by splitting components into small services that could be place and operated independently from each other. So I have decided to implement it involving in development, and design of an ride sharing platform console application to store collection of passenger and driver account information upon user creation. As it primary level, each individual microservices of 'ridingplatform.go' and 'userconsole.go' acts as an application in itself. With the structure console of a collection to communicate between each services, I have selected a suitable infrastructure and persistent storage to store data into MYSQL database environment. As with this database engine it's ideal as it provide reliable and fast storage that could support storing of my data. In additional with the intention of this database it allows retrieves data collection and insert information when a user creates or update an account into the various console.  
  
5.1.3.2 Architecture diagram  
  
   
     
       
         
           
       
![Architecture diagram](https://user-images.githubusercontent.com/75166174/208292927-a457571f-9cd7-4184-bb8f-a0a501aa4304.JPG)














  

5.1.3.3 Instruction for setting up and running your microservices  
  
Step 1: Open up the Microservices of 'userconsole.go' and 'ridingplatform.go' and MYSQL database 'Storage'.  
Step 2: With the various database 'Storage' statement execute all of it, so that the table of passenger and driver will be created into the database accordingly.  
Step 3: For now let's simulate the user console as a passenger, first run the microservices of 'ridingplatform.go' with that it allows the 'userconsole.go' microservices to call the end point when user requst the option and trigger towards the server to retrieved the specific information.  
Step 4: In order to execute the option in the passenger console, you will have to open up command prompt to run 'userconsole.go' microservice file to open up to start with.  
Step 5: Upon opening up the console, go ahead to select 'Proceed to create a Passsenger Account' that is the 1st option.  
Step 6: In there you will be welcome with various options to choose from to execute your desire needs in the passenger main page.  
  
Below are the brief explaination of each options.  
Option 1) List all passengers will eventually list everything from microservices that have been created by different user.  
Option 2) Create new passengers allows you to create an passenger account. It will do it by calling the REST API then pass forward the information, and it will packet it into json and send it over.  
Option 3) Update passenger allows users to change and be updated with the latest information accordingly.  
Option 4) It brings to back to the main page where it will display the options of passenger and driver.  
Option 5) Breaks the outer case and quit the console.  
  
Step 7: Now let's simulate the user console as a driver with starting at the main console page. And before we proceed we will just required to only run the 'userconsole.go' microservices this time round. For this option 2 (Proceed to create a driver account), I have developed it by calling the end point of user request option and trigger MYSQL database to retrieve or insert information that user have requested specifically.  
  
 
Below are the brief explaination of each options. (Using MYSQL Database storage)  
Option 1) List all drivers will allows us to be retrieving a set of results that are from our database that we have stored.  
Option 2) Create new driver will enable execution to db query statement to insert the related information that we have entered into the console. To verify if our execution are achieved, we could check by proceeding to the MYSQL database to execute the 'SELECT * FROM my_db.driver;' to see if the latest user created is stored in the database.
Option 3) Update driver it allows users to change and be updated with the latest information accordingly.  
Option 4) It brings to back to the main page where it will display the options of passenger and driver.  
Option 5) Breaks the outer case and quit the console.




