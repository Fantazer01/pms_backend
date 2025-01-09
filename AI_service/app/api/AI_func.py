import g4f

def AI_func(request: str):
    request = "Напиши план для данной задачи: " + request
    response = g4f.ChatCompletion.create(
        model="gpt-3.5-turbo",
        messages=[{"role": "user", "content": request}],
        stream=True,
    )

    return ''.join(response)