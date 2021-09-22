from flask import Flask
import os
import time

app = Flask(__name__)

@app.route("/")
def hello():
    print ("Connect")
    time.sleep(3)
    return "Hello 3 sec."


if __name__ == "__main__":
    port = int(os.environ.get("PORT", 5000))
    app.run(debug=True,host='0.0.0.0',port=port)
