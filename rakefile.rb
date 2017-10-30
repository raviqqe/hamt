task :deps do
  sh %w[
    go get -u
    github.com/client9/misspell/cmd/misspell
    github.com/golang/lint/golint
    github.com/kisielk/errcheck
    github.com/opennota/check/cmd/aligncheck
    github.com/opennota/check/cmd/structcheck
    github.com/opennota/check/cmd/varcheck
    golang.org/x/tools/cmd/goimports
    mvdan.cc/interfacer
    honnef.co/go/tools/...
  ].join ' '

  sh 'go get -d -t ./...'

  sh 'gem install rake rubocop'
end

task :lint do
  [
    'go vet',
    'golint',
    'gosimple',
    'unused',
    'staticcheck',
    'interfacer',
    'errcheck',
    'aligncheck',
    'structcheck',
    'varcheck'
  ].each do |command|
    sh "#{command} ./..."
  end

  sh 'misspell -error .'
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
  sh 'go test -cover'
end

task bench: :deps do
  sh 'go test -bench .'
end
