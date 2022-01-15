project = "producer"

app "server" {
  build {
    use "pack" {}
  }

  deploy {
    use "docker" {}
  }
}