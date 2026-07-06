#!/usr/bin/env python3

import sys
import json
import uuid

def collect_ids(obj, ids):
    if isinstance(obj, dict):
        if "id" in obj and isinstance(obj["id"], str):
            ids.add(obj["id"])
        for v in obj.values():
            collect_ids(v, ids)
    elif isinstance(obj, list):
        for item in obj:
            collect_ids(item, ids)

def apply_id_mapping(obj, id_map):
    if isinstance(obj, dict):
        new_obj = {}
        for k, v in obj.items():
            if k == "id" and isinstance(v, str):
                new_obj[k] = id_map.get(v, v)
            elif isinstance(v, str) and v in id_map:
                new_obj[k] = id_map[v]
            else:
                new_obj[k] = apply_id_mapping(v, id_map)
        return new_obj
    elif isinstance(obj, list):
        return [apply_id_mapping(item, id_map) for item in obj]
    elif isinstance(obj, str) and obj in id_map:
        return id_map[obj]
    else:
        return obj

def main():
    if len(sys.argv) != 2:
        print(f"Usage: {sys.argv[0]} <input-json-file>")
        sys.exit(1)

    file_path = sys.argv[1]

    print(f"Loading JSON from {file_path}...")
    with open(file_path, "r") as f:
        data = json.load(f)

    old_ids = set()
    collect_ids(data, old_ids)
    print(f"Found {len(old_ids)} IDs to replace")

    id_map = {old: str(uuid.uuid4()) for old in old_ids}

    new_data = apply_id_mapping(data, id_map)

    with open(file_path, "w") as f:
        json.dump(new_data, f, indent=2)

    print(f"Rewritten IDs in {file_path}")

if __name__ == "__main__":
    main()
