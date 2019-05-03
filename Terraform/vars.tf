variable "aws_region" {
  description = "Region for the VPC"
  default = "us-east-2"
}

variable "vpc_cidr" {
  description = "CIDR for the VPC"
  default = "10.1.0.0/16"
}

variable "public_subnet_cidr" {
  description = "CIDR for the public subnet"
  default = "10.1.0.0/24"
}

variable "public_subnet_az" {
  description = "AZ for the public subnet"
  default = "us-east-2a"
}

variable "private_subnet_1_cidr" {
  description = "CIDR for the private subnet"
  default = "10.1.1.0/24"
}

variable "private_subnet_1_az" {
  description = "AZ for the private subnet 1"
  default = "us-east-2b"
}

variable "private_subnet_2_cidr" {
  description = "CIDR for the private subnet 2"
  default = "10.1.2.0/24"
}

variable "private_subnet_2_az" {
  description = "AZ for the private subnet 2"
  default = "us-east-2c"
}

variable "aws_key_name" {
    description = "Key name for NAT instance"
    default = "mayur"
}
