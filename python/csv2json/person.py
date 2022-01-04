def process_person(item):
    return {
        "id": int(item["id"]),
        "name": item["name"],
        "age": int(item["age"]),
        "height": float(item["height"]),
        "is_active": item["is_active"] == "true",
        "created_at": item["created_at"]
    }
