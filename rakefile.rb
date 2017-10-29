task :build do
  sh 'go build'
end

task :test do
  sh 'go test -cover'
end
