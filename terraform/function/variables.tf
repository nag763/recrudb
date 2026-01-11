variable "region" {
  type    = string
  default = "europe-west9"
}

variable "name" {
  type = string
}

variable "entrypoint" {
  type = string
}

variable "roles" {
  type    = set(string)
  default = []
}

variable "has_acces_to_mcp" {
  type    = bool
  default = true
}

variable "environment_variables" {
  type = map(string)
  default = {
  }
}

variable "bucket" {
  type = string

}

variable "object" {
  type = string
}

variable "network" {
  type = string
}

variable "subnetwork" {
  type = string
}

variable "public_access" {
  type    = bool
  default = false
}
