import sys
import time

if len(sys.argv) != 3:
    print("Нужно указать два файла")
    sys.exit(1)

with open(sys.argv[1], "r") as f1, open(sys.argv[2], "r") as f2:
    content = f1.read() + "\n" + f2.read()
    time.sleep(2)

with open("result.txt", "w") as fout:
    fout.write(content)

print("Файлы объединены в result.txt")