# CMPE281 - Team Hackathon Project : LinkedIn
## Team : AWSome
## Team Members

- Avinav Tyagi
- Mayur Barge
- Nikita Bairagi
- Dharma Dheeraj Chintala

## Project Journal
Project Journal available at following locations for **AWSome** project.

[Week1 Journal](Project%20Journal/Week1.md)

[Week2 Journal](Project%20Journal/Week2.md)

[Week3 Journal](Project%20Journal/Week3.md)

[Week4 Journal](Project%20Journal/Week4.md)

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

## Docker Swarm Cluster:
For replication of services we have used docker swarm as container orchestrater.
You will need to initialize the swarm on the Manager node. Following command is used to initialize docker in swarm mode
```
docker swarm init --advertise-addr <PRIVATE_IP_ADDRESS>
```
You should see following in the output
```
docker swarm join --token SWMTKN-1-4b9w5yrifkwqhs32zdsfg1ml3y12277qz1u1x24c5gpbt8gg80-1e1r92ksw2mzohcj0ffymwsb3 <PRIVATE_IP_ADDRESS>:2377
```
Run above command in worker node to join other nodes to cluster
You can verify swarm cluster status using `docker info` and `docker node ls` command

### CAP Principle: 

## Technology :
* **BackEnd** - GoLang 
* **Cloud Services** - Amazon EC2, API Gateway, Network Load Balancer, Elastic Load Balancer
* **Data Store** - MongoDB, Redis for cache. 
* **FrontEnd** - React,Redux, HTML5, CSS


