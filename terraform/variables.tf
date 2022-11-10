data "aws_caller_identity" "current" {}

variable "env" {
  description = "Environement name"
  type        = string
  default     = "dev"
}

variable "stack-name" {
  description = "Stack name"
  type        = string
  default     = "g-password-manager"
}
