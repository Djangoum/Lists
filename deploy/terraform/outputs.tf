output "app_name" {
    value = aws_elastic_beanstalk_application.application.name
}

output "app_environment_name" {
    value = aws_elastic_beanstalk_environment.environment.name
}