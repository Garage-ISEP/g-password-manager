resource "aws_api_gateway_rest_api" "api" {
  name = "${local.prefix}-http-api-${local.suffix}"

  endpoint_configuration {
    types = ["REGIONAL"]
  }

  tags = local.tags
}

resource "aws_api_gateway_authorizer" "auth" {
  name          = "${local.prefix}-cognito-auth-${local.suffix}"
  rest_api_id   = aws_api_gateway_rest_api.api.id
  type          = "COGNITO_USER_POOLS"
  provider_arns = ["arn:aws:cognito-idp:eu-west-3:671149233773:userpool/eu-west-3_wUMdKQSFg"]
}

resource "aws_api_gateway_resource" "secrets" {
  path_part   = "secrets"
  parent_id   = aws_api_gateway_rest_api.api.root_resource_id
  rest_api_id = aws_api_gateway_rest_api.api.id
}

resource "aws_api_gateway_resource" "context-ids" {
  path_part   = "context-ids"
  parent_id   = aws_api_gateway_rest_api.api.root_resource_id
  rest_api_id = aws_api_gateway_rest_api.api.id
}

resource "aws_api_gateway_deployment" "api" {
  rest_api_id = aws_api_gateway_rest_api.api.id

  triggers = {
    redeployment = sha1(jsonencode([
      aws_api_gateway_resource.secrets.id,
      aws_api_gateway_resource.context-ids.id,
      filebase64sha256(data.archive_file.lambda-zip-get.output_path),
      filebase64sha256(data.archive_file.lambda-zip-put.output_path),
      filebase64sha256(data.archive_file.lambda-zip-delete.output_path)
    ]))
  }

  depends_on = [
    aws_api_gateway_integration.secrets-get-integration
  ]

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_api_gateway_stage" "v1" {
  deployment_id = aws_api_gateway_deployment.api.id
  rest_api_id   = aws_api_gateway_rest_api.api.id
  stage_name    = var.env

  tags = local.tags
}
