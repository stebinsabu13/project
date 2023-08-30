# project

# the project contains 1 api gateway, where all the api's are residing and corresponding microservice is invoked based on the request.

# for cloning the repository

run the following command on the terminal   git clone https://github.com/stebinsabu13/project.git

# Configure the .env files for binding the necessary environment variable for the api-gateway and for both the microservices

# Requesting endpoints

1) CREATING A NEW USER - POST localhost:4000/user ,request this endpoint with this following json binding 
   {
      "firstname":"abc",
      "lastname":"def",
      "email":"abc123@gmail.com",
      "mobilenum":"1234567890"
  }
  here the email and mobilenum fields should be in the valid format of an email and mobile number 

2) GET USER BY ID  - GET localhost:4000/user?id=1, request this endpoint. The user is retreived if user exists or else an error message is shown

3) UPDATE USER BY ID - PATCH localhost:4000/user?id=1, request this endpoint with this following json binding 
   {
      "firstname":"abc",
      "lastname":"def",
      "email":"abc123@gmail.com",
      "mobilenum":"1234567890"
    }

4) DELETE USER BY ID - DELETE localhost:4000/user?id=1, request this endpoint. The user is deleted.
5) API FOR DEMONSTRATING CONCURRENCY AND SEQUANTIAL EXECUTION  -  POST  localhost:4000/methods, request this endpoin with the following json binding
  {
    "method":1,
    "waitTime":30
  } 
