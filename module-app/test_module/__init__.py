from flask import Flask, jsonify
import os

app = Flask(__name__)

@app.route('/', methods=['GET'])
def test():
    return jsonify({'message': 'This endpoint is running!'})


if __name__ == '__main__':
    app.run(host=os.environ['FLASK_HOST'], port=int(os.environ['FLASK_PORT']))
    