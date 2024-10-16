import sys, os
sys.path = []
path = os.getcwd()
sys.path.append('')
sys.path.append(os.path.join(path, "python39"))
sys.path.append(os.path.join(path, "Scripts"))
sys.path.append(path)
sys.path.append(os.path.join(path, "lib", "site-packages"))

