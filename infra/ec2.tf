resource "aws_security_group" "vm_sg" {
  vpc_id = aws_vpc.vpc.id
  ingress = [
    {
      cidr_blocks      = ["0.0.0.0/0"]
      description      = "everything"
      from_port        = 0
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = -1
      security_groups  = []
      self             = false
      to_port          = 0
    }
  ]
  egress = [
    {
      cidr_blocks      = ["0.0.0.0/0"]
      description      = "everything"
      from_port        = 0
      ipv6_cidr_blocks = []
      prefix_list_ids  = []
      protocol         = -1
      security_groups  = []
      self             = false
      to_port          = 0
    }
  ]

  timeouts {
    delete = "2m"
  }
}

resource "aws_iam_role" "vm_role" {
  name = "role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Service = "ec2.amazonaws.com"
        }
        Action = "sts:AssumeRole"
      }
    ]
  })
}

resource "aws_iam_role_policy_attachment" "vm_role_policy_attachment" {
  role       = aws_iam_role.vm_role.name
  policy_arn = "arn:aws:iam::aws:policy/AdministratorAccess"
}

resource "aws_iam_instance_profile" "vm_profile" {
  name = "profile"
  role = aws_iam_role.vm_role.name
}

resource "aws_instance" "vm" {
  ami                         = "ami-07c1b39b7b3d2525d"
  instance_type               = "t2.micro"
  subnet_id                   = aws_subnet.public_subnet_a.id
  vpc_security_group_ids      = [aws_security_group.vm_sg.id]
  iam_instance_profile        = aws_iam_instance_profile.vm_profile.name
  user_data_replace_on_change = true
  user_data                   = <<-EOF
    #!/bin/bash
    apt-get update -y
    apt-get install -y curl wget unzip vim tmux git
    curl -fsSl https://get.docker.com | sh
    gpasswd -a ubuntu docker
    docker swarm init

    docker service create \
      --name cloudflared \
      --network host \
      --replicas 1 \
      --update-parallelism 1 \
      --update-delay 10s \
      --update-order start-first \
      --rollback-parallelism 1 \
      --rollback-delay 10s \
      --rollback-order stop-first \
      --restart-condition on-failure \
      --restart-max-attempts 3 \
      --env TUNNEL_TOKEN=${random_string.cf_tunnel_secret.result} \
      cloudflare/cloudflared:2025.2.1 \
      tunnel run --protocol http2
    
    docker service create \
      --name go-cr \
      --network host \
      --port 8080:8080 \
      --replicas 1 \
      --update-parallelism 1 \
      --update-delay 10s \
      --update-order start-first \
      --rollback-parallelism 1 \
      --rollback-delay 10s \
      --rollback-order stop-first \
      --restart-condition on-failure \
      --restart-max-attempts 3 \
      --env APP_ENV="prod" \
      --env APP_DOMAIN="cr.ip812.com" \
      --env APP_PORT="8080" \
      --env AWS_ACCESS_KEY_ID=${var.aws_access_key} \ 
      --env AWS_SECRET_ACCESS_KEY=${var.aws_secret_key} \
      --env AWS_REGION=${var.aws_region} \
      iypetrov/go-cr:1.0.0
  EOF

  lifecycle {
    replace_triggered_by = [
      aws_security_group.vm_sg.name,
      aws_security_group.vm_sg.ingress,
      aws_security_group.vm_sg.egress
    ]
  }
}
