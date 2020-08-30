# helper functions...
# -----------------------------------------------------------------------------
def is_repo_clean?
  current_branch = `git rev-parse --abbrev-ref HEAD`.strip
  any_changes = `git status -s | wc -l`.strip.to_i
  if any_changes == 0
    true
  else
    false
  end
end

def command_exists?(cmd)
  cmd_check = `command -v #{cmd} > /dev/null 2>&1 && echo $?`.chomp
  cmd_check.length == 0 ? false : true
end

def current_version(lookup_file=".bumpversion.cfg")
  file = File.open(lookup_file, "r")
  data = file.read
  file.close
  match = /current_version = (\d+).(\d+).(\d+)/.match(data)
  "#{match[1]}.#{match[2]}.#{match[3]}"
end
# -----------------------------------------------------------------------------

# tasks
# -----------------------------------------------------------------------------
desc "Default task, show avaliable tasks"
task :default do
  system("rake -sT")
end

AVAILABLE_REVISIONS = ["major", "minor", "patch"]
task :bump, [:revision] do |_, args|
  args.with_defaults(revision: "patch")
  abort "bumpversion command not found..." unless command_exists?("bumpversion")
  abort "Please provide valid revision: #{AVAILABLE_REVISIONS.join(',')}" unless AVAILABLE_REVISIONS.include?(args.revision)

  system "bumpversion #{args.revision}"
end

desc "Run tests"
task :test, [:verbose] do |_, args|
  args.with_defaults(verbose: "")
  system "go test #{args.verbose} ./..."
end

desc "Run doc server"
task :serve_doc, [:port] do |_, args|
  args.with_defaults(port: 6060)
  url = "http://127.0.0.1:#{args.port}"

  pages = [
    'bin-checker-go', 
  ]

  puts "Doc server is running at: #{url}...\n\n"
  pages.each do |u|
    puts "#{url}/pkg/github.com/promptapi/#{u}/"
  end
  puts "\n"

  system "godoc -http=:#{args.port}"
end

ORGANIZATION_NAME = "promptapi"
REPO_NAME = "bin-checker-go"

namespace :release do
  desc "Do release check"
  task :check do
    puts "-> go mod tidy"
    system "go mod tidy"
    
    puts "-> run tests"
    Rake::Task["test"].invoke
    
    puts "-> is repo clean ?"
    if is_repo_clean?
      puts "you are good to go... run:"
      puts "\trake release:publish[REVISON]"
      puts "\trake release:publish        # patch"
      puts "\trake release:publish[minor] # minor"
      puts "\trake release:publish[major] # major"
      puts "current version: #{current_version}"
    else
      puts "<- please commit your changes first..."
    end
  end

  desc "Publish project with revision: #{AVAILABLE_REVISIONS.join(',')}, default: patch"
  task :publish, [:revision] do |_, args|
    args.with_defaults(revision: "patch")
    abort "please commit your changes first..." unless is_repo_clean?

    Rake::Task["bump"].invoke(args.revision)

    current_git_tag = "v#{current_version}"
    puts "-> new version: #{current_git_tag}"
    puts "-> pushing tag #{current_git_tag} to remote..."
    system %{
      git push origin #{current_git_tag} &&
      go list -m github.com/#{ORGANIZATION_NAME}/#{REPO_NAME}@#{current_git_tag}
    }

    current_branch = `git rev-parse --abbrev-ref HEAD`.strip

    puts "-> updating/pushing #{current_branch} branch"
    system "git push origin #{current_branch}"
    
    puts "-> all complete!"
  end
end
# -----------------------------------------------------------------------------
