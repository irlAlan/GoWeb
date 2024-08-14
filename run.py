import os
import sys

GoMain = os.getcwd()+"/main.go";
ExecName = "main";
WebPageDir = os.getcwd()+"/pages";
BinDir = os.getcwd()+"/bin";

command = sys.argv; 


def build() -> bool:
    os.system(f"cd {WebPageDir} && templ generate")
    os.system(f"go build -o {BinDir} {GoMain}")
    return True


def run() -> bool:
    if build():
        os.system(f"{BinDir}/{ExecName}")
    return True



if command[1].lower() == 'b' or command[1].lower() == "build":
        build()
elif command[1].lower() == 'r' or command[1].lower() == "run":
        run()
