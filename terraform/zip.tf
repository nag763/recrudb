data "archive_file" "api_function_zip" {
  type        = "zip"
  source_dir  = "../services/agents"
  output_path = "/tmp/recrudb-agents.zip"
}

resource "terraform_data" "go_binary" {
  provisioner "local-exec" {
    command = "GOOS=linux GOARCH=amd64 go build -C ../services/mcp/ -o /tmp/mcp"
  }
}

data "archive_file" "mcp_function_zip" {
  depends_on  = [terraform_data.go_binary]
  type        = "tar.gz"
  source_file = "/tmp/mcp"
  output_path = "/tmp/recrudb-mcp.zip"
}
