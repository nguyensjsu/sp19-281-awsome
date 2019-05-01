# CMPE281 - Team Hackathon Project : LinkedIn
## Team : AWSome
## Team Members

- Avinav Tyagi
- Mayur Barge
- Nikita Bairagi
- Dharma Dheeraj Chintala

## Project Journal
[Link to Project Journal](https://github.com/nguyensjsu/sp19-281-awsome/blob/master/Project%20Journal.md)

## Architecture:
### Microservices
* **Auth** 
```
- /auth/login
- /auth/signup
- /auth/logout
```
* **Search**
```
- /search/jobs
- /search/users
```
* **Users**
```
- /ping -> "GET"
- /users -> "POST"
- /users -> "GET"
- /users/{id} -> "GET"
```
* **Jobs** 
```
- /ping -> "GET"
- /jobs  -> "POST"
- /jobs -> "GET"
- /job/{id} -> "GET"
```
## Design Principles:
### AKF Scaling Cube
* **Scaling With X-Axis** - Horizontal Scaling using ELBs and multiple servers for each service.
* **Scaling With Y-Axis** - MicroService Architecture comprising of 4 replicated services running independently.
* **Scaling With Z-Axis** - The system uses sharded MongoDB databases for each service.

### CAP Principle: 

## Technology :
* **BackEnd** - GoLang 
* **Cloud Services** - Amazon EC2, API Gateway, Network Load Balancer, Elastic Load Balancer
* **Data Store** - MongoDB, Redis for cache. 
* **FrontEnd** - React,Redux, HTML5, CSS


