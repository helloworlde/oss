# OSS upload CLI tool

> This is a CLI tool for upload files to OSS, now support Aliyun only

## Quick start

- Install

```bash
go get github.com/helloworlde/oss
```

- Config

Add configuration file in `$HOME/.oss.yaml` like this:

```yaml
aliyun:
  accessKeyId: YOUR_AK
  accessSecretId: YOUR_SK
  oss:
    endPoint: YOUR_ENDPOINT
    bucketName: YOUR_BUCKET
    host: YOUR_HOST
    picturePath: /picture
```

For more convenient upload picture to OSS, so there have an `picturePath` config

## Example

- Upload file

```bash
oss upload file.zip
```

And then will print url after success like below:

```bash
Transfer Rate:100% TotalSize: 2M FinishedSize: 2M
go.mod https://hellowoodes.oss-cn-beijing.aliyuncs.com/file.zip
```

- Upload file to specify directory 

Just add `-d directory` when upload 

```bash
oss upload file.zip -d archive/file/zip
```

- Upload picture to picture dicrectory

Just add `-p` for upload to picture directory

```bash
oss upload hello.jpg -p
```

- Print as Markdown format when upload picture

Just add `-m` for print in Markdown format

```bash
oss upload hello.jpg -p -m
Transfer Rate:100% TotalSize: 0.20K FinishedSize: 0.20K
![hello.jpg](https://hellowoodes.oss-cn-beijing.aliyuncs.com/picture/hello.jpg)
```

- Upload folder

Upload folder is also support

```bash
oss upload .idea/ -d idea
```
