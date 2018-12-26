"""
    Example command:
        curl -F 'dbsql=@/home/xvzf/Downloads/DHL_label_2018-12-12_16_9_50.pdf' http://127.0.0.1:1337/test
"""
from sanic import Sanic
from sanic.response import json
from datetime import datetime
from os import environ, path

app = Sanic()


@app.route("/upload/obl_sqlbackup", methods=["POST"])
async def test(request):
    # test.strftime("%d%m%y_%H_%M_%S")

    with open(
        path.join(
            environ.get("BACKUP_DIR") or "./backup",
            f"backup-{datetime.now().strftime('%d%m%y_%H_%M_%S')}.sql"
        ),
        "wb"
    ) as f:
        f.write(request.files.get("backup.sql").body)

    return json(True)

if __name__ == "__main__":
    app.run("127.0.0.1", 1337)
