provider "aws" {
    region = "${var.aws_region}"
}

resource "aws_vpc" "my-vpc" {
    cidr_block = "${var.vpc_cidr}"
    enable_dns_hostnames = true
    tags = {
        Name = "Terraform"
    } 
}

# Define the public subnet
resource "aws_subnet" "public-subnet" {
    vpc_id = "${aws_vpc.my-vpc.id}"
    cidr_block = "${var.public_subnet_cidr}"
    availability_zone = "${var.public_subnet_az}"
    tags {
        Name = "Public subnet - Terraform"
    }
}

# Define the private subnet 1
resource "aws_subnet" "private-subnet-1" {
    vpc_id = "${aws_vpc.my-vpc.id}"
    cidr_block = "${var.private_subnet_1_cidr}"
    availability_zone = "${var.private_subnet_1_az}"
    tags {
        Name = "Private subnet - Terraform"
    }
}

# Define the private subnet 2
resource "aws_subnet" "private-subnet-2" {
    vpc_id = "${aws_vpc.my-vpc.id}"
    cidr_block = "${var.private_subnet_2_cidr}"
    availability_zone = "${var.private_subnet_2_az}"
    tags {
        Name = "Private subnet - Terraform"
    }
}

# Define the internet gateway
resource "aws_internet_gateway" "gw" {
  vpc_id = "${aws_vpc.my-vpc.id}"
}

# Define NAT instance security group
resource "aws_security_group" "nat" {
    name = "vpc_nat"
    description = "Allow traffic to pass from the private subnet to the internet"
    ingress {
        from_port = 0
        to_port = 0
        protocol = -1
        cidr_blocks = ["0.0.0.0/0"]
    }
    egress {
        from_port = 0
        to_port = 0
        protocol = -1
        cidr_blocks = ["0.0.0.0/0"]
    }
    vpc_id = "${aws_vpc.my-vpc.id}"
    tags {
        Name = "NATSG"
    }
}

# Define NAT instance
resource "aws_instance" "nat" {
    ami = "ami-00d1f8201864cc10c" # this is a special ami preconfigured to do NAT
    availability_zone = "${var.public_subnet_az}"
    instance_type = "t2.micro"
    key_name = "${var.aws_key_name}"
    vpc_security_group_ids = ["${aws_security_group.nat.id}"]
    subnet_id = "${aws_subnet.public-subnet.id}"
    associate_public_ip_address = true
    source_dest_check = false
    tags {
        Name = "Terraform NAT"
    }
}

resource "aws_eip" "nat" {
    instance = "${aws_instance.nat.id}"
    vpc = true
}

resource "aws_route_table" "us-east-public" {
    vpc_id = "${aws_vpc.my-vpc.id}"
    route {
        cidr_block = "0.0.0.0/0"
        gateway_id = "${aws_internet_gateway.gw.id}"
    }
    tags {
        Name = "Public subnet - Terraform"
    }
}

resource "aws_route_table_association" "us-east-public" {
    subnet_id = "${aws_subnet.public-subnet.id}"
    route_table_id = "${aws_route_table.us-east-public.id}"
}

resource "aws_route_table" "us-east-private" {
    vpc_id = "${aws_vpc.my-vpc.id}"
    route {
        cidr_block = "0.0.0.0/0"
        instance_id = "${aws_instance.nat.id}"
    }
    tags {
        Name = "Private Subnet - Terraform"
    }
}
resource "aws_route_table_association" "us-east-2b-private" {
    subnet_id = "${aws_subnet.private-subnet-1.id}"
    route_table_id = "${aws_route_table.us-east-private.id}"
}

resource "aws_route_table_association" "us-east-2c-private" {
    subnet_id = "${aws_subnet.private-subnet-2.id}"
    route_table_id = "${aws_route_table.us-east-private.id}"
}