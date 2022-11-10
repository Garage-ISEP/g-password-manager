resource "aws_api_gateway_method" "secrets-get" {
  rest_api_id   = aws_api_gateway_rest_api.api.id
  resource_id   = aws_api_gateway_resource.secrets.id
  http_method   = "GET"
  authorization = "COGNITO_USER_POOLS"
  authorizer_id = aws_api_gateway_authorizer.auth.id
}

resource "aws_api_gateway_integration" "secrets-get-integration" {
  rest_api_id             = aws_api_gateway_rest_api.api.id
  resource_id             = aws_api_gateway_resource.secrets.id
  http_method             = aws_api_gateway_method.secrets-get.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = aws_lambda_function.secrets-get-lambda.invoke_arn
}

resource "aws_iam_role" "secrets-get-role" {
  name               = "${local.prefix}-get-secret-lambda-role-${local.suffix}"
  assume_role_policy = data.aws_iam_policy_document.lambda-assume-role-policy.json
  managed_policy_arns = [
    aws_iam_policy.crud.arn,
    aws_iam_policy.lambda-logging-policy.arn,
    "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  ]

  tags = local.tags
}

resource "aws_lambda_permission" "secrets-get-permission" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.secrets-get-lambda.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "arn:aws:execute-api:ap-southeast-2:${data.aws_caller_identity.current.account_id}:${aws_api_gateway_rest_api.api.id}/*/${aws_api_gateway_method.secrets-get.http_method}${aws_api_gateway_resource.secrets.path}"
}

data "archive_file" "lambda-zip-get" {
  type             = "zip"
  source_dir       = "${path.module}/../api/secrets/get"
  output_file_mode = "0666"
  output_path      = "${path.module}/archives/secrets-get.zip"
}

resource "aws_lambda_function" "secrets-get-lambda" {
  filename      = data.archive_file.lambda-zip-get.output_path
  function_name = "${local.prefix}-get-secret-lambda-${local.suffix}"
  role          = aws_iam_role.secrets-get-role.arn
  handler       = "get"

  source_code_hash = filebase64sha256(data.archive_file.lambda-zip-get.output_path)

  runtime     = "go1.x"
  memory_size = 128
  timeout     = 3

  environment {
    variables = {
      DYNAMO_TABLE = aws_dynamodb_table.table.name
    }
  }

  tags = local.tags
}

resource "aws_cloudwatch_log_group" "secrets-get-logs" {
  name = "/aws/lambda/${aws_lambda_function.secrets-get-lambda.function_name}"

  retention_in_days = 30

  tags = local.tags
}
