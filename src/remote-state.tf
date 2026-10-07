module "vpc" {
  source  = "cloudposse/stack-config/yaml//modules/remote-state"
  version = "2.0.1"

  component = "vpc"

  context = module.this.context
}
