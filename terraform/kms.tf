resource "aws_kms_key" "encrypt" {
  description              = "${local.prefix}-kms-${local.suffix}"
  deletion_window_in_days  = 10
  key_usage                = "ENCRYPT_DECRYPT"
  customer_master_key_spec = "RSA_2048"
  tags                     = local.tags
}
