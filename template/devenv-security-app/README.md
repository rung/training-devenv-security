# devenv-security-rung-application
- Run on local machine
```bash
$ docker build -t training-helloworld .
$ docker run -p 8080:8080 training-helloworld
```

- Terraform can manage deployment, but for simplicity of this training, the actions workflows build and deply this application