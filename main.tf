provider "aws" {
  region = "ap-south-1"  
}

resource "aws_vpc" "TestCicdVpc" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_instance" "example" {
  ami           = "ami-0287a05f0ef0e9d9a"
  instance_type = "t2.micro"

  tags = {
    Name = "Rajveer"
  }
}
