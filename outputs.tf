output "ecr_arn" {
  value = aws_ecr_repository.image.arn
}
output "ecr_id" {
  value = aws_ecr_repository.image.id
}
output "ecr_name" {
  value = aws_ecr_repository.image.name
}
output "ecr_repository" {
  value = aws_ecr_repository.image.repository_url
}