# github_scanner.rb
require 'net/http'
require 'json'

GITHUB_TOKEN = 'your_github_token'  # Replace with your token
REPO = 'owner_name/repo_name'  # Replace with owner/repo

def get_repository_contents(repo)
  uri = URI("https://api.github.com/repos/#{repo}/contents")
  request = Net::HTTP::Get.new(uri)
  request['Authorization'] = "token #{GITHUB_TOKEN}"

  response = Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) do |http|
    http.request(request)
  end

  if response.is_a?(Net::HTTPSuccess)
    JSON.parse(response.body)
  else
    puts "Error fetching contents: #{response.code}"
    []
  end
end

def scan_files(contents)
  contents.each do |content|
    if content['type'] == 'file' && content['name'].end_with?('.md')
      puts "Found markdown file: #{content['name']}"
    end
  end
end

contents = get_repository_contents(REPO)
scan_files(contents)
