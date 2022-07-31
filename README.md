# S3CLI


### Build

Clone

`git clone https://github.com/mutashim/s3cli.git`

Build

`go build`

### Install

`sudo ln -s /path/to/s3cli /usr/local/bin/s3cli`

### Usage

Upload from local to S3

`s3cli cp example.pdf s3://bucket/dir/dir/example.pdf` 

Download from S3 to local

`s3cli cp s3://bucket/dir/dir/example.pdf example.pdf`

Change accessibility to public

`s3cli chmod s3://bucket/dir/dir/example.pdf public`

Change accessibility to private

`s3cli chmod s3://bucket/dir/dir/example.pdf private`

Delete file from S3

`s3cli rm s3://bucket/dir/dir/example.pdf`

List file

`s3cli ls s3://bucket/dir/dir/`

List bucket

`s3cli lsbuck`

Make new bucket

`s3cli mkbuck newbucket`

Delete bucket

`s3cli rmbuck oldbucket`

## Coming soon

- Configuration file