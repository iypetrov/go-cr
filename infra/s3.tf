resource "aws_vpc_endpoint" "s3" {
  vpc_id            = aws_vpc.vpc.id
  service_name      = "com.amazonaws.${var.aws_region}.s3"
  vpc_endpoint_type = "Gateway"
  route_table_ids   = [aws_route_table.rt_a.id]
}

resource "aws_s3_bucket" "image_storage" {
  bucket = "image-storage-202504091111"
}

resource "aws_s3_bucket_public_access_block" "image_storage_public_access_block" {
  bucket = aws_s3_bucket.image_storage.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket_policy" "image_storage_policy" {
  bucket = aws_s3_bucket.image_storage.id

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Sid    = "AllowFullAccess",
        Effect = "Allow",
        Principal = {
          AWS = "arn:aws:iam::678468774710:root"
        },
        Action = "*",
        Resource = [
          aws_s3_bucket.image_storage.arn,
          "${aws_s3_bucket.image_storage.arn}/*",
        ],
      },
    ],
  })
}
