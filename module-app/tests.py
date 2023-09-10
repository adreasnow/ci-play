import pytest
from test_module.__main__ import app

@pytest.fixture
def client():
    app.config['TESTING'] = True
    with app.test_client() as client:
        yield client

def test_hello_world(client):
    response = client.get('/')
    assert response.status_code == 200
    assert response.data.decode('utf-8') == '{"message":"This endpoint is running!"}\n'
