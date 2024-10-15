# github_scanner.py
import requests

GITHUB_TOKEN = 'your_github_token'  # Replace with your token
REPO = 'owner_name/repo_name'  # Replace with owner/repo

def get_repository_contents(repo):
    url = f'https://api.github.com/repos/{repo}/contents'
    headers = {'Authorization': f'token {GITHUB_TOKEN}'}
    response = requests.get(url, headers=headers)
    
    if response.status_code == 200:
        return response.json()
    else:
        print(f"Error fetching contents: {response.status_code}")
        return []

def scan_files(contents):
    for content in contents:
        if content['type'] == 'file' and content['name'].endswith('.txt'):
            print(f"Found text file: {content['name']}")

if __name__ == "__main__":
    contents = get_repository_contents(REPO)
    scan_files(contents)
