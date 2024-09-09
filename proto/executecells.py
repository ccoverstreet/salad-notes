with open("simplepythonplotting.md") as f:
    cells = []
    content = f.read()
    i = 0
    capture = False
    buffer = ""
    while i < len(content) - 3:
        #print(type(content[i:i+3]), content[i:i+3])
        if content[i:i+3] == "```":
            if capture and len(buffer) > 0:
                cells.append(buffer)
                buffer = ""

            if not capture:
                while content[i] != "\n":
                    i += 1


            capture = not capture

            continue

        if capture:
            buffer += content[i]
        i = i + 1

    print(cells)

for cell in cells:
    exec(cell)
