resource "aws_vpc_endpoint" "dynamodb" {
  vpc_id            = aws_vpc.vpc.id
  service_name      = "com.amazonaws.${var.aws_region}.dynamodb"
  vpc_endpoint_type = "Gateway"
  route_table_ids   = [aws_route_table.rt_a.id]
}

resource "aws_dynamodb_table" "image_metadata" {
  name         = "image-metadata"
  billing_mode = "PAY_PER_REQUEST"
  attribute {
    name = "image_id"
    type = "S"
  }
  hash_key = "image_id"
}

resource "aws_iam_policy" "dynamodb_access_policy" {
  name        = "DynamoDBFullAccessPolicy"
  description = "Full access to image metadata table"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Sid      = "AllowFullAccess",
        Effect   = "Allow",
        Action   = "dynamodb:*",
        Resource = aws_dynamodb_table.image_metadata.arn
      },
    ],
  })
}
