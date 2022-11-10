resource "aws_dynamodb_table" "table" {
  name           = "${local.prefix}-table-${local.suffix}"
  billing_mode   = "PROVISIONED"
  read_capacity  = 20
  write_capacity = 20
  hash_key       = "pk"
  range_key      = "sk"

  attribute {
    name = "pk"
    type = "S"
  }

  attribute {
    name = "sk"
    type = "S"
  }

  tags = local.tags
}

data "aws_iam_policy_document" "dynamo-crud-doc" {
  statement {
    actions = [
      "dynamodb:BatchGetItem",
      "dynamodb:BatchWriteItem",
      "dynamodb:ConditionCheckItem",
      "dynamodb:PutItem",
      "dynamodb:DescribeTable",
      "dynamodb:DeleteItem",
      "dynamodb:GetItem",
      "dynamodb:Scan",
      "dynamodb:Query",
      "dynamodb:UpdateItem"
    ]
    resources = [aws_dynamodb_table.table.arn]
    effect    = "Allow"
  }
}

resource "aws_iam_policy" "crud" {
  name   = "${local.prefix}-dynamo-crud-${local.suffix}"
  policy = data.aws_iam_policy_document.dynamo-crud-doc.json

  tags = local.tags
}
