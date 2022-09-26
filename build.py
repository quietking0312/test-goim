import os
import shutil

os.environ["GO111MODULE"] = "on"
os.environ["CGO_ENABLED"] = "0"
os.environ["GOARCH"] = "amd64"

# darwin, windows, linux
os.environ["GOOS"] = "linux"

out_path = "target"

if os.path.isdir(out_path):
    shutil.rmtree(out_path)

os.mkdir(out_path)

shutil.copyfile("cmd/comet/comet-example.toml", "target/comet.toml")
shutil.copyfile("cmd/logic/logic-example.toml", "target/logic.toml")
shutil.copyfile("cmd/job/job-example.toml", "target/job.toml")
shutil.copyfile("examples/javascript/client.js", "target/client.js")
shutil.copyfile("examples/javascript/index.html", "target/index.html")
os.system(f"go build -o {out_path}/comet cmd/comet/main.go")
os.system(f"go build -o {out_path}/logic cmd/logic/main.go")
os.system(f"go build -o {out_path}/job cmd/job/main.go")
os.system(f"go build -o {out_path}/examples examples/javascript/main.go")