data "aws_iam_policy_document" "cognito-admin-doc" {
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
    resources = ["*"]
    effect    = "Allow"
  }
}

resource "aws_iam_policy" "cognito-admin" {
  name   = "${local.prefix}-cognito-admin-${local.suffix}"
  policy = data.aws_iam_policy_document.cognito-admin-doc.json

  tags = local.tags
}


resource "aws_iam_role" "group-sync-role" {
  name               = "${local.prefix}-sync-group-google-cognito-${local.suffix}"
  assume_role_policy = data.aws_iam_policy_document.lambda-assume-role-policy.json
  managed_policy_arns = [
    aws_iam_policy.cognito-admin.arn,
    aws_iam_policy.lambda-logging-policy.arn,
    "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  ]

  tags = local.tags
}

data "archive_file" "lambda-zip-cognito-sync" {
  type             = "zip"
  source_file      = "${path.module}/../dist/cognito/sync/groups"
  output_file_mode = "0666"
  output_path      = "${path.module}/../dist/archives/sync-groups.zip"
}

resource "aws_lambda_function" "group-sync-lambda" {
  filename      = data.archive_file.lambda-zip-cognito-sync.output_path
  function_name = "${local.prefix}-cognito-group-sync-lambda-${local.suffix}"
  role          = aws_iam_role.group-sync-role.arn
  handler       = "groups"

  source_code_hash = filebase64sha256(data.archive_file.lambda-zip-cognito-sync.output_path)

  runtime     = "go1.x"
  memory_size = 128
  timeout     = 3

  environment {
    variables = {
      GOOGLE_KEY = "hello",
    }
  }

  tags = local.tags
}

resource "aws_cloudwatch_log_group" "sync-logs" {
  name = "/aws/lambda/${aws_lambda_function.group-sync-lambda.function_name}"

  retention_in_days = 30

  tags = local.tags
}
