output "base_url" {
  description = "Base URL for API Gateway stage."
  value       = aws_api_gateway_stage.v1.invoke_url
}

output "front_endpoint" {
  description = "S3 Front Bucket endpoint."
  value       = aws_s3_bucket_website_configuration.front.website_endpoint
}

output "kms_key_id" {
  description = "KMS Key Id"
  value       = aws_kms_key.encrypt.key_id
}
