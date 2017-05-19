resource "aws_iam_user" "registrar" {
  name = "registrar-${var.env}"
}

resource "aws_iam_access_key" "registrar" {
  user = "${aws_iam_user.registrar.name}"
}

output "registrar-aws_access_key_id" {
  value = "${aws_iam_access_key.registrar.id}"
}

output "registrar-aws_secret_access_key" {
  value = "${aws_iam_access_key.registrar.secret}"
}
