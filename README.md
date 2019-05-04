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
<img src="https://github.com/nguyensjsu/sp19-281-awsome/blob/master/cmpe281-architecture.png" width="100%">

### Microservices
* **Auth** 
```
- /auth/ping
- /auth/login
- /auth/signup
- /auth/logout
```
* **Users**
```
- /users/ping -> "GET"
- /users -> "POST"
- /users -> "GET"
- /users/{id} -> "GET"
```
* **Jobs** 
```
- /jobs/ping -> "GET"
- /jobs  -> "POST"
- /jobs -> "GET"
- /job/{id} -> "GET"
```
## Design Principles:
### AKF Scaling Cube
* **Scaling With X-Axis** - For horizontal scaling we have created a replica set of 3 applicaton servers.Each Application servers is runnning on Amazon EC2 instances using docker. 

* **Scaling With Y-Axis** - We are achieving scaling across Y axix by decomposing the code base into 4 indepedent microservices. Each microservice is running on docker swarm cluster.

* **Scaling With Z-Axis** - We are using Mongo DB sharded cluster for /jobs and /user services. Sharded cluster consist of 2 config serviers and 2 shard servers and 1 mongos query router.

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

## Extra Credit - Terraform
We have used Terraform to launch API gateway. Following is the output for `terraform apply`
```
➜  Terraform git:(master) ✗ terraform plan
Refreshing Terraform state in-memory prior to plan...
The refreshed state will be used to calculate this plan, but will not be
persisted to local or remote state storage.
------------------------------------------------------------------------
An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create
Terraform will perform the following actions:
  + aws_api_gateway_integration.integration
      id:                                                <computed>
      cache_namespace:                                   <computed>
      connection_type:                                   "INTERNET"
      http_method:                                       "ANY"
      integration_http_method:                           "ANY"
      passthrough_behavior:                              <computed>
      request_parameters.%:                              "1"
      request_parameters.integration.request.path.proxy: "method.request.path.proxy"
      resource_id:                                       "${aws_api_gateway_resource.resource.id}"
      rest_api_id:                                       "${aws_api_gateway_rest_api.api.id}"
      timeout_milliseconds:                              "29000"
      type:                                              "HTTP_PROXY"
      uri:                                               "http://your.domain.com/{proxy}"
  + aws_api_gateway_method.method
      id:                                                <computed>
      api_key_required:                                  "false"
      authorization:                                     "NONE"
      http_method:                                       "ANY"
      request_parameters.%:                              "1"
      request_parameters.method.request.path.proxy:      "1"
```
## Technology :
* **BackEnd** - GoLang 
* **Cloud Services** - Amazon EC2, API Gateway, Network Load Balancer, Elastic Load Balancer
* **Data Store** - MongoDB, Redis for cache. 
* **FrontEnd** - React,Redux, HTML5, CSS


