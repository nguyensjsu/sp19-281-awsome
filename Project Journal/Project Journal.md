# CMPE281-Team Hackathon Project
# Project Journal
## WEEK1
* Decided to build a clone of LinkedIn Application.
* Discussed the architecture and flow of the application.
* Discussed services required for application backend.
* Decided to implement following Microservices and each team member was assigned one service:\
   **Auth** \
   **Users** \
   **Jobs** \
   **Search** 
* Decided on technologies - 
  - **'GoLang'** for backend services
  - **'React'** for frontend

## WEEK2
* Each team member worked on learning GoLang and implementing the services.
* Worked on implementation of Login,Logout and Signup functionality **(/auth/login,/auth/signup, /auth/logout)** in Auth Service.
* Implemented API's for posting a job and fetching the list of all job **(/jobs - GET and POST)** in Jobs service.
* Implemented API's for creating user and fetching all users **(/create/user ,/getUsers)** in Users Service.
* Discussed on the Data store options for the Application. Decided to go with MongoDB and Redis for cache.

## WEEK3
* Implemented Front End - 
1. Login, Signup page
2. Profile Page
3. Application home page
* Discussed on cloud scaling strategies for the applicaton. Following the **AKF scaling cube** - \
  Scaling with X-axis : Building the load balanced replica set of 3 application servers for each microservice.\
  Scaling with Y-axis : Application is splitted in 4 Microservices \
  Scaling with Z-axis : Going with data sharding across muliple availability zones. 

## WEEK 4
* Each team member started working on integrating their services.
* Front End Integration of Services to the go API's and check Response and Request Payload.
* Configuring Micro service Instances, Database Instances and API Gateway.
* Finalize Go API and Buildi the Docker images and deploy it to docker hub.
* Run the Docker containers in respective EC2 instances and link them to the mongo Cluster.
* Deploying Frontend in Heroku and changing the Respective constant URL's.
