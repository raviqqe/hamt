task :deps do
  sh 'go get github.com/alecthomas/gometalinter'
  sh 'gometalinter --install'
  sh 'go get -d -t ./...'
  sh 'gem install rake rubocop'
end

task :lint do
  sh 'gometalinter --enable gofmt --enable goimports --enable misspell ./...'
end

task :format do
  sh 'go fix ./...'
  sh 'go fmt ./...'

  Dir.glob '**/*.go' do |file|
    sh "goimports -w #{file}"
  end

  sh 'rubocop -a'
end

task :build do
  sh 'go build'
end

task :test do
  sh 'go test -covermode atomic -coverprofile coverage.txt'
end

task :bench do
  sh 'go test -bench .'
end
