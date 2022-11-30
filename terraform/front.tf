resource "aws_s3_bucket" "front" {
  bucket = "${var.env != "prod" ? "${var.env}-" : ""}vault.garageisep.com"
}

resource "aws_s3_bucket_acl" "example_bucket_acl" {
  bucket = aws_s3_bucket.front.id
  acl    = "public-read"
}

resource "aws_s3_bucket_website_configuration" "front" {
  bucket = aws_s3_bucket.front.bucket

  routing_rule {
    redirect {
      replace_key_with = "index.html"
    }
  }
}

resource "aws_s3_bucket_policy" "front" {
  bucket = aws_s3_bucket.front.id
  policy = data.aws_iam_policy_document.front.json
}

data "aws_iam_policy_document" "front" {
  statement {
    sid    = "PublicReadGetObject"
    effect = "Allow"

    principals {
      type        = "*"
      identifiers = ["*"]
    }

    actions = [
      "s3:GetObject",
    ]

    resources = [
      "${aws_s3_bucket.front.arn}/*",
    ]

    condition {
      test     = "IpAddress"
      variable = "aws:SourceIp"
      values = [
        "173.245.48.0/20",
        "103.21.244.0/22",
        "103.22.200.0/22",
        "103.31.4.0/22",
        "141.101.64.0/18",
        "108.162.192.0/18",
        "190.93.240.0/20",
        "188.114.96.0/20",
        "197.234.240.0/22",
        "198.41.128.0/17",
        "162.158.0.0/15",
        "104.16.0.0/13",
        "104.24.0.0/14",
        "172.64.0.0/13",
        "131.0.72.0/22"
      ]
    }
  }
}

module "template-files" {
  source   = "hashicorp/dir/template"
  base_dir = "${path.module}/../dist/app"
}


resource "aws_s3_object" "static-files" {
  for_each = module.template-files.files

  bucket       = aws_s3_bucket.front.bucket
  key          = each.key
  content_type = each.value.content_type

  source  = each.value.source_path
  content = each.value.content

  etag = each.value.digests.md5
}
