resource "aws_dynamodb_table" "buildings" {
  name           = "liszt-buildings-${var.env}"
  read_capacity  = 1
  write_capacity = 1
  hash_key       = "building_id"

  attribute {
    name = "building_id"
    type = "S"
  }
}

resource "aws_dynamodb_table" "units" {
  name           = "liszt-units-${var.env}"
  read_capacity  = 1
  write_capacity = 1
  hash_key       = "unit_id"

  attribute {
    name = "unit_id"
    type = "S"
  }

  attribute {
    name = "building_id"
    type = "S"
  }

  global_secondary_index {
    name            = "building_unit_gsi"
    hash_key        = "building_id"
    range_key       = "unit_id"
    read_capacity   = 1
    write_capacity  = 1
    projection_type = "ALL"
  }
}

resource "aws_dynamodb_table" "residents" {
  name           = "liszt-residents-${var.env}"
  read_capacity  = 1
  write_capacity = 1
  hash_key       = "resident_id"

  attribute {
    name = "resident_id"
    type = "S"
  }
}

resource "aws_iam_policy" "registrar-dynamodb-rw" {
  name        = "registrar-dynamdob-rw-${var.env}"
  description = "r/w access to liszt dynamodb tables"
  policy      = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "dynamodb:BatchGetItem",
        "dynamodb:DeleteItem",
        "dynamodb:GetItem",
        "dynamodb:PutItem",
        "dynamodb:Query",
        "dynamodb:UpdateItem"
      ],
      "Effect": "Allow",
      "Resource": [
        "${aws_dynamodb_table.buildings.arn}",
        "${aws_dynamodb_table.units.arn}",
        "${aws_dynamodb_table.units.arn}/index/*",
        "${aws_dynamodb_table.residents.arn}"
      ]
    }
  ]
}
EOF
}

resource "aws_iam_role" "registrar-rw" {
  name               = "registrar-rw-${var.env}"
  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_policy_attachment" "registry-dynamo-rw" {
  name       = "registry-dynamo-rw-${var.env}"
  roles      = ["${aws_iam_role.registrar-rw.name}"]
  users      = ["${aws_iam_user.registrar.name}"]
  policy_arn = "${aws_iam_policy.registrar-dynamodb-rw.arn}"
}
