module github.com/aws/aws-sdk-go-v2/service/s3

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v0.31.0
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v0.4.0
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v0.2.0
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v0.4.0
	github.com/aws/smithy-go v0.5.0
	github.com/google/go-cmp v0.5.4
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding => ../../service/internal/accept-encoding/

replace github.com/aws/aws-sdk-go-v2/service/internal/presigned-url => ../../service/internal/presigned-url/

replace github.com/aws/aws-sdk-go-v2/service/internal/s3shared => ../../service/internal/s3shared/
