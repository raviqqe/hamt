task :deps do
  sh 'go get github.com/stretchr/testify/assert'
end

task build: :deps do
  sh 'go build'
end

task test: :deps do
  sh 'go test -cover'
end
