from pydantic import BaseModel


class Request(BaseModel):
    request: str


class Response(Request):
    id: int
    response: str