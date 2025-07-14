class GraphdbCli < Formula
  desc "GraphDB CLI tool"
  homepage "https://github.com/your/repo"
  url "file:///tmp/graphdb-cli.tar.gz"
  version "0.3.0"
  license "MIT"

  depends_on "go" => :build

  def install
    system "ls", "-la", buildpath
    system "cat", "#{buildpath}/go.mod"
    system "go", "build", "-o", bin/"gdb-cli", "main.go"
  end
end